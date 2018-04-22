package main

import (
	"net/http"
	"log"
	"github.com/umeframework/gear/httpd"
	_ "github.com/umeframework/gear/test/services"
)

func main() {
	//services.Service1Test()
	//return

	handler := makeHandler()
	server := http.Server{
		Addr: ":8090",
		Handler: handler,
	}
	log.Fatal(server.ListenAndServe())
}

func makeHandler() http.Handler {
	propertyBag := httpd.NewPropertyBag()
	handler := httpd.NewHttpHandler(
		[]interface{}{
			&httpd.AuthenticationInterceptor{},
			&httpd.AuthorizationInterceptor{},
			&httpd.InvocationInterceptor{},
			&httpd.ResultRenderInterceptor{},
		},
		&httpd.HttpInterceptorExceptionHandlerBase{},
		propertyBag)

	return handler
}