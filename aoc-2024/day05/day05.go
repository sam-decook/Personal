package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Part2("input.txt"))
}

func parse(file string) (map[int][]int, map[int][]int, [][]int) {
	f, _ := os.Open(file)
	scanner := bufio.NewScanner(f)

	// Parse the orderings
	before := make(map[int][]int)
	after := make(map[int][]int)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		a, b, _ := strings.Cut(line, "|")
		x, _ := strconv.Atoi(a)
		y, _ := strconv.Atoi(b)

		after[x] = append(after[x], y)
		before[y] = append(before[y], x)
	}

	// Parse the updates
	updates := make([][]int, 0, 16)
	for scanner.Scan() {
		pages := strings.Split(scanner.Text(), ",")

		update := make([]int, len(pages))
		for i, page := range pages {
			n, _ := strconv.Atoi(page)
			update[i] = n
		}

		updates = append(updates, update)
	}

	return before, after, updates
}

func isBefore(n int, rest []int, after map[int][]int) bool {
	for _, page := range rest {
		if !slices.Contains(after[n], page) {
			return false
		}
	}
	return true
}

func isAfter(n int, rest []int, before map[int][]int) bool {
	for _, page := range rest {
		if !slices.Contains(before[n], page) {
			return false
		}
	}
	return true
}

func inOrder(update []int, before, after map[int][]int) bool {
	for i, page := range update {
		belowGood := true
		if i > 0 {
			belowGood = isAfter(page, update[:i], before)
		}
		// Fun fact, you don't need to check above...
		// not sure if that's a mistake, or I'm just reading it wrong
		aboveGood := true
		if i < len(update)-1 {
			aboveGood = isBefore(page, update[i+1:], after)
		}
		if !belowGood || !aboveGood {
			return false
		}
	}
	return true
}

func Part1(file string) int {
	before, after, updates := parse(file)

	sum := 0
	for _, update := range updates {
		if inOrder(update, before, after) {
			mid := len(update) / 2
			sum += update[mid]
		}
	}

	return sum
}

// Optimization: just find the middle element and return it!
func ordered(update []int, before map[int][]int) []int {
	ordered := make([]int, len(update))
	for _, page := range update {
		idx := matches(update, before[page])
		ordered[idx] = page
	}
	return ordered
}

func matches(update, before []int) int {
	amt := 0
	for _, n := range update {
		if slices.Contains(before, n) {
			amt++
		}
	}
	return amt
}

func Part2(file string) int {
	before, after, updates := parse(file)

	sum := 0
	for _, update := range updates {
		if !inOrder(update, before, after) {
			mid := len(update) / 2
			sum += ordered(update, before)[mid]
		}
	}

	return sum
}
