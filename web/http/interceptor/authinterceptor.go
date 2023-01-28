package interceptor

import (
	"fmt"
	"net/http"
)

type AuthInterceptor struct {
	paths []string
}

func NewAuthInterceptor(paths ...string) Interceptor {
	return &AuthInterceptor{
		paths: paths,
	}
}

func (auth AuthInterceptor) preRequest(w http.ResponseWriter, r *http.Request) bool {
	uri := r.RequestURI
	fmt.Println(uri)
	return true
}

func (auth AuthInterceptor) afterRequest(w http.ResponseWriter, r *http.Request) {

}
