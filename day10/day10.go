package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cast"
)

// Knotter is the data and operations for producing a knot hash
type Knotter struct {
	list    []int
	current int
	skip    int
}

// Step performs a single step of a given length.
func (k *Knotter) Step(length int) {
	if length > len(k.list) {
		return
	}

	k.reverseSegment(length)

	// This will become the new current after we reverse this segment
	newCurrent := (k.current + length + k.skip) % len(k.list)

	k.current = newCurrent
	k.skip++
}

func (k *Knotter) reverseSegment(length int) {
	if length > len(k.list) {
		return
	}

	first := k.current
	last := (k.current + length - 1) % len(k.list)

	for i := 0; i < length/2; i++ {
		a := k.list[first]
		k.list[first] = k.list[last]
		k.list[last] = a

		first = (first + 1) % len(k.list)

		last--
		if last < 0 {
			last = len(k.list) - 1
		}
	}
}

// GenerateSparseHash generates a sparse hash
func (k *Knotter) GenerateSparseHash(steps []int) {
	for i := 0; i < 64; i++ {
		for _, s := range steps {
			k.Step(s)
		}
	}
}

// GetDenseHash builds a dense hash from a given list
func (k *Knotter) GetDenseHash() []int {
	var denseHash []int

	for i := 0; i < len(k.list)/16; i++ {
		base := i * 16
		elm := k.list[base]

		for j := 1; j < 16; j++ {
			elm = elm ^ k.list[base+j]
		}

		denseHash = append(denseHash, elm)
	}

	return denseHash
}

// GetKnotHash turns a denseHash into a Knot Hash
func (k *Knotter) GetKnotHash(denseHash []int) string {
	hash := ""

	for _, h := range denseHash {
		hash = fmt.Sprintf("%s%02x", hash, h)
	}

	return hash

}

// NewKnotter returns a newly initialized Knotter
func NewKnotter(size int) *Knotter {
	return &Knotter{makeList(size), 0, 0}
}

func makeList(size int) []int {
	list := make([]int, size)

	for i := 0; i < size; i++ {
		list[i] = i
	}

	return list
}

func main() {
	steps := stepsToInt(loadFromFile("./input.txt"))
	k := NewKnotter(256)

	for _, s := range steps {
		k.Step(s)
	}

	checksum := k.list[0] * k.list[1]

	fmt.Printf("Part 1: Checksum: %d\n\n", checksum)

	k = NewKnotter(256)
	steps = stepsToASCII(loadFromFile("./input.txt"))

	k.GenerateSparseHash(steps)
	dense := k.GetDenseHash()
	knot := k.GetKnotHash(dense)

	fmt.Printf("Part 2: Knot Hash: %s\n\n", knot)
}

func loadFromFile(filename string) []byte {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(bytes) <= 0 {
		fmt.Println("No Input")
		os.Exit(1)
	}

	return bytes
}

func stepsToInt(bytes []byte) []int {
	var steps []int

	for _, s := range strings.Split(string(bytes), ",") {
		steps = append(steps, cast.ToInt(s))
	}

	return steps
}

func stepsToASCII(bytes []byte) []int {
	var steps []int

	chars := string(bytes)

	for i := 0; i < len(bytes); i++ {
		steps = append(steps, cast.ToInt(chars[i]))
	}

	steps = append(steps, []int{17, 31, 73, 47, 23}...)

	return steps
}
