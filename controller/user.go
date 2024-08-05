package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-web-app/logic"
	"go-web-app/models"
	"go.uber.org/zap"
)

func Alarmsetting(c *gin.Context) {
	p := new(models.ParamAlarmSetting)
	if err := c.ShouldBindJSON(&p); err != nil {
		//请求参数有误,直接返回响应

		errs, ok := err.(validator.ValidationErrors)
		if !ok {
			ResopnseError(c, CodeServerApiType)
			return
		}
		ResponseErrorwithMsg(c, CodeServerApiType, removeTopStruct(errs.Translate(trans)))
		return
	}
	s, err := logic.AlarmOption(p)
	if err != nil {
		zap.L().Error("hostlitdata with invalid param", zap.String("ParameterType", p.AlarmSettingOption), zap.Error(err))
		ResopnseError(c, CodeAlarminfo)

		return
	}
	//3.返回响应

	ResopnseSystemDataSuccess(c, s)
}
