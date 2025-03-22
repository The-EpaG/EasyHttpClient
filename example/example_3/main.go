package main

import (
	"fmt"

	easyhttpclient "github.com/The-EpaG/EasyHttpClient"
	"github.com/The-EpaG/EasyHttpClient/interceptor"
	"github.com/The-EpaG/EasyHttpClient/transport"
)

func main() {
	httpClient := easyhttpclient.NewHttpClient(&easyhttpclient.HttpClientOptions{
		BaseUrl: "http://example.com",
		ResponseInterceptors: []transport.EasyHttpClientResponseInterceptor{
			interceptor.NewResponseErrorInterceptor(),
		}})

	body, err := httpClient.Get("/Not-Existing-Path")
	if err != nil {
		fmt.Println("Failed successfully: ", err)
		panic(err)
	}

	fmt.Println(string(body))
}