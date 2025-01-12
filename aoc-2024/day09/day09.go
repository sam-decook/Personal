package main

import (
	"fmt"
	"os"
	"slices"
)

func main() {
	fmt.Println(Part1("input.txt"))
}

// Allows me to work with each individual file/free space, and also
// modify the underlying blocks array when I need to move a file.
// I don't know why a normal slice did not work.
type Slice struct {
	Idx, Length int
}

// The files and free Slice slices are for pt2
func parse(file string) ([]int, []Slice, []Slice) {
	b, _ := os.ReadFile(file)

	blocks := make([]int, 0, 32)
	files := make([]Slice, 0, 32)
	free := make([]Slice, 0, 32)

	last := len(b) - 1 //disregard the trailing \n
	id := 0
	for i, n := range b[:last] {
		// ASCII to decimal
		n -= 48

		if i&1 == 0 { //it's a file
			files = append(files, Slice{Idx: len(blocks), Length: int(n)})
			for range n {
				blocks = append(blocks, id)
			}
			id++
		} else { //it's free space
			free = append(free, Slice{Idx: len(blocks), Length: int(n)})
			for range n {
				blocks = append(blocks, -1)
			}
		}
	}

	return blocks, files, free
}

// Compact the disk, whether or not the file is fragmented
func Part1(f string) int {
	blocks, _, _ := parse(f)

	// Pointer to next free block (-1)
	free := slices.Index(blocks, -1)
	// Pointer to next file block, starting last working forwards
	file := len(blocks) - 1

	// Put last file block into first free block. Then move each to the next
	for free < file {
		blocks[free], blocks[file] = blocks[file], blocks[free]
		for blocks[free] != -1 {
			free++
		}
		for blocks[file] == -1 {
			file--
		}
	}

	return checksum(blocks)
}

func checksum(blocks []int) int {
	cs := 0
	for i, n := range blocks {
		if n != -1 {
			cs += i + n
		}
	}
	return cs
}

// Compact the disk, but keep files contiguous
func Part2(file string) int {
	blocks, files, frees := parse(file)

	// Try to move each file
	for fileIdx := len(files) - 1; fileIdx >= 0; fileIdx-- {
		file := files[fileIdx]

		for i, free := range frees {
			if free.Idx < file.Idx && free.Length >= file.Length {
				// Copy the file into the free space
				for i := range file.Length {
					a := free.Idx + i
					b := file.Idx + i
					blocks[a], blocks[b] = blocks[b], blocks[a]
				}

				// Update the free space
				frees[i].Idx += file.Length
				frees[i].Length -= file.Length

				// Move to the next file
				break
			}
		}
	}

	return checksum(blocks)
}

// Returns indices such that the next free section is blocks[lo:hi]
func nextFree(blocks []int16, lo int) (int, int) {
	for blocks[lo] != -1 {
		lo++
		if lo == len(blocks) {
			return lo, lo
		}
	}

	hi := lo
	for blocks[hi] == -1 {
		hi++
	}

	return lo, hi
}

// Returns indices such that the previous file from the end is blocks[lo:hi]
func prevFile(blocks []int16, lo int) (int, int) {
	if lo == 0 {
		return 0, 0
	}

	hi := lo
	for blocks[hi-1] == -1 {
		hi--
		if hi == 0 {
			return 0, 0
		}
	}

	id := blocks[hi-1]

	lo = hi
	for blocks[lo-1] == id {
		lo--
	}

	return lo, hi
}
