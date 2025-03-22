package easyhttpclient


type HttpClient interface {
	Get(path string) ([]byte, error)
	Post(path string, body []byte, contentType ...string) ([]byte, error)
	Put(path string, body []byte, contentType ...string) ([]byte, error)
	Delete(path string) ([]byte, error)
	Patch(path string, body []byte, contentType ...string) ([]byte, error)
	Options(path string) ([]byte, error)
}
