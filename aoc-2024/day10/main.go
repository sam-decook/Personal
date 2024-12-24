package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println(Part1("test.txt"))
}

type Loc struct {
	i, j int
}

func parse(file string) ([][]byte, []Loc) {
	f, _ := os.Open(file)
	scanner := bufio.NewScanner(f)

	trails := make([][]byte, 0, 16)
	heads := make([]Loc, 0, 16)
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Bytes()
		row := make([]byte, len(line))
		for j := range line {
			if line[j] == '0' {
				heads = append(heads, Loc{i, j})
			}
			if line[j] >= '0' && line[j] <= '9' {
				row[j] = line[j] - '0'
			} else {
				row[j] = 128
			}
		}
		trails = append(trails, row)
	}

	return trails, heads
}

func Part1(file string) int {
	trails, heads := parse(file)

	score := 0
	for _, head := range heads {
		score += scoreOf(head, 0, trails)
		fmt.Println()
	}

	return score
}

// Returns the amount of trails that rise from 0-9 in increments of 1.
func scoreOf(loc Loc, height int, trails [][]byte) int {
	fmt.Printf("%d,%d ", loc.i, loc.j)
	if height == 9 {
		return 1
	}

	score := 0
	start := trails[loc.i][loc.j]
	if i, j := loc.i-1, loc.j; i >= 0 && trails[i][j]-start == 1 {
		score += scoreOf(Loc{i, j}, height+1, trails)
	}
	if i, j := loc.i, loc.j-1; j >= 0 && trails[i][j]-start == 1 {
		score += scoreOf(Loc{i, j}, height+1, trails)
	}
	if i, j := loc.i+1, loc.j; i < len(trails) && trails[i][j]-start == 1 {
		score += scoreOf(Loc{i, j}, height+1, trails)
	}
	if i, j := loc.i, loc.j+1; j < len(trails[0]) && trails[i][j]-start == 1 {
		score += scoreOf(Loc{i, j}, height+1, trails)
	}

	return score
}
