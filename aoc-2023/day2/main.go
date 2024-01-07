package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// The puzzle revolves around the classic game of picking marbles from a jar
// For each game, the elf grabs some marbles (colored cubes) out of the bag
// many times (colors separated by commas, sets separated by semi-colons)
//
// In part 1, we find the games that are possible if there were 12
//   red, 13 green, and 14 blue cubes. Then we sum the IDs together.
// In part 2, we find the smallest number of cubes that would make the game
//   possible. The we sum the powers (the product of the amount of each cube)
//   of all games together.

type cubes struct {
	Red   int
	Green int
	Blue  int
}

// Updates the fewest required cubes
func (c *cubes) updateFewest(f cubes) {
	if f.Red > c.Red {
		c.Red = f.Red
	}
	if f.Green > c.Green {
		c.Green = f.Green
	}
	if f.Blue > c.Blue {
		c.Blue = f.Blue
	}
}

func main() {
	input, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(input)

	sumPt1 := 0
	sumPt2 := 0
	for scanner.Scan() {
		line := scanner.Text()

		sumPt1 += pt1(line)
		sumPt2 += pt2(line)
	}

	fmt.Println("Answer to part 1:", sumPt1)
	fmt.Println("Answer to part 2:", sumPt2)
}

// Returns the ID of the game if it was valid, which means:
// the amount of cubes shown in a set never exceeds the max (12 R, 13 G, 14 B)
func pt1(line string) int {
	max := cubes{12, 13, 14}

	id, sets := parse(line)
	for _, set := range sets {
		if set.Red > max.Red || set.Green > max.Green || set.Blue > max.Blue {
			return 0
		}
	}

	return id
}

// Returns the power of the game, which is the product of the fewest cubes of
// each color
func pt2(line string) int {
	fewest := cubes{}

	_, sets := parse(line)
	for _, set := range sets {
		fewest.updateFewest(set)
	}

	return fewest.Red * fewest.Green * fewest.Blue
}

// Returns the ID and a slice of the cubes revealed in each set
func parse(line string) (int, []cubes) {
	// Preprocessing to make splitting cleaner
	line = strings.Replace(line, ": ", ":", 1)
	line = strings.Replace(line, "; ", ";", -1)
	line = strings.Replace(line, ", ", ",", -1)

	// First divide the label from the information formatted: "Game n: sets"
	a := strings.Split(line, ":")
	label := a[0]
	sets := a[1]

	// Get the ID
	sId := strings.Split(label, " ")[1]
	id, _ := strconv.Atoi(sId)

	// Make an array of the sets of cubes revealed
	parsed := make([]cubes, 0)
	for _, set := range strings.Split(sets, ";") {
		cSet := parseSet(strings.Split(set, ","))
		parsed = append(parsed, cSet)
	}

	return id, parsed
}

// Takes a set (semi-colon delimited) and counts the cubes revealed
func parseSet(set []string) cubes {
	shown := cubes{}
	for _, grab := range set {
		// Each grab is formatted: "number color"
		a := strings.Split(grab, " ")
		n, _ := strconv.Atoi(a[0])
		switch a[1] {
		case "red":
			shown.Red = n
		case "green":
			shown.Green = n
		case "blue":
			shown.Blue = n
		}
	}
	return shown
}
