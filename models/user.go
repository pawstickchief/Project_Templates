package models

type User struct {
	Id     int64  `db:"id"`
	Name   string `db:"name"`
	UserId int    `db:"userid"`
	Status int    `db:"status"`
}
type LoginUserinfo struct {
	UserName string `json:"userName" binding:"required"`
	UserCode int    `json:"UserCode"`
}
