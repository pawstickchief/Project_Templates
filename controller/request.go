package controller

import (
	"errors"
	"github.com/gin-gonic/gin"
)

var ErrorUserNotLogin = errors.New("用户未登录")

const ContextUserIdKey = "userId"

// 获取当前用户id
func GetCurrentUser(c *gin.Context) (userId int64, err error) {

	uid, ok := c.Get(ContextUserIdKey)
	if !ok {
		err = ErrorUserNotLogin
		return
	}
	userId, ok = uid.(int64)
	if !ok {
		err = ErrorUserNotLogin
		return
	}

	return
}
