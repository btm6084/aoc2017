package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/spf13/cast"
)

func main() {
	bytes, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(bytes) <= 0 {
		fmt.Println("No Input")
		os.Exit(1)
	}

	fmt.Printf("Part 1 Sum: %d\n", sumByNext(bytes))
	fmt.Printf("Part 2 Sum: %d\n", sumByHalf(bytes))
}

func sumByNext(bytes []byte) int {
	last := len(bytes) - 1
	next := 0
	sum := 0

	for i := range bytes {
		if i >= last {
			next = 0
		} else {
			next = i + 1
		}

		v := cast.ToInt(string(bytes[i]))
		n := cast.ToInt(string(bytes[next]))

		if v == n {
			sum += v
		}
	}

	return sum
}

func sumByHalf(bytes []byte) int {
	next := 0
	step := len(bytes) / 2
	sum := 0

	for i := range bytes {
		if (i + step) >= len(bytes) {
			next = (i + step) - len(bytes)
		} else {
			next = i + step
		}

		v := cast.ToInt(string(bytes[i]))
		n := cast.ToInt(string(bytes[next]))

		if v == n {
			sum += v
		}
	}

	return sum
}
