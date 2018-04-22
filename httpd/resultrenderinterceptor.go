package httpd

import (
	"net/http"
	"encoding/json"
)

type ResultRenderInterceptor struct {

}

func (this *ResultRenderInterceptor) Initialize(PropertyBag) {
	// foo
}

func (this *ResultRenderInterceptor) Destroy() {
	// foo
}

func (this *ResultRenderInterceptor) Intercept(chain HttpInterceptorChain, request *http.Request,
	response http.ResponseWriter, context HttpRequestContext) {
	if result, ok := context.GetInterface(HttpRequestResultType); ok {
		this.Render(chain, request, response, context, result)
	}
	chain.DoChain(request, response, context)
}

func (this *ResultRenderInterceptor) Render(chain HttpInterceptorChain, request *http.Request,
	response http.ResponseWriter, context HttpRequestContext, result interface{}) {
	encoder := json.NewEncoder(response)
	encoder.Encode(result)
}
