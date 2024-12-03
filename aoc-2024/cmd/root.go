package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

// Flags
var (
	all bool
	day int
)

var rootCmd = &cobra.Command{
	Use:   "aoc-2024",
	Short: "A CLI for helpful Advent of Code automations",
	Long: `A CLI to generate your boilerplate and run your code.
    You can scope either action to a day or all available days.
    Omit to run the current day, if you are in December 2024.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&all, "all", "a", false, "Init all days available")
	rootCmd.PersistentFlags().IntVarP(&day, "day", "d", 0, "Choose a specific day")
	rootCmd.MarkFlagsMutuallyExclusive("all", "day")
}
