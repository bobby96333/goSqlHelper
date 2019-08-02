package HelperError

type Error interface {
	error
	Code() uint
	IsEmpty() bool
}