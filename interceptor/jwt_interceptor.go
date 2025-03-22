package interceptor

import (
	"net/http"
)

type JWTRequestInterceptor struct {
	Token *string
}

func (i *JWTRequestInterceptor) Intercept(req *http.Request) (*http.Request, error) {
	if i.Token != nil {
		req.Header.Add("Authorization", "Bearer "+*i.Token)
	}

	return req, nil
}

func NewJWTRequestInterceptor(token *string) *JWTRequestInterceptor {
	return &JWTRequestInterceptor{
		Token: token,
	}
}