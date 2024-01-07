package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type card struct {
	Id      int
	Numbers []string
	Winning map[string]bool
}

func main() {
	input, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(input)

	cards := make([]card, 8)
	for scanner.Scan() {
		cards = append(cards, parse(scanner.Text()))
	}

	sumPt1 := 0
	for _, c := range cards {
		sumPt1 += pt1(c)
	}

	sumPt2 := pt2(cards)

	fmt.Println("Answer to part 1:", sumPt1)
	fmt.Println("Answer to part 2:", sumPt2)
}

// Format of a line:
// Game #: num ... num | num ... num
func parse(line string) card {
	a := strings.Split(line, ":")
	b := strings.Split(a[1], " | ")

	id, _ := strconv.Atoi(strings.Fields(a[0])[1])
	winning := strings.Fields(b[0])
	numbers := strings.Fields(b[1])

	// Use a map for O(1) lookup
	winningMap := make(map[string]bool, len(winning))
	for i := 0; i < len(winning); i++ {
		winningMap[winning[i]] = true
	}

	return card{id, numbers, winningMap}
}

// A card is worth 1 point for one matched number, and doubles each additional
// number matches
func pt1(c card) int {
	matches := getMatches(c)

	points := 0
	if matches > 0 {
		points = 1 << (matches - 1)
	}

	return points
}

// Counts the number of winning numbers on a card
func getMatches(c card) int {
	matches := 0
	for _, number := range c.Numbers {
		if c.Winning[number] {
			matches++
		}
	}
	return matches
}

func pt2(cards []card) int {
	// Hold the amount of each card you have (idx -> Card #{idx+1})
	// You start with one of each card
	count := make([]int, len(cards))
	for i := 0; i < len(count); i++ {
		count[i] = 1
	}

	// For each card, you win {matches} more cards for the next {matches} cards
	// ^ that's a confusing statement, but it's accurate
	for _, c := range cards {
		matches := getMatches(c)

		if matches > 0 {
			for i := c.Id; i < c.Id+matches; i++ {
				count[i] += count[c.Id-1]
			}
		}
	}

	// Sum the amount of cards
	total := 0
	for i := 0; i < len(count); i++ {
		total += count[i]
	}
	return total
}
