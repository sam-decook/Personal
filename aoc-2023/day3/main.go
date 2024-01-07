package main

import (
	"bufio"
	"fmt"
	"os"
)

// To store the location of a token in the schematic
type coor struct {
	I int
	J int
}

// To store each number and the location of each digit to check for adjacent symbols
type number struct {
	Val  int
	Locs []coor
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)
	input := make([]string, 0)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}

	numbers := parse(input)

	sumPt1 := 0
	for _, num := range numbers {
		adjacentSymbols := searchSurrounding(input, num.Locs, isSymbol)
		if len(adjacentSymbols) > 0 {
			sumPt1 += num.Val
		}
	}

	// Make a map from the individual locations to the number
	m := make(map[coor]number)
	for _, num := range numbers {
		for _, loc := range num.Locs {
			m[loc] = num
		}
	}

	sumPt2 := 0
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			gear, ratio := getGearRatio(input, i, j, m)
			if gear {
				sumPt2 += ratio
			}
		}
	}

	fmt.Println("Answer to part 1:", sumPt1)
	fmt.Println("Answer to part 2:", sumPt2)
}

func parse(input []string) []number {
	// Pass 1: fill in the locations
	numbers := make([]number, 0)
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input[i]); j++ {
			if isNumeric(input[i][j]) {
				n := number{}
				for j < len(input[i]) && isNumeric(input[i][j]) {
					n.Locs = append(n.Locs, coor{i, j})
					j++
				}
				numbers = append(numbers, n)
			}
		}
	}

	// Pass 2: calculate the numbers
	for i := 0; i < len(numbers); i++ {
		val := 0
		for _, loc := range numbers[i].Locs {
			val *= 10
			val += int(input[loc.I][loc.J]) - int("0"[0])
		}
		numbers[i].Val = val
	}
	return numbers
}

// Searches the areas surrounding the search coordinates
// Returns locations according to the found function
func searchSurrounding(input []string, search []coor, found func(byte) bool) []coor {
	// Find bounds, ensure all indices are in range
	top := search[0].I - 1
	if top < 0 {
		top = 0
	}

	bottom := search[0].I + 1
	if bottom >= len(input) {
		bottom--
	}

	left := search[0].J - 1
	if left < 0 {
		left = 0
	}

	last := len(search) - 1
	right := search[last].J + 1
	if right >= len(input[0]) {
		right--
	}

	// Add locations to slice if found() evaluates true
	locs := make([]coor, 0)
	for a := top; a <= bottom; a++ {
		for b := left; b <= right; b++ {
			if found(input[a][b]) {
				locs = append(locs, coor{a, b})
			}
		}
	}

	return locs
}

func isNumeric(c byte) bool {
	return c >= 48 && c <= 57
}

// The input consists of numbers 0-9, ".", and symbols
// Thus, if it isn't a number or a ".", it is a symbol
func isSymbol(c byte) bool {
	return !isNumeric(c) && c != "."[0]
}

// A gear is an asterisk ("*") with two adjacent numbers
// It's gear ratio is found by multiopling the two numbers together
func getGearRatio(input []string, i, j int, m map[coor]number) (bool, int) {
	if input[i][j] != "*"[0] {
		return false, 0
	}

	// Get locations of surrounding numbers
	var loc []coor = []coor{{i, j}}
	adjacentNums := searchSurrounding(input, loc, isNumeric)

	// Find the number each location is a part of
	numbers := make([]number, 0)
	for _, num := range adjacentNums {
		numbers = append(numbers, m[num])
	}

	// Remove duplicates
	numbers = dedup(numbers)

	// Check if there are two
	if len(numbers) != 2 {
		return false, 0
	}

	// Find gear ratio by multiplying the numbers
	return true, numbers[0].Val * numbers[1].Val
}

// Compares numbers by first location
func dedup(numbers []number) []number {
	dedup := make([]number, 0)
	for i := 0; i < len(numbers); i++ {
		found := false
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i].Locs[0] == numbers[j].Locs[0] {
				found = true
			}
		}
		if !found {
			dedup = append(dedup, numbers[i])
		}
	}
	return dedup
}
