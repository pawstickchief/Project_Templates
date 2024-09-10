package mysql

import (
	"go-web-app/models"
	"go.uber.org/zap"
)

func SelectSwitchUplinkInfo() ([]models.SelectNeighbors, error) {
	query := `
SELECT 
    @rownum := @rownum + 1 AS switchnumber,
    switchname
FROM 
    (SELECT DISTINCT switchname FROM switchlinkinformation) AS unique_switches,
    (SELECT @rownum := 0) AS r
ORDER BY 
    switchname;
`
	var switchLinks []models.SelectNeighbors

	// 执行查询并获取多行结果
	rows, err := db.Query(query)
	if err != nil {
		zap.L().Info("执行查询失败:", zap.Error(err))
		return nil, err
	}
	defer rows.Close()

	// 遍历结果集
	for rows.Next() {
		var switchLink models.SelectNeighbors
		err := rows.Scan(
			&switchLink.SwitchNumber,
			&switchLink.SwitchName,
		)
		if err != nil {
			zap.L().Info("读取数据失败:", zap.Error(err))
			return nil, err
		}
		// 将每一行的数据添加到切片中
		switchLinks = append(switchLinks, switchLink)
	}

	// 检查是否有迭代错误
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return switchLinks, nil
}
