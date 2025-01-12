package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	fmt.Println(Part2("input.txt"))
}

func parseLines(file string) [][]int {
	f, _ := os.Open(file)
	scanner := bufio.NewScanner(f)

	lines := make([][]int, 0, 16)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		line := make([]int, len(fields))

		for i, field := range fields {
			n, _ := strconv.Atoi(field)
			line[i] = n
		}

		lines = append(lines, line)
	}
	return lines
}

func Part1(file string) int {
	lines := parseLines(file)

	safe := 0
	for _, line := range lines {
		if isSlowDecreasing(line) || isSlowIncreasing(line) {
			safe++
		}
	}

	return safe
}

// The difference can be negative and recovered from.
// If the difference is greater than 3, you can't remove either and make it safe.
// It's monotonically increasing, so only a non-increasing section can be recovered from.
//
// Ugh, what if the first is the one to remove, and that switches it's direction?
//
// I give up. I'm going to check each line for all combinations of a single removal.
// But... brute forcing it isn't that brute-y!
// - The input is only 1k lines and each line only has 5-8 numbers, roughly.
// - You're only adding 4-7 more numbers, so max 7k lines. That's small
//
// 326 too low
func Part2(file string) int {
	lines := parseLines(file)

	start := time.Now()
	safe := 0
	for _, line := range lines {
		if isSlowDecreasing(line) || isSlowIncreasing(line) {
			safe++
		} else if removeAndCheck(line) {
			safe++
		}
	}
	fmt.Println("Took: ", time.Since(start))
	return safe
}

// All lines have multiple numbers

func isSlowIncreasing(nums []int) bool {
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		curr := nums[i]
		diff := curr - prev
		if diff <= 0 || diff > 3 {
			return false
		}
		prev = curr
	}
	return true
}

func isSlowDecreasing(nums []int) bool {
	prev := nums[0]
	for i := 1; i < len(nums); i++ {
		curr := nums[i]
		diff := prev - curr
		if diff <= 0 || diff > 3 {
			return false
		}
		prev = curr
	}
	return true
}

// Doing this instead of making a new array each time: 210.417µs -> 55.709µs
func removeAndCheck(nums []int) bool {
	skipped := nums[1:]
	last := len(nums) - 1
	for i := range nums {
		if isSlowDecreasing(skipped) || isSlowIncreasing(skipped) {
			return true
		}
		if i != last {
			skipped[i] = nums[i]
		}
	}
	return false
}

// 1 2 3 4 5
// 2 3 4 5
// 1 3 4 5
// 1 2 4 5
// 1 2 3 5
// 1 2 3 4
//
