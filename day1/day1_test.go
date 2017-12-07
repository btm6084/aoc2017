package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSumByNext(t *testing.T) {
	assert.Equal(t, 3, sumByNext([]byte("1122")))
	assert.Equal(t, 4, sumByNext([]byte("1111")))
	assert.Equal(t, 0, sumByNext([]byte("1234")))
	assert.Equal(t, 9, sumByNext([]byte("91212129")))
}

func TestSumByHalf(t *testing.T) {
	assert.Equal(t, 6, sumByHalf([]byte("1212")))
	assert.Equal(t, 0, sumByHalf([]byte("1221")))
	assert.Equal(t, 4, sumByHalf([]byte("123425")))
	assert.Equal(t, 12, sumByHalf([]byte("123123")))
	assert.Equal(t, 4, sumByHalf([]byte("12131415")))
}
