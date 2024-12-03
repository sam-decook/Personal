package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Part1("input.txt"))
}

func Part1(file string) int {
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
// 326 too low
func Part2(file string) int {
	f, _ := os.Open(file)
	scanner := bufio.NewScanner(f)

	// Lines stores the differences now
	lines := make([][]int, 0, 16)
	increasing := make([]bool, 0, 16)
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		line := make([]int, 0, len(fields)-1)
		inc := 0

		prev, _ := strconv.Atoi(fields[0])
		for i := 1; i < len(fields); i++ {
			n, _ := strconv.Atoi(fields[i])
			line = append(line, n-prev)
			if n-prev > 0 {
				inc++
			} else if n-prev < 0 {
				inc--
			}
			prev = n
		}

		lines = append(lines, line)
		increasing = append(increasing, inc > 0)
	}

	safe := 0
	for i, line := range lines {
		unsafe := 0

		if increasing[i] {
			for _, diff := range line {
				fmt.Printf("%d ", diff)
				if diff <= 0 {
					unsafe += 1
					fmt.Print("tolerable ")
				} else if diff > 3 {
					unsafe += 2
					fmt.Print("intolerable ")
				}
			}
		} else {
			for _, diff := range line {
				fmt.Printf("%d ", diff)
				if diff >= 0 {
					unsafe += 1
					fmt.Print("tolerable ")
				} else if diff < -3 {
					unsafe += 2
					fmt.Print("intolerable ")
				}
			}
		}

		if unsafe < 2 {
			safe++
		}
		fmt.Printf("#: %d\n", safe)
		if unsafe >= 2 {
			fmt.Println("\t", line)
		}
	}

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
