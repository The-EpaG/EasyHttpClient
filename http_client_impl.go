package easyhttpclient

import (
	"bytes"
	"io"
	"net/http"
	"net/url"

	"github.com/The-EpaG/EasyHttpClient/transport"
)

type HttpClientImpl struct {
	BaseUrl    string
	HttpClient *http.Client
	Headers    *http.Header
}

type HttpClientOptions struct {
	RequestInterceptors  []transport.EasyHttpClientRequestInterceptor
	ResponseInterceptors []transport.EasyHttpClientResponseInterceptor
	BaseUrl              string
	Headers              *http.Header
}

func NewHttpClient(opts *HttpClientOptions) *HttpClientImpl {

	return &HttpClientImpl{
		HttpClient: &http.Client{
			Transport: &transport.InterceptorManagerTransport{
				Proxied:              http.DefaultTransport,
				RequestInterceptors:  opts.RequestInterceptors,
				ResponseInterceptors: opts.ResponseInterceptors,
			},
		},
		Headers: func() *http.Header {
			if opts.Headers == nil {
				return &http.Header{}
			}
			return opts.Headers
		}(),
		BaseUrl: opts.BaseUrl,
	}
}

func (c *HttpClientImpl) getFullUrl(path string) *url.URL {
	parsedURL, err := url.Parse(c.BaseUrl + path)
	if err != nil {
		
		return nil
	}
	return parsedURL
}

func (c *HttpClientImpl) createHeader(body []byte, contentType ...string) http.Header {
	header := c.Headers.Clone()
	if len(contentType) > 0 {
		ct := contentType[0]
		header.Set("Content-Type", ct)
	} else {
		header.Set("Content-Type", http.DetectContentType(body))
	}
	return header
}

func (c *HttpClientImpl) Get(path string) ([]byte, error) {
	req := &http.Request{
		Method: http.MethodGet,
		URL:    c.getFullUrl(path),
		Header: *c.Headers,
	}

	response, err := c.HttpClient.Do(req)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}

func (c *HttpClientImpl) Post(path string, body []byte, contentType ...string) ([]byte, error) {
	response, err := c.HttpClient.Do(&http.Request{
		Method: http.MethodPost,
		URL:    c.getFullUrl(path),
		Body:   io.NopCloser(bytes.NewReader(body)),
		Header: c.createHeader(body, contentType...),
	})

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return resBody, nil
}

func (c *HttpClientImpl) Patch(path string, body []byte, contentType ...string) ([]byte, error) {
	response, err := c.HttpClient.Do(&http.Request{
		Method: http.MethodPatch,
		URL:    c.getFullUrl(path),
		Body:   io.NopCloser(bytes.NewReader(body)),
		Header: c.createHeader(body, contentType...),
	})

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	
	return resBody, nil
}

func (c *HttpClientImpl) Put(path string, body []byte, contentType ...string) ([]byte, error) {
	response, err := c.HttpClient.Do(&http.Request{
		Method: http.MethodPut,
		URL:    c.getFullUrl(path),
		Body:   io.NopCloser(bytes.NewReader(body)),
		Header: c.createHeader(body, contentType...),
	})

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	
	return resBody, nil
}

func (c *HttpClientImpl) Delete(path string) ([]byte, error) {
	response, err := c.HttpClient.Do(&http.Request{
		Method: http.MethodDelete,
		URL:    c.getFullUrl(path),
		Header: *c.Headers,
	})

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	
	return resBody, nil
}

func (c *HttpClientImpl) Options(path string) ([]byte, error) {
	response, err := c.HttpClient.Do(&http.Request{
		Method: http.MethodOptions,
		URL:    c.getFullUrl(path),
		Header: *c.Headers,
	})

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	resBody, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	
	return resBody, nil
}
