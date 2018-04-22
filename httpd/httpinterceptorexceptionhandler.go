package httpd

import (
	"net/http"
	"encoding/json"
	"runtime/debug"
	"strings"
	"log"
	"fmt"
)

type HttpInterceptorExceptionHandlerBase struct {

}

func (this *HttpInterceptorExceptionHandlerBase) HandleException(request *http.Request, response http.ResponseWriter,
	context HttpRequestContext, exception interface{}) {
	this.PrintStackTrace(request, response, context, exception)
	if err, ok := exception.(HttpInterceptorException); ok {
		this.HandleInterceptorException(request, response, context, err)
	} else {
		this.HandleOtherException(request, response, context, exception)
	}
}

func (this *HttpInterceptorExceptionHandlerBase) PrintStackTrace(request *http.Request, response http.ResponseWriter,
	context HttpRequestContext, exception interface{}) {
	builder := strings.Builder{}
	builder.WriteString("Exception detected:\n")
	builder.WriteString(fmt.Sprintf("URL: %v\n", request.URL.String()))
	builder.WriteString(fmt.Sprintf("Method: %v\n", request.Method))
	if err, ok := exception.(error); ok {
		builder.WriteString(fmt.Sprintf("Error: %v\n", err.Error()))
	} else {
		builder.WriteString(fmt.Sprintf("Error: %v\n", exception))
	}
	builder.WriteString("Call Stack: \n")
	builder.Write(debug.Stack())
	log.Println(builder.String())
}

func (this *HttpInterceptorExceptionHandlerBase) HandleInterceptorException(request *http.Request, response http.ResponseWriter,
	context HttpRequestContext, exception HttpInterceptorException) {
	outputMap := make(map[string]interface{})

	if statusCode, ok := exception.Status(); ok {
		response.WriteHeader(statusCode)
		outputMap["statusCode"] = statusCode
	}

	if content := exception.Content(); content != nil {
		outputMap["content"] = content
	}

	if cause := exception.Cause(); cause != nil {
		outputMap["cause"] = cause
	}

	stackBytes := debug.Stack()
	stackText := string(stackBytes)
	outputMap["stack"] = strings.Split(stackText, "\n")

	encoder := json.NewEncoder(response)
	encoder.Encode(outputMap)
}

func (this *HttpInterceptorExceptionHandlerBase) HandleOtherException(request *http.Request, response http.ResponseWriter,
	context HttpRequestContext, exception interface{}) {
	outputMap := make(map[string]interface{})

	outputMap["exception"] = exception

	stackBytes := debug.Stack()
	stackText := string(stackBytes)
	outputMap["stack"] = strings.Split(stackText, "\n")

	encoder := json.NewEncoder(response)
	encoder.Encode(outputMap)
}