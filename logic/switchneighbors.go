package logic

import (
	"bufio"
	"github.com/ziutek/telnet"
	"go-web-app/dao/mysql"
	"go-web-app/models"
	"go-web-app/settings"
	"strings"
)

func SelectNeighborsLogic(p *models.SelectNeighbors) (interface{}, error) {
	username := settings.Conf.Username
	password := settings.Conf.Passtoken
	switchconn, switchtype, err := connectAndValidate(buildSwitchAddress(p.SwitchName), username, password)
	if err != nil {
		return nil, err
	}
	neighbors, err := NeighborsDetail(switchconn, switchtype, p.SwitchName)
	return neighbors, err
}
func NeighborsDetail(conn *telnet.Conn, switchtype string, switchname string) (neighbors []models.HuaweiNeighbor, err error) {
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)

	if switchtype == "huawei" {
		selectswitinfoformdb, err := mysql.SelectSwitchUpLinkInfo(switchname)
		if err != nil {
			return neighbors, err
		}
		cmdout, err := SendCommand(conn, reader, writer, "screen-length 0 temporary")
		if err != nil {
			return neighbors, err
		}
		cmdout, err = SendCommand(conn, reader, writer, "display cdp neighbor")
		if err != nil {
			return neighbors, err
		}

		neighbors = HuaweiNeighborPoint(cmdout, selectswitinfoformdb.UplinkPort)
		return neighbors, err
	} else {
		selectswitinfoformdb, err := mysql.SelectSwitchUpLinkInfo(switchname)
		if err != nil {
			return neighbors, err
		}
		cmdout, err := SendCommand(conn, reader, writer, "show cdp neighbors detail")
		if err != nil {
			return neighbors, err
		}

		neighbors = CiscoNeighborPoint(cmdout, selectswitinfoformdb.UplinkPort)
		return neighbors, err
	}
}
func SelectInterface(p *models.SelectNeighbors) (interface{}, error) {
	username := settings.Conf.Username
	password := settings.Conf.Passtoken
	switchconn, switchtype, err := connectAndValidate(buildSwitchAddress(p.SwitchName), username, password)
	if err != nil {
		return nil, err
	}
	combineintface, err := SelectInterfaceDetail(switchconn, switchtype)
	return combineintface, err
}
func SelectInterfaceDetail(conn *telnet.Conn, switchtype string) (combineintface []models.InterfaceVlanInfo, err error) {
	reader := bufio.NewReader(conn)
	writer := bufio.NewWriter(conn)
	var interfacebrief []models.HuaweiInterfaceBrief
	var interfacevlan []models.HuaweiPortVlan
	if switchtype == "huawei" {
		// 执行第一个命令
		cmdout, err := SendCommand(conn, reader, writer, "screen-length 0 temporary")
		if err != nil {
			return combineintface, err
		}
		// 执行第二个命令
		cmdout, err = SendCommand(conn, reader, writer, "display interface brief")
		if err != nil {
			return combineintface, err
		}
		// 解析接口状态
		interfacebrief = ParseInterfaceBrief(cmdout)
		// 执行第三个命令
		vlancmdout, err := SendCommand(conn, reader, writer, "display port vlan")
		if err != nil {
			return combineintface, err
		}
		interfacebrief = ParseInterfaceBrief(vlancmdout)
		vlancmdout, err = SendCommand(conn, reader, writer, "display port vlan")
		if err != nil {
			return combineintface, err
		}
		interfacevlan = ParsePortVlanCommand(vlancmdout)
		combineintface = CombineInterfaceAndVlan(interfacebrief, interfacevlan)
		return combineintface, err
	} else {
		cmdout, err := SendCommandAll(conn, reader, writer, "terminal length 0")
		if err != nil {
			return combineintface, err
		}
		cmdout, err = SendCommand(conn, reader, writer, "show interfaces status")
		if err != nil {
			return combineintface, err
		}
		combineintface = ParseInterfaceStatus(cmdout)
		return combineintface, err
	}
}
func HuaweiNeighborPoint(input string, uplinkport string) (neighbors []models.HuaweiNeighbor) {
	lines := strings.Split(input, "\n")
	var neighbor models.HuaweiNeighbor
	counter := 1 // 用于序号计数

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.Contains(line, " has ") && strings.Contains(line, "neighbor(s):") {
			// 解析接口名并进行格式转换
			parts := strings.Split(line, " ")
			neighbor.SInterface = parts[0]

			neighbor.SInterface = strings.Replace(neighbor.SInterface, "XGigabitEthernet", "XGE", 1)
			neighbor.SInterface = strings.Replace(neighbor.SInterface, "GigabitEthernet", "Gi", 1)
		} else if strings.HasPrefix(line, "Device ID") {
			// 解析设备名称
			neighbor.SwitchName = strings.TrimSpace(strings.Split(line, ":")[1])
		} else if strings.HasPrefix(line, "Port ID") {
			// 解析端口并进行格式转换
			neighbor.DInterface = strings.TrimSpace(strings.Split(line, ":")[1])
			neighbor.DInterface = strings.Replace(neighbor.DInterface, "XGigabitEthernet", "XGE", 1)
			neighbor.DInterface = strings.Replace(neighbor.DInterface, "GigabitEthernet", "Gi", 1)

		} else if strings.HasPrefix(line, "Platform") {
			// 解析平台信息
			neighbor.SwitchPlatform = strings.TrimSpace(strings.Split(line, ":")[1])
		} else if strings.HasPrefix(line, "MacAddress") {
			// 解析MAC地址
			neighbor.SwitchMac = strings.TrimSpace(strings.Split(line, ":")[1])
			// 设置序号
			neighbor.SelectNumber = counter

			// 判断 SInterface 是否为上联端口
			if neighbor.SInterface == uplinkport {
				neighbor.IsUpperDevice = true
			}

			// 完成一个 neighbor 的解析，加入结果集
			neighbors = append(neighbors, neighbor)
			// 重置 neighbor 结构体，以解析下一个
			neighbor = models.HuaweiNeighbor{}
			// 序号自增
			counter++
		}
	}
	return neighbors
}

func CiscoNeighborPoint(input string, uplinkport string) (neighbors []models.HuaweiNeighbor) {
	lines := strings.Split(input, "\n")
	var neighbor models.HuaweiNeighbor
	counter := 1 // 用于序号计数

	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "Device ID") {
			// 如果当前 neighbor 不为空且有效，将其添加到结果中
			if neighbor.SwitchName != "" && neighbor.SInterface != "" {
				// 设置序号
				neighbor.SelectNumber = counter
				// 添加到结果集
				neighbors = append(neighbors, neighbor)
				// 序号自增
				counter++
			}
			// 重置 neighbor 结构体
			neighbor = models.HuaweiNeighbor{}

			// 解析设备名称
			neighbor.SwitchName = strings.TrimSpace(strings.Split(line, ":")[1])
		} else if strings.HasPrefix(line, "Interface:") && strings.Contains(line, "Port ID") {
			// 解析接口名和 Port ID 信息在同一行的情况
			interfaceParts := strings.Split(line, ",")
			neighbor.SInterface = strings.TrimSpace(strings.Split(interfaceParts[0], ":")[1])
			neighbor.SInterface = strings.Replace(neighbor.SInterface, "GigabitEthernet", "Gi", 1)

			neighbor.DInterface = strings.TrimSpace(strings.Split(interfaceParts[1], ":")[1])
			neighbor.DInterface = strings.Replace(neighbor.DInterface, "GigabitEthernet", "Gi", 1)
		} else if strings.HasPrefix(line, "Platform:") {
			// 解析平台信息
			platformParts := strings.Split(line, ",  Capabilities")
			neighbor.SwitchPlatform = strings.TrimSpace(platformParts[0][len("Platform:"):])
		}
	}
	// 判断 SInterface 是否为上联端口
	if neighbor.SInterface == uplinkport {
		neighbor.IsUpperDevice = true
	}
	// 添加最后一个 neighbor（如果有效）
	if neighbor.SwitchName != "" && neighbor.SInterface != "" {
		neighbor.SelectNumber = counter
		neighbors = append(neighbors, neighbor)
	}

	return neighbors
}
