package dao

import (
	"errors"
	"taskmanager/web/entity/model"
	"taskmanager/web/store"
	"time"
)

type AccountDao interface {
	SelectByUsername(string) (model.UserModel, error)
}

type accountdaoimpl struct {
	store store.SqlStore
}

func NewAccountDao() (AccountDao, error) {
	store, err := store.NewSingleSqliteStore(".", "db.sql")
	if err != nil {
		return nil, err
	}
	return &accountdaoimpl{
		store: store,
	}, nil
}

func (impl accountdaoimpl) SelectByUsername(username string) (model.UserModel, error) {
	// 根据用户名查询数据库
	sql := "select * from userinfo where username=?"
	data, err := impl.store.Query(sql, username)
	if err != nil {
		return model.UserModel{}, err
	}

	if len(data) == 0 {
		return model.UserModel{}, errors.New("user not found")
	}

	record := data[0]
	// 组装数据返回
	return model.UserModel{
		Id:         record["id"].(int),
		Username:   record["username"].(string),
		Password:   record["password"].(string),
		CreateTime: record["create_time"].(time.Time),
	}, nil
}
