package main

import (
	"bufio"
	"fmt"
	"os"
)

// For this puzzle, we are given lines of text. We need to find the first and
// last number in each line, combine them to form one number (eg. 4, 5 -> 45),
// and return the sum of every line.
//
// In part 1, you simply look for numbers.
// In part 2, the numbers can be written out as well (eg. "six")

func main() {
	input, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(input)

	sumPt1 := 0
	sumPt2 := 0
	for scanner.Scan() {
		sumPt1 += pt1(scanner.Text())
		sumPt2 += pt2(scanner.Text())
	}

	fmt.Println("Answer to part 1:", sumPt1)
	fmt.Println("Answer to part 2:", sumPt2)
}

// Scans from either end until it finds a number
// It converts them from ASCII to int and returns the combined number
func pt1(line string) int {
	first := 0
	last := len(line) - 1

	for !isNumeric(line[first]) {
		first++
	}
	for !isNumeric(line[last]) {
		last--
	}

	f := int(line[first]) - int("0"[0])
	l := int(line[last]) - int("0"[0])
	return f*10 + l
}

// Simple check if a byte is an ASCII number 0-9
func isNumeric(c byte) bool {
	return c >= 48 && c <= 57
}

// Uses [first/last]Num, which checks a slice and returns a number if the first
// digit is an ASCII numeral or the next letter spell out a number
func pt2(line string) int {
	return firstNum(line)*10 + lastNum(line)
}

var writtenToInt = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

// Returns the first number or written number in the string
func firstNum(str string) int {
	for i := 0; i < len(str); i++ {
		if found, num := findNum(str[i:]); found {
			return num
		}
	}
	return -1 //never happens by problem contraints
}

// Returns the last number or written number in the string
func lastNum(str string) int {
	for i := len(str) - 1; i >= 0; i-- {
		if found, num := findNum(str[i:]); found {
			return num
		}
	}
	return -1 //never happens by problem contraints
}

// Returns the result of trying to find a number at the start of a string
func findNum(str string) (bool, int) {
	if str[0] >= 48 && str[0] <= 57 {
		return true, int(str[0]) - int("0"[0])
	}
	for k, v := range writtenToInt {
		if len(k) <= len(str) && str[:len(k)] == k {
			return true, v
		}
	}
	return false, 0
}
