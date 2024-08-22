package medium

import (
	"bytes"
	"encoding/json"
	"fmt"
	"go-web-app/models"
	"io"
	"log"
	"net/http"
)

func WXWork(key *models.NotiAPI) (err error) {
	dataJsonStr := fmt.Sprintf(`{"msgtype": "text", "text": {"content": "%s", "mentioned_list": [%s]}}`, key.Text, key.WorkAtuser)
	resp, err := http.Post(
		*key.WorkApiUrl,
		"application/json",
		bytes.NewBuffer([]byte(dataJsonStr)))
	if err != nil {
		fmt.Println("weworkAlarm request error")
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)
	return
}

func SendMessage(token, content string, maxBytes int) error {
	url := fmt.Sprintf("https://qyapi.weixin.qq.com/cgi-bin/webhook/send?key=%s", token)

	// 检查消息内容是否超出限制
	if len(content) > maxBytes {
		// 超出限制，拆分消息内容并依次发送
		segments := SplitMessage(content, maxBytes)
		for _, segment := range segments {
			err := SendMessageSegmentMarkDown(url, segment)
			if err != nil {
				return err
			}
		}
	} else {
		// 未超出限制，直接发送消息
		err := SendMessageSegmentMarkDown(url, content)
		if err != nil {
			return err
		}
	}

	return nil
}
func SendMessageSegmentMarkDown(url, content string) error {
	// 构造消息结构体
	message := struct {
		MsgType  string `json:"msgtype"`
		Markdown struct {
			Content string `json:"content"`
		} `json:"markdown"`
	}{
		MsgType: "markdown",
		Markdown: struct {
			Content string `json:"content"`
		}{
			Content: content,
		},
	}
	payload, err := json.Marshal(message)
	if err != nil {
		return err
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payload))
	if err != nil {
		return err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Println("reboot mag err:", err)
		}
	}(resp.Body)

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to send message, status code: %d", resp.StatusCode)
	}

	return nil

}

// SplitMessage 将消息内容按照指定的长度拆分成多个小片段
func SplitMessage(content string, maxBytes int) []string {
	var segments []string
	runes := []rune(content)

	for len(runes) > 0 {
		// 计算当前片段的长度
		var segmentLength int
		for i, r := range runes {
			segmentLength += len(string(r))
			if segmentLength > maxBytes {
				// 当前片段长度超过了最大字节数，截取并保存当前片段
				segments = append(segments, string(runes[:i]))
				runes = runes[i:]
				break
			}
		}
		if segmentLength <= maxBytes {
			// 当前片段长度未超过最大字节数，保存整个内容并退出循环
			segments = append(segments, string(runes))
			break
		}
	}

	return segments
}

func DingDing(key *models.NotiAPI) (err error) {
	dataJsonStr := fmt.Sprintf(
		`{"at": {"atMobiles":["%s"],"atUserIds":["user123"],"isAtAll": false},"text": {"content":"%s"},"msgtype":"text"}`,
		key.DingAtuser, key.Text)
	resp, err := http.Post(
		*key.DingApiUrl,
		"application/json",
		bytes.NewBuffer([]byte(dataJsonStr)))
	if err != nil {
		fmt.Println("weworkAlarm request error")
		return
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(resp.Body)
	return
}
