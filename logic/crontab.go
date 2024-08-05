package logic

import (
	"fmt"
	"go-web-app/dao/mysql"
	"go-web-app/models"
	"runtime"
)

func CrontabInit() {
	runtime.GOMAXPROCS(runtime.NumCPU())
}

func Crond(p *models.ParameCrontab) (Reply interface{}, err error) {
	switch {
	case p.ParameOption == "add":
		err = mysql.CheckJob(p)
		if err != nil {
			fmt.Println("任务已存在")
			return
		}
		Reply, err := mysql.CrontabAdd(p)
		if err != nil {
			return Reply, err
		}
		return Reply, err
	case p.ParameOption == "del":
		Reply, err := mysql.CrontabDel(p)
		if err != nil {
			return Reply, err
		}
		return Reply, err
	case p.ParameOption == "edit":
		Reply, err := mysql.CrontabEdit(p)
		if err != nil {
			return Reply, err
		}
		return Reply, err
	case p.ParameOption == "init":
		Reply, err := mysql.CrontabSelect(p)
		if err != nil {
			return Reply, err
		}
		return Reply, err
	case p.ParameOption == "killjob":
		Reply, err := mysql.KillJob(p)
		if err != nil {
			return Reply, err
		}
		return Reply, err
	case p.ParameOption == "jobtotal":
		Reply, err := mysql.CrontabTotal(p)
		if err != nil {
			return Reply, err
		}
		return Reply, err
	case p.ParameOption == "jobonline":
		Reply, err := mysql.CrontabOnline(p)
		if err != nil {
			return Reply, err
		}
		return Reply, err
	case p.ParameOption == "jobtodaytotal":
		Reply, err := mysql.CrontabTodayTotal(p)
		if err != nil {
			return Reply, err
		}
		return Reply, err
	case p.ParameOption == "jobaddtoday":
		Reply, err := mysql.CrontabAddToday(p)
		if err != nil {
			return Reply, err
		}
		return Reply, err
	case p.ParameOption == "taskjoblog":
		Reply, err := mysql.TaskJobLog(p)
		if err != nil {
			return Reply, err
		}
		return Reply, err
	case p.ParameOption == "taskjobselect":
		Reply, err := mysql.TaskJobLogSelect(p)
		if err != nil {
			return Reply, err
		}
		return Reply, err
	case p.ParameOption == "logmsgget":
		Reply, err := mysql.LogmsgGet(p)
		if err != nil {
			return Reply, err
		}
		return Reply, err
	case p.ParameOption == "logmsgsystem":
		Reply, err := mysql.SystemLogGet(p)
		if err != nil {
			return Reply, err
		}
		return Reply, err
	case p.ParameOption == "logmsgfile":
		Reply, err := mysql.FileLogGet(p)
		if err != nil {
			return Reply, err
		}
		return Reply, err
	case p.ParameOption == "fileremove":
		Reply, err := mysql.FileLogDel(p)
		if err != nil {
			return Reply, err
		}
		return Reply, err
	case p.ParameOption == "logoption":
		Reply, err := mysql.FileOptionLogGet(p)
		if err != nil {
			return Reply, err
		}
		return Reply, err
	case p.ParameOption == "systemlogclean":
		Reply, err := mysql.LogClean("systemlog")
		if err != nil {
			return Reply, err
		}
		return Reply, err
	case p.ParameOption == "uploadlogclean":
		Reply, err := mysql.LogClean("filedata")
		if err != nil {
			return Reply, err
		}
		return Reply, err
	case p.ParameOption == "alarmlogclean":
		Reply, err := mysql.LogClean("alarmstatistics")
		if err != nil {
			return Reply, err
		}
		return Reply, err
	case p.ParameOption == "tasklogclean":
		Reply, err := mysql.LogClean("jobdata")
		if err != nil {
			return Reply, err
		}
		return Reply, err
	case p.ParameOption == "sshshell":
		Reply, err := mysql.SshShell(p)
		if err != nil {
			return Reply, err
		}
		return Reply, err
	}

	return
}
