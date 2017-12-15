package main

import (
	"fmt"
	"math"
)

// Point is a point in a 2d cartesian system.
type Point struct {
	X int
	Y int
}

func main() {
	pos := getCoords(325489)
	dist := getTaxiDist(Point{0, 0}, pos)

	fmt.Printf("Part 1: (X: %d, Y: %d, D: %d)\n", pos.X, pos.Y, dist)

	larger := getFirstLargerThan(325489)

	fmt.Printf("Part 2: First value written larger than %d: %d\n", 325489, larger)
}

func getTaxiDist(a, b Point) int {
	x := math.Abs(float64(a.X - b.X))
	y := math.Abs(float64(a.Y - b.Y))

	return int(x + y)
}

func getCoords(nodes int) Point {
	x := 0
	y := 0

	dx := 1
	dy := 0

	stepX := 1
	stepY := 1

	for i := 1; i < nodes; i++ {
		if dx > 0 && dy == 0 {
			// East
			x++
			dx--

			if dx == 0 {
				dy = stepY
				stepX++
			}
		} else if dx < 0 && dy == 0 {
			// West
			x--
			dx++

			if dx == 0 {
				dy = -1 * stepY
				stepX++
			}
		} else if dx == 0 && dy > 0 {
			// North
			y++
			dy--

			if dy == 0 {
				dx = -1 * stepX
				stepY++
			}
		} else {
			// South
			y--
			dy++

			if dy == 0 {
				dx = stepX
				stepY++
			}
		}
	}

	return Point{x, y}
}

// Find the highest value without going over
func getFirstLargerThan(target int) int {
	x := 0
	y := 0

	dx := 1
	dy := 0

	stepX := 1
	stepY := 1

	s := make(map[int]map[int]int)
	s[0] = make(map[int]int)

	s[0][0] = 1

	for s[x][y] <= target {
		if dx > 0 && dy == 0 {
			// East
			x++
			dx--

			s[x][y] = getValueForPoint(Point{x, y}, s)

			if dx == 0 {
				dy = stepY
				stepX++
			}
		} else if dx < 0 && dy == 0 {
			// West
			x--
			dx++

			s[x][y] = getValueForPoint(Point{x, y}, s)

			if dx == 0 {
				dy = -1 * stepY
				stepX++
			}
		} else if dx == 0 && dy > 0 {
			// North
			y++
			dy--

			s[x][y] = getValueForPoint(Point{x, y}, s)

			if dy == 0 {
				dx = -1 * stepX
				stepY++
			}
		} else {
			// South
			y--
			dy++

			s[x][y] = getValueForPoint(Point{x, y}, s)

			if dy == 0 {
				dx = stepX
				stepY++
			}
		}

		if s[x][y] > target {
			return s[x][y]
		}
	}

	return 0
}

func getValueForPoint(p Point, s map[int]map[int]int) int {
	v := 0

	x := p.X
	y := p.Y

	if _, isset := s[x]; !isset {
		s[x] = make(map[int]int)
	}

	v += s[x][y-1]
	v += s[x][y+1]

	if _, isset := s[x-1]; !isset {
		s[x-1] = make(map[int]int)
	}

	v += s[x-1][y-1]
	v += s[x-1][y]
	v += s[x-1][y+1]

	if _, isset := s[x+1]; !isset {
		s[x+1] = make(map[int]int)
	}

	v += s[x+1][y-1]
	v += s[x+1][y]
	v += s[x+1][y+1]

	return v
}
