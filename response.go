package ape

import (
	"net/http"
)

type Response struct {
	StatusCode int
	Header     http.Header
	Body       []byte
}

func newResponse() *Response {
	return &Response{
		StatusCode: 0,
		Header:     make(http.Header),
	}
}
