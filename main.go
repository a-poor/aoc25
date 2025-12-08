package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"unicode"
)

type operation rune

const (
	Mul operation = '*'
	Add operation = '+'
	Div operation = '/'
	Sub operation = '-'
)

func main() {
	var total int
	nrs, ors := readInput()

	// Get the grid width and height
	// for re-use
	w, h := len(ors), len(nrs)

	// Track the numbers in the current equation
	var ns []int

	// Loop across the grid right-to-left and
	// top-to-bottom
	for i := w - 1; i >= 0; i-- {
		// Start a text number for the column
		var nr []rune

		// Loop top-to-bottom
		for j := 0; j < h; j++ {
			if r := nrs[j][i]; unicode.IsDigit(r) {
				nr = append(nr, r)
			}
		}

		// Empty column? Skip it
		if len(nr) == 0 {
			continue
		}

		// We got a number! Convert it and store it
		sr := string(nr)
		n, err := strconv.Atoi(sr)
		if err != nil {
			panic(fmt.Errorf("failed to parse %q (col=%d) as a number", sr, i))
		}
		ns = append(ns, n)

		// Check if there's an operation in
		// that column
		if o := ors[i]; o == '+' || o == '*' {
			total += equation{
				ns: ns,
				op: operation(o),
			}.solve()
			ns = make([]int, 0)
		}
	}

	fmt.Printf("total = %d\n", total)
}

type equation struct {
	ns []int
	op operation
}

func (e equation) solve() int {
	n := e.ns[0]
	for i := 1; i < len(e.ns); i++ {
		switch e.op {
		case Mul:
			n *= e.ns[i]
		case Add:
			n += e.ns[i]
		case Div:
			n /= e.ns[i]
		case Sub:
			n -= e.ns[i]
		default:
			panic(fmt.Sprintf("Unknown rune %q", string(e.op)))
		}
	}
	return n
}

func readInput() ([][]rune, []rune) {
	s := bufio.NewScanner(os.Stdin)

	var nrs [][]rune
	var ors []rune

	for s.Scan() {
		// Read the next line
		line := s.Text()

		// Handle different line cases
		switch {
		case isNumberLine(line):
			nrs = append(nrs, []rune(line))

		default:
			ors = []rune(line)
		}
	}
	if err := s.Err(); err != nil {
		panic(err)
	}
	return nrs, ors
}

func isNumberLine(ln string) bool {
	return !slices.Contains([]rune(ln), '+')
}
