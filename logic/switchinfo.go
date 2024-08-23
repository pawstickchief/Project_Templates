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
			phyIndex := -1
			for i := 1; i < len(fields); i++ {
				if fields[i] == "connected" || fields[i] == "notconnect" {
					phy = fields[i]
					phyIndex = i
					break
				}
			}

			// 如果在字段中没有找到有效的 PHY 值，则标记为 "unknown"
			if phy == "" {
				phy = "unknown"
			}

			// 确保索引不越界并且正确获取网段和链路类型信息
			pvid := "unknown"
			linkType := "unknown"
			if phyIndex != -1 && len(fields) > phyIndex+1 {
				pvid = fields[phyIndex+1] // 网段字段在 PHY 后面
			}
			if phyIndex != -1 && len(fields) > phyIndex+2 {
				linkType = fields[phyIndex+2] // 链路类型在网段后面
			}

			interfaceInfo := models.InterfaceVlanInfo{
				Interface: fields[0],
				PVID:      pvid,
				PHY:       phy,
				LinkType:  linkType,
			}

			interfaces = append(interfaces, interfaceInfo)
		} else {
			fmt.Println("解析行时遇到问题，字段数不足:", line)
		}
	}
	return interfaces
}
