package mysql

import (
	"fmt"
	"go-web-app/models"
	"go-web-app/pkg/snowflake"
	"go-web-app/pkg/todaytime"
)

var (
	optionok       = "成功"
	optionadd      = "上传文件"
	optiondel      = "删除文件"
	optiondownload = "下载文件"
)

func FileLogAdd(file *models.Filelog) (err error) {
	fileid := snowflake.IdNum()
	uploadtime := todaytime.NowTimeFull()
	sqlStr := "insert into filelog(fileid,filename,filesize,filedir,uploadtime) values (?,?,?,?,?)"
	ret, err := db.Exec(sqlStr,
		fileid,
		file.FileName,
		file.FileSize,
		file.FileDir,
		uploadtime)
	if err != nil {
		return
	}
	theId, err := ret.LastInsertId()
	if err != nil {
		return
	} else {
		fmt.Printf("插入数据的id 为 %d. \n", theId)
	}
	filedata := &models.FileOption{
		FileId:     fileid,
		FileName:   file.FileName,
		FileInfo:   optionok,
		FileOption: optionadd,
		OptionTime: uploadtime,
	}
	_, err = FileOption(filedata)
	if err != nil {
		return err
	}
	return
}

func FileName(fileid int64) (filename string) {
	sqlStr := `select filename  from filelog where fileid = ?`
	if err := db.Get(&filename, sqlStr, fileid); err != nil {
		return
	}
	return
}
func FileDir(fileid int64) (filedir string) {
	sqlStr := `select filedir  from filelog where fileid = ?`
	if err := db.Get(&filedir, sqlStr, fileid); err != nil {
		return
	}
	return
}
func FileOption(host *models.FileOption) (Reply int64, err error) {
	sqlStr := "insert into filedata(fileid,filename,fileoption,fileinfo,optiontime) values (?,?,?,?,?)"
	ret, err := db.Exec(sqlStr,
		host.FileId,
		host.FileName,
		host.FileOption,
		host.FileInfo,
		todaytime.NowTimeFull(),
	)
	Reply, err = ret.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("更新数据为 %d 条\n", Reply)
	}
	return
}
