package dao

import (
	"errors"
	"taskmanager/web/entity/model"
	"taskmanager/web/store"
)

type MediaFileDao interface {
	SaveFile(model.MediaModel) error
}

type MediaFileDaoImpl struct {
	store store.SqlStore
}

func NewMediaFileDao() (MediaFileDao, error) {
	store, err := store.NewSingleSqliteStore(".", "db.sql")
	if err != nil {
		return nil, err
	}
	return &MediaFileDaoImpl{
		store: store,
	}, nil
}

func (impl MediaFileDaoImpl) SaveFile(media model.MediaModel) error {
	sql := "insert into `media_file` (file_name, store_path, file_create_time, media_type, create_time) values (?, ?, ?, ?, ?)"
	affected, err := impl.store.Insert(sql, media.FileName, media.StorePath, media.FileCreateTime, media.MediaType, media.CreateTime)
	if err != nil {
		return err
	}

	if affected != 1 {
		return errors.New("insert data failed")
	}

	return nil
}
