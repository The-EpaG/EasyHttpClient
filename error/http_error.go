package httpError

type HttpError struct {
	error
	StatusCode int
	Status     string
}