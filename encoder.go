package ape

type Encoder interface {
	Encode(data interface{}) ([]byte, error)
}
