package errors

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInterpreterError(t *testing.T) {
	err1 := NewError(UnexpectedError, "something went wrong", "foo:", 1, "bar:", 2)
	assert.Error(t, err1, "it should be an error")
	assert.Contains(t, err1.Error(), UnexpectedError, "it should contain the error type")
	assert.Contains(t, err1.Error(), "something went wrong", "it should contain the error description")
	assert.Contains(t, err1.Error(), "foo: 1", "it should contain the key value irritant")
	assert.Contains(t, err1.Error(), "bar: 2", "it should contain the key value irritant")
}
