package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRunInstructionSet(t *testing.T) {
	raw := `
	b inc 5 if a > 1
	a inc 1 if b < 5
	c dec -10 if a >= 1
	c inc -20 if c == 10
`

	input := strings.Split(raw, "\n")
	for k, v := range input {
		input[k] = strings.Trim(v, "\t")
	}

	instructions := buildInstructions(input)
	r := NewRegisterBank()
	r.RunInstructionSet(instructions)

	assert.Equal(t, 1, r.Get("a"), "a")
	assert.Equal(t, 0, r.Get("b"), "b")
	assert.Equal(t, -10, r.Get("c"), "c")
}

func TestFindLargestValue(t *testing.T) {
	raw := `
	b inc 5 if a > 1
	a inc 1 if b < 5
	c dec -10 if a >= 1
	c inc -20 if c == 10
`

	input := strings.Split(raw, "\n")
	for k, v := range input {
		input[k] = strings.Trim(v, "\t")
	}

	instructions := buildInstructions(input)
	r := NewRegisterBank()
	r.RunInstructionSet(instructions)

	assert.Equal(t, 1, r.FindLargestValue())
}

func TestGetHistoricHigh(t *testing.T) {
	raw := `
	b inc 5 if a > 1
	a inc 1 if b < 5
	c dec -10 if a >= 1
	c inc -20 if c == 10
`

	input := strings.Split(raw, "\n")
	for k, v := range input {
		input[k] = strings.Trim(v, "\t")
	}

	instructions := buildInstructions(input)
	r := NewRegisterBank()
	r.RunInstructionSet(instructions)

	assert.Equal(t, 10, r.GetHistoricHigh())
}
