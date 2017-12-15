package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValidSimple(t *testing.T) {
	v := "aa bb cc dd ee"
	assert.Equal(t, true, isValidSimple(v), v)

	v = "aa bb cc dd aaa"
	assert.Equal(t, true, isValidSimple(v), v)

	v = "aa bb cc dd aa"
	assert.Equal(t, false, isValidSimple(v), v)
}

func TestIsValidComplex(t *testing.T) {
	v := "abcde fghij"
	assert.Equal(t, true, isValidComplex(v), v)

	v = "a ab abc abd abf abj"
	assert.Equal(t, true, isValidComplex(v), v)

	v = "iiii oiii ooii oooi oooo"
	assert.Equal(t, true, isValidComplex(v), v)

	v = "abcde xyz ecdab"
	assert.Equal(t, false, isValidComplex(v), v)

	v = "oiii ioii iioi iiio"
	assert.Equal(t, false, isValidComplex(v), v)
}
