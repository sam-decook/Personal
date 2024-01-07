package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

const (
	NORTH = iota
	SOUTH
	EAST
	WEST
)

type direction int

type location struct {
	kind     rune
	distance int
	i, j     int
	inLoop   bool
	inside   bool
}

// Useful for part 1
// func (l *location) String() string {
// 	d := "."
// 	if l.distance != math.MaxInt {
// 		d = fmt.Sprintf("%d", l.distance)
// 	}
// 	return string(l.kind) + " " + d
// }

// Useful for part2
func (l *location) String() string {
	d := string(l.kind)

	if l.kind == '.' {
		if l.inside {
			d = "I"
		} else {
			d = "O"
		}
	}

	return d
}

func main() {
	input, _ := os.Open("input.txt")

	field := parse(input)
	start := findStart(field)
	startUpdate(start, field)

	fmt.Println("Answer to part 1:", part1(field))
	// 50 is wrong
	fmt.Println("Answer to part 2:", part2(field))
}

func parse(input *os.File) [][]location {
	scanner := bufio.NewScanner(input)
	field := make([][]location, 0)

	i := 0
	for scanner.Scan() {
		row := parseLine(scanner.Text(), i)
		field = append(field, row)
		i++
	}

	return field
}

func parseLine(line string, i int) []location {
	row := make([]location, 0)
	for j, c := range line {
		row = append(row, location{c, math.MaxInt, i, j, false, false})
	}
	return row
}

func part1(field [][]location) int {
	steps := 0
	for i := range field {
		for j := range field[i] {
			dist := field[i][j].distance
			if dist != math.MaxInt && dist > steps {
				steps = dist
			}
		}
	}
	return steps
}

func findStart(field [][]location) location {
	for i := range field {
		for j := range field[i] {
			if field[i][j].kind == 'S' {
				updateStartKind(field, i, j)
				return field[i][j]
			}
		}
	}
	// Unreachable
	return location{}
}

// github.com/omotto/AdventOfCode2023/blob/main/src/day10/main.go
// is the inspiration for this solution.
func updateStartKind(field [][]location, i, j int) {
	// Is there a pipe that connects _ to the start?
	down := false
	if i-1 >= 0 {
		down = field[i-1][j].kind == 'F' ||
			field[i-1][j].kind == '7' ||
			field[i-1][j].kind == '|'
	}

	up := false
	if i+1 < len(field) {
		up = field[i+1][j].kind == 'L' ||
			field[i+1][j].kind == 'J' ||
			field[i+1][j].kind == '|'
	}

	left := false
	if j-1 >= 0 {
		left = field[i][j-1].kind == 'F' ||
			field[i][j-1].kind == 'L' ||
			field[i][j-1].kind == '-'
	}

	right := false
	if j+1 < len(field[i]) {
		right = field[i][j+1].kind == 'J' ||
			field[i][j+1].kind == '7' ||
			field[i][j+1].kind == '-'
	}

	switch {
	case left && right:
		field[i][j].kind = '-'
	case up && down:
		field[i][j].kind = '|'
	case up && right:
		field[i][j].kind = 'F'
	case up && left:
		field[i][j].kind = 'J'
	case down && right:
		field[i][j].kind = 'L'
	case down && left:
		field[i][j].kind = '7'
	}
}

func startUpdate(start location, field [][]location) {
	field[start.i][start.j].distance = 0
	field[start.i][start.j].inLoop = true

	i := start.i
	j := start.j

	// Up
	if i-1 >= 0 {
		next := field[i-1][j]
		switch next.kind {
		case '7', 'F', '|':
			updateDistance(next, start, field, 1)
		}
	}
	// Down
	if i+1 < len(field) {
		next := field[i+1][j]
		switch next.kind {
		case 'L', 'J', '|':
			updateDistance(next, start, field, 1)
		}
	}
	// Left
	if j-1 >= 0 {
		next := field[i][j-1]
		switch next.kind {
		case 'L', 'F', '-':
			updateDistance(next, start, field, 1)
		}
	}
	// Right
	if j+1 < len(field[0]) {
		next := field[i][j+1]
		switch next.kind {
		case 'J', '7', '-':
			updateDistance(next, start, field, 1)
		}
	}
}

func updateDistance(curr, prev location, field [][]location, dist int) {
	if curr.distance == 0 {
		return
	}

	field[curr.i][curr.j].distance = min(curr.distance, dist)
	field[curr.i][curr.j].inLoop = true
	i := curr.i
	j := curr.j

	// Assume the next direction is valid
	switch nextDirection(curr, prev) {
	case NORTH:
		if i-1 >= 0 {
			next := field[i-1][j]
			updateDistance(next, curr, field, dist+1)
		}
	case SOUTH:
		if i+1 < len(field) {
			next := field[i+1][j]
			updateDistance(next, curr, field, dist+1)
		}
	case WEST:
		if j-1 >= 0 {
			next := field[i][j-1]
			updateDistance(next, curr, field, dist+1)
		}
	case EAST:
		if j+1 < len(field[0]) {
			next := field[i][j+1]
			updateDistance(next, curr, field, dist+1)
		}
	}
}

func nextDirection(curr, prev location) direction {
	switch curr.kind {
	case '|':
		if prev.i > curr.i {
			return NORTH
		} else {
			return SOUTH
		}
	case '-':
		if prev.j < curr.j {
			return EAST
		} else {
			return WEST
		}
	case 'L':
		if curr.i == prev.i {
			return NORTH
		} else {
			return EAST
		}
	case 'J':
		if curr.i == prev.i {
			return NORTH
		} else {
			return WEST
		}
	case '7':
		if curr.i == prev.i {
			return SOUTH
		} else {
			return WEST
		}
	case 'F':
		if curr.i == prev.i {
			return SOUTH
		} else {
			return EAST
		}
	}
	panic("Issue with nextDirection")
}

func part2(field [][]location) int {
	inside := 0
	for i := range field {
		for j := range field[i] {
			if field[i][j].kind == '.' && insideLoop(field, i, j) {
				field[i][j].inside = true
				inside++
			}
			//fmt.Print(field[i][j].String())
		}
		//fmt.Println()
	}
	return inside
}

// From Algorithms: Computational Geometry: Point inside polygon
// Extend a horizontal line and count intersections with loop
// - even: you are outside the path
// - odd: you are inside the path
//
// Edge case where it incrememts many times for a straight line
// - eg: F7, this would increment twice (even -> outside) when you are inside
func insideLoop(field [][]location, i, j int) bool {
	intersections := 0
	for k := 0; k < j; k++ {
		if field[i][k].inLoop {
			intersections++

			// If the type is 'F' or 'L', a 'J' or '7' is guaranteed to follow
			// Set k to that index so that section is only counted once
			if field[i][k].kind == 'F' {
				k++
				for field[i][k].kind != 'J' &&
					field[i][k].kind != '7' {
					k++
				}
				if field[i][k].kind == '7' {
					intersections++
				}
			}

			if field[i][k].kind == 'L' {
				k++
				for field[i][k].kind != 'J' &&
					field[i][k].kind != '7' {
					k++
				}
				if field[i][k].kind == 'J' {
					intersections++
				}
			}
		}
	}
	return intersections%2 == 1
}
