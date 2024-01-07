package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Okay, this makes me like recursion a lot

func main() {
	input, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(input)

	sumPt1 := 0
	sumPt2 := 0
	for scanner.Scan() {
		numbers := parse(scanner.Text())
		differences := getDifferences([][]int{numbers})
		sumPt1 += part1(differences)
		sumPt2 += part2(differences)
	}

	fmt.Println("Answer to part 1:", sumPt1)
	fmt.Println("Answer to part 2:", sumPt2)
}

// Parse a line of numbers into a slice of ints
func parse(line string) []int {
	numbers := make([]int, 0)
	for _, field := range strings.Fields(line) {
		n, _ := strconv.Atoi(field)
		numbers = append(numbers, n)
	}
	return numbers
}

// Recurse down until you find the rate of growth
// Each line contains the difference between two numbers
// Go until the whole line is 0
func getDifferences(history [][]int) [][]int {
	last := len(history) - 1
	if allZeros(history[last]) {
		return history
	}

	// Get the differences
	diffsLen := len(history[last]) - 1
	diffs := make([]int, diffsLen)
	for i := 0; i < diffsLen; i++ {
		diffs[i] = history[last][i+1] - history[last][i]
	}

	// Add to the history and keep going
	history = append(history, diffs)
	return getDifferences(history)
}

// Add up the last numbers in each line to predict the next value
func part1(differences [][]int) int {
	if len(differences) == 1 {
		return 0
	}

	last := len(differences[0]) - 1
	return differences[0][last] + part1(differences[1:])
}

// Subtract the first numbers to predict the value before the history started
func part2(differences [][]int) int {
	if len(differences) == 1 {
		return 0
	}

	return differences[0][0] - part2(differences[1:])
}

func allZeros(numbers []int) bool {
	for _, n := range numbers {
		if n != 0 {
			return false
		}
	}
	return true
}
