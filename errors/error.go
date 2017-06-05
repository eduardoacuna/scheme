package errors

import (
	"fmt"
	"runtime"
	"strings"
)

// ErrorName is a type string for the name of an error
type ErrorName string

const (
	// UnexpectedError is used when the error is unknown
	UnexpectedError ErrorName = "unexpected error"
	// TypeError is used when any given type isn't correct
	TypeError = "type error"
	// NilError is used when a reference is expected to be non nil
	NilError = "nil error"
	// ValueError is used when a value is expected to be different
	ValueError = "value error"
	// OutOfBoundsError is used when indexing an object outside of it's range
	OutOfBoundsError = "out of bounds error"
)

// InterpreterError is the error type for the implementation of scheme
type InterpreterError struct {
	Name        ErrorName
	Description string
	Irritants   []interface{}
	Stack       []byte
}

// Error makes a descriptive string from an InterpreterError
func (err *InterpreterError) Error() string {
	strs := make([]string, len(err.Irritants))
	for i, irr := range err.Irritants {
		strs[i] = fmt.Sprintf("%v", irr)
	}
	return fmt.Sprintf("%s (%s) %s", err.Name, strings.Join(strs, " "), err.Description)
}

// NewError is an InterpreterError constructor
func NewError(name ErrorName, description string, irritants ...interface{}) error {
	err := &InterpreterError{
		Name:        name,
		Description: description,
		Irritants:   irritants,
	}
	err.Stack = make([]byte, 800)
	runtime.Stack(err.Stack, false)
	return err
}
