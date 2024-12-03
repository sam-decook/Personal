package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// Flags
var (
	initAll bool
	initDay int
)

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Initializes a folder with boilerplate",
	Long: `Inializes a folder with boilerplate.
	You can specify a day, or do all at once.
	Otherwise, it will just do the current day if in Dec 2024.`,
	Run: initFunc,
}

func init() {
	rootCmd.AddCommand(initCmd)
}

func initFunc(cmd *cobra.Command, args []string) {
	// Set the name to dayxx, single digits padded with zero
	name := fmt.Sprintf("day%02d", day)
	file := fmt.Sprintf("/%s/%s.go", name, name)

	// TODO: doesn't work
	_, err := os.Open(file)
	if err == nil {
		fmt.Println(file + " already exists.")
		os.Exit(1)
	}

	os.Mkdir(name, 0755)
	os.Chdir(name)
	os.Create("test.txt")
	os.Create("input.txt")
	f, err := os.Create(name + ".go")
	f.Chmod(0755)

	fmt.Fprintf(f, "package %s\n%s", name, template)
}

const template = `
import (
	"bufio"
	"os"
)

func Part1(file string) int {
	f, _ := os.Open(file)
	scanner := bufio.NewScanner(f)

	return 0
}

func Part2(file string) int {
	f, _ := os.Open(file)
	scanner := bufio.NewScanner(f)

	return 0
}
`
