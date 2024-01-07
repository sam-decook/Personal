package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	input, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(input)

	// Holds each pattern
	patterns := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		pattern := make([]string, 0)
		for line != "" {
			pattern = append(pattern, line)
			scanner.Scan()
			line = scanner.Text()
		}
		patterns = append(patterns, pattern)
	}

	// Holds each pattern turned 90 degrees
	orthPatterns := make([][]string, 0)
	for _, pattern := range patterns {
		lines := make([]string, 0)
		for i := range pattern[0] {
			var b strings.Builder
			for _, row := range pattern {
				b.WriteByte((row[i]))
			}
			lines = append(lines, b.String())
		}
		orthPatterns = append(orthPatterns, lines)
	}

	fmt.Println("Answer to part 1:", part1(patterns, orthPatterns))
	// Part 2 is another problem I don't know how to solve
	// "Find the one square that would make a reflection if it was the opposite"
	fmt.Println("Answer to part 2:", 0)
}

func part1(patterns, orthPatterns [][]string) int {
	total := 0

	// Horizontal reflection
	for _, pattern := range patterns {
		found, i := findReflection(pattern)
		if found {
			total += i * 100
		}
	}

	// Vertical reflection
	for _, pattern := range orthPatterns {
		found, j := findReflection(pattern)
		if found {
			total += j
		}
	}

	return total
}

func findReflection(pattern []string) (bool, int) {
	for i := 1; i < len(pattern); i++ {
		if isRelflection(i-1, i, pattern) {
			return true, i
		}
	}

	return false, 0
}

func isRelflection(a, b int, pattern []string) bool {
	if pattern[a] != pattern[b] {
		return false
	}
	for a >= 0 && b < len(pattern) {
		if pattern[a] != pattern[b] {
			return false
		}
		a--
		b++
	}
	return true
}
