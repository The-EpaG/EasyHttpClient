package interceptor

import (
	"bytes"
	"io"
	"net/http"

	"github.com/The-EpaG/EasyHttpClient/error"
)

type ResponseErrorInterceptor struct{}

func (i *ResponseErrorInterceptor)Intercept(res *http.Response) (*http.Response, error){
	if res.StatusCode >= http.StatusBadRequest {
		return nil, &httpError.HttpError{
			StatusCode: res.StatusCode,
			Status:     res.Status,
		}
	}

	if res.Body != nil {
		respBody, err := io.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		res.Body = io.NopCloser(bytes.NewBuffer(respBody)) //Set the body to the new buffer
	}

	return res, nil
}

func NewResponseErrorInterceptor() *ResponseErrorInterceptor {
	return &ResponseErrorInterceptor{}
}