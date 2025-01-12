package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Part2("input.txt"))
}

func parse(file string) []byte {
	b, _ := os.ReadFile(file)
	return b
}

// Slice multiply(x,y) -> x,y
// Split, parse, and return x * y
func multiply(match []byte) int {
	start := 4
	end := len(match) - 1
	nums := match[start:end]

	x, y, _ := strings.Cut(string(nums), ",")

	n1, _ := strconv.Atoi(x)
	n2, _ := strconv.Atoi(y)

	return n1 * n2
}

func Part1(file string) int {
	input := parse(file)

	// Looks for mul(x,y)
	re := regexp.MustCompile(`mul\(\d+,\d+\)`).FindAll(input, -1)
	matches := re

	sum := 0
	for _, match := range matches {
		sum += multiply(match)
	}

	return sum
}

func Part2(file string) int {
	input := parse(file)

	// Looks for mul(x,y), do(), and don't()
	re := regexp.MustCompile(`mul\(\d+,\d+\)|do\(\)|don\'t\(\)`)
	matches := re.FindAll(input, -1)

	mul := []byte("mul(")
	do := []byte("do()")
	dont := []byte("don't()")

	enabled := true

	sum := 0
	for _, match := range matches {
		switch {
		case bytes.Equal(match, do):
			enabled = true
		case bytes.Equal(match, dont):
			enabled = false
		case enabled && bytes.HasPrefix(match, mul):
			sum += multiply(match)
		}
	}

	return sum
}
