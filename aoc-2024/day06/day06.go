package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(Part2("input.txt"))
}

type Guard struct {
	i, j, iStep, jStep int
}

func (g *Guard) Move(lab []string) {
	nextI := g.i + g.iStep
	nextJ := g.j + g.jStep

	if inBounds(nextI, nextJ, lab) && lab[nextI][nextJ] == '#' {
		g.Turn()
	} else {
		g.i = g.i + g.iStep
		g.j = g.j + g.jStep
	}
}

// U: -1  0
// R:  0  1
// D:  1  0
// L:  0 -1
func (g *Guard) Turn() {
	g.iStep, g.jStep = g.jStep, -g.iStep
}

func parse(file string) (Guard, []string) {
	f, _ := os.Open(file)
	scanner := bufio.NewScanner(f)

	lab := make([]string, 0, 8)
	guard := Guard{iStep: -1, jStep: 0}

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		lab = append(lab, line)

		if idx := strings.Index(line, "^"); idx != -1 {
			guard.i = i
			guard.j = idx
		}
	}

	return guard, lab
}

func inBounds(i, j int, lab []string) bool {
	return i >= 0 &&
		j >= 0 &&
		i < len(lab) &&
		j < len(lab[0])
}

func Part1(file string) int {
	guard, lab := parse(file)

	// Keep track of visited squares
	visited := make(map[Loc]bool)

	for inBounds(guard.i, guard.j, lab) {
		visited[Loc{guard.i, guard.j}] = true
		guard.Move(lab)
	}

	return len(visited)
}

type Loc struct {
	i, j int
}

func visited(guard Guard, lab []string) map[Loc]bool {
	visited := make(map[Loc]bool)

	for inBounds(guard.i, guard.j, lab) {
		visited[Loc{guard.i, guard.j}] = true
		guard.Move(lab)
	}

	return visited
}

// Run it normally and get all of the positions.
// Add a block to each position and try it out.
func Part2(file string) int {
	guard, lab := parse(file)

	// Get all locations (except starting)
	locs := visited(guard, lab)
	delete(locs, Loc{guard.i, guard.j})

	// Try adding a block at each location
	loops := 0
	for loc := range locs {
		// Add a block
		orig := lab[loc.i]
		lab[loc.i] = orig[:loc.j] + "#" + orig[loc.j+1:]

		if loop(Guard{guard.i, guard.j, guard.iStep, guard.jStep}, lab) {
			loops++
		}

		// Remove the block
		lab[loc.i] = orig
	}

	return loops
}

// Some guy on reddit said there is a loop if the guard was in the same
// position and direction. Also to make `Move()` only move or turn, not both.
// Don't know why the fast/slow algorithm didn't work.
func loop(guard Guard, lab []string) bool {
	locs := make(map[Guard]bool)
	locs[guard] = true

	for inBounds(guard.i, guard.j, lab) {
		guard.Move(lab)

		if locs[guard] {
			return true
		} else {
			locs[guard] = true
		}
	}

	return false
}

// func loop(guard Guard, lab []string) bool {
// 	// Initialize new fast and slow guards (pointers) for loop detection
// 	slow := guard
// 	fast := guard

// 	for inBounds(fast.i, fast.j, lab) {
// 		fast.Move(lab)
// 		slow.Move(lab)

// 		if !inBounds(fast.i, fast.j, lab) {
// 			return false
// 		}
// 		fast.Move(lab)

// 		if fast.i == slow.i && fast.j == slow.j {
// 			return true
// 		}
// 	}

// 	return false
// }
