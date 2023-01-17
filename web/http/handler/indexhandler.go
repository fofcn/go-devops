package handler

import (
	"fmt"
	"net/http"
)

type IndexHandler struct {
	pattern string
	handler http.Handler
}

func NewIndexHandler() HttpHandler {
	index := &IndexHandler{
		pattern: "/",
	}

	index.handler = index
	return index
}

func (index IndexHandler) Pattern() string {
	return index.pattern
}

func (index IndexHandler) Handler() http.Handler {
	return index.handler
}

func (index *IndexHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to CCloud!")
}
