package macswitch

import (
	"fmt"
	"strings"
)

func FormatMACAddress(mac string) string {
	// 去除分隔符并将所有字母转为小写
	mac = strings.ToLower(strings.ReplaceAll(strings.ReplaceAll(mac, "-", ""), ":", ""))

	// 按照需要的格式插入分隔符
	if len(mac) == 12 {
		return fmt.Sprintf("%s.%s.%s", mac[0:4], mac[4:8], mac[8:12])
	}

	return "Invalid MAC address format"
}
