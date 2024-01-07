package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const (
	UP = iota
	DOWN
	LEFT
	RIGHT
)

type direction int

type line struct {
	dir direction
	amt int
}

type point struct {
	i, j int
}

type pair struct {
	lo, hi int
}

func main() {
	lines, start, size := getLinesAndGrid("input.txt")

	// Make the floor map
	floor := make([][]bool, size.i)
	for row := range floor {
		floor[row] = make([]bool, size.j)
	}

	floor[start.i][start.j] = true

	sumPt2 := 0

	fmt.Println("Answer to part 1:", part1(start, lines, floor))
	fmt.Println("Answer to part 2:", sumPt2)
}

// Parses lines (easy)
// Finds grid size and starting position (not easy)
func getLinesAndGrid(file string) (lines []line, start, size point) {
	input, _ := os.Open(file)
	scanner := bufio.NewScanner(input)

	lines = make([]line, 0)

	// Find the boundaries of the dig plan
	i := 0
	j := 0
	maxI := 0
	maxJ := 0
	minI := 0
	minJ := 0

	for scanner.Scan() {
		dir, amt := parse(scanner.Text())
		lines = append(lines, line{dir, amt})

		switch dir {
		case UP:
			i -= amt
			if i < minI {
				minI = i
			}
		case DOWN:
			i += amt
			if i > maxI {
				maxI = i
			}
		case LEFT:
			j -= amt
			if j < minJ {
				minJ = j
			}
		case RIGHT:
			j += amt
			if j > maxJ {
				maxJ = j
			}
		}
	}

	start = point{-minI, -minJ}
	size = point{maxI - minI + 1, maxJ - minJ + 1}

	return lines, start, size
}

func parse(line string) (direction, int) {
	fields := strings.Fields(line)

	var dir direction = UP
	switch fields[0] {
	case "D":
		dir = DOWN
	case "L":
		dir = LEFT
	case "R":
		dir = RIGHT
	}

	amt, _ := strconv.Atoi(fields[1])

	return dir, amt
}

func display(floor [][]bool) {
	for i := range floor {
		for j := range floor[i] {
			if floor[i][j] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

// Could solve this by casting it into something the code from day 10 can solve
//
// Excavate the perimeter
// Then the middle
// Return the number of squares dug out
// 37785 too high
func part1(start point, lines []line, floor [][]bool) int {
	// Excavate perimeter
	x := start.i
	y := start.j
	for _, line := range lines {
		x, y = dig(x, y, line.amt, line.dir, floor)
	}
	// display(floor)
	// fmt.Println()

	//Excavate middle
	for i := range floor {
		excavateRow(i, floor)
	}
	display(floor)
	fmt.Println()

	// Count cubic meters
	n := 0
	for i := range floor {
		for j := range floor[i] {
			if floor[i][j] {
				n++
			}
		}
	}

	return n
}

// From (x, y), dig amt in dir
func dig(x, y, amt int, dir direction, floor [][]bool) (int, int) {
	switch dir {
	case UP:
		for i := x - 1; i >= x-amt; i-- {
			floor[i][y] = true
		}
		return x - amt, y
	case DOWN:
		for i := x + 1; i <= x+amt; i++ {
			floor[i][y] = true
		}
		return x + amt, y
	case LEFT:
		for j := y - 1; j >= y-amt; j-- {
			floor[x][j] = true
		}
		return x, y - amt
	case RIGHT:
		for j := y + 1; j <= y+amt; j++ {
			floor[x][j] = true
		}
		return x, y + amt
	default:
		return -1, -1
	}
}

func excavateRow(i int, floor [][]bool) {
	ranges := getDigRange(i, floor)

	for _, pair := range ranges {
		for j := pair.lo; j < pair.hi; j++ {
			floor[i][j] = true
		}
	}
}

// Returns a slice of ranges to dig in a row
// Edge cases to consider:                            ######   (2)
// - straight edges                                   #    ### (1)
//   - either delay start of range until it ends (1)  #      #
//   - or don't even start range (2)                  ######## (2)
//   - you have no freaking idea
func getDigRange(i int, floor [][]bool) []pair {
	ranges := make([]pair, 0)
	start := -1
	for j := 1; j < len(floor[i]); j++ {
		if start == -1 {
			// Looks for a start pattern "#."
			if floor[i][j-1] && !floor[i][j] {
				start = j
			}
		} else {
			// Looks for an end pattern ".#"
			if !floor[i][j-1] && floor[i][j] {
				if i == 0 {
					if floor[i-1][start-1] != floor[i-1][j] {
						ranges = append(ranges, pair{start, j})
					}
				} else {
					if floor[i+1][start-1] != floor[i+1][j] {
						ranges = append(ranges, pair{start, j})
					}
				}
				start = -1
			}
		}
	}

	return ranges
}
