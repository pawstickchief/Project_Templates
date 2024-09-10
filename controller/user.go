package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go-web-app/dao/mysql"
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

func SelectNeighbors(c *gin.Context) {
	p := new(models.SelectNeighbors)
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
	s, err := logic.SelectNeighborsLogic(p)
	if err != nil {
		zap.L().Error("switchneighbors is fail", zap.String("ParameterType", p.SwitchName), zap.Error(err))
		ResopnseError(c, CodeAlarminfo)

		return
	}
	//3.返回响应

	ResopnseSystemDataSuccess(c, s)
}
func InterfaceDetail(c *gin.Context) {
	p := new(models.SelectNeighbors)
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
	s, err := logic.SelectInterface(p)
	if err != nil {
		zap.L().Error("switchneighbors is fail", zap.String("ParameterType", p.SwitchName), zap.Error(err))
		ResopnseError(c, CodeAlarminfo)

		return
	}
	//3.返回响应

	ResopnseSystemDataSuccess(c, s)
}
func SelectSwitchTotal(c *gin.Context) {
	p := new(models.SelectNeighbors)
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
	if p.SwitchName != "8.66" {
		var errs validator.ValidationErrors
		ResponseErrorwithMsg(c, CodeServerApiType, removeTopStruct(errs.Translate(trans)))
		return
	}
	s, err := mysql.SelectTotalSwitch()
	if err != nil {
		zap.L().Error("switchneighbors is fail", zap.String("ParameterType", p.SwitchName), zap.Error(err))
		ResopnseError(c, CodeAlarminfo)

		return
	}
	//3.返回响应

	ResopnseSystemDataSuccess(c, s)
}

func LoginUserVerif(c *gin.Context) {
	p := new(models.LoginUserinfo)
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
	var userinfo models.User
	userinfo.Name = p.UserName
	userinfo.UserId = p.UserCode
	err := mysql.LoginCode(&userinfo)
	if err != nil {
		zap.L().Error("用户信息核对失败", zap.String("ParameterType", p.UserName), zap.Error(err))
		ResopnseError(c, CodeUserNotExist)

		return
	}
	//3.返回响应

	ResopnseSystemDataSuccess(c, "用户信息核对正确")
}

func SelectUplinkInfo(c *gin.Context) {
	p := new(models.SwitchUplinkInfo)
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
	s, err := logic.SelectSwitchUplink(p)
	if err != nil {
		ResopnseError(c, CodeSelectSwitch)
		return
	}

	ResopnseSystemDataSuccess(c, s)
}
