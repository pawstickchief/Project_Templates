package controller

type ResCode int64

// 定义错误码
const (
	CodeSuccess ResCode = 1000 + iota
	CodeInvalidParam
	CodeUserExist
	CodeUserNotExist
	CodeInvalidPassword
	CodeServerBusy
	CodeNeedLogin
	CodeInvalidAuth
	CodeServerApiType
	CodeHostlist
	CodeAlarminfo
)

var codeMsgMap = map[ResCode]string{
	CodeSuccess:         "success",
	CodeInvalidParam:    "请求参数错误",
	CodeUserExist:       "用户已存在",
	CodeUserNotExist:    "用户不存在",
	CodeInvalidPassword: "用户名或密码输入错误",
	CodeServerBusy:      "服务器忙",
	CodeNeedLogin:       "需要登陆",
	CodeInvalidAuth:     "无效的token",
	CodeServerApiType:   "接口参数错误",
	CodeHostlist:        "主机已存在",
	CodeAlarminfo:       "报警接口参数错误",
}

func (c ResCode) Msg() string {
	msg, ok := codeMsgMap[c]
	if !ok {
		msg = codeMsgMap[CodeServerBusy]
	}
	return msg
}
