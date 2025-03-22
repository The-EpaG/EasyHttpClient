package main

import (
	"fmt"

	easyhttpclient "github.com/The-EpaG/EasyHttpClient"
	"github.com/The-EpaG/EasyHttpClient/interceptor"
	"github.com/The-EpaG/EasyHttpClient/transport"
)

func main() {
	token := "my-jwt-token"
	httpClient := easyhttpclient.NewHttpClient(&easyhttpclient.HttpClientOptions{
		BaseUrl: "http://example.com",
		RequestInterceptors: []transport.EasyHttpClientRequestInterceptor{
			interceptor.NewJWTRequestInterceptor(&token),
		}})

	body, err := httpClient.Get("/")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}