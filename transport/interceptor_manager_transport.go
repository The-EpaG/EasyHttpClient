package transport

import (
	"net/http"
)

type InterceptorManagerTransport struct {
	Proxied              http.RoundTripper
	RequestInterceptors  []EasyHttpClientRequestInterceptor
	ResponseInterceptors []EasyHttpClientResponseInterceptor
}

func (lrt *InterceptorManagerTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	var err error
	for _, interceptor := range lrt.RequestInterceptors {
		req, err = interceptor.Intercept(req)
		if err != nil {
			return nil, err
		}
	}

	resp, err := lrt.Proxied.RoundTrip(req)

	if err != nil {
		return nil, err
	}

	for _, interceptor := range lrt.ResponseInterceptors {
		resp, err = interceptor.Intercept(resp)
		if err != nil {
			return nil, err
		}
	}

	return resp, nil
}
