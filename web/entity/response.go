package entity

import (
	"encoding/json"
	"taskmanager/web/constant"
)

type Response struct {
	Success   bool        `json:"success"`
	ErrorCode string      `json:"errorCode"`
	ErrorMsg  string      `json:"errorMsg"`
	Data      interface{} `json:"data"`
}

func OK() Response {
	return Response{
		Success:   true,
		ErrorCode: constant.OK.Code,
		ErrorMsg:  constant.OK.Msg,
		Data:      nil,
	}
}

func OKWithData(data interface{}) Response {
	return Response{
		Success:   true,
		ErrorCode: constant.OK.Code,
		ErrorMsg:  constant.OK.Msg,
		Data:      data,
	}
}

func Fail(errCode constant.ErrorCode) Response {
	return Response{
		Success:   false,
		ErrorCode: errCode.Code,
		ErrorMsg:  errCode.Msg,
		Data:      nil,
	}
}

func (resp *Response) ToString() string {
	// r := &struct {
	// 	Success   bool        `json:"success"`
	// 	ErrorCode string      `json:"errorCode"`
	// 	ErrorMsg  string      `json:"errorMsg"`
	// 	Data      interface{} `json:"data"`
	// }{
	// 	Success:   resp.Success,
	// 	ErrorCode: resp.ErrorCode,
	// 	ErrorMsg:  resp.ErrorMsg,
	// 	Data:      resp.Data,
	// }

	bytes, _ := json.Marshal(resp)
	return string(bytes)
}
