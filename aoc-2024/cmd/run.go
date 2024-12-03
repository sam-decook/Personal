package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// Flags
var (
	part  int
	input string
)

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs the code for a day",
	Long: `Runs the code for a problem in Advent of Code.
	Chooses the current day, or you can specify which or run all available.
	Runs test and full input for part 1 and 2, or you can specify.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("run called")
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	runCmd.Flags().IntVarP(&part, "part", "p", 0, "Run part 1 or 2; omit for both")
	runCmd.Flags().StringVarP(&input, "input", "i", "", "Use test or full input; omit for both")
}

// fmt.Println("Day 1 results:")
// fmt.Printf("             Test      Full\n")
// fmt.Printf("Part 1 %10d%10d\n", day01.Part1("test.txt"), day01.Part1("input.txt"))
// fmt.Printf("Part 2 %10d%10d\n", day01.Part2("test.txt"), day01.Part2("input.txt"))
