package main

import (
	"bytes"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test(t *testing.T) {
	input := io.NopCloser(bytes.NewReader([]byte(``)))

	v := InputTo{{ .Name }}(input)

	assert.Equal(t, "", v)
}