package ape

import (
	"net/http"
)

type Response struct {
	http.ResponseWriter

	StatusCode int
	Body       []byte
	written    bool
}

func newResponse(w http.ResponseWriter) *Response {
	return &Response{
		ResponseWriter: w,
		StatusCode:     0,
		written:        false,
	}
}

func (res *Response) Write(body []byte) (int, error) {
	if !res.written {
		res.written = true
		return res.ResponseWriter.Write(body)
	}
	return 0, nil
}

func (res *Response) Halt(statusCode int) {
	res.StatusCode = statusCode
	panic(&RequestHaltedError{})
}

func (res *Response) Pass() (Any, error) {
	return nil, &RequestPassedError{}
}
