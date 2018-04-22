package httpd

import (
	"net/http"
	"fmt"
)

type AuthorizationInterceptor struct {

}

func (this *AuthorizationInterceptor) Initialize(PropertyBag) {
	fmt.Println("AuthorizationInterceptor::Initialize")
}

func (this *AuthorizationInterceptor) Destroy() {
	fmt.Println("AuthorizationInterceptor::Destroy")
}

func (this *AuthorizationInterceptor) Intercept(chain HttpInterceptorChain, request *http.Request, response http.ResponseWriter, context HttpRequestContext) {
	fmt.Println("AuthorizationInterceptor::Intercept, begin")
	chain.DoChain(request, response, context)
	fmt.Println("AuthorizationInterceptor::Intercept, end")
}

