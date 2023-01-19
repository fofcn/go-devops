package handler

import (
	"io/ioutil"
	"net/http"
	"taskmanager/web/http/util"
)

type loginhandler struct {
	pattern string
	handler http.Handler
}

func NewLoginHandler() HttpHandler {
	login := &loginhandler{
		pattern: "/account/login",
	}

	login.handler = login
	return login
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
		w.Write([]byte(err.Error()))
	}

	// read account info from file

	util.WriteJsonBytes(w, buf)
}
