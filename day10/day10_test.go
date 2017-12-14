package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeList(t *testing.T) {
	testCases := []struct {
		name     string
		size     int
		expected []int
	}{
		{"Size 0", 0, []int{}},
		{"Size 1", 1, []int{0}},
		{"Size 3", 3, []int{0, 1, 2}},
		{"Size 5", 5, []int{0, 1, 2, 3, 4}},
		{"Size 8", 8, []int{0, 1, 2, 3, 4, 5, 6, 7}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			k := NewKnotter(tc.size)
			assert.Equal(t, tc.expected, k.list, tc.name)
		})
	}
}

func TestReverseSegment(t *testing.T) {
	testCases := []struct {
		name     string
		current  int
		length   int
		list     []int
		expected []int
	}{
		{"Run 1", 3, 5, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{1, 2, 3, 8, 7, 6, 5, 4, 9, 10}},
		{"Run 2", 0, 10, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1}},
		{"Run 3", 9, 2, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{10, 2, 3, 4, 5, 6, 7, 8, 9, 1}},
		{"Run 4", 8, 3, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{9, 2, 3, 4, 5, 6, 7, 8, 1, 10}},
		{"Run 5", 8, 6, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{2, 1, 10, 9, 5, 6, 7, 8, 4, 3}},
		{"Run 6", 9, 10, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{8, 7, 6, 5, 4, 3, 2, 1, 10, 9}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			k := NewKnotter(len(tc.list))
			k.list = tc.list
			k.current = tc.current
			k.reverseSegment(tc.length)

			assert.Equal(t, tc.expected, k.list, tc.name)
		})
	}
}

func TestStep(t *testing.T) {

	k := NewKnotter(5)

	testCases := []struct {
		name    string
		length  int
		list    []int
		current int
		skip    int
	}{
		{"Step 1", 3, []int{2, 1, 0, 3, 4}, 3, 1},
		{"Step 2", 4, []int{4, 3, 0, 1, 2}, 3, 2},
		{"Step 3", 1, []int{4, 3, 0, 1, 2}, 1, 3},
		{"Step 4", 5, []int{3, 4, 2, 1, 0}, 4, 4},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			k.Step(tc.length)
			assert.Equal(t, tc.list, k.list, tc.name+" List")
			assert.Equal(t, tc.current, k.current, tc.name+" Current")
			assert.Equal(t, tc.skip, k.skip, tc.name+" Skip Size")
		})
	}
}

func TestStepsToASCII(t *testing.T) {
	testCases := []struct {
		name   string
		input  []byte
		output []int
	}{
		{"Run 1", []byte("1,2,3"), []int{49, 44, 50, 44, 51, 17, 31, 73, 47, 23}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.output, stepsToASCII(tc.input), tc.name)
		})
	}
}

func TestStepsToInt(t *testing.T) {
	testCases := []struct {
		name   string
		input  []byte
		output []int
	}{
		{"Run 1", []byte("1,2,3"), []int{1, 2, 3}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.output, stepsToInt(tc.input), tc.name)
		})
	}
}

func TestDenseHash(t *testing.T) {

	testCases := []struct {
		name   string
		input  []int
		output []int
	}{
		{"Run 1", []int{65, 27, 9, 1, 4, 3, 40, 50, 91, 7, 6, 0, 2, 5, 68, 22}, []int{64}},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			k := NewKnotter(256)
			k.list = tc.input

			dense := k.GetDenseHash()

			assert.Equal(t, tc.output, dense, tc.name)
		})
	}
}

func TestSimpleKnotHash(t *testing.T) {

	testCases := []struct {
		name   string
		input  []int
		output string
	}{
		{"Run 1", []int{64, 7, 255}, "4007ff"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			k := NewKnotter(256)
			assert.Equal(t, tc.output, k.GetKnotHash(tc.input), tc.name)
		})
	}
}

func TestFullKnotHash(t *testing.T) {

	testCases := []struct {
		name   string
		input  []byte
		output string
	}{
		{"Run 1: ", []byte(""), "a2582a3a0e66e6e86e3812dcb672a272"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {

			k := NewKnotter(256)
			steps := stepsToASCII(tc.input)

			k.GenerateSparseHash(steps)
			assert.Equal(t, 256, len(k.list), tc.name+"List Length")

			sparseExpected := []int{38, 171, 116, 63, 70, 137, 168, 29, 198, 55, 160, 15, 34, 95, 58, 7, 188, 189, 238, 141, 30, 31, 124, 241, 20, 1, 244, 203, 234, 73, 236, 211, 122, 197, 94, 227, 142, 57, 72, 239, 54, 81, 154, 217, 10, 13, 186, 161, 6, 17, 128, 105, 106, 69, 44, 51, 248, 23, 136, 173, 52, 39, 40, 5, 254, 195, 64, 187, 192, 37, 230, 153, 56, 177, 84, 147, 96, 249, 252, 121, 166, 143, 62, 169, 90, 99, 196, 155, 132, 159, 162, 229, 76, 117, 164, 127, 150, 21, 88, 27, 242, 67, 114, 115, 226, 191, 190, 53, 2, 65, 206, 205, 24, 251, 14, 75, 74, 247, 80, 11, 50, 181, 46, 101, 100, 179, 48, 131, 32, 97, 102, 201, 170, 93, 104, 103, 182, 125, 12, 43, 220, 113, 158, 167, 68, 47, 66, 33, 112, 135, 194, 185, 218, 219, 8, 245, 130, 253, 204, 243, 202, 109, 92, 209, 156, 133, 250, 107, 4, 183, 60, 215, 172, 231, 240, 83, 98, 193, 82, 139, 210, 91, 146, 85, 184, 163, 140, 145, 178, 35, 232, 151, 214, 213, 200, 199, 18, 221, 212, 9, 152, 123, 78, 3, 228, 25, 26, 225, 0, 61, 138, 255, 222, 233, 110, 129, 208, 207, 176, 235, 108, 77, 148, 19, 180, 79, 28, 149, 224, 237, 86, 157, 216, 111, 22, 89, 16, 41, 144, 71, 134, 59, 246, 165, 174, 223, 118, 119, 36, 175, 126, 87, 120, 45, 42, 49}
			assert.Equal(t, sparseExpected, k.list, tc.name+"Sparse List")

			dense := k.GetDenseHash()
			assert.Equal(t, 16, len(dense), tc.name+"Dense Length")

			denseExpected := []int{162, 88, 42, 58, 14, 102, 230, 232, 110, 56, 18, 220, 182, 114, 162, 114}
			assert.Equal(t, denseExpected, dense, tc.name+"Dense List")

			knot := k.GetKnotHash(dense)
			assert.Equal(t, 32, len(knot), tc.name+"Knot Length")

			assert.Equal(t, tc.output, knot, tc.name+"Knot Hash")
		})
	}
}
