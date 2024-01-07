package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const HAND_SIZE = 5

type HandType int

func (t HandType) String() string {
	symbol := ""
	switch t {
	case FIVE_OF_A_KIND:
		symbol = "Five of a kind"
	case FOUR_OF_A_KIND:
		symbol = "Four of a kind"
	case FULL_HOUSE:
		symbol = "Full house"
	case THREE_OF_A_KIND:
		symbol = "Three of a kind"
	case TWO_PAIR:
		symbol = "Two pair"
	case ONE_PAIR:
		symbol = "One pair"
	case HIGH_CARD:
		symbol = "High card"
	}
	return symbol
}

const (
	TWO = iota
	THREE
	FOUR
	FIVE
	SIX
	SEVEN
	EIGHT
	NINE
	TEN
	JACK
	QUEEN
	KING
	ACE
)

type card int

const (
	HIGH_CARD = iota
	ONE_PAIR
	TWO_PAIR
	THREE_OF_A_KIND
	FULL_HOUSE
	FOUR_OF_A_KIND
	FIVE_OF_A_KIND
)

func (c card) String() string {
	symbol := ""
	switch c {
	case TWO:
		symbol = "2"
	case THREE:
		symbol = "3"
	case FOUR:
		symbol = "4"
	case FIVE:
		symbol = "5"
	case SIX:
		symbol = "6"
	case SEVEN:
		symbol = "7"
	case EIGHT:
		symbol = "8"
	case NINE:
		symbol = "9"
	case TEN:
		symbol = "T"
	case JACK:
		symbol = "J"
	case QUEEN:
		symbol = "Q"
	case KING:
		symbol = "K"
	case ACE:
		symbol = "A"
	}
	return symbol
}

type hand struct {
	Cards [HAND_SIZE]card
	Count map[card]int
	Type  HandType
	Bid   int
}

func (h *hand) String() string {
	s := h.Type.String() + ": "
	for _, card := range h.Cards {
		s += card.String()
	}
	return s
}

func (left *hand) gt(right *hand) bool {
	if left.Type > right.Type {
		return true
	}
	if left.Type < right.Type {
		return false
	}

	// If they are the same, hand with higher card (left to right) is stronger
	for i := 0; i < HAND_SIZE; i++ {
		diff := left.Cards[i] - right.Cards[i]
		switch {
		case diff > 0:
			return true
		case diff < 0:
			return false
		}
	}

	return false
}

func (left *hand) eq(right *hand) bool {
	equal := true
	for i := 0; i < HAND_SIZE; i++ {
		if left.Cards[i] != right.Cards[i] {
			equal = false
		}
	}
	return equal
}

func (left *hand) gte(right *hand) bool {
	return left.eq(right) || left.gt(right)
}

func (left *hand) lte(right *hand) bool {
	return left.eq(right) || !left.gt(right)
}

func main() {
	input, _ := os.Open("input.txt")

	hands := parse(input)

	orderHands(hands)

	sumPt1 := 0
	for i := range hands {
		rank := i + 1
		sumPt1 += hands[i].Bid * rank
	}

	sumPt2 := 0

	fmt.Println("Answer to part 1:", sumPt1)
	fmt.Println("Answer to part 2:", sumPt2)
}

func parse(input *os.File) []hand {
	scanner := bufio.NewScanner(input)

	hands := make([]hand, 0)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)

		// Parse each card into a hand
		cards := parseHand(fields[0])

		// Count the amount of each card
		count := make(map[card]int)
		for i := range cards {
			count[cards[i]]++
		}

		// Get the type of the hand
		handType := findType(count)

		// Parse the bid amount
		bid, _ := strconv.Atoi(fields[1])

		hands = append(hands, hand{cards, count, handType, bid})
	}

	return hands
}

func parseHand(input string) [5]card {
	cards := [5]card{}
	for i := range input {
		// It's gross, but the 0th index of an untyped string constant is a byte
		switch input[i] {
		case '2':
			cards[i] = TWO
		case '3':
			cards[i] = THREE
		case '4':
			cards[i] = FOUR
		case '5':
			cards[i] = FIVE
		case '6':
			cards[i] = SIX
		case '7':
			cards[i] = SEVEN
		case '8':
			cards[i] = EIGHT
		case '9':
			cards[i] = NINE
		case 'T':
			cards[i] = TEN
		case 'J':
			cards[i] = JACK
		case 'Q':
			cards[i] = QUEEN
		case 'K':
			cards[i] = KING
		case 'A':
			cards[i] = ACE
		}
	}
	return cards
}

func findType(count map[card]int) HandType {
	if len(count) == 1 {
		return FIVE_OF_A_KIND
	}

	four := false
	three := false
	two := 0

	for _, n := range count {
		switch n {
		case 4:
			four = true
		case 3:
			three = true
		case 2:
			two++
		}
	}

	switch {
	case four:
		return FOUR_OF_A_KIND
	case two > 0 && three:
		return FULL_HOUSE
	case three:
		return THREE_OF_A_KIND
	case two == 2:
		return TWO_PAIR
	case two == 1:
		return ONE_PAIR
	default:
		return HIGH_CARD
	}
}

// Quicksort
func orderHands(hands []hand) {
	if len(hands) > 1 {
		mid := partition(hands)
		orderHands(hands[:mid])
		orderHands(hands[mid+1:])
	}
}

func partition(hands []hand) int {
	lo := 0
	hi := len(hands) - 1

	// Use the middle element as the pivot
	hands[hi/2], hands[hi] = hands[hi], hands[hi/2]
	pivot := len(hands) - 1

	// Advance two pointers from either side until both elements need to be swapped
	for lo <= hi {
		for lo <= hi && hands[hi].gte(&hands[pivot]) {
			hi -= 1
		}
		for lo <= hi && hands[lo].lte(&hands[pivot]) {
			lo += 1
		}

		if lo < hi {
			hands[lo], hands[hi] = hands[hi], hands[lo]
		}
	}

	// Return pivot to middle
	hands[lo], hands[pivot] = hands[pivot], hands[lo]
	return lo
}
