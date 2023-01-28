package service

import (
	"io"
	"os"
	"strings"
	"taskmanager/web/constant"
	"taskmanager/web/dao"
	"taskmanager/web/entity"
	"taskmanager/web/entity/cmd"
	"taskmanager/web/entity/model"
	"time"

	"github.com/google/uuid"
)

type UploadFileService interface {
	Upload(cmd *cmd.UploadCmd) entity.Response
}

type uploadfileserviceimpl struct {
	dao dao.MediaFileDao
}

func NewUploadFileService() (UploadFileService, error) {
	mediaFileDao, err := dao.NewMediaFileDao()
	if err != nil {
		return nil, err
	}

	return &uploadfileserviceimpl{
		dao: mediaFileDao,
	}, nil
}

func (impl uploadfileserviceimpl) Upload(cmd *cmd.UploadCmd) entity.Response {
	// 保留文件后缀名
	originFileName := cmd.FileHeader.Filename
	suffixIdx := strings.LastIndex(originFileName, ".")
	var suffix string = ""
	if suffixIdx != -1 {
		suffix = string(originFileName[suffixIdx])
	}

	// 生成新名字
	uuid := uuid.New()
	newFileName := strings.Join([]string{uuid.String(), suffix}, "")

	// 补充目录
	absPath := "/app/images/" + newFileName

	// 存储到本地文件系统
	f, err := os.Create(absPath)
	defer f.Close()
	if err != nil {
		return entity.Fail(constant.StoreFileFailed)
	}
	io.Copy(f, cmd.File)

	// 转换时间字符串为Time
	createTime, err := time.ParseInLocation("20060102150405", cmd.CreateTime, time.Local)
	if err != nil {
		return entity.Fail(constant.TimeFormatError)
	}

	// 获取文件Mime类型
	var mediaType int = 0
	if cmd.MediaType == "photo" {
		mediaType = 0
	} else if cmd.MediaType == "video" {
		mediaType = 1
	} else {
		mediaType = 2
	}

	// 记录到数据库
	// todo 一致性问题处理：需要处理存储到fs成功，但是插入到数据库失败场景
	model := model.MediaModel{
		FileName:       cmd.FileHeader.Filename,
		StorePath:      absPath,
		FileCreateTime: createTime,
		MediaType:      mediaType,
	}
	err = impl.dao.SaveFile(model)
	if err != nil {
		return entity.Fail(constant.DBInsertError)
	}

	return entity.OK()
}
