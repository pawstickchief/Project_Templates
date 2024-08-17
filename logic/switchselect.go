package logic

import (
	"fmt"
	"go-web-app/models"
)

func SelectSwitchInfoOption(p *models.SelectSwitchMac) (interface{}, error) {

	switch p.SwitchLevel {
	case 123:
		return nil, nil
	default:
		return nil, fmt.Errorf("未知的 SwitchLevel: %d", p.SwitchLevel)
	}
}
