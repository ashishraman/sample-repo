package main

import (
	"testing"

	"github.com/stretchr/testify/master/assert"
)

func TestDisplayHello(t *testing.T) {
	str := displayHello()
	assert.Equal(t, str, "Hello Go", "Expected Hello World")
}
