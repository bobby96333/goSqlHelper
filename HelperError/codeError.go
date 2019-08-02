package HelperError

import "errors"

const(
	ErrorEmpty=iota
	ErrorParent

)



// return new instance of code
func New(code uint) Error {
	return &codeError{code,nil}
}
// return new instance of message
func NewString(msg string) Error {

	return &codeError{ErrorParent,errors.New(msg) }
}
// return new instance of parent code
func NewParent(err error) Error {
	return &codeError{ErrorParent,err}
}

type codeError struct {
	code uint
	err error
}

func (this *codeError) Error() string {
	switch this.code {
	case ErrorEmpty: return "no found data"
	case ErrorParent: return this.err.Error()
	default: panic("no found code:"+string(this.code))
	}
}

func (this *codeError) Code() uint {
	return this.code
}
func (this *codeError) IsEmpty() bool {
	return this.code==ErrorEmpty
}

