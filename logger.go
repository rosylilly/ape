package ape

type Logger interface {
	Printf(string, ...interface{})
}
