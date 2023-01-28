package model

import "time"

type UserModel struct {
	Id         int
	Username   string
	Password   string
	CreateTime time.Time
}

type MediaModel struct {
	Id             int
	FileName       string
	StorePath      string
	FileCreateTime time.Time
	MediaType      int
	CreateTime     time.Time
}
