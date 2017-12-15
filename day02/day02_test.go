package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalcByRow(t *testing.T) {
	low, high := getHighLow([]int{})
	assert.Equal(t, 0, low)
	assert.Equal(t, 0, high)

	low, high = getHighLow([]int{1, 2, 3, 4})
	assert.Equal(t, 1, low)
	assert.Equal(t, 4, high)

	low, high = getHighLow([]int{7})
	assert.Equal(t, 7, low)
	assert.Equal(t, 7, high)

	low, high = getHighLow([]int{7, 43243, -2932734, -1, 12})
	assert.Equal(t, -2932734, low)
	assert.Equal(t, 43243, high)
}

func TestGetChecksum(t *testing.T) {
	sheet := [][]int{
		[]int{5, 1, 9, 5},
		[]int{7, 5, 3},
		[]int{2, 4, 6, 8},
	}

	assert.Equal(t, 18, getChecksum(sheet))
}

func TestCanDivide(t *testing.T) {
	assert.Equal(t, true, canDivide(9, 3), "9, 3")
	assert.Equal(t, true, canDivide(3, 9), "3, 9")
	assert.Equal(t, true, canDivide(178, 2), "178, 2")
	assert.Equal(t, true, canDivide(178, 178), "178, 178")
	assert.Equal(t, true, canDivide(2, 178), "2, 178")

	assert.Equal(t, false, canDivide(9, 4), "9, 4")
	assert.Equal(t, false, canDivide(3, 0), "3, 0")
	assert.Equal(t, false, canDivide(178, 3), "178, 3")
	assert.Equal(t, false, canDivide(178, 7), "178, 7")
	assert.Equal(t, false, canDivide(2, 163), "2, 163")
	assert.Equal(t, false, canDivide(0, 163), "0, 163")
}

func TestGetDivisionValue(t *testing.T) {
	assert.Equal(t, 2, getDivisionValue([]int{1, 2, 3, 4}))
	assert.Equal(t, 4, getDivisionValue([]int{5, 9, 2, 8}))
	assert.Equal(t, 3, getDivisionValue([]int{9, 4, 7, 3}))
	assert.Equal(t, 2, getDivisionValue([]int{3, 8, 6, 5}))
}

func TestGetDivisionChecksum(t *testing.T) {
	sheet := [][]int{
		[]int{5, 9, 2, 8},
		[]int{9, 4, 7, 3},
		[]int{3, 8, 6, 5},
	}

	assert.Equal(t, 9, getDivisionChecksum(sheet))
}
