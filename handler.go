package ape

type Marshalable interface{}

type Handler interface {
	Serve(*Request, *Response) (Marshalable, error)
}

type HandlerFunc func(*Request, *Response) (Marshalable, error)

func (f HandlerFunc) Serve(req *Request, res *Response) (Marshalable, error) {
	return f(req, res)
}
