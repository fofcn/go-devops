package handler

import "net/http"

type HttpHandler interface {
	Pattern() string
	Handler() http.Handler
}
