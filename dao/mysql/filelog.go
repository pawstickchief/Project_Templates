package mysql

import (
	"fmt"
	"go-web-app/models"
	"go-web-app/pkg/codeconversion"
	"go-web-app/pkg/snowflake"
	"go-web-app/pkg/todaytime"
	"golang.org/x/crypto/ssh"
	"os"
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

func FileLogGet(host *models.ParameCrontab) (data []*models.Filelog, err error) {
	sqlStr := "select fileid,filename,uploadtime,filedir,filesize from filelog ORDER BY `uploadtime` DESC;"
	if err = db.Select(&data, sqlStr); err != nil {
		return
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
func FileLogDel(host *models.ParameCrontab) (delid int64, err error) {
	filename := FileDir(host.FileId)
	err = os.Remove(filename)
	if err == nil {
		filedata := &models.FileOption{
			FileId:     host.FileId,
			FileName:   FileName(host.FileId),
			FileInfo:   optionok,
			FileOption: optiondel,
			OptionTime: todaytime.NowTimeFull(),
		}
		_, err = FileOption(filedata)
	}
	sqlStr := "delete  from filelog where fileid=?"
	ret, err := db.Exec(sqlStr, host.FileId)
	if err != nil {
		return
	}
	delid, err = ret.RowsAffected()
	if err != nil {
		return
	} else {
		fmt.Printf("文件删除%d 个\n", delid)
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
func FileOptionLogGet(host *models.ParameCrontab) (data []*models.FileOption, err error) {
	sqlStr := "select fileid,filename,fileoption,fileinfo,optiontime from filedata ORDER BY `optiontime` DESC;"
	if err = db.Select(&data, sqlStr); err != nil {
		return
	}
	return
}
func LogClean(tablename string) (Reply int64, err error) {
	sqlStr := "truncate table " + tablename
	ret, err := db.Exec(sqlStr)
	Reply, err = ret.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Printf("更新数据为 %d 条\n", Reply)
	}
	return
}
func SshShell(sshparam *models.ParameCrontab) (result string, err error) {

	config := &ssh.ClientConfig{
		User: sshparam.UserName,
		Auth: []ssh.AuthMethod{
			ssh.Password(sshparam.Password),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}

	client, err := ssh.Dial("tcp", fmt.Sprintf("%s:%d", sshparam.Address, sshparam.Port), config)
	if err != nil {
		panic("Failed to dial: " + err.Error())
	}
	defer func(client *ssh.Client) {
		err := client.Close()
		if err != nil {

		}
	}(client)

	// 执行命令
	session, err := client.NewSession()
	if err != nil {
		fmt.Println("Failed to create session: " + err.Error())
	}
	defer func(session *ssh.Session) {
		err := session.Close()
		if err != nil {

		}
	}(session)

	if buf, err := session.CombinedOutput(sshparam.Shell); err != nil {
		result = codeconversion.ConvertByte2String(buf, "GB18030")

	} else {
		result = codeconversion.ConvertByte2String(buf, "GB18030")
	}

	return
}
