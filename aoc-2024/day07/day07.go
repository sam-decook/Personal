package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println(Part2("input.txt"))
}

type Equation struct {
	Test    int
	Factors []int
}

func parse(file string) []Equation {
	f, _ := os.Open(file)
	scanner := bufio.NewScanner(f)

	eqs := make([]Equation, 0, 16)

	for scanner.Scan() {
		a, b, _ := strings.Cut(scanner.Text(), ": ")
		test, _ := strconv.Atoi(a)

		factors := make([]int, 0, 16)
		for _, n := range strings.Fields(b) {
			num, _ := strconv.Atoi(n)
			factors = append(factors, num)
		}

		eqs = append(eqs, Equation{Test: test, Factors: factors})
	}

	return eqs
}

func Part1(file string) int {
	eqs := parse(file)

	sum := 0
	for _, eq := range eqs {
		if solve(eq.Factors[0], eq.Test, eq.Factors[1:]) {
			sum += eq.Test
		}
	}

	return sum
}

func solve(num, test int, factors []int) bool {
	if len(factors) == 1 {
		return num+factors[0] == test || num*factors[0] == test
	}
	return solve(num+factors[0], test, factors[1:]) || solve(num*factors[0], test, factors[1:])
}

func Part2(file string) int {
	eqs := parse(file)

	sum := 0
	for _, eq := range eqs {
		if solve2(eq.Factors[0], eq.Test, eq.Factors[1:]) {
			sum += eq.Test
		}
	}

	return sum
}

func solve2(num, test int, factors []int) bool {
	if len(factors) == 1 {
		return num+factors[0] == test ||
			num*factors[0] == test ||
			concat(num, factors[0]) == test
	}
	return solve2(num+factors[0], test, factors[1:]) ||
		solve2(num*factors[0], test, factors[1:]) ||
		solve2(concat(num, factors[0]), test, factors[1:])
}

func concat(a, b int) int {
	num, _ := strconv.Atoi(strconv.Itoa(a) + strconv.Itoa(b))
	return num
}
