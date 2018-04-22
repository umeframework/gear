package httpd

import (
	"errors"
	"fmt"
	"net/http"
)

type CreateInterceptorFunc func() HttpInterceptor

var (
	interceptorMap map[string]CreateInterceptorFunc
	ErrorInterceptorNameNotFound = errors.New("interceptor name not found")
)

type HttpInterceptorChainBase struct {
	interceptors []HttpInterceptor
	exceptionHandler HttpInterceptorExceptionHandler
}

func (this *HttpInterceptorChainBase) ServeHTTP(response http.ResponseWriter, request *http.Request) {
	// Create request context
	context := NewHttpRequestContext()

	defer func() {
		if exception := recover(); exception != nil {
			this.HandleException(request, response, context, exception)
		}
	}()

	// Call interceptor chain
	this.DoChain(request, response, context)
}

func (this *HttpInterceptorChainBase) Initialize(propertyBag PropertyBag) {
	for _, interceptor := range this.interceptors {
		interceptor.Initialize(propertyBag)
	}
}

func (this *HttpInterceptorChainBase) Destroy() {
	for _, interceptor := range this.interceptors {
		interceptor.Destroy()
	}
}

func (this *HttpInterceptorChainBase) DoChain(request *http.Request, response http.ResponseWriter, context HttpRequestContext) {
	if len(this.interceptors) > 0 {
		interceptor := this.interceptors[0]
		chain := this.CloneNext()
		interceptor.Intercept(chain, request, response, context)
	}
}

func (this *HttpInterceptorChainBase) HandleException(request *http.Request, response http.ResponseWriter, context HttpRequestContext, exception interface{}) {
	if this.exceptionHandler != nil {
		this.exceptionHandler.HandleException(request, response, context, exception)
	}
}

func (this *HttpInterceptorChainBase) CloneNext() HttpInterceptorChain {
	var interceptors []HttpInterceptor
	if len(this.interceptors) > 1 {
		interceptors = this.interceptors[1:]
	}
	chain := HttpInterceptorChainBase{
		interceptors: interceptors,
	}
	return &chain
}

func NewHttpInterceptorChain(interceptors []HttpInterceptor, exceptionHandler HttpInterceptorExceptionHandler) HttpInterceptorChain {
	chain := HttpInterceptorChainBase{
		interceptors,
		exceptionHandler,
	}
	return &chain
}

func NewHttpInterceptorChainByNames(names []string, exceptionHandler HttpInterceptorExceptionHandler) (HttpInterceptorChain, error) {
	interceptors := make([]HttpInterceptor, len(names))
	for i, name := range names {
		if fnCreate, found := interceptorMap[name]; found {
			interceptor := fnCreate()
			interceptors[i] = interceptor
		} else {
			return nil, errors.New(fmt.Sprintf("interceptor not found for name: %s", name))
		}
	}
	return NewHttpInterceptorChain(interceptors, exceptionHandler), nil
}

