package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
)

/* Implementing the HASHMAP for part2 */
type lens struct {
	label string
	focal int
}

type node struct {
	item lens
	next *node
}

// Walks the linked list to find the lens using its label
//
// Returns a pointer to the previous node to the lens,
// else the last node in the list if not found
func (n *node) find(label string) *node {
	switch {
	case n.next == nil:
		// End of list
		return n
	case n.next.item.label == label:
		// The lens has been foud
		return n
	default:
		// Lens not found
		return n.next.find(label)
	}
}

// Backed by an array of slices ~linked lists
type hashmap struct {
	array [256]node
	//array [256][]lens
}

func (h *hashmap) String() string {
	b := strings.Builder{}

	for i, list := range h.array {
		if list.item.label != "" {
			b.WriteString("Box " + fmt.Sprint(i) + ":")
			node := &list
			for node != nil {
				l := node.item.label
				f := fmt.Sprint(node.item.focal)
				b.WriteString(" [" + l + " " + f + "]")
				node = node.next
			}
			if node != &h.array[i] {
				b.WriteByte('\n')
			}
		}
	}

	return b.String()
}

func (h *hashmap) AddOrUpdate(label string, focal int) {
	i := hash([]byte(label))
	list := &h.array[i]

	// First bucket empty or a match
	if list.item.label == "" || list.item.label == label {
		h.array[i].item = lens{label, focal}
		return
	}

	// Search for place to insert
	prev := list.find(label)

	if prev.next != nil {
		// Not at the end of the list -> label found
		prev.next.item.focal = focal
	} else {
		// Label not found
		prev.next = &node{lens{label, focal}, nil}
	}
}

func (h *hashmap) Remove(label string) {
	i := hash([]byte(label))
	list := &h.array[i]

	// Check first slot
	if list.item.label == label {
		if list.next != nil {
			h.array[i] = *list.next
		} else {
			h.array[i].item = lens{}
		}
		return
	}

	// Get the node before the node with the lens being removed
	prev := list.find(label)

	// Label not in bucket, our work here is done
	if prev.next == nil {
		return
	}

	prev.next = prev.next.next
}

func main() {
	file := "input.txt"
	input, _ := os.Open(file)
	fmt.Println("Answer to part 1:", part1(input))

	input, _ = os.Open(file)
	fmt.Println("Answer to part 2:", part2(input))
}

// Looking at the bufio source helped a bunch
// This returns tokens separated by commas from a line
func scanCommas(data []byte, atEOF bool) (advance int, token []byte, err error) {
	// Check if we are done
	if atEOF && len(data) == 0 {
		return 0, nil, nil
	}

	// Return the last field before the line ends
	if atEOF {
		return len(data), data, nil
	}

	// Return the data up until the comma
	if i := bytes.IndexByte(data, ','); i >= 0 {
		return i + 1, data[0:i], nil

	}

	// Request more data
	return 0, nil, nil
}

// Sums the hashes of each token
func part1(input *os.File) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(scanCommas)

	sum := 0
	for scanner.Scan() {
		sum += hash(scanner.Bytes())
	}

	return sum
}

func hash(b []byte) int {
	h := 0
	for _, b := range b {
		h += int(b)
		h *= 17
		h %= 256
	}
	return h
}

// Parses a token into the label, operation, and lens focal length (if op == '=')
func parse(token []byte) (string, byte, int) {
	if eq := bytes.Index(token, []byte("=")); eq != -1 {
		n, _ := strconv.Atoi(string(token[eq+1]))
		return string(token[:eq]), token[eq], n
	}
	if dash := bytes.Index(token, []byte("-")); dash != -1 {
		return string(token[:dash]), token[dash], 0
	}
	panic("Error: token does not contain '=' or '-'")
}

// 773844 too high
func part2(input *os.File) int {
	scanner := bufio.NewScanner(input)
	scanner.Split(scanCommas)

	var hm hashmap

	for scanner.Scan() {
		label, op, focal := parse(scanner.Bytes())
		switch op {
		case '=':
			hm.AddOrUpdate(label, focal)
		case '-':
			hm.Remove(label)
		}
	}

	sum := 0
	for i, list := range hm.array {
		if list.item.label != "" {
			node := &list
			j := 1
			for node != nil {
				sum += (i + 1) * j * node.item.focal
				node = node.next
				j++
			}
		}
	}
	return sum
}
