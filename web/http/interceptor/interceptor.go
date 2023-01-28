package interceptor

import (
	"net/http"

	"github.com/vibrantbyte/go-antpath/antpath"
)

type Interceptor interface {
	preRequest(w http.ResponseWriter, r *http.Request) bool
	afterRequest(w http.ResponseWriter, r *http.Request)
}

type BaseInterceptor struct {
	paths       []string
	interceptor Interceptor
}

func (base BaseInterceptor) preRequest(w http.ResponseWriter, r *http.Request) bool {
	uri := r.RequestURI
	var matcher antpath.PathMatcher = antpath.New()
	for _, path := range base.paths {
		matched := matcher.Match(path, uri)
		if matched {
			return base.interceptor.preRequest(w, r)
		}
	}

	return true
}

func (base BaseInterceptor) afterRequest(w http.ResponseWriter, r *http.Request) {

}
