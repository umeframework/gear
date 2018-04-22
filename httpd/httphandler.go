package httpd

import (
	"net/http"
)

type SimpleHttpHandler struct {
	interceptors []interface{}
	exceptionHandler interface{}
	propertyBag PropertyBag
}

func (this *SimpleHttpHandler) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	if chain := this.CreateHttpHandler(); chain != nil {
		chain.Initialize(this.propertyBag)
		defer chain.Destroy()
		chain.ServeHTTP(response, request)
	}
}

func (this *SimpleHttpHandler) CreateHttpHandler() HttpInterceptorChain {
	realInterceptors := make([]HttpInterceptor, 0)
	for _, interceptor := range this.interceptors {
		if realInterceptor := this.CreateHttpInterceptor(interceptor); realInterceptor != nil {
			realInterceptors = append(realInterceptors, realInterceptor)
		}
	}
	realExceptionHandler := this.CreateExceptionHandler()
	chain := NewHttpInterceptorChain(realInterceptors, realExceptionHandler)
	return chain
}

func (this *SimpleHttpHandler) CreateHttpInterceptor(object interface{}) HttpInterceptor {
	var interceptor HttpInterceptor = nil
	var ok = false
	if interceptor, ok = object.(HttpInterceptor); !ok {
		var fnCreate CreateInterceptorMethod = nil
		if fnCreate, ok = this.exceptionHandler.(CreateInterceptorMethod); ok {
			interceptor = fnCreate()
		}
	}
	return interceptor
}

func (this *SimpleHttpHandler) CreateExceptionHandler() HttpInterceptorExceptionHandler {
	var handler HttpInterceptorExceptionHandler = nil
	var ok = false
	if handler, ok = this.exceptionHandler.(HttpInterceptorExceptionHandler); !ok {
		var fnCreate CreateInterceptorExceptionHandlerMethod = nil
		if fnCreate, ok = this.exceptionHandler.(CreateInterceptorExceptionHandlerMethod); ok {
			handler = fnCreate()
		}
	}
	return handler
}

func NewHttpHandler(interceptors []interface{}, exceptionHandler interface{}, propertyBag PropertyBag) http.Handler {
	handler := SimpleHttpHandler{
		interceptors,
		exceptionHandler,
		propertyBag,
	}
	return &handler
}
