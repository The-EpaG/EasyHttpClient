package transport

import "net/http"

type EasyHttpClientRequestInterceptor interface {
	Intercept(req *http.Request) (*http.Request, error)
}

type EasyHttpClientResponseInterceptor interface {
	Intercept(res *http.Response) (*http.Response, error)
}