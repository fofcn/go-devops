package handler

import (
	"net/http"
	"taskmanager/web/constant"
	"taskmanager/web/entity"
	"taskmanager/web/entity/cmd"
	"taskmanager/web/http/util"
	"taskmanager/web/service"
)

type uploadfilehandler struct {
	pattern       string
	handler       http.Handler
	uploadservice service.UploadFileService
}

func (upload uploadfilehandler) Pattern() string {
	return upload.pattern
}

func (upload uploadfilehandler) Handler() http.Handler {
	return upload.handler
}

func (upload *uploadfilehandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		resp := entity.Fail(constant.FileReadError)
		util.WriteJson(w, resp.ToString())
		return
	}

	createTime := r.FormValue("createTime")
	cmd := cmd.UploadCmd{
		File:       file,
		FileHeader: header,
		CreateTime: createTime,
	}
	defer file.Close()
	response := upload.uploadservice.Upload(&cmd)
	util.WriteJson(w, response)
}
