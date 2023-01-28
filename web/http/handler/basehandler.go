package handler

import "net/http"

type BaseHandler struct {
	handler *HttpHandler
}

func (base *BaseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

}
