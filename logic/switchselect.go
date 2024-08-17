package logic

import (
	"fmt"
	"go-web-app/models"
	"go-web-app/settings"
	"go.uber.org/zap"
)

func SelectSwitchInfoOption(p *models.SelectSwitchMac) (interface{}, error) {
	username := settings.Conf.Username
	password := settings.Conf.Passtoken
	var linkswitch string
	switch p.SwitchLevel {
	case 502:
		linkswitch = "5.6"
		return SwitchProcess(linkswitch, username, password, p.ShortMAC)
	case 801:
		linkswitch = "8.1"
		return SwitchProcess(linkswitch, username, password, p.ShortMAC)

	default:
		return nil, fmt.Errorf("未知的 SwitchLevel: %d", p.SwitchLevel)
	}
}

func SwitchProcess(linkswitch, username, password, ShortMAC string) (ClientSwitchInfo models.ClientSwitchInfo, err error) {

	switchconn, switchtype, err := connectAndValidate(buildSwitchAddress(linkswitch), username, password)
	if err != nil {
		return ClientSwitchInfo, err
	}

	for {
		s, err, switchinfo := SelectSwitch(switchconn, switchtype, ShortMAC, linkswitch)
		if err != nil {
			zap.L().Named(switchinfo.SwitchName)
			return ClientSwitchInfo, err
		}
		if switchinfo.SwitchNote == "" {
			switchvlan, switchport, _ := SelectSwitchClient(switchconn, switchtype, ShortMAC, buildSwitchAddress(linkswitch))
			zap.L().Info("工位接入", zap.String(switchvlan, switchport))
			ClientSwitchInfo.Vlan = switchvlan
			ClientSwitchInfo.SwitchPort = switchport
			ClientSwitchInfo.SwitchName = linkswitch
			return ClientSwitchInfo, err
		}

		linkswitch = switchinfo.DownLinkSwitch
		switchconn, switchtype, err = connectAndValidate(buildSwitchAddress(linkswitch), username, password)
		if err != nil {
			return ClientSwitchInfo, err
		}

		if switchinfo.SwitchName == switchinfo.DownLinkSwitch {

			ClientSwitchInfo.SwitchName, err = processSwitchOutput(switchtype, s, switchinfo)

			return ClientSwitchInfo, err
		}

	}
}
