package ape

type RequestPassedError struct{}

func (e *RequestPassedError) Error() string {
	return "Request is passed"
}

type RequestHaltedError struct{}

func (e *RequestHaltedError) Error() string {
	return "Request is halted"
}
