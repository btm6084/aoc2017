package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetSqSpiralCoords(t *testing.T) {
	var pos Point

	pos = getCoords(1)
	assert.Equal(t, 0, pos.X, "1x")
	assert.Equal(t, 0, pos.Y, "1y")

	pos = getCoords(2)
	assert.Equal(t, 1, pos.X, "2x")
	assert.Equal(t, 0, pos.Y, "2y")

	pos = getCoords(13)
	assert.Equal(t, 2, pos.X, "13x")
	assert.Equal(t, 2, pos.Y, "13y")

	pos = getCoords(19)
	assert.Equal(t, -2, pos.X, "19x")
	assert.Equal(t, 0, pos.Y, "19y")

	pos = getCoords(22)
	assert.Equal(t, -1, pos.X, "22x")
	assert.Equal(t, -2, pos.Y, "22y")

	pos = getCoords(325489)
	assert.Equal(t, -267, pos.X, "325489x")
	assert.Equal(t, -285, pos.Y, "325489y")
}

func TestGetTaxiDist(t *testing.T) {
	pos := getCoords(1)
	dist := getTaxiDist(Point{0, 0}, pos)
	assert.Equal(t, 0, dist, "1")

	pos = getCoords(12)
	dist = getTaxiDist(Point{0, 0}, pos)
	assert.Equal(t, 3, dist, "12")

	pos = getCoords(23)
	dist = getTaxiDist(Point{0, 0}, pos)
	assert.Equal(t, 2, dist, "23")

	pos = getCoords(1024)
	dist = getTaxiDist(Point{0, 0}, pos)
	assert.Equal(t, 31, dist, "1024")
}

func TestGetNearestValue(t *testing.T) {
	v := getFirstLargerThan(1)
	assert.Equal(t, 2, v, "1")

	v = getFirstLargerThan(2)
	assert.Equal(t, 4, v, "2")

	v = getFirstLargerThan(3)
	assert.Equal(t, 4, v, "3")

	v = getFirstLargerThan(4)
	assert.Equal(t, 5, v, "4")

	v = getFirstLargerThan(5)
	assert.Equal(t, 10, v, "5")

	v = getFirstLargerThan(6)
	assert.Equal(t, 10, v, "6")

	v = getFirstLargerThan(10)
	assert.Equal(t, 11, v, "10")

	v = getFirstLargerThan(11)
	assert.Equal(t, 23, v, "11")

	v = getFirstLargerThan(23)
	assert.Equal(t, 25, v, "23")

	v = getFirstLargerThan(26)
	assert.Equal(t, 54, v, "26")

	v = getFirstLargerThan(747)
	assert.Equal(t, 806, v, "747")
}
