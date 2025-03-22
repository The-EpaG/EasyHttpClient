package main

import (
	"fmt"

	easyhttpclient "github.com/The-EpaG/EasyHttpClient"
)

func main() {
	httpClient := easyhttpclient.NewHttpClient(&easyhttpclient.HttpClientOptions{BaseUrl: "http://example.com"})

	body, err := httpClient.Get("/")
	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))
}