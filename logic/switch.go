package logic

import (
	"bufio"
	"fmt"
	"github.com/ziutek/telnet"
	"go-web-app/dao/mysql"
	"go-web-app/models"
	"go-web-app/pkg/medium"
	"go-web-app/settings"
	"go.uber.org/zap"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func SelectSwitchOption(p *models.SelectSwitchMac) (interface{}, error) {
	username := settings.Conf.Username
	password := settings.Conf.Passtoken
	token := settings.Conf.ApiToken
	var ClientSwtichInfo models.ClientSwitchInfo

	switch p.SwitchLevel {
	case 801:
		linkswitch := "8.1"
		switchconn, switchtype, err := connectAndValidate(buildSwitchAddress(linkswitch), username, password)
		if err != nil {
			return nil, err
		}

		for {
			s, err, switchinfo := SelectSwitch(switchconn, switchtype, p.ShortMAC, linkswitch)
			if err != nil {
				zap.L().Named(switchinfo.SwitchName)
				return ClientSwtichInfo, err
			}
			if switchinfo.SwitchNote == "" {
				var recodchang strings.Builder
				fmt.Println(buildSwitchAddress(linkswitch))
				switchvlan, switchport, _ := SelectSwitchClient(switchconn, switchtype, p.ShortMAC, buildSwitchAddress(linkswitch))
				zap.L().Info("工位接入", zap.String(switchvlan, switchport))
				switchvlan, switchport, err = ChangStitchPort(switchconn, switchtype, p.ShortMAC, switchport, p.ChangVlan)
				ClientSwtichInfo.Vlan = switchvlan
				ClientSwtichInfo.SwitchPort = switchport
				ClientSwtichInfo.SwitchName = linkswitch
				recodchang.WriteString("本次修改的交换机端口配置如下：\n")
				recodchang.WriteString("Vlan：" + switchvlan)
				recodchang.WriteString("\n设备连接端口：" + switchport)
				recodchang.WriteString("\n设备所属交换机：" + linkswitch)
				err = medium.SendMessage(token, recodchang.String(), 4096)
				zap.L().Info("已修改信息", zap.String(switchvlan, switchport))
				recodchang.Reset()
				return ClientSwtichInfo, err
			}

			linkswitch = switchinfo.DownLinkSwitch
			switchconn, switchtype, err = connectAndValidate(buildSwitchAddress(linkswitch), username, password)
			if err != nil {
				return nil, err
			}

			if switchinfo.SwitchName == switchinfo.DownLinkSwitch {
				return processSwitchOutput(switchtype, s, switchinfo)
			}
		}

	case 502:
		linkswitch := "5.6"
		switchconn, switchtype, err := connectAndValidate(buildSwitchAddress(linkswitch), username, password)
		if err != nil {
			return nil, err
		}

		for {
			s, err, switchinfo := SelectSwitch(switchconn, switchtype, p.ShortMAC, linkswitch)
			if err != nil {
				zap.L().Named(switchinfo.SwitchName)
				return ClientSwtichInfo, err
			}
			if switchinfo.SwitchNote == "" {
				var recodchang strings.Builder
				switchvlan, switchport, _ := SelectSwitchClient(switchconn, switchtype, p.ShortMAC, buildSwitchAddress(linkswitch))
				zap.L().Info("工位接入", zap.String(switchvlan, switchport))
				switchvlan, switchport, err = ChangStitchPort(switchconn, switchtype, p.ShortMAC, switchport, p.ChangVlan)
				fmt.Println(switchvlan, switchport, linkswitch)
				ClientSwtichInfo.Vlan = switchvlan
				ClientSwtichInfo.SwitchPort = switchport
				ClientSwtichInfo.SwitchName = linkswitch
				recodchang.WriteString("本次修改的交换机端口配置如下：\n")
				recodchang.WriteString("Vlan：" + switchvlan)
				recodchang.WriteString("\n设备连接端口：" + switchport)
				recodchang.WriteString("\n设备所属交换机：" + linkswitch)
				err = medium.SendMessage(token, recodchang.String(), 4096)
				zap.L().Info("已修改信息", zap.String(switchvlan, switchport))
				recodchang.Reset()
				return ClientSwtichInfo, err
			}

			linkswitch = switchinfo.DownLinkSwitch
			switchconn, switchtype, err = connectAndValidate(buildSwitchAddress(linkswitch), username, password)
			if err != nil {
				return nil, err
			}

			if switchinfo.SwitchName == switchinfo.DownLinkSwitch {
				return processSwitchOutput(switchtype, s, switchinfo)
			}
		}

	default:
		return nil, fmt.Errorf("未知的 SwitchLevel: %d", p.SwitchLevel)
	}
}
func ConnectSwitch(DIP, DUser, DPwd string) (err error, switchtype string, conn *telnet.Conn) {
	conn, err = telnet.Dial("tcp", DIP)
	if err != nil {
		return fmt.Errorf("连接到 Telnet 服务器时发生错误: %v", err), "unknown", conn
	}

	conn.SetUnixWriteMode(true)

	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	// 登录步骤
	if err := LoginSwtich(conn, reader, writer, DUser, DPwd); err != nil {
		conn.Close()
		return err, "unknown", conn
	}

	// 查询交换机时间信息
	output, err := SendCommand(conn, reader, writer, "display clock")
	if err != nil {
		conn.Close()
		return err, "unknown", conn
	}

	// 检测交换机类型
	switchType, err := DetectSwitchType(output)
	if err != nil {
		conn.Close()
		return err, "unknown", conn
	}

	return nil, switchType, conn
}
func LoginSwtich(conn *telnet.Conn, reader *bufio.Reader, writer *bufio.Writer, user, password string) error {
	fmt.Println("等待用户名提示符...")
	if err := WaitForPrompts(conn, reader, []string{"Username:", "login:"}, 20*time.Second); err != nil {
		return err
	}
	fmt.Println("发送用户名...")
	if err := SendAndReceive(writer, user); err != nil {
		return err
	}
	fmt.Println("等待密码提示符...")
	if err := WaitForPrompts(conn, reader, []string{"Password:"}, 20*time.Second); err != nil {
		return err
	}
	fmt.Println("发送密码...")
	if err := SendAndReceive(writer, password); err != nil {
		return err
	}
	fmt.Println("等待登录提示符...")
	return WaitForPrompts(conn, reader, []string{">", "#"}, 20*time.Second)
}

func SendCommand(conn *telnet.Conn, reader *bufio.Reader, writer *bufio.Writer, command string) (string, error) {
	if err := SendAndReceive(writer, command); err != nil {
		return "", err
	}
	return ReadOutput(conn, reader)
}
func SendCommandAll(conn *telnet.Conn, reader *bufio.Reader, writer *bufio.Writer, command string) (string, error) {
	if err := SendAndReceive(writer, command); err != nil {
		return "", err
	}
	return ReadOutputAll(conn, reader)
}

func SendCommandChang(conn *telnet.Conn, reader *bufio.Reader, writer *bufio.Writer, command string) (string, error) {
	if err := SendAndReceive(writer, command); err != nil {
		return "", err
	}
	return ReadOutputAll(conn, reader)
}

func SendAndReceive(writer *bufio.Writer, command string) error {
	if command != "" {
		_, err := fmt.Fprintf(writer, "%s\n", command)
		if err != nil {
			return err
		}
		return writer.Flush()
	}
	return nil
}

func WaitForPrompts(conn *telnet.Conn, reader *bufio.Reader, prompts []string, timeout time.Duration) error {
	deadline := time.Now().Add(timeout)
	conn.SetReadDeadline(deadline)
	var output []byte
	for {
		buf := make([]byte, 1024)
		_, err := reader.Read(buf)
		if err != nil {
			return fmt.Errorf("从 Telnet 服务器读取时发生错误: %v", err)
		}
		output = append(output, buf...)
		for _, prompt := range prompts {
			if strings.Contains(string(output), prompt) {
				return nil
			}
		}
	}
}

func ReadOutputAll(conn *telnet.Conn, reader *bufio.Reader) (string, error) {
	var output []byte
	for {
		buf := make([]byte, 1)
		_, err := reader.Read(buf)
		if err != nil {
			return "", fmt.Errorf("从 Telnet 服务器读取时发生错误: %v", err)
		}
		output = append(output, buf...)
		if strings.Contains(string(output), ">") || strings.Contains(string(output), "#") || strings.Contains(string(output), "]") {
			break
		}
	}
	return string(output), nil
}
func ReadOutput(conn *telnet.Conn, reader *bufio.Reader) (string, error) {
	var output []byte
	var result []byte
	inResult := false
	for {
		buf := make([]byte, 1)
		_, err := reader.Read(buf)
		if err != nil {
			return "", fmt.Errorf("从 Telnet 服务器读取时发生错误: %v", err)
		}
		output = append(output, buf...)

		// 检测到命令提示符变化，标识结果的开始
		if strings.Contains(string(output), "\n") && !inResult {
			inResult = true
			output = []byte{} // 清空 output，只保留结果部分
		}

		if inResult {
			result = append(result, buf...)
		}

		// 检测到结束提示符
		if strings.Contains(string(result), ">") || strings.Contains(string(result), "#") || strings.Contains(string(result), "]") || strings.Contains(string(result), "EOF") {
			break
		}
	}

	// 去掉结束提示符
	resultStr := string(result)
	if strings.Contains(resultStr, ">") {
		resultStr = strings.Split(resultStr, ">")[0]
	}
	if strings.Contains(resultStr, "#") {
		resultStr = strings.Split(resultStr, "#")[0]
	}

	// 按行拆分并去掉最后一行
	lines := strings.Split(resultStr, "\n")
	if len(lines) > 1 {
		resultStr = strings.Join(lines[:len(lines)-1], "\n")
	}

	return strings.TrimSpace(resultStr), nil
}

func DetectSwitchType(output string) (string, error) {
	if strings.Contains(output, "Time") {
		return "huawei", nil
	} else if strings.Contains(output, "%") {

		return "cisco", nil
	} else if strings.Contains(output, "#") {

		return "cisco", nil
	}
	return "", fmt.Errorf("无法检测交换机类型")
}
func ExtractContentBetweenDashes(input string) (string, error) {
	// 使用正则表达式匹配第二条和第三条 `-----` 之间的内容
	re := regexp.MustCompile(`(?s)MAC Address\s+VLAN/VSI/BD\s+Learned-From\s+Type\s*-------------------------------------------------------------------------------\s*(.*?)\s*-------------------------------------------------------------------------------`)
	matches := re.FindStringSubmatch(input)
	if len(matches) < 2 {
		return "", fmt.Errorf("未找到匹配的内容")
	}
	return matches[1], nil
}
func ExtractFields(input string) ([]string, error) {
	// 使用正则表达式匹配每列字段
	re := regexp.MustCompile(`\s*(\S+)\s+(\S+)\s+(\S+)\s+(\S+)`)
	matches := re.FindStringSubmatch(input)
	if len(matches) < 5 {
		return nil, fmt.Errorf("未找到匹配的内容")
	}

	// 返回匹配的字段
	return matches[1:], nil
}
func AnalyzeHuaweiSwtich(input string) (vlanport, linkport string) {
	columns := ExtractColumns(input)
	if len(columns) > 0 {
		vlanport = columns[1]
		linkport = columns[len(columns)-2]
		return vlanport, linkport
	}
	return vlanport, linkport
}
func AnalyzeCiscoSwtich(input string) (vlanport, linkport string) {
	columns := ExtractColumns(input)
	if len(columns) > 0 {
		vlanport = columns[0]
		linkport = columns[len(columns)-1]
		return vlanport, linkport
	}
	return vlanport, linkport
}

func ExtractColumns(input string) []string {
	// 使用 strings.Fields 以空白字符为分隔符分割字符串
	return strings.Fields(input)
}
func SelectSwitchClient(conn *telnet.Conn, switchtype string, shortmac string, switchuplink string) (vlanport, linkport string, err error) {
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	if switchtype == "huawei" {
		cmdout, err := SendCommand(conn, reader, writer, "display mac-address | in "+shortmac)
		if err != nil {
			return vlanport, linkport, err
		}
		zap.L().Info("交换机输出:", zap.String("cmdout", cmdout))
		cmdout, err = ExtractContentBetweenDashes(cmdout)
		if err != nil {
			return vlanport, linkport, err
		}
		vlanport, linkport := AnalyzeHuaweiSwtich(cmdout)
		zap.L().Named("MAC地址当前vlan" + vlanport)
		zap.L().Named("MAC地址当前port" + linkport)
		linkport = FilterCharacter(linkport, 'E')
		return vlanport, linkport, err
	} else {
		cmdout, err := SendCommand(conn, reader, writer, "show mac address-table | in "+shortmac)
		if err != nil {
			return vlanport, linkport, err
		}
		vlanport, linkport := AnalyzeCiscoSwtich(cmdout)
		zap.L().Named("MAC地址当前vlan" + vlanport)
		zap.L().Named("MAC地址当前port" + linkport)

		return vlanport, linkport, err
	}
}
func FilterCharacter(input string, charToFilter rune) string {
	return strings.Map(func(r rune) rune {
		if r == charToFilter {
			return -1 // 返回 -1 表示要过滤掉该字符
		}
		return r
	}, input)
}
func SelectSwitch(conn *telnet.Conn, switchtype string, shortmac string, switchuplink string) (cmdout string, err error, selectswitinfo models.SwitchLinkInfo) {
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	if switchtype == "huawei" {
		cmdout, err = SendCommand(conn, reader, writer, "display mac-address | in "+shortmac)
		if err != nil {
			return cmdout, err, selectswitinfo
		}
		zap.L().Info("交换机输出:", zap.String("cmdout", cmdout))
		cmdout, err := ExtractContentBetweenDashes(cmdout)
		if err != nil {
			return cmdout, err, selectswitinfo
		}
		vlanport, linkport := AnalyzeHuaweiSwtich(cmdout)
		zap.L().Named("MAC地址当前vlan" + vlanport)
		zap.L().Named("MAC地址当前port" + linkport)
		selectswitinfoformdb, err := mysql.SelectSwitchLinkInfo(switchuplink, linkport)
		if err != nil {
			zap.L().Named("无记录")
			return cmdout, err, selectswitinfoformdb
		}
		return cmdout, err, selectswitinfoformdb
	} else {
		cmdout, err = SendCommand(conn, reader, writer, "show mac address-table | in "+shortmac)
		if err != nil {
			return cmdout, err, selectswitinfo
		}
		vlanport, linkport := AnalyzeCiscoSwtich(cmdout)
		zap.L().Named("MAC地址当前vlan" + vlanport)
		zap.L().Named("MAC地址当前port" + linkport)
		selectswitinfoformdb, err := mysql.SelectSwitchLinkInfo(switchuplink, linkport)
		if err != nil {
			return cmdout, err, selectswitinfo
		}
		return cmdout, err, selectswitinfoformdb
	}
}
func ChangStitchPort(conn *telnet.Conn, switchtype string, shortmac string, switchport string, vlan int) (vlanport, linkport string, err error) {
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	if switchtype == "huawei" {
		commands := []string{
			"sys",
			fmt.Sprintf("interface %s", switchport),
			fmt.Sprintf("port default vlan %d", vlan),
		}
		savecommands := []string{
			"save",
			"Y",
		}

		for _, cmd := range commands {
			cmdout, err := SendCommandChang(conn, reader, writer, cmd)
			if err != nil {
				return vlanport, linkport, err
			}
			zap.L().Named("保存执行命令结果:" + cmdout)

		}
		for _, cmd := range savecommands {
			cmdout, err := SendCommandChang(conn, reader, writer, cmd)
			if err != nil {
				return vlanport, linkport, err
			}
			zap.L().Named("执行命令结果:" + cmdout)
		}
		cmdout, err := SendCommand(conn, reader, writer, "display mac-address | in "+shortmac)
		if err != nil {
			conn.Close()
			return vlanport, linkport, err
		}
		zap.L().Info("交换机输出:", zap.String("cmdout", cmdout))
		cmdout, err = ExtractContentBetweenDashes(cmdout)
		if err != nil {
			return vlanport, linkport, err
		}
		vlanport, linkport := AnalyzeHuaweiSwtich(cmdout)
		zap.L().Named("MAC地址当前vlan" + vlanport)
		zap.L().Named("MAC地址当前port" + linkport)
		return vlanport, linkport, err
	} else {
		commands := []string{
			"configure terminal",
			fmt.Sprintf("interface %s", switchport),
			fmt.Sprintf("switchport access vlan %d", vlan),
			"exit",
			"exit",
			"wr",
		}

		for _, cmd := range commands {
			cmdout, err := SendCommandChang(conn, reader, writer, cmd)
			if err != nil {
				return vlanport, linkport, err
			}
			zap.L().Info("执行命令结果:", zap.String("cmdout", cmdout))
			time.Sleep(2 * time.Second) // 增加延迟以避免超时
		}
		zap.L().Named("MAC地址当前vlan" + strconv.Itoa(vlan))
		zap.L().Named("MAC地址当前port" + switchport)
		return strconv.Itoa(vlan), switchport, err
	}
}

func buildSwitchAddress(linkswitch string) string {
	return fmt.Sprintf("100.100.%s:23", linkswitch)
}

func connectAndValidate(address, username, password string) (*telnet.Conn, string, error) {
	err, switchtype, switchconn := ConnectSwitch(address, username, password)
	if err != nil {
		return nil, "", err
	}
	return switchconn, switchtype, nil
}

func processSwitchOutput(switchtype string, s string, switchinfo models.SwitchLinkInfo) (string, error) {
	var vlanport, linkport string
	var err error

	switch switchtype {
	case "huawei":
		cmdout, err := ExtractContentBetweenDashes(s)
		if err != nil {
			return "", err
		}
		vlanport, linkport = AnalyzeHuaweiSwtich(cmdout)

	case "cisco":
		vlanport, linkport = AnalyzeCiscoSwtich(s)

	default:
		return "", fmt.Errorf("未知的交换机类型: %s", switchtype)
	}

	zap.L().Named("MAC地址当前vlan" + vlanport)
	zap.L().Named("MAC地址当前port" + linkport)

	return switchinfo.SwitchName, err
}
