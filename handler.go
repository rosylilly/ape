package ape

type Any interface{}

type Handler interface {
	Serve(*Request, *Response) (Any, error)
}

type HandlerFunc func(*Request, *Response) (Any, error)

func (f HandlerFunc) Serve(req *Request, res *Response) (Any, error) {
	return f(req, res)
}
