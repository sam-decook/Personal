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

func parse(file string) [][]byte {
	f, _ := os.Open(file)
	scanner := bufio.NewScanner(f)

	board := make([][]byte, 0, 16)
	for scanner.Scan() {
		sBytes := scanner.Bytes()
		cBytes := make([]byte, len(sBytes))
		copy(cBytes, sBytes)
		board = append(board, cBytes)
	}
	return board
}

// 2745 too high
// 2727 too high
func Part1(file string) int {
	board := parse(file)

	sum := 0
	lookUp := false
	lookDown := true
	for i := range board {
		if i == 3 {
			lookUp = true
		}
		if i+3 == len(board) {
			lookDown = false
		}

		lookLeft := false
		lookRight := true
		for j := range board[0] {
			if j == 3 {
				lookLeft = true
			}
			if j+3 == len(board[i]) {
				lookRight = false
			}
			if board[i][j] != 'X' {
				continue
			}

			old := sum
			if lookLeft && check(i, j, STRAIGHT, LEFT, board) {
				sum++
			}
			if lookUp && lookLeft && check(i, j, UP, LEFT, board) {
				sum++
			}
			if lookUp && check(i, j, UP, STRAIGHT, board) {
				sum++
			}
			if lookUp && lookRight && check(i, j, UP, RIGHT, board) {
				sum++
			}
			if lookRight && check(i, j, STRAIGHT, RIGHT, board) {
				sum++
			}
			if lookDown && lookRight && check(i, j, DOWN, RIGHT, board) {
				sum++
			}
			if lookDown && check(i, j, DOWN, STRAIGHT, board) {
				sum++
			}
			if lookDown && lookLeft && check(i, j, DOWN, LEFT, board) {
				sum++
			}

			if sum-old == 5 {
				printAround(sum-old, i, j, board)
			}
			// if sum > old {
			// 	printAround(sum-old, i, j, board)
			// }
		}
	}

	return sum
}

func printAround(matches, x, y int, board [][]byte) {
	iStart := max(0, x-3)
	iEnd := min(x+4, len(board))
	jStart := max(0, y-3)
	jEnd := min(y+4, len(board[0]))

	for i := iStart; i < iEnd; i++ {
		for j := jStart; j < jEnd; j++ {
			fmt.Printf("%s ", string(board[i][j]))
		}
		if x == i {
			fmt.Print("<")
		}
		fmt.Println()
	}
	fmt.Println(strings.Repeat("  ", y-jStart) + "^")
	fmt.Printf("Matches at %d,%d: %d\n", x, y, matches)
	fmt.Scanln()
}

const (
	STRAIGHT = 0
	UP       = -1
	DOWN     = 1
	LEFT     = -1
	RIGHT    = 1
)

func check(i, j, down, right int, board [][]byte) bool {
	return board[i+(1*down)][j+(1*right)] == 'M' &&
		board[i+(2*down)][j+(2*right)] == 'A' &&
		board[i+(3*down)][j+(3*right)] == 'S'
}

func Part2(file string) int {
	board := parse(file)

	sum := 0
	for i := 1; i < len(board)-1; i++ {
		for j := 1; j < len(board[0])-1; j++ {
			if board[i][j] == 'A' && check2(i, j, board) {
				sum++
			}
		}
	}

	return sum
}

// Looks for two `mas`'s in an X, in any order
// [MS] . [MS]
//
//	.  X   .
//
// [MS] . [MS]
func check2(i, j int, board [][]byte) bool {
	downLeft := board[i-1][j-1] == 'M' && board[i+1][j+1] == 'S' ||
		board[i-1][j-1] == 'S' && board[i+1][j+1] == 'M'

	upRight := board[i+1][j-1] == 'M' && board[i-1][j+1] == 'S' ||
		board[i+1][j-1] == 'S' && board[i-1][j+1] == 'M'

	return downLeft && upRight
}
