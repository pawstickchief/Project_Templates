package mysql

import (
	"crypto/md5"
	"database/sql"
	"encoding/hex"
	"errors"
	"go-web-app/models"
)

// 每一步数据库操作封装为单独的函数
// 等待logic层进行调用

const secret = "wiki.52bucky.cn"

var (
	ErrorUserExist    = errors.New("用户已存在")
	ErrorUserPassword = errors.New("用户名或密码错误")
	ErrorUserNoExist  = errors.New("用户不存在")
)

// 加密明文
func encryptPassword(oPassword string) string {
	h := md5.New()
	h.Write([]byte(secret))
	h.Sum([]byte(oPassword))
	return hex.EncodeToString(h.Sum([]byte(oPassword)))
}

func LoginCode(userinfo *models.User) (err error) {
	// 确保传入的 userinfo 不是 nil
	if userinfo == nil {
		return errors.New("userinfo cannot be nil")
	}

	// 保存传入的姓名
	UserName := userinfo.Name

	// SQL 查询语句
	sqlStr := `SELECT name, userid FROM userinfo WHERE userid = ?`

	// 执行查询
	err = db.Get(userinfo, sqlStr, userinfo.UserId)
	if errors.Is(err, sql.ErrNoRows) {
		return ErrorUserNoExist
	}
	if err != nil {
		return err
	}

	// 检查姓名是否匹配
	if UserName != userinfo.Name {
		return ErrorUserPassword
	}

	return err
}
