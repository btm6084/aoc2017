package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cast"
)

// Program represents a node in a program tree
type Program struct {
	Name         string
	Parent       string
	Weight       int
	ChildSum     int
	ChildList    []string
	Children     []Program
	BranchWeight int
}

// ProgramTree stores the relationships between various programs
type ProgramTree struct {
	Root Program
}

// Load populates the program tree from an input file
func (t *ProgramTree) Load(inputFile string) {
	bytes, err := ioutil.ReadFile(inputFile)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(bytes) <= 0 {
		fmt.Println("No Input")
		os.Exit(1)
	}

	t.Parse(strings.Split(string(bytes), "\n"))
}

// Parse populates the tree based on the input
func (t *ProgramTree) Parse(lines []string) {
	var nodes []Program
	parents := make(map[string]string)

	//
	for _, line := range lines {
		var p Program
		pieces := strings.Split(line, " ")

		if len(pieces) < 2 || len(line) == 0 {
			continue
		}

		p.Name = pieces[0]
		p.Weight = cast.ToInt(strings.Trim(pieces[1], "()"))

		for i := 3; i < len(pieces); i++ {
			p.ChildList = append(p.ChildList, strings.Trim(pieces[i], ","))
			parents[strings.Trim(pieces[i], ",")] = p.Name
		}

		nodes = append(nodes, p)
	}

	// Apply Parental Markings to Children
	for cn, childNode := range nodes {
		if parentName, isset := parents[childNode.Name]; isset {
			nodes[cn].Parent = parentName
		}
	}

	for _, v := range nodes {
		if v.Parent == "" {
			t.Root = v
		}
	}

	t.Root.Children = t.buildChildren(t.Root, nodes)
	t.Root = t.weightBranches(t.Root)
}

// buildChildren requires that parse has been run successfully.
func (t *ProgramTree) buildChildren(root Program, nodes []Program) []Program {

	var children []Program

	if len(root.ChildList) <= 0 {
		return children
	}

	for _, v := range nodes {
		if v.Parent == root.Name {
			v.Children = t.buildChildren(v, nodes)
			children = append(children, v)
		}
	}

	return children

}

// weightBranches recursively traverses all children and finds the branch weight for each node.
func (t *ProgramTree) weightBranches(node Program) Program {
	if len(node.Children) == 0 {
		node.BranchWeight = node.Weight
		return node
	}

	weight := node.Weight
	childWeights := 0

	for i, n := range node.Children {
		node.Children[i] = t.weightBranches(n)
		weight += node.Children[i].BranchWeight

		childWeights += node.Children[i].BranchWeight
	}

	node.BranchWeight = weight

	return node
}

func (t *ProgramTree) findImbalancedNode(node Program) bool {
	if len(node.Children) == 0 {
		return true
	}

	sum := 0

	for _, n := range node.Children {
		sum += n.BranchWeight

		if !t.findImbalancedNode(n) {
			return false
		}
	}

	imbalanced := sum%len(node.Children) != 0

	if imbalanced {
		seen := make(map[int]int)
		index := make(map[int]int)
		for i, n := range node.Children {
			seen[n.BranchWeight]++
			index[n.BranchWeight] = i
		}

		diff := 0
		sum := 0
		var unbalanced Program

		for i, x := range seen {
			sum += i
			if x == 1 {
				diff = i * 2
				unbalanced = node.Children[index[i]]
			}
		}

		fmt.Printf("Part 2: `%s` should be %d\n", unbalanced.Name, unbalanced.Weight+(sum-diff))

		return false
	}

	return true
}

// Print out the structure for inspection
func (t *ProgramTree) Print(node Program, tabs string, depth int) {
	if depth == 0 {
		return
	}

	fmt.Printf("%s%s (%d) [%d]\n", tabs, node.Name, node.BranchWeight, node.Weight)

	for _, n := range node.Children {
		t.Print(n, tabs+"\t", depth-1)
	}
}

func main() {
	var t ProgramTree

	t.Load("./input.txt")
	fmt.Printf("Part 1: Root is `%s`\n", t.Root.Name)

	t.findImbalancedNode(t.Root)
}
