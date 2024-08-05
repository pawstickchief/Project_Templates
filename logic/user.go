package logic

import (
	"fmt"
	"go-web-app/dao/mysql"
	"go-web-app/models"
)

// 存放业务逻辑代码

func AlarmOption(p *models.ParamAlarmSetting) (s interface{}, err error) {
	switch {
	case p.AlarmSettingOption == "add":
		s, err := mysql.AlarmAdd(p)
		if err != nil {
			return s, err
		}
		return s, err
	case p.AlarmSettingOption == "alarmedit":
		s, err := mysql.AlarmEdit(p)
		if err != nil {
			return s, err
		}
		return s, err
	case p.AlarmSettingOption == "optioninit":
		s, err := mysql.AlarmSettingInit(p)
		if err != nil {
			return s, err
		}
		fmt.Println(s)
		return s, err
	case p.AlarmSettingOption == "updatenoti":
		s, err := mysql.AlarmUpdateNoti(p)
		if err != nil {
			return s, err
		}
		return s, err
	case p.AlarmSettingOption == "updatethreshold":
		s, err := mysql.AlarmUpdateThreshold(p)
		if err != nil {
			return s, err
		}
		return s, err
	}
	return
}
