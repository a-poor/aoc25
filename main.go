package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	puzzle := readInput()
	beamLine := makeStartLine(puzzle.w, puzzle.start)
	var splitCount int
	for _, s := range puzzle.splitters {
		next, n := stepBeams(beamLine, s)
		splitCount += n
		beamLine = next
	}
	fmt.Printf("Split count: %d\n", splitCount)
}

type puzzle struct {
	w, h      int      // Grid size
	start     int      // x-pos of start on row=1
	splitters [][]bool // Beam splitters
}

func readInput() puzzle {
	s := bufio.NewScanner(os.Stdin)

	start := -1
	var w int
	var splitters [][]bool

	for s.Scan() {
		// Read the next line
		line := s.Text()

		// Is it the first line?
		// (No splitters)
		if start == -1 {
			start = strings.IndexRune(line, 'S')
			w = len([]rune(line))
			continue
		}

		// Make an empty line
		sl := make([]bool, w)
		for i, r := range []rune(line) {
			if r == '^' {
				sl[i] = true
			}
		}
		splitters = append(splitters, sl)
	}
	if err := s.Err(); err != nil {
		panic(err)
	}
	return puzzle{
		w:         w,
		h:         len(splitters),
		start:     start,
		splitters: splitters,
	}
}

func makeStartLine(w, s int) []bool {
	line := make([]bool, w)
	line[s] = true
	return line
}

func stepBeams(current, splits []bool) ([]bool, int) {
	var count int
	next := make([]bool, len(current))
	for i, b := range current {
		// No beam? No-op
		if !b {
			continue
		}

		// If there isn't a splitter, it just moves
		// one down
		if !splits[i] {
			next[i] = true
			continue
		}

		// Otherwise, we split
		count++
		// ...split left
		if i > 0 {
			next[i-1] = true
		}
		// ...split right
		if i < len(next)-1 {
			next[i+1] = true
		}
	}
	return next, count
}
