package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sort"
	"strings"
)

func main() {
	input := getInput()

	numValidSimple := 0
	numValidComplex := 0

	for _, l := range input {
		if isValidSimple(l) {
			numValidSimple++
		}

		if isValidComplex(l) {
			numValidComplex++
		}
	}

	fmt.Printf("Part 1: Number of Valid Simple Passphrases: %d\n", numValidSimple)
	fmt.Printf("Part 2: Number of Valid Complex Passphrases: %d\n", numValidComplex)
}

func getInput() []string {
	bytes, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(bytes) <= 0 {
		fmt.Println("No Input")
		os.Exit(1)
	}

	return strings.Split(string(bytes), "\n")
}

func isValidSimple(phrase string) bool {
	seen := make(map[string]int)

	words := strings.Split(phrase, " ")

	for _, w := range words {
		seen[w]++

		if seen[w] > 1 {
			return false
		}
	}

	return true
}

func isValidComplex(phrase string) bool {
	seen := make(map[string]int)

	words := strings.Split(phrase, " ")

	for _, w := range words {
		pieces := strings.Split(w, "")
		sort.Strings(pieces)

		w = strings.Join(pieces, "")

		seen[w]++

		if seen[w] > 1 {
			return false
		}
	}

	return true
}
