package ape

type ErrorHandler func(*Request, *Response, error)

var (
	defaultErrorHandler ErrorHandler = func(req *Request, res *Response, err error) {
		res.StatusCode = 500
	}
)
