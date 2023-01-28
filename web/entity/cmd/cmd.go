package cmd

import "mime/multipart"

type LoginCmd struct {
	User string `json:"username"`
	Pass string `json:"password"`
}

type UploadCmd struct {
	File       multipart.File
	FileHeader *multipart.FileHeader
	CreateTime string
	MediaType  string
}
