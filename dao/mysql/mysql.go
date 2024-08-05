package mysql

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"go-web-app/settings"
	"go.uber.org/zap"
)

var db *sqlx.DB

func Init(cfg *settings.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True",
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DbName)
	//初始化全局的DB对象
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("connect to db failed", zap.Error(err))
		return
	}
	db.SetMaxOpenConns(cfg.MaxConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	return

}

func Close() {
	_ = db.Close()
}
