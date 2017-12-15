package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetGroupCount(t *testing.T) {

	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{"Run 1", "{}", 1},
		{"Run 2", "{{{}}}", 3},
		{"Run 3", "{{},{}}", 3},
		{"Run 4", "{{{},{},{{}}}}", 6},
		{"Run 5", "{<{},{},{{}}>}", 1},
		{"Run 6", "{<a>,<a>,<a>,<a>}", 1},
		{"Run 7", "{{<a>},{<a>},{<a>},{<a>}}", 5},
		{"Run 8", "{{<!>},{<!>},{<!>},{<a>}}", 2},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := processStream(tc.input)
			assert.Equal(t, tc.expected, s.NumGroups, tc.name)
		})
	}
}

func TestGetScore(t *testing.T) {

	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{"Run 1", "{}", 1},
		{"Run 2", "{{{}}}", 6},
		{"Run 3", "{{},{}}", 5},
		{"Run 4", "{{{},{},{{}}}}", 16},
		{"Run 5", "{<a>,<a>,<a>,<a>}", 1},
		{"Run 6", "{{<ab>},{<ab>},{<ab>},{<ab>}}", 9},
		{"Run 7", "{{<!!>},{<!!>},{<!!>},{<!!>}}", 9},
		{"Run 8", "{{<a!>},{<a!>},{<a!>},{<ab>}}", 3},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := processStream(tc.input)
			assert.Equal(t, tc.expected, s.Score, tc.name)
		})
	}
}

func TestGarbageCount(t *testing.T) {

	testCases := []struct {
		name     string
		input    string
		expected int
	}{
		{"Run 1", "<>", 0},
		{"Run 2", "<random characters>", 17},
		{"Run 3", "<<<<>", 3},
		{"Run 4", "<{!>}>", 2},
		{"Run 5", "<!!>", 0},
		{"Run 6", "<!!!>>", 0},
		{"Run 7", `<{o"i!a,<{i<a>`, 10},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			s := processStream(tc.input)
			assert.Equal(t, tc.expected, s.NumGarbage, tc.name)
		})
	}
}
