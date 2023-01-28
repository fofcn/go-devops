package interceptor

import (
	"sync"
)

type InterceptorChain interface {
	RegisterInterceptor(Interceptor)
}

type DefaultInterceptorChain struct {
	list []Interceptor
}

var instance *InterceptorChain
var once sync.Once

func GetInstance() InterceptorChain {
	// once.Do(func() {
	// 	instance = &DefaultInterceptorChain{
	// 		list: make([]Interceptor, 0),
	// 	}
	// })

	return *instance
}

func (chain DefaultInterceptorChain) RegisterInterceptor(interceptor Interceptor) {

}
