package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// Stream  holds metadata about a given stream
type Stream struct {
	Groups     map[int]int
	Score      int
	NumGroups  int
	NumGarbage int
}

func main() {
	s := processStream(loadFromFile("./input.txt"))

	fmt.Printf("Number of Groups: %d\n", s.NumGroups)
	fmt.Printf("Score: %d\n", s.Score)
	fmt.Printf("Garbage: %d\n", s.NumGarbage)
}

func processStream(input string) Stream {
	var s Stream

	s.Groups, s.NumGarbage = parseInput(input)
	s.Score = getScore(s.Groups)
	s.NumGroups = getGroupCount(s.Groups)

	return s
}

func getScore(groups map[int]int) int {
	score := 0
	for d, c := range groups {
		score += d * c
	}

	return score
}

func getGroupCount(groups map[int]int) int {
	count := 0
	for _, g := range groups {
		count += g
	}

	return count
}

// parseInput returns group inclusion data and garbage data
func parseInput(input string) (map[int]int, int) {
	groups := make(map[int]int)
	open := make(map[int]bool)

	stream := strings.Split(input, "")

	depth := 0
	maxDepth := 0
	garbage := false
	consume := false
	garbageCount := 0

	for _, s := range stream {
		switch true {
		case consume:
			consume = false
		case s == "!":
			consume = true
		case garbage && s == ">":
			garbage = false
		case garbage:
			garbageCount++
		case s == "<":
			garbage = true
		case s == "{":
			depth++
			open[depth] = true
			if depth > maxDepth {
				maxDepth = depth
			}
		case s == "}":
			open[depth] = false
			groups[depth]++
			depth--
		}
	}

	return groups, garbageCount
}

func loadFromFile(filename string) string {
	bytes, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(bytes) <= 0 {
		fmt.Println("No Input")
		os.Exit(1)
	}

	return string(bytes)
}
