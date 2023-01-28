package handler

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"taskmanager/web/constant"
	"taskmanager/web/entity"
	"taskmanager/web/entity/cmd"
	"taskmanager/web/http/util"
	"taskmanager/web/service"
)

type loginhandler struct {
	pattern      string
	handler      http.Handler
	loginservice service.LoginService
}

func NewLoginHandler() (HttpHandler, error) {
	loginservice, err := service.NewLoginService()
	if err != nil {
		return nil, err
	}

	login := &loginhandler{
		pattern:      "/account/login",
		loginservice: loginservice,
	}

	login.handler = login
	return login, nil
}

func (login loginhandler) Pattern() string {
	return login.pattern
}

func (login loginhandler) Handler() http.Handler {
	return login.handler
}

func (login *loginhandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	buf, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resp := entity.Fail(constant.RequestReadErr)
		util.WriteJson(w, resp.ToString())
		return
	}

	var cmd cmd.LoginCmd
	err = json.Unmarshal(buf, cmd)
	if err != nil {
		resp := entity.Fail(constant.JsonStringToInterfaceErr)
		util.WriteJson(w, resp.ToString())
		return
	}

	resp := login.loginservice.Login(&cmd)
	util.WriteJson(w, resp.ToString())
}
