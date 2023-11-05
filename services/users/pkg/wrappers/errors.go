package wrappers

import "fmt"

type WrappedError string

var DoesNotExist WrappedError = "does not exist"
var	NotFound  WrappedError = "not found"
var AlreadyExists WrappedError = "already exists"
var NotValid WrappedError = "not valid"

func newErrWithString(s string, w WrappedError) error {
	return fmt.Errorf("%s %s", s, w)
}


func NewErrDoesNotExist(s string) error {
	return newErrWithString(s, DoesNotExist)
}

func NewErrNotFound(s string) error {
	return newErrWithString(s, NotFound)
}

func NewErrAlreadyExists(s string) error {
	return newErrWithString(s, AlreadyExists)
}

func NewErrNotValid(s string) error {
	return newErrWithString(s, NotValid)
}