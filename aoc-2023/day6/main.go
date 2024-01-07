package main

import (
	"fmt"
)

// Holds the duration of a race and the best distance ever recorded
type race struct {
	Duration,
	Furthest int
}

var test1 []race = []race{
	{7, 9},
	{15, 40},
	{30, 200},
}

var input1 []race = []race{
	{42, 308},
	{89, 1170},
	{91, 1291},
	{89, 1467},
}

var test2 race = race{71530, 940200}

var input2 race = race{42899189, 308117012911467}

func main() {
	pt1 := 1
	for _, race := range input1 {
		pt1 *= numberFaster(race)
	}

	// Brute force works b/c each loop is cheap, likely a few instructions
	pt2 := numberFaster(input2)

	fmt.Println("Answer to part 1:", pt1)
	fmt.Println("Answer to part 2:", pt2)
}

// Each ms a boat is held, its speed increases by 1 mm/ms
// -> the time it is held == its speed
// After it is released, it travels that speed for the rest of the race
func calculateDistance(holdTime, raceDuration int) int {
	return holdTime * (raceDuration - holdTime)
}

// Finds the amount of ways to win a race, r
func numberFaster(r race) int {
	faster := 0
	// Exclude the end cases, they travel 0 mm
	for holdTime := 1; holdTime < r.Duration-1; holdTime++ {
		distance := calculateDistance(holdTime, r.Duration)
		if distance > r.Furthest {
			faster++
		}
	}
	return faster
}
