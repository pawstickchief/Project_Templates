package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-web-app/logic"
	"go-web-app/models"
	"go-web-app/pkg/macswitch"
	"go.uber.org/zap"
	"strconv"
)

// SelectSwitchMac 查询交换机端口
// @Summary 查询交换机端口
// @Description 查询交换机端口
// @Accept  json
// @Produce  json
// @Param data body models.SelectSwitchMac true "查询交换机的参数"
// @Success 200 {object} models.ClientSwitchInfo "成功"
// @Failure 500 {object} models.ErrorResponse "内部错误"
// @Router /selectswitch [post]
func SelectSwitchMac(c *gin.Context) {
	p := new(models.SelectSwitchMac)
	if err := c.ShouldBindJSON(&p); err != nil {
		//请求参数有误,直接返回响应
		var errs validator.ValidationErrors
		ok := errors.As(err, &errs)
		if !ok {
			ResopnseError(c, CodeServerApiType)
			return
		}
		ResponseErrorwithMsg(c, CodeServerApiType, removeTopStruct(errs.Translate(trans)))
		return
	}
	p.ShortMAC = macswitch.FormatMACAddress(p.ShortMAC)
	s, err := logic.SelectSwitchInfoOption(p)
	if err != nil {
		zap.L().Error("selectswitchm with invalid param", zap.String("ParameterType", strconv.Itoa(p.SwitchLevel)), zap.Error(err))
		ResopnseError(c, CodeAlarminfo)

		return
	}
	//3.返回响应

	ResopnseSystemDataSuccess(c, s)
}
func SelectSwitchChangeVlan(c *gin.Context) {
	p := new(models.SelectSwitchMac)
	if err := c.ShouldBindJSON(&p); err != nil {
		//请求参数有误,直接返回响应
		var errs validator.ValidationErrors
		ok := errors.As(err, &errs)
		if !ok {
			ResopnseError(c, CodeServerApiType)
			return
		}
		ResponseErrorwithMsg(c, CodeServerApiType, removeTopStruct(errs.Translate(trans)))
		return
	}
	p.ShortMAC = macswitch.FormatMACAddress(p.ShortMAC)
	s, err := logic.SelectSwitchOption(p)
	if err != nil {
		zap.L().Error("selectswitchm with invalid param", zap.String("ParameterType", strconv.Itoa(p.SwitchLevel)), zap.Error(err))
		ResopnseError(c, CodeAlarminfo)

		return
	}
	//3.返回响应

	ResopnseSystemDataSuccess(c, s)
}
