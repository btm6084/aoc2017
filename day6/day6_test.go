package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindFullest(t *testing.T) {

	banks := []int{0, 2, 7, 0}
	assert.Equal(t, 2, findFullest(banks), "1")

	banks = []int{0, 5, 10, 0, 11, 14, 13, 4, 11, 8, 8, 7, 1, 4, 12, 11}
	assert.Equal(t, 5, findFullest(banks), "2")
}

func TestRedistribute(t *testing.T) {
	banks := []int{0, 2, 7, 0}
	redistribute(banks)
	assert.Equal(t, []int{2, 4, 1, 2}, banks, "1")

	banks = []int{2, 4, 1, 2}
	redistribute(banks)
	assert.Equal(t, []int{3, 1, 2, 3}, banks, "2")

	banks = []int{3, 1, 2, 3}
	redistribute(banks)
	assert.Equal(t, []int{0, 2, 3, 4}, banks, "3")

	banks = []int{0, 2, 3, 4}
	redistribute(banks)
	assert.Equal(t, []int{1, 3, 4, 1}, banks, "4")

	banks = []int{1, 3, 4, 1}
	redistribute(banks)
	assert.Equal(t, []int{2, 4, 1, 2}, banks, "5")
}

func TestFindRepeat(t *testing.T) {
	banks := []int{0, 2, 7, 0}
	assert.Equal(t, 5, findRepeat(banks), "1")

	banks = []int{2, 1, 4, 2}
	assert.Equal(t, 5, findRepeat(banks), "2")
}

func TestFindLoopSize(t *testing.T) {
	banks := []int{0, 2, 7, 0}
	assert.Equal(t, 4, findLoopSize(banks), "1")

	banks = []int{2, 1, 4, 2}
	assert.Equal(t, 4, findLoopSize(banks), "2")
}
