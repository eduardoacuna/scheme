package main

import (
	"bufio"
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestREPL(t *testing.T) {
	input := bufio.NewReader(strings.NewReader("hello"))
	output := repl(t, input)
	assert.Equal(t, "hello", output, "should be equal")
}

func repl(t *testing.T, iport *bufio.Reader) string {
	obuff := bytes.NewBuffer(nil)
	oport := bufio.NewWriter(obuff)

	inputData, err := read(iport)
	assert.NoError(t, err, "it shouldn't be an error")

	outputData, err := eval(inputData)
	assert.NoError(t, err, "it shouldn't be an error")

	err = print(outputData, oport)
	assert.NoError(t, err, "it shouldn't be an error")

	return obuff.String()
}
