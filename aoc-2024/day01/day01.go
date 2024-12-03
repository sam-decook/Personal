package day01

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

func Part1(file string) int {
	f, _ := os.Open(file)
	scanner := bufio.NewScanner(f)

	left := make([]int, 0, 32)
	right := make([]int, 0, 32)

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		n1, _ := strconv.Atoi(fields[0])
		n2, _ := strconv.Atoi(fields[1])
		left = append(left, n1)
		right = append(right, n2)
	}

	slices.Sort(left)
	slices.Sort(right)

	sum := 0
	for i := range left {
		if left[i] > right[i] {
			sum += left[i] - right[i]
		} else {
			sum += right[i] - left[i]
		}
	}

	return sum
}

func Part2(file string) int {
	f, _ := os.Open(file)
	scanner := bufio.NewScanner(f)

	left := make(map[int]int)
	right := make(map[int]int)

	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		n1, _ := strconv.Atoi(fields[0])
		n2, _ := strconv.Atoi(fields[1])
		left[n1]++
		right[n2]++
	}

	sum := 0
	for n, amt := range left {
		sum += n * amt * right[n]
	}

	return sum
}
