package main

import (
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type line struct {
	low,
	high,
	diff int
}

func main() {
	input, _ := os.ReadFile("input.txt")
	// I should probably use a scanner and build the groups "manually"
	groups := strings.Split(string(input), "\r\n\r\n")

	// Get the seeds and make the table to hold each category
	seedLine := groups[0]
	seeds := strings.Fields(seedLine)[1:]

	to := make([][8]int, len(seeds))
	for i := range to {
		n, _ := strconv.Atoi(seeds[i])
		to[i] = [8]int{n}
	}

	// Parse out the maps
	maps := make([][]line, 0, 7)
	for _, group := range groups[1:] {
		lines := strings.Split(group, "\r\n")
		m := make([]line, 0)
		for i := 1; i < len(lines); i++ {
			m = append(m, parse(lines[i]))
		}
		maps = append(maps, m)
	}

	// Loop through each map (moving from one category to another)
	// Then loop through each seed
	// Finally, loop through each line in the category map
	// If the number wasn't mapped, carry it over from the previous
	// - for this reason, you need to loop through the seeds before the map
	for j := range maps {
		for i := range to {
			for _, m := range maps[j] {
				if to[i][j] >= m.low && to[i][j] < m.high {
					to[i][j+1] = to[i][j] + m.diff
				}
			}
			if to[i][j+1] == 0 {
				to[i][j+1] = to[i][j]
			}
		}
	}

	minLoc := math.MaxInt
	for i := range to {
		if to[i][7] < minLoc {
			minLoc = to[i][7]
		}
	}

	sumPt2 := 0

	fmt.Println("Answer to part 1:", minLoc)
	fmt.Println("Answer to part 2:", sumPt2)
}

func parse(str string) line {
	nums := strings.Fields(str)
	dst, _ := strconv.Atoi(nums[0])
	src, _ := strconv.Atoi(nums[1])
	amount, _ := strconv.Atoi(nums[2])

	hi := src + amount
	diff := dst - src
	return line{src, hi, diff}
}
