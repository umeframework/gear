package httpd

import (
	"net/http"
	"reflect"
	"strings"
	"encoding/json"
	"net/url"
	"errors"
	"bytes"
	"io"
	"fmt"
	"strconv"
)

var (
	servicePoints []ServicePoint
)

type combinedParamsType map[string]interface{}

func NewServicePoint(urlPattern string, methods []string, handler interface{}) {
	servicePoint := ServicePoint{
		urlPattern,
		methods,
		handler,
	}
	servicePoints = append(servicePoints, servicePoint)
}

type InvocationInterceptor struct {
	
}

func (this *InvocationInterceptor) Initialize(PropertyBag) {
	// foo
}

func (this InvocationInterceptor) Destroy() {
	// foo
}

func (this *InvocationInterceptor) Intercept(chain HttpInterceptorChain, request *http.Request, response http.ResponseWriter, context HttpRequestContext) {
	if servicePoint := this.FindServicePoint(request); servicePoint != nil {
		// Run services point handler
		if result, ok := this.Invocate(request, response, context, servicePoint); ok {
			// Save result to context (for succeeding interceptor)
			context.SetInterface(HttpRequestResultType, result)
		}
	}

	chain.DoChain(request, response, context)
}

func (this *InvocationInterceptor) FindServicePoint(request *http.Request) *ServicePoint {
	var ret *ServicePoint = nil

	for _, servicePoint := range servicePoints {
		if this.MatchServicePoint(request, &servicePoint) {
			ret = &servicePoint
			break
		}
	}

	return ret
}

func (this *InvocationInterceptor) MatchServicePoint(request *http.Request, servicePoint *ServicePoint) bool {
	// Match URL path only at present.
	// TODO: Extend to a full match check
	url := request.URL.Path
	if strings.ToLower(url) != strings.ToLower(servicePoint.UrlPattern) {
		return false
	}

	// Exit
	return true
}

func (this *InvocationInterceptor) Invocate(request *http.Request, response http.ResponseWriter,
	context HttpRequestContext, servicePoint *ServicePoint) (interface{}, bool) {
	var ret interface{} = nil
	var hasReturn bool = false
	handler := servicePoint.handler

	if httpHandler, ok := handler.(http.Handler); ok {
		httpHandler.ServeHTTP(response, request)
	} else if httpHandlerFunc, ok := handler.(http.HandlerFunc); ok {
		httpHandlerFunc(response, request)
	} else {
		handlerWrapper := reflect.ValueOf(handler)
		if handlerWrapper.Kind() != reflect.Func {
			// TODO: throws some exception here
		}

		hasReturn = true
		args := this.PrepareInvocationArgs(request, response, context, servicePoint, handlerWrapper)
		if outputs := handlerWrapper.Call(args); outputs != nil && len(outputs) > 0 {
			ret = outputs[0].Interface()
		}

	}

	// Exit
	return ret, hasReturn
}

func (this *InvocationInterceptor) PrepareInvocationArgs(request *http.Request, response http.ResponseWriter,
	context HttpRequestContext, servicePoint *ServicePoint, handler reflect.Value) []reflect.Value {
	var args []reflect.Value = nil

	params := this.PrepareParams(request, servicePoint)
	switch(request.Method) {
	case http.MethodGet:
		args = this.PrepareInvocationArgs_DTO(request, context, servicePoint, handler, params, nil, nil)
	default:
		requestBodyBuffer := this.BackupRequestBody(request)
		requestBody := this.GetRequestBody(requestBodyBuffer)
		requestBodyValue := reflect.ValueOf(requestBody)
		if requestBodyValue.Kind() == reflect.Slice || requestBodyValue.Kind() == reflect.Array {
			args = this.PrepareInvocationArgs_Array(request, context, servicePoint, handler, params, requestBody,
				requestBodyBuffer)
		} else {
			args = this.PrepareInvocationArgs_DTO(request, context, servicePoint, handler, params, requestBody,
				requestBodyBuffer)
		}
	}

	// Exit
	return args
}

func (this *InvocationInterceptor) PrepareInvocationArgs_DTO(request *http.Request,
	context HttpRequestContext, servicePoint *ServicePoint, handler reflect.Value,
	params combinedParamsType, requestBody interface{}, requestBodyBuffer *bytes.Buffer) []reflect.Value {
	// Prepare for each argument
	handlerType := handler.Type()
	argCount := handlerType.NumIn()
	args := make([]reflect.Value, argCount)

	// Set context for the 1st argument
	args[0] = reflect.ValueOf(context)

	// Convert 2nd argument as DTO
	if argCount != 2 {
		panic(errors.New("only 2 arguments are supported"))
	}
	argType := handlerType.In(1)
	if argType.Kind() == reflect.Map {
		// Combine params
		this.MergeRequestBody(params, requestBody)

		// Convert combined params to target
		args[1] = reflect.ValueOf(params)
	} else if argType.Kind() == reflect.Struct {
		// Create instance
		arg := reflect.New(argType)

		// Combine params (map) to arg
		value := arg.Elem()
		t := value.Type()
		fieldCount := value.NumField()
		for i := 0; i < fieldCount; i++ {
			fieldInfo := t.Field(i)
			fieldName := strings.ToUpper(fieldInfo.Name)
			if fieldParam, ok := params[fieldName]; ok {
				fieldValue := this.ConvertType(reflect.ValueOf(fieldParam), fieldInfo.Type)
				fmt.Println("fieldValue: ", fieldValue.Interface())
				value.Field(i).Set(fieldValue)
			}
		}

		// Combine from Json
		if requestBodyBuffer != nil {
			decoder := json.NewDecoder(bytes.NewReader(requestBodyBuffer.Bytes()))
			decoder.Decode(arg.Interface())
		}

		// Add to args array
		args[1] = arg.Elem()
	} else {
		panic(errors.New("unspported type for 2nd argument: " + argType.Name()))
	}

	// Exit
	return args
}

func (this *InvocationInterceptor) MergeRequestBody(params combinedParamsType, requestBody interface{}) {
	requestBodyValue := reflect.ValueOf(requestBody)
	if requestBodyValue.Kind() == reflect.Map {
		for _, key := range requestBodyValue.MapKeys() {
			if keyText, ok := key.Interface().(string); ok {
				value := requestBodyValue.MapIndex(key)
				params[keyText] = value.Interface()
			}
		}
	}
}

var (
	interfaceType = reflect.TypeOf((*interface{})(nil)).Elem()
)

func (this *InvocationInterceptor) PrepareInvocationArgs_Array(request *http.Request,
	context HttpRequestContext, servicePoint *ServicePoint, handler reflect.Value,
	params combinedParamsType, requestBody interface{}, requestBodyBuffer *bytes.Buffer) []reflect.Value {
	// Prepare for each argument
	handlerType := handler.Type()
	argCount := handlerType.NumIn()
	args := make([]reflect.Value, argCount)

	// Set context for the 1st argument
	args[0] = reflect.ValueOf(context)

	// Set 2nd argument for slice/array types
	if argCount == 2 {
		// Only one argument left, check what it is
		argType := handlerType.In(1)
		if argType.Kind() == reflect.Slice || argType.Kind() == reflect.Array {
			// Distinct []interface{} with []type
			elemType := argType.Elem()
			if elemType == interfaceType {
				// Use request body direct
				args[1] = reflect.ValueOf(requestBody)
			} else {
				// Re-parse from request body buffer
				decoder := json.NewDecoder(bytes.NewReader(requestBodyBuffer.Bytes()))
				list := reflect.New(argType)
				decoder.Decode(list.Interface())
				args[1] = list.Elem()
			}
			return args
		}
	}

	// Set other arguments as ordered ones
	requestBodyValue := reflect.ValueOf(requestBody)
	for i := 1; i < argCount; i++ {
		argType := handlerType.In(i)
		requestBodyItemValue := requestBodyValue.Index(i - 1).Elem()
		arg := this.ConvertType(requestBodyItemValue, argType)
		args[i] = arg
	}

	// Exit
	return args
}

func (this *InvocationInterceptor) ConvertType(value reflect.Value, t reflect.Type) (ret reflect.Value) {
	defer func() {
		if err := recover(); err != nil {
			ret = this.ConvertAsJson(value, t)
		}
	}()

	fmt.Println("value.type = ", value.Type())
	fmt.Println("t (target type) = ", t)

	// TODO: implement a conversion system
	if value.Type() == t {
		ret = value
	} else if value.Kind() == reflect.String {
		// Special handling for string conversion
		str := value.Interface().(string)
		ret = this.ConvertStringType(str, t)
	} else {
		ret = value.Convert(t)
	}
	return ret
}

func (this *InvocationInterceptor) ConvertStringType(str string, t reflect.Type) reflect.Value {
	var ret reflect.Value
	var value interface{} = nil
	var useValue = true
	switch(t.Kind()) {
	case reflect.Int:
		value, _ = strconv.Atoi(str)
	case reflect.Uint:
		value, _ = strconv.ParseUint(str, 10, 0)
	}
	if useValue {
		ret = reflect.ValueOf(value)
	}
	return ret
}

func (this *InvocationInterceptor) ConvertAsJson(value reflect.Value, t reflect.Type) (ret reflect.Value) {
	// Encode to json
	buffer := bytes.NewBuffer(nil)
	encoder := json.NewEncoder(buffer)
	encoder.Encode(value.Interface())

	// Decode as specified type
	temp := reflect.New(t)
	decoder := json.NewDecoder(bytes.NewReader(buffer.Bytes()))
	decoder.Decode(temp.Interface())

	// Exit
	ret = temp
	return ret
}

func (this *InvocationInterceptor) NormalizeParams(params url.Values) combinedParamsType {
	ret := make(combinedParamsType)
	for key, value := range params {
		key = strings.ToUpper(key)
		len := len(value)
		switch len {
		case 0:
			ret[key] = nil
		case 1:
			ret[key] = value[0]
		default:
			ret[key] = value
		}
	}
	return ret
}

func (this *InvocationInterceptor) PrepareParams(request *http.Request,
	servicePoint *ServicePoint) combinedParamsType {
	params := url.Values{}

	request.ParseForm()
	queryParams := this.GetQueryParams(request)
	//formParams := this.GetFormParams(request)
	postFormParams := this.GetPostFormParams(request)
	pathParams := this.GetPathParams(request, servicePoint)

	this.MergeParams2(pathParams, params)
	this.MergeParams(queryParams, params)
	//this.MergeParams(formParams, params)
	this.MergeParams(postFormParams, params)

	normalizedParams := this.NormalizeParams(params)
	return normalizedParams
}

func (this *InvocationInterceptor) MergeParams(source url.Values, target url.Values) {
	for key, value := range source {
		target[key] = value[:]
	}
}

func (this *InvocationInterceptor) MergeParams2(source map[string]string, target url.Values) {
	for key, value := range source {
		target[key] = []string{value}
	}
}

func (this *InvocationInterceptor) GetPathParams(request *http.Request, servicePoint *ServicePoint) HttpRequestPathParam {
	return HttpRequestPathParam{}
}

func (this *InvocationInterceptor) GetQueryParams(request *http.Request) url.Values {
	return request.URL.Query()
}

func (this *InvocationInterceptor) GetFormParams(request *http.Request) url.Values {
	return request.Form
}

func (this *InvocationInterceptor) GetPostFormParams(request *http.Request) url.Values {
	return request.PostForm
}

func (this *InvocationInterceptor) BackupRequestBody(request *http.Request) *bytes.Buffer {
	buffer := bytes.NewBuffer(nil)
	io.Copy(buffer, request.Body)

	// Exit
	return buffer
}

func (this *InvocationInterceptor) GetRequestBody(buffer *bytes.Buffer) interface{} {
	var requestBody interface{}
	reader := bytes.NewReader(buffer.Bytes())
	decoder := json.NewDecoder(reader)
	decoder.Decode(&requestBody)
	return requestBody
}

func (this *InvocationInterceptor) PrepareInvocationArg(request *http.Request, response http.ResponseWriter,
	context HttpRequestContext, servicePoint *ServicePoint, handler reflect.Value,
	params url.Values, requestBody interface{}, requestBodyBuffer *bytes.Buffer,
	argIndex int, argType reflect.Type, reservedParamsCount *int) reflect.Value {
	var arg reflect.Value

	if argIndex == 0 {
		arg = reflect.ValueOf(context)
	} else {
		requestBodyValue := reflect.ValueOf(requestBody)
		if requestBodyValue.Kind() != reflect.Slice && requestBodyValue.Kind() != reflect.Array {
			// throw some exception
			panic(errors.New("request body must be an array"))
		}

		arg = requestBodyValue.Index(argIndex - 1).Elem().Convert(argType)
	}

	// Exit
	return arg
}