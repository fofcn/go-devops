package model

import "time"

type UserModel struct {
	Id         int
	Username   string
	Password   string
	CreateTime time.Time
}
