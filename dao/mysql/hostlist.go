package mysql

import (
	"errors"
	"fmt"
	"go-web-app/models"
	"go-web-app/pkg/snowflake"
	"go-web-app/pkg/todaytime"
	"strconv"
	"time"
)

var (
	ErrorHostExist    = errors.New("主机已存在")
	timeLayoutday     = "2006-01-02"
	timeLayoutdayfull = "2006-01-02 00:00:00"
)

func Hostlistalarm(host *models.ParamHostDateGet) (hostgetdata []models.Hostlist, err error) {
	sqlStr := `select hostid,hostname,systemtype,hoststatus,hostip,hostlocation,hostowner,hostaddtime,hostnote,hostuptime,
       (select count(hostid)  from alarmstatistics where alarmstatus=1 and hostid = host.hostid) as hostissues from hostlist host;`
	if err := db.Select(&hostgetdata, sqlStr); err != nil {
		return hostgetdata, err
	}
	return
}

func Hostlistdataget(host *models.ParamHostDateGet) (hostgetdata []models.Hostlist, err error) {

	sqlStr := `select hostid,hostname,systemtype,hoststatus,hostip,hostlocation,hostowner,hostaddtime,hostnote from hostlist`
	if err := db.Select(&hostgetdata, sqlStr); err != nil {
		return hostgetdata, err
	}
	return
}
func Hostinfo(host *models.ParamHostDateGet) (hostgetdata interface{}, err error) {
	hostlistid := host.Hostid
	sqlStr := `select hostid,hostname,systemtype,hoststatus,hostip,hostlocation,hostowner,hostaddtime,hostnote,hostysteminfo from hostlist where hostip = ?`
	if err := db.Get(hostgetdata, sqlStr, hostlistid); err != nil {
		return hostgetdata, err
	}
	return
}
func HostName(hostid string) (hostname string) {
	sqlStr := `select hostname  from hostlist where hostid = ?`
	if err := db.Get(&hostname, sqlStr, hostid); err != nil {
		return
	}

	return
}
func Hostedit(host *models.ParamHostDateGet) (n int64, err error) {
	sqlStr := "update hostlist set hostname=?,systemtype=?,hoststatus=?,hostip=?,hostlocation=?,hostowner=?,hostnote=? where hostid=?"
	ret, err := db.Exec(sqlStr,
		host.HostName,
		host.SystemType,
		host.HostStatus,
		host.HostIP,
		host.HostLocation,
		host.HostOwner,
		host.HostNote,
		host.Hostid,
	)
	if err != nil {
		fmt.Println(err)
		return
	}
	n, err = ret.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("更新数据为 %d 条\n", n)
	}
	clinet := models.SystemLog{
		SystemlogHostName:  host.HostName,
		SystemlogType:      "修改主机信息",
		SystemlogInfo:      sqlStr,
		SystemlogStartTime: todaytime.NowTimeFull(),
	}
	if err == nil {
		clinet.SystemlogNote = "成功"
	} else {
		clinet.SystemlogNote = err.Error()
	}
	_, err = SystemLogInsert(clinet)
	if err != nil {
		return
	}
	return

}

func Hostadd(host *models.ParamHostDateGet) (theId int64, err error) {

	hosttime, _ := strconv.ParseInt(host.HostAddTime, 10, 64)
	hostidadd := snowflake.IdNum()
	addtime := time.Unix(hosttime/1000, 0).Format(timeLayoutday) + todaytime.NowTime()

	sqlStr := "insert into hostlist(hostid,hostname,systemtype,hoststatus,hostip,hostlocation,hostowner,hostnote,hostaddtime) values (?,?,?,?,?,?,?,?,?)"
	ret, err := db.Exec(sqlStr,
		hostidadd,
		host.HostName,
		host.SystemType,
		host.HostStatus,
		host.HostIP,
		host.HostLocation,
		host.HostOwner,
		host.HostNote,
		addtime)
	if err != nil {
		return
	}
	theId, err = ret.LastInsertId()
	if err != nil {
		return theId, err
	} else {
		fmt.Printf("主机表插入数据的id 为 %d. \n", theId)
	}
	clinet := models.SystemLog{
		SystemlogHostName:  host.HostName,
		SystemlogType:      "添加主机",
		SystemlogInfo:      sqlStr,
		SystemlogStartTime: addtime,
	}
	var alarmtype = 4011
	sqlStrAlarm := "insert into alarmsetting(hostid,alarmtype,alarmstatus,alarmhostonwer) values (?,?,?,?)"
	ret, err = db.Exec(sqlStrAlarm, hostidadd, alarmtype, 1, host.HostOwner)
	if err != nil {
		return
	}
	theId, err = ret.LastInsertId()
	if err != nil {

		return theId, err
	} else {
		fmt.Printf("主机设置表插入数据的id 为 %d. \n", theId)
	}
	if err == nil {
		clinet.SystemlogNote = "成功"
	} else {
		clinet.SystemlogNote = err.Error()
	}
	_, err = SystemLogInsert(clinet)
	if err != nil {
		return
	}
	return

}
func Hostdel(host *models.ParamHostDateGet) (n int64, err error) {
	sqlStr := "delete  from hostlist where hostid=?"
	clinet := models.SystemLog{
		SystemlogHostName:  HostName(strconv.FormatInt(host.Hostid, 10)),
		SystemlogType:      "删除主机",
		SystemlogInfo:      sqlStr,
		SystemlogStartTime: todaytime.NowTimeFull(),
	}
	if err == nil {
		clinet.SystemlogNote = "成功"
	} else {
		clinet.SystemlogNote = err.Error()
	}
	_, err = SystemLogInsert(clinet)
	if err != nil {
		return
	}

	ret, err := db.Exec(sqlStr, host.Hostid)
	if err != nil {
		return
	}
	n, err = ret.RowsAffected()
	if err != nil {
		return
	} else {
		fmt.Printf("主机表删除数据为 %d 条\n", n)
	}
	return

}
func Hostcheck(host *models.ParamHostDateGet) (err error) {
	sqlStr := `select count(hostid)  from hostlist where hostip = ?`
	var count int
	if err := db.Get(&count, sqlStr, host.HostIP); err != nil {
		return err
	}
	if count > 0 {
		return ErrorHostExist
	}

	return
}
