package models

type User struct {
	Id     int64  `db:"id"`
	Name   string `db:"name"`
	UserId int    `db:"userid"`
	Status int    `db:"status"`
}
