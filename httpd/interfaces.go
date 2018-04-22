package httpd

import (
	"reflect"
	"net/http"
)

//------------------------------
// Common interfaces

type PropertyBag interface {
	GetValue(key string) (interface{}, bool)
	SetValue(key string, value interface{})
	DeleteValue(key string)
	GetAllKeys() []string
	GetInterface(t reflect.Type) (interface{}, bool)
	SetInterface(t reflect.Type, value interface{})
}

type LifeCycleManaged interface {
	Initialize(PropertyBag)
	Destroy()
}

//------------------------------
// Http Request Interceptor interfaces

type HttpRequestContext interface {
	PropertyBag
}

type HttpInterceptorChain interface {
	LifeCycleManaged
	http.Handler
	DoChain(request *http.Request, response http.ResponseWriter, context HttpRequestContext)
}

type CreateInterceptorMethod func() HttpInterceptor;

type HttpInterceptor interface {
	LifeCycleManaged
	Intercept(chain HttpInterceptorChain, request *http.Request, response http.ResponseWriter, context HttpRequestContext)
}

type CreateInterceptorExceptionHandlerMethod func() HttpInterceptorExceptionHandler;

type HttpInterceptorExceptionHandler interface {
	HandleException(request *http.Request, response http.ResponseWriter, context HttpRequestContext, exception interface{})
}

type HttpInterceptorException interface {
	error
	Cause() error
	Status() (int, bool)
	Content() interface{}
}

//------------------------------
// Service Point interfaces

type ServicePoint struct {
	UrlPattern string
	Methods []string
	handler interface{}
}

type HttpRequestResult interface {
}

type HttpRequestPathParam map[string]string

//------------------------------
// Known types

var (
	PropertyBagType = reflect.TypeOf((*PropertyBag)(nil)).Elem()
	HttpRequestContextType = reflect.TypeOf((*HttpRequestContext)(nil)).Elem()
	ServicePointType = reflect.TypeOf((*ServicePoint)(nil)).Elem()
	HttpRequestResultType = reflect.TypeOf((*HttpRequestResult)(nil)).Elem()
)
