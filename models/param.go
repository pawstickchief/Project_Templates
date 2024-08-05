package models

import clientv3 "go.etcd.io/etcd/client/v3"

//定义请求的参数结构体

// 用户注册参数
type ParamSignUp struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	RePassword string `json:"re_password" binding:"required,eqfield=Password"`
	Email      string `json:"email"`
}

// 用户登陆参数
type ParamLogin struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type ParamSystemGet struct {
	ParameterType string `json:"parametertype" binding:"required"`
	ClientParame  `json:"clientparame"`
}
type ClientParame struct {
	Hostid             int64  `json:"hostid"`
	Hostname           string `json:"hostname"`
	OptionTime         string `json:"optiontime"`
	OptionNote         string `json:"optionnote"`
	OptionIp           string `json:"optionip"`
	OpitonParame       string `json:"opitonparame"`
	OptionParameCpu    string `json:"optionparamecpu"`
	OptionParameMemory string `json:"optionparamememory"`
	OptionParameDisk   string `json:"optionparamedisk"`
	OptionParameUns    string `json:"optionparameuns"`
	OptionParameDns    string `json:"optionparamedns"`
}

type ParamHostDateGet struct {
	TypeOperation string `json:"typeoperation" binding:"required"`
	Hostlist      `json:"hostlist"`
}
type ParamStatistics struct {
	StatisticsType string `json:"statisticstype" binding:"required"`
	Hostline       int    `json:"hostonline" `
	Alarmline      int    `json:"alarmonline" `
	Alarmlist      `json:"alarmlist"`
}

type ParamAlarmSetting struct {
	AlarmSettingOption string `json:"alarmoption" binding:"required"`
	//若数据为空值使用指针
	CpuOption        *int `json:"cpuoption"`
	MemoryOption     *int `json:"memoryoption"`
	SystemDiskOption *int `json:"systemdiskoption"`
	ThresholdStatus  *int `json:"thresholadstatus"`
	Alarmlist        `json:"alarmlist"`
	NotiAPI          `json:"notiapi"`
}
type Alarmlist struct {
	Alarmid        int64  `json:"alarmid"`
	Hostid         int64  `json:"hostid" `
	AlarmStatus    int    `json:"alarmstatus"`
	AlarmType      int    `json:"alarmtype"`
	AlarmInfo      string `json:"alarminfo"`
	AlarmNote      string `json:"alarmnote"`
	AlarmStartTime string `json:"alarmstarttime"`
	AlarmStopTime  string `json:"alarmstoptime"`
	AlarmHostOnwer string `json:"alarmhostonwer"`
	AlarmHostName  string `json:"alarmhostname"`
	AlarmHostIp    string `json:"alarmhostip"`
}
type Hostlist struct {
	Hostid         int64   `json:"hostid"`
	HostName       string  `json:"hostname" bindding:"required"`
	SystemType     string  `json:"systemtype" bindding:"required"`
	HostStatus     int     `json:"hoststatus" bindding:"required"`
	HostIP         string  `json:"hostip" bindding:"required"`
	HostLocation   string  `json:"hostlocation" bindding:"required"`
	HostOwner      string  `json:"hostowner" bindding:"required"`
	HostAddTime    string  `json:"hostaddtime" bindding:"required"`
	HostNote       string  `json:"hostnote" bindding:"required"`
	HostSystemInfo string  `json:"hostsysteminfo"`
	HostUptime     *string `json:"hostuptime"`
	HostIssues     int     `json:"hostissues"`
}

type NotiAPI struct {
	WorkApiUrl *string `json:"workapiurl"`
	DingApiUrl *string `json:"dingapiurl"`
	DingAtuser *string `json:"dingatuser"`
	WorkAtuser *string `json:"workatuser"`
	Text       string  `json:"content"`
}

type ParameCrontab struct {
	ParameOption string `json:"parameoption" bindding:"required"`
	CrontabJob   `json:"crontabjob"`
	Job          `json:"job"`
	JobMgr       `json:"jobmgr"`
	SystemLog    `json:"systemlog"`
	Filelog      `json:"filelog"`
	SshTool      `json:"sshtool"`
}

type Job struct {
	Name     string `json:"name"`
	Command  string `json:"command"`
	CronExpr string `json:"cronexpr"`
}
type JobMgr struct {
	Kv     clientv3.KV
	Lease  clientv3.Lease
	Clinet *clientv3.Client
}

type CrontabJob struct {
	JobId        int    `json:"jobid"`
	JobCronExpr  string `json:"jobcronexpr"`
	JobName      string `json:"jobname"`
	JobShell     string `json:"jobshell"`
	JobStatus    int    `json:"jobstatus"`
	JobStartTime string `json:"jobstarttime"`
	JobStopTime  string `json:"jobstoptime"`
	JobInfo      string `json:"jobinfo"`
	JobRunning   string `json:"jobrunning"`
	JobErr       string `json:"joberr"`
}

type SystemLog struct {
	SystemlogId        int64  `json:"systemlogid"`
	SystemlogHostName  string `json:"systemloghostname" `
	SystemlogType      string `json:"systemlogtype"`
	SystemlogInfo      string `json:"systemloginfo"`
	SystemlogNote      string `json:"systemlognote"`
	SystemlogStartTime string `json:"systemlogstarttime"`
}

type Filelog struct {
	FileName   string `json:"filename"`
	FileId     int64  `json:"fileid"`
	Uploadtime string `json:"uploadtime"`
	FileSize   int64  `json:"filesize"`
	FileDir    string `json:"filedir"`
}

type FileOption struct {
	FileId     int64  `json:"fileid"`
	FileName   string `json:"filename"`
	FileInfo   string `json:"fileinfo"`
	FileOption string `json:"fileoption"`
	OptionTime string `json:"optiontime"`
}

type SshTool struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	Address  string `json:"address"`
	Port     int    `json:"port"`
	Shell    string `json:"shell"`
}
