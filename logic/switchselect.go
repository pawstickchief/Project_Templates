package logic

import (
	"fmt"
	"go-web-app/models"
	"go.uber.org/zap"
)

func SelectSwitchInfoOption(p *models.SelectSwitchMac) (interface{}, error) {
	username := "gyop"
	password := "E@2wYZ!Asa"
	var ClientSwtichInfo models.ClientSwitchInfo

	switch p.SwitchLevel {
	case 503:
		linkswitch := "5.27"
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
				switchvlan, switchport, _ := SelectSwitchClient(switchconn, switchtype, p.ShortMAC, buildSwitchAddress(linkswitch))
				zap.L().Info("工位接入", zap.String(switchvlan, switchport))
				ClientSwtichInfo.Vlan = switchvlan
				ClientSwtichInfo.SwitchPort = switchport
				ClientSwtichInfo.SwitchName = linkswitch
				return ClientSwtichInfo, err
			}

			linkswitch = switchinfo.DownLinkSwitch
			switchconn, switchtype, err = connectAndValidate(buildSwitchAddress(linkswitch), username, password)
			if err != nil {
				return nil, err
			}

			if switchinfo.SwitchName == switchinfo.DownLinkSwitch {
				return processSwitchOutput(switchconn, switchtype, s, switchinfo)
			}
		}

	case 502:
		linkswitch := "5.1"
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
				switchvlan, switchport, _ := SelectSwitchClient(switchconn, switchtype, p.ShortMAC, buildSwitchAddress(linkswitch))
				zap.L().Info("工位接入", zap.String(switchvlan, switchport))
				ClientSwtichInfo.Vlan = switchvlan
				ClientSwtichInfo.SwitchPort = switchport
				ClientSwtichInfo.SwitchName = linkswitch
				return ClientSwtichInfo, err
			}

			linkswitch = switchinfo.DownLinkSwitch
			switchconn, switchtype, err = connectAndValidate(buildSwitchAddress(linkswitch), username, password)
			if err != nil {
				return nil, err
			}

			if switchinfo.SwitchName == switchinfo.DownLinkSwitch {
				return processSwitchOutput(switchconn, switchtype, s, switchinfo)
			}
		}
	case 50201:
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
				switchvlan, switchport, err := SelectSwitchClient(switchconn, switchtype, p.ShortMAC, buildSwitchAddress(linkswitch))
				zap.L().Info("工位接入", zap.String(switchvlan, switchport))
				ClientSwtichInfo.Vlan = switchvlan
				ClientSwtichInfo.SwitchPort = switchport
				ClientSwtichInfo.SwitchName = linkswitch
				return ClientSwtichInfo, err

			}

			linkswitch = switchinfo.DownLinkSwitch
			switchconn, switchtype, err = connectAndValidate(buildSwitchAddress(linkswitch), username, password)
			if err != nil {
				return nil, err
			}

			if switchinfo.SwitchName == switchinfo.DownLinkSwitch {
				return processSwitchOutput(switchconn, switchtype, s, switchinfo)
			}
		}
	case 801:
		return nil, nil

	default:
		return nil, fmt.Errorf("未知的 SwitchLevel: %d", p.SwitchLevel)
	}
}
