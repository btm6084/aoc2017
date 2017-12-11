package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/spf13/cast"
)

func main() {
	fmt.Printf("First Exit: %d\n", findFirstExit(getInput()))
	fmt.Printf("Second Exit: %d\n\n", findSecondExit(getInput()))
}

func findFirstExit(maze []int) int {
	current := 0
	steps := 0

	for current < len(maze) {
		steps++

		// Jump forward `current` number.
		start := current
		move := maze[start]
		current = start + move

		// Increment previous position by 1.
		maze[start]++
	}

	return steps
}

func findSecondExit(maze []int) int {
	current := 0
	steps := 0

	for current < len(maze) {
		steps++

		// Jump forward `current` number.
		start := current
		move := maze[start]
		current = start + move

		// Increment previous position by 1 if less than 3, otherwise decrement.
		if move >= 3 {
			maze[start]--
		} else {
			maze[start]++
		}
	}

	return steps
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

	lines := strings.Split(string(bytes), "\n")

	var out []int

	for _, v := range lines {
		out = append(out, cast.ToInt(v))
	}

	return out
}
