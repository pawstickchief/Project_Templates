package mysql

import (
	"errors"
	"fmt"
	"go-web-app/models"
)

var (
	ErrorHostNoExist = errors.New("主机不存在")
)

func ClientConfirm(client *models.ParamSystemGet) (hostid int64, err error) {
	sqlStr1 := "select count(hostid) from hostlist where hostname=?"
	var count int64
	if err := db.Get(&count, sqlStr1, client.Hostname); err != nil {
		return count, err
	}

	if count == 0 {
		return count, ErrorHostNoExist
	}
	sqlStr2 := "select hostid from hostlist where hostname=?"
	err = db.Get(&hostid, sqlStr2, client.Hostname)
	if err != nil {
		return hostid, err
	}
	return hostid, err

}

func ClientSystemInfoGet(client *models.ParamSystemGet) (Reply string, err error) {
	sqlStr := "select hostsysteminfo from hostlist where hostip=?"
	err = db.Get(&Reply, sqlStr, client.OptionIp)
	if err != nil {
		return
	}
	return Reply, err
}
func ClientUptime(client *models.ParamSystemGet) (Reply int64, err error) {
	sqlStr := "update hostlist set hostuptime=? where hostid=?"
	ret, err := db.Exec(sqlStr,
		client.OpitonParame,
		client.Hostid,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	Reply, err = ret.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("更新数据为 %d 条\n", Reply)
	}
	return

}
func ClientSystemInfo(client *models.ParamSystemGet) (Reply int64, err error) {
	sqlStr := "update hostlist set hostsysteminfo=? where hostid=?"
	ret, err := db.Exec(sqlStr,
		client.OpitonParame,
		client.Hostid,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	Reply, err = ret.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("更新数据为 %d 条\n", Reply)
	}
	return

}
func BaseMonitoring(client *models.ParamSystemGet) (Reply int64, err error) {
	sqlStr := "insert into hostdata(hostid,cpupart,rampart,uns,dns,insertdatatime) values (?,?,?,?,?,?)"
	ret, err := db.Exec(sqlStr,
		client.Hostid,
		client.OptionParameCpu,
		client.OptionParameMemory,
		client.OptionParameUns,
		client.OptionParameDns,
		client.OptionTime,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	Reply, err = ret.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("更新数据为 %d 条\n", Reply)
	}
	return

}
