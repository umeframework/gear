package httpd

import (
	"net/http"
	"fmt"
)

type AuthenticationInterceptor struct {

}

func (this *AuthenticationInterceptor) Initialize(PropertyBag) {
	fmt.Println("AuthenticationInterceptor::Initialize")
}

func (this *AuthenticationInterceptor) Destroy() {
	fmt.Println("AuthenticationInterceptor::Destroy")
}

func (this *AuthenticationInterceptor) Intercept(chain HttpInterceptorChain, request *http.Request, response http.ResponseWriter, context HttpRequestContext) {
	fmt.Println("AuthenticationInterceptor::Intercept, begin")
	chain.DoChain(request, response, context)
	fmt.Println("AuthenticationInterceptor::Intercept, end")
}

