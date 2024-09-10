package logic

import (
	"fmt"
	"go-web-app/dao/mysql"
	"go-web-app/models"
)

func SelectSwitchUplink(p *models.SwitchUplinkInfo) (interface{}, error) {

	switch p.SwitchOption {
	case "select":
		info, err := mysql.SelectSwitchUplinkInfo()
		if err != nil {
			return nil, err
		}
		return info, err
	case "insert":
		info, err := mysql.SelectSwitchUplinkInfo()
		if err != nil {
			return nil, err
		}
		return info, err
	case "modifiy":
		info, err := mysql.SelectSwitchUplinkInfo()
		if err != nil {
			return nil, err
		}
		return info, err
	case "delete":
		info, err := mysql.SelectSwitchUplinkInfo()
		if err != nil {
			return nil, err
		}
		return info, err
	default:
		return nil, fmt.Errorf("未知的 SwitchOption: %d", p.SwitchOption)
	}
}
