package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindFirstExit(t *testing.T) {
	maze := []int{0, 3, 0, 1, -3}
	assert.Equal(t, 5, findFirstExit(maze))
}

func TestFindSecondExit(t *testing.T) {
	maze := []int{0, 3, 0, 1, -3}
	assert.Equal(t, 10, findSecondExit(maze))
}
