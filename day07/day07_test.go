package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	raw := `
	pbga (66)
	xhth (57)
	ebii (61)
	havc (66)
	ktlj (57)
	fwft (72) -> ktlj, cntj, xhth
	qoyq (66)
	padx (45) -> pbga, havc, qoyq
	tknk (41) -> ugml, padx, fwft
	jptl (61)
	ugml (68) -> gyxo, ebii, jptl
	gyxo (61)
	cntj (57)
`

	input := strings.Split(raw, "\n")
	for k, v := range input {
		input[k] = strings.Trim(v, "\t")
	}

	var tree ProgramTree
	tree.Parse(input)

	assert.Equal(t, "tknk", tree.Root.Name)
	assert.Equal(t, 41, tree.Root.Weight)
	assert.Equal(t, 778, tree.Root.BranchWeight)

	assert.Equal(t, "fwft", tree.Root.Children[0].Name)
	assert.Equal(t, 72, tree.Root.Children[0].Weight)
	assert.Equal(t, 243, tree.Root.Children[0].BranchWeight)

	assert.Equal(t, "padx", tree.Root.Children[1].Name)
	assert.Equal(t, 45, tree.Root.Children[1].Weight)
	assert.Equal(t, 243, tree.Root.Children[1].BranchWeight)

	assert.Equal(t, "ugml", tree.Root.Children[2].Name)
	assert.Equal(t, 68, tree.Root.Children[2].Weight)
	assert.Equal(t, 251, tree.Root.Children[2].BranchWeight)
}
