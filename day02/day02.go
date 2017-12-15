package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"os"
	"strings"

	"github.com/spf13/cast"
)

func main() {
	sheet := loadInput()
	fmt.Printf("Part 1 Checksum: %d\n", getChecksum(sheet))
	fmt.Printf("Part 2 Checksum: %d\n", getDivisionChecksum(sheet))
}

func loadInput() [][]int {
	bytes, err := ioutil.ReadFile("./input.txt")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	if len(bytes) <= 0 {
		fmt.Println("No Input")
		os.Exit(1)
	}

	input := strings.Split(string(bytes), "\n")

	var sheet [][]int
	for _, r := range input {
		var row []int
		rb := strings.Split(r, "\t")

		for _, i := range rb {
			row = append(row, cast.ToInt(i))
		}

		sheet = append(sheet, row)
	}

	return sheet
}

func getHighLow(row []int) (int, int) {
	if len(row) <= 0 {
		return 0, 0
	}

	high := row[0]
	low := row[0]

	for _, v := range row {
		if v > high {
			high = v
		}
		if v < low {
			low = v
		}
	}

	return low, high
}

func getChecksum(sheet [][]int) int {
	checksum := 0

	for _, row := range sheet {
		low, high := getHighLow(row)
		checksum += high - low
	}

	return checksum
}

func getDivisionChecksum(sheet [][]int) int {
	checksum := 0

	for _, row := range sheet {
		checksum += getDivisionValue(row)
	}

	return checksum
}

func getDivisionValue(row []int) int {
	for i, a := range row {
		for k, b := range row {
			if i == k {
				continue
			}

			if canDivide(a, b) {
				high := math.Max(float64(a), float64(b))
				low := math.Min(float64(a), float64(b))
				return int(high / low)
			}
		}
	}

	return 0
}

func canDivide(a, b int) bool {
	if a == 0 || b == 0 {
		return false
	}
	return (a%b == 0) || (b%a == 0)
}
