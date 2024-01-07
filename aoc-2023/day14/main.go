package main

// Holy crap Sam, you have to test your code!
// It took you two days and two tries to get rolling in each direction correct

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

const (
	NONE = iota
	ROUND
	SQUARE
)

type rock int8

func (r rock) String() string {
	switch r {
	case NONE:
		return "."
	case ROUND:
		return "O"
	default:
		return "#"
	}
}

func main() {
	input, _ := os.Open("input.txt")

	platform := parse(input)

	fmt.Println("Answer to part 1:", part1(platform))
	fmt.Println("Answer to part 2:", part2(platform))
}

func parse(input *os.File) [][]rock {
	platform := make([][]rock, 0)
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		line := make([]rock, 0)
		for _, c := range scanner.Text() {
			switch c {
			case '.':
				line = append(line, NONE)
			case 'O':
				line = append(line, ROUND)
			case '#':
				line = append(line, SQUARE)
			}
		}
		platform = append(platform, line)
	}
	return platform
}

func part1(platform [][]rock) int {
	rollNorth(platform)
	return load(platform)
}

func part2(platform [][]rock) int {
	mem := make(map[string]int)

	var i int
	for i = 1; i < 1_000_000_000; i++ {
		rollNorth(platform)
		rollWest(platform)
		rollSouth(platform)
		rollEast(platform)

		ps := toString(platform)
		if mem[ps] != 0 {
			loop := i - mem[ps]
			cyclesLeft := 1_000_000_000 - i
			i += loop * (cyclesLeft / loop)
			break
		} else {
			mem[ps] = i
		}
	}

	for ; i < 1_000_000_000; i++ {
		rollNorth(platform)
		rollWest(platform)
		rollSouth(platform)
		rollEast(platform)
	}
	return load(platform)
}

func toString(platform [][]rock) string {
	var b strings.Builder
	for i := range platform {
		for j := range platform[i] {
			b.WriteString(platform[i][j].String())
		}
	}
	return b.String()
}

func rollNorth(platform [][]rock) {
	for j := range platform[0] {
		empty := -1
		for i := range platform {
			r := platform[i][j]
			switch {
			case r == SQUARE:
				empty = -1
			case r == ROUND && empty != -1:
				platform[i][j], platform[empty][j] =
					platform[empty][j], platform[i][j]
				empty++
			case r == NONE && empty == -1:
				empty = i
			}
		}
	}
}

func rollSouth(platform [][]rock) {
	for j := range platform[0] {
		empty := -1
		for i := len(platform) - 1; i >= 0; i-- {
			r := platform[i][j]
			switch {
			case r == SQUARE:
				empty = -1
			case r == ROUND && empty != -1:
				platform[i][j], platform[empty][j] =
					platform[empty][j], platform[i][j]
				empty--
			case r == NONE && empty == -1:
				empty = i
			}
		}
	}
}

func rollEast(platform [][]rock) {
	for i := range platform {
		empty := -1
		for j := len(platform[i]) - 1; j >= 0; j-- {
			r := platform[i][j]
			switch {
			case r == SQUARE:
				empty = -1
			case r == ROUND && empty != -1:
				platform[i][j], platform[i][empty] =
					platform[i][empty], platform[i][j]
				empty--
			case r == NONE && empty == -1:
				empty = j
			}
		}
	}
}

func rollWest(platform [][]rock) {
	for i := range platform {
		empty := -1
		for j := range platform[i] {
			r := platform[i][j]
			switch {
			case r == SQUARE:
				empty = -1
			case r == ROUND && empty != -1:
				platform[i][j], platform[i][empty] =
					platform[i][empty], platform[i][j]
				empty++
			case r == NONE && empty == -1:
				empty = j
			}
		}
	}
}

func load(platform [][]rock) int {
	load := 0
	last := len(platform)
	for i := range platform {
		for j := range platform[0] {
			if platform[i][j] == ROUND {
				load += last - i
			}
		}
	}
	return load
}
