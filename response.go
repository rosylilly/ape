package ape

type Response struct {
	StatusCode int
	Body       []byte
}

func newResponse() *Response {
	return &Response{
		StatusCode: 0,
	}
}
