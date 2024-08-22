package logic

import (
	"fmt"
	"go-web-app/models"
	"strconv"
	"strings"
)

func ParseInterfaceBrief(input string) (interfaces []models.HuaweiInterfaceBrief) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		// 跳过标题行、说明行和无关行
		if strings.HasPrefix(line, "Interface") || strings.HasPrefix(line, "----") || strings.Contains(line, ":") || line == "" {
			continue
		}

		fields := strings.Fields(line)
		if len(fields) >= 7 {
			// 解析 inErrors 和 outErrors
			interfaceName := strings.Replace(fields[0], "GigabitEthernet", "G", 1)
			inErrors, _ := strconv.Atoi(fields[5])  // 转换为整数
			outErrors, _ := strconv.Atoi(fields[6]) // 转换为整数

			interfaceBrief := models.HuaweiInterfaceBrief{
				Interface: interfaceName,
				PHY:       fields[1],
				Protocol:  fields[2],
				InUti:     fields[3],
				OutUti:    fields[4],
				InErrors:  inErrors,
				OutErrors: outErrors,
			}
			interfaces = append(interfaces, interfaceBrief)
		} else {
			fmt.Println("解析行时遇到问题:", line)
		}
	}
	return interfaces
}

// ParsePortVlanCommand 解析 display port vlan 命令的输出
func ParsePortVlanCommand(input string) (portVlans []models.HuaweiPortVlan) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		// 跳过标题行
		if strings.HasPrefix(line, "Port") || strings.HasPrefix(line, "----") || line == "" {
			continue
		}
		fields := strings.Fields(line)
		if len(fields) >= 4 {
			interfaceName := strings.Replace(fields[0], "GigabitEthernet", "G", 1)
			portVlan := models.HuaweiPortVlan{
				Interface:     interfaceName,
				LinkType:      fields[1],
				PVID:          fields[2],
				TrunkVLANList: strings.Join(fields[3:], " "), // 将剩余部分拼接为 Trunk VLAN 列表
			}
			portVlans = append(portVlans, portVlan)
		}
	}
	return portVlans
}

// CombineInterfaceAndVlan 组合两个解析结果
func CombineInterfaceAndVlan(interfaces []models.HuaweiInterfaceBrief, portVlans []models.HuaweiPortVlan) []models.InterfaceVlanInfo {
	combined := []models.InterfaceVlanInfo{}
	// 建立一个 map 以快速查找 VLAN 信息
	vlanMap := make(map[string]models.HuaweiPortVlan)
	for _, vlan := range portVlans {
		vlanMap[vlan.Interface] = vlan
	}
	// 组合数据
	for _, intf := range interfaces {
		if vlanInfo, found := vlanMap[intf.Interface]; found {
			combined = append(combined, models.InterfaceVlanInfo{
				Interface: intf.Interface,
				PHY:       intf.PHY,
				PVID:      vlanInfo.PVID,
				LinkType:  vlanInfo.LinkType,
			})
		}
	}
	return combined
}
func ParseInterfaceStatus(input string) (interfaces []models.InterfaceVlanInfo) {
	lines := strings.Split(input, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		// 跳过标题行、分隔行和空行
		if strings.HasPrefix(line, "Port") || strings.HasPrefix(line, "----") || line == "" {
			continue
		}
		fields := strings.Fields(line)
		// 确保字段长度足够，避免数组越界
		if len(fields) >= 6 {
			phy := ""
			for i := 1; i < len(fields) && i < 6; i++ {
				if fields[i] == "connected" || fields[i] == "notconnect" {
					phy = fields[i]
					break
				}
			}

			// 如果在 6 次检查内没有找到有效的 PHY 值，则填空
			if phy == "" {
				phy = "unknown"
			}

			// 处理字段，跳过已经检查的字段，确保索引不越界
			pvidIndex := 2
			linkTypeIndex := 3
			if phy != fields[1] {
				pvidIndex++
				linkTypeIndex++
			}

			interfaceInfo := models.InterfaceVlanInfo{
				Interface: fields[0],
				PVID:      fields[pvidIndex],
				PHY:       phy,
				LinkType:  fields[linkTypeIndex],
			}

			interfaces = append(interfaces, interfaceInfo)
		} else {
			fmt.Println("解析行时遇到问题，字段数不足:", line)
		}
	}
	return interfaces
}
