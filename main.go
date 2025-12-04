package main

import (
	"bufio"
	"fmt"
	"os"
)

const ndigits = 12

var (
	North     = point{x: 0, y: -1}
	NorthWest = point{x: -1, y: -1}
	NorthEast = point{x: 1, y: -1}
	South     = point{x: 0, y: 1}
	SouthWest = point{x: -1, y: 1}
	SouthEast = point{x: 1, y: 1}
	East      = point{x: 1, y: 0}
	West      = point{x: -1, y: 0}
)

func main() {
	var grid [][]bool

	// Scan the input
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		grid = append(grid, []bool{})
		n := len(grid)

		// Get the next line
		ln := s.Text()
		for _, c := range []rune(ln) {
			switch c {
			case '.':
				grid[n-1] = append(grid[n-1], false)
			case '@':
				grid[n-1] = append(grid[n-1], true)
			default:
				panic("Unknown character: '" + string(c) + "'")
			}
		}
	}
	if err := s.Err(); err != nil {
		panic(err)
	}

	// Go back through and count
	var total int

	// We'll need this to track what's
	// in and out of bounds
	h, w := len(grid), len(grid[0])

	for i, row := range grid {
		for j, cell := range row {
			// If it isn't toilet paper, skip it
			if !cell {
				continue
			}

			// Otherwise, get it's neighbors
			p := point{x: j, y: i}
			ns := neighbors(p, w, h)

			// And count how many have rolls
			var nrc int
			for _, n := range ns {
				if grid[n.y][n.x] {
					nrc++
				}
			}

			// If the neighbor-row-count is less
			// than 4, it can be collected
			if nrc < 4 {
				total++
			}
		}
	}

	// Done!
	fmt.Printf("total: %d\n", total)
}

type point struct{ x, y int }

func (p point) add(o point) point {
	return point{
		x: p.x + o.x,
		y: p.y + o.y,
	}
}

func (p point) in(w, h int) bool {
	if p.x < 0 {
		return false
	}
	if p.y < 0 {
		return false
	}
	if p.x >= w {
		return false
	}
	if p.y >= h {
		return false
	}
	return true
}

func neighbors(p point, w, h int) []point {
	var ns []point
	if n := p.add(North); n.in(w, h) {
		ns = append(ns, n)
	}
	if n := p.add(NorthWest); n.in(w, h) {
		ns = append(ns, n)
	}
	if n := p.add(NorthEast); n.in(w, h) {
		ns = append(ns, n)
	}
	if n := p.add(South); n.in(w, h) {
		ns = append(ns, n)
	}
	if n := p.add(SouthWest); n.in(w, h) {
		ns = append(ns, n)
	}
	if n := p.add(SouthEast); n.in(w, h) {
		ns = append(ns, n)
	}
	if n := p.add(East); n.in(w, h) {
		ns = append(ns, n)
	}
	if n := p.add(West); n.in(w, h) {
		ns = append(ns, n)
	}
	return ns
}
