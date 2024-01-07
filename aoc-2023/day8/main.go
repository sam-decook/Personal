package main

import (
	"bufio"
	"fmt"
	"os"
)

const (
	LEFT = iota
	RIGHT
)

type direction int

func (d direction) String() string {
	if d == LEFT {
		return "Left"
	} else {
		return "Right"
	}
}

func main() {
	input, _ := os.Open("input.txt")

	directions, m, starting := parse(input)

	fmt.Println("Answer to part 1:", getSteps("AAA", directions, m, onZZZ))

	pt2 := part2(directions, m, starting)
	fmt.Println("Answer to part 2:", pt2)
}

func parse(input *os.File) ([]direction, map[string][2]string, []string) {
	scanner := bufio.NewScanner(input)

	// Turn the first line into an array of directions (LEFT | RIGHT)
	scanner.Scan()
	turns := make([]direction, 0)
	for _, c := range scanner.Text() {
		var d direction = LEFT
		if c == 'R' {
			d = RIGHT
		}
		turns = append(turns, d)
	}

	scanner.Scan() // Skip blank line

	// Make a map from each node to an array of the next node to travel to
	// given a direction
	m := make(map[string][2]string)
	// Also find all starting nodes for part 2
	starting := make([]string, 0)

	for scanner.Scan() {
		line := scanner.Text()
		node := line[0:3]
		next := [2]string{line[7:10], line[12:15]}
		m[node] = next

		if node[2] == 'A' {
			starting = append(starting, node)
		}
	}

	return turns, m, starting
}

// Find the least common multiple of the number of steps it takes for each
// starting position to reach a node whose last character is 'Z'
func part2(directions []direction, m map[string][2]string, starting []string) int {
	steps := make([]int, 0)

	for _, node := range starting {
		steps = append(steps, getSteps(node, directions, m, onZ))
	}

	return lcm(steps)
}

func getSteps(node string, directions []direction, m map[string][2]string, cond func(string) bool) int {
	idx := 0 //into direction array
	steps := 0

	for !cond(node) {
		turn := directions[idx]

		node = m[node][turn]

		steps++
		idx++
		if idx >= len(directions) {
			idx = 0
		}
	}

	return steps
}

// Condition for part 1
func onZZZ(node string) bool {
	return node == "ZZZ"
}

// Condition for part 2
func onZ(node string) bool {
	return node[2] == 'Z'
}

// Find LCM using GCD: lcm(a, b) = a * b / gcd(a, b)
// Find LCM of multiple by recursion: lcm(a, b, c) = lcm(a, lcm(b, c))
// - but do it iteratively with an accumulated value
func lcm(nums []int) int {
	last := len(nums) - 1
	acc := nums[last]

	for i := last - 1; i >= 0; i-- {
		acc = nums[i] * acc / gcd(nums[i], acc)
	}

	return acc
}

// Using the Euclidean algorithm
func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}
