package httpd

import "fmt"

type HttpInterceptorExceptionBase struct {
	cause error
	statusCode int
	content interface{}
}

func (this *HttpInterceptorExceptionBase) Error() string {
	return fmt.Sprintf("http interceptor exception, conent = %v, cause = %v", this.Content(), this.Cause())
}

func (this *HttpInterceptorExceptionBase) Cause() error {
	return this.cause
}

func (this *HttpInterceptorExceptionBase) Status() (int, bool) {
	if this.statusCode > 0 {
		return this.statusCode, true
	} else {
		return -1, false
	}
}

func (this *HttpInterceptorExceptionBase) Content() interface{} {
	return this.content
}

func NewInterceptorException(cause error, statusCode int, content interface{}) HttpInterceptorException {
	return &HttpInterceptorExceptionBase{
		cause: cause,
		statusCode: statusCode,
		content: content,
	}
}