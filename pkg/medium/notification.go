package medium

import (
	"bytes"
	"fmt"
	"go-web-app/models"
	"io"
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
