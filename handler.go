package ape

type Handler interface {
	Handle(*Request, *Response) (interface{}, error)
}

type HanderFunc func(*Request, *Response) (interface{}, error)

func (f HanderFunc) Handle(req *Request, res *Response) (interface{}, error) {
	return f(req, res)
}
