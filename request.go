package ape

import (
	"net/http"
)

type Request struct {
	*http.Request

	Verb        string
	Path        string
	Format      string
	RouteParams map[string]string
}

func newRequestFromHTTPRequest(req *http.Request) *Request {
	request := &Request{
		Request: req,
	}
	request.Verb = req.Method
	request.Path = req.URL.Path

	return request
}
