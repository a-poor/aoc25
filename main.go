package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var mem = map[point]int{}

func main() {
	puzzle := readInput()
	count := countPaths(puzzle.splitters, point{y: 0, x: puzzle.start})
	// fmt.Printf("mem=%+v\n", mem)
	fmt.Printf("count: %d\n", count)
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

type point struct{ x, y int }

func (p point) down() point {
	return point{x: p.x, y: p.y + 1}
}

func (p point) left() point {
	return point{x: p.x - 1, y: p.y}
}

func (p point) right() point {
	return point{x: p.x + 1, y: p.y}
}

func countPaths(grid [][]bool, p point) int {
	if n, ok := mem[p]; ok {
		return n
	}
	n := _countPaths(grid, p)
	mem[p] = n
	return n
}

func _countPaths(grid [][]bool, p point) int {
	// If we're on the last line, we're done
	// (there aren no splitters on the last line)
	if p.y >= len(grid)-1 {
		return 1
	}

	// Otherwise we need to step the beam
	p2 := p.down()

	// If it didn't land on a splitter, continue
	if !grid[p2.y][p2.x] {
		return countPaths(grid, p2)
	}

	// Otherwise, split time
	return countPaths(grid, p2.left()) + countPaths(grid, p2.right())
}
