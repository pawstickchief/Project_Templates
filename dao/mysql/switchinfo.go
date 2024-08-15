package mysql

import (
	"database/sql"
	"go-web-app/models"
	"go.uber.org/zap"
)

func SelectSwitchLinkInfo(switchname, downlinkport string) (models.SwitchLinkInfo, error) {
	query := `
SELECT
downlinkswitch, uplinkport, downlinkport, switchname, switchnote, switchlocation, switchtype
FROM
switchlinkinformation
WHERE
downlinkport = ? AND
switchname = ?;
`
	var switchLink models.SwitchLinkInfo
	row := db.QueryRow(query, downlinkport, switchname)

	err := row.Scan(
		&switchLink.DownLinkSwitch,
		&switchLink.UplinkPort,
		&switchLink.DownLinkPort,
		&switchLink.SwitchName,
		&switchLink.SwitchNote,
		&switchLink.SwitchLocation,
		&switchLink.SwitchType,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			zap.L().Info("没有找到记录")
			switchLink.SwitchName = switchname
			return switchLink, nil
		}
		zap.L().Info("读取数据失败:", zap.Error(err))
		return switchLink, err
	}

	return switchLink, nil
}
