package httpd

type HttpRequestContextBase struct {
	PropertyBag
}

func NewHttpRequestContext() HttpRequestContext {
	context := HttpRequestContextBase{
		NewPropertyBag(),
	}
	return &context
}