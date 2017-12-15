package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cast"
)

func main() {
	banks := getInput()
	fmt.Printf("Part 1 Steps: %d\n", findRepeat(banks))
	fmt.Printf("Part 2 Steps: %d\n", findLoopSize(banks))
}

func findLoopSize(banks []int) int {
	findRepeat(banks)
	return findRepeat(banks) - 1
}

func findRepeat(banks []int) int {
	steps := 0
	seen := make(map[string]int)

	for {
		steps++
		redistribute(banks)
		key := getCacheKey(banks)

		seen[key]++

		if seen[key] > 1 {
			return steps
		}
	}
}

func getCacheKey(banks []int) string {
	var key []string
	for _, v := range banks {
		key = append(key, cast.ToString(v))
	}

	return strings.Join(key, "")
}

func redistribute(banks []int) {
	fullest := findFullest(banks)
	blocks := banks[fullest]
	current := fullest + 1

	banks[fullest] = 0

	for i := 0; i < blocks; i++ {
		if current >= len(banks) {
			current = 0
		}

		banks[current]++
		current++
	}
}

func findFullest(banks []int) int {
	if len(banks) <= 0 {
		return 0
	}

	pos := 0
	max := banks[0]
	for i, v := range banks {
		if v > max {
			max = v
			pos = i
		}
	}

	return pos
}

func getInput() []int {
	bytes, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(bytes) <= 0 {
		fmt.Println("No Input")
		os.Exit(1)
	}

	lines := strings.Split(string(bytes), "\t")

	var out []int

	for _, v := range lines {
		out = append(out, cast.ToInt(v))
	}

	return out
}
