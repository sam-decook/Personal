package main

import (
	"bufio"
	"fmt"
	"os"
)

type coor struct {
	i, j int
}

func main() {
	input, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(input)

	image := make([]string, 0)
	for scanner.Scan() {
		image = append(image, scanner.Text())
	}

	fmt.Println("Answer to part 1:", part1(image))
	fmt.Println("Answer to part 2:", part2(image))
}

// Sum the distance of each galaxy to every other galaxy
// Uses the manhattan distance (not euclidean)
func part1(image []string) int {
	// Expand the empty rows and columns
	image = expandGalaxy(image)

	// Find the galaxies
	galaxies := findGalaxies(image)

	total := 0

	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			vert := galaxies[j].i - galaxies[i].i
			if vert < 0 {
				vert = -vert
			}
			horiz := galaxies[j].j - galaxies[i].j
			if horiz < 0 {
				horiz = -horiz
			}
			total += vert + horiz
		}
	}

	return total
}

func expandGalaxy(image []string) []string {
	expanded := make([]string, 0, len(image))

	// Expand vertically
	for i := range image {
		expanded = append(expanded, image[i])
		if emptyRow(image, i) {
			expanded = append(expanded, image[i])
		}
	}

	// Expand horizontally
	// Iterate in reverse to avoid problems appending to the strings
	for j := len(expanded[0]) - 1; j >= 0; j-- {
		if emptyCol(expanded, j) {
			expanded = addEmptyCol(expanded, j)
		}
	}

	return expanded
}

func emptyRow(image []string, row int) bool {
	for _, c := range image[row] {
		if c == '#' {
			return false
		}
	}
	return true
}

func emptyCol(image []string, col int) bool {
	for i := range image {
		if image[i][col] == '#' {
			return false
		}
	}
	return true
}

func addEmptyCol(image []string, col int) []string {
	expanded := make([]string, 0, len(image))
	for _, line := range image {
		expanded = append(expanded, line[:col]+"."+line[col:])
	}
	return expanded
}

func findGalaxies(image []string) []coor {
	galaxies := make([]coor, 0)
	for i := range image {
		for j := range image[i] {
			if image[i][j] == '#' {
				galaxies = append(galaxies, coor{i, j})
			}
		}
	}
	return galaxies
}

// For part 2, every blank row or column expands into 1 million instead of 1
// That won't fit into memory
// Find the empty rows and columns
// If one lies on the path between two galaxies, add 999,999 (in addition to 1)
//
// The code for part 2 can be parameterized and used to solve part 1 as well!
// - Multiply by 1 instead of 999,999 -> double empty space
func part2(image []string) int {
	galaxies := findGalaxies(image)

	rowsEmpty := make(map[int]bool, len(image))
	for i := range image {
		if emptyRow(image, i) {
			rowsEmpty[i] = true
		}
	}

	colsEmpty := make(map[int]bool, len(image))
	for j := range image[0] {
		if emptyCol(image, j) {
			colsEmpty[j] = true
		}
	}

	total := 0

	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			total += vertDist(galaxies[i], galaxies[j], rowsEmpty) +
				horizDist(galaxies[i], galaxies[j], colsEmpty)
		}
	}

	return total
}

func vertDist(a, b coor, rowsEmpty map[int]bool) int {
	if a.i > b.i {
		a, b = b, a
	}

	empties := 0
	for i := a.i; i < b.i; i++ {
		if rowsEmpty[i] {
			empties++
		}
	}

	return b.i - a.i + empties*999_999
}

func horizDist(a, b coor, colsEmpty map[int]bool) int {
	if a.j > b.j {
		a, b = b, a
	}

	empties := 0
	for j := a.j; j < b.j; j++ {
		if colsEmpty[j] {
			empties++
		}
	}

	return b.j - a.j + empties*999_999
}
