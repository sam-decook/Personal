package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println(Part2("input.txt"))
}

type Loc struct {
	i, j int
}

func parse(file string) ([][]byte, map[byte][]Loc) {
	f, _ := os.Open(file)
	scanner := bufio.NewScanner(f)

	// A map of the city's antennas
	city := make([][]byte, 0)

	for scanner.Scan() {
		city = append(city, []byte(scanner.Text()))
	}

	// Keeps track of the locations of antennas with the same frequency
	antennas := make(map[byte][]Loc)

	for i := range city {
		for j := range city[i] {
			if square := city[i][j]; square != '.' {
				antennas[square] = append(antennas[square], Loc{i, j})
			}

		}
	}

	return city, antennas
}

func Part1(file string) int {
	return core(file, extend)
}

func Part2(file string) int {
	return core(file, extend2)
}

func core(file string, extend func(Loc, Loc, [][]byte) []Loc) int {
	city, antennas := parse(file)

	antinodes := make([][]bool, len(city))
	for i := range city {
		row := make([]bool, len(city[i]))
		antinodes[i] = row
	}

	for _, locs := range antennas {
		for i := 0; i < len(locs)-1; i++ {
			for j := i + 1; j < len(locs); j++ {
				for _, loc := range extend(locs[i], locs[j], city) {
					antinodes[loc.i][loc.j] = true
				}
			}
		}
	}

	sum := 0
	for i := range antinodes {
		for j := range antinodes {
			if antinodes[i][j] {
				sum++
			} else {
			}
		}
	}

	return sum
}

func inBounds(i, j int, city [][]byte) bool {
	return i >= 0 && j >= 0 && i < len(city) && j < len(city[0])
}

// Just extend it once.
func extend(a, b Loc, city [][]byte) []Loc {
	antinodes := make([]Loc, 0, 16)

	di, dj := a.i-b.i, a.j-b.j
	i, j := a.i+di, a.j+dj
	if inBounds(i, j, city) {
		antinodes = append(antinodes, Loc{i, j})
	}

	di, dj = b.i-a.i, b.j-a.j
	i, j = b.i+di, b.j+dj
	if inBounds(i, j, city) {
		antinodes = append(antinodes, Loc{i, j})
	}
	return antinodes
}

// Extend it repeatedly. There will be antinodes on the two antennas.
func extend2(a, b Loc, city [][]byte) []Loc {
	antinodes := make([]Loc, 0, 16)

	di, dj := a.i-b.i, a.j-b.j
	i, j := b.i+di, b.j+dj
	for inBounds(i, j, city) {
		antinodes = append(antinodes, Loc{i, j})
		i += di
		j += dj
	}

	di, dj = b.i-a.i, b.j-a.j
	i, j = a.i+di, a.j+dj
	for inBounds(i, j, city) {
		antinodes = append(antinodes, Loc{i, j})
		i += di
		j += dj
	}

	return antinodes
}
