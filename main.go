package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

const (
	IterCount = 1_000
	// IterCount = 10
	MaxInt = int((^uint(0) >> 1))
)

func main() {
	// Read in the points
	ps := readInput()

	// Track the wires between points
	wires := map[wire]bool{}

	// Doo the looping
	for it := range IterCount {
		// Create an array to track point distances
		bestni, bestnj, bestnd := -1, -1, MaxInt
		// neighbors := make([]struct{
		// 	ni int // neighbor's index
		// 	nd int // neighbor's distance
		// }, len(ps))

		for i, p1 := range ps {
			// Init best point's index and max distance
			bestj, bestd := -1, MaxInt

			// Compare against other points
			for j, p2 := range ps {
				// If it's the same point, skip
				if i == j {
					continue
				}

				// // If they're in the same group, skip
				// if p1.g == p2.g {
				// 	continue
				// }
				//
				// Actually...
				//
				// If they're already connected, skip
				if wires[makeWire(p1.i, p2.i)] {
					continue
				}

				// Otherwise, get the distance to the
				// other point
				d := p1.dist(p2)

				// If it's greater than the current
				// best, then skip
				if d > bestd {
					continue
				}

				// Otherwise, we have a new winner
				bestj = j
				bestd = d
			}

			// Is it further than the previous best?
			if bestd >= bestnd {
				continue
			}

			// Is it valid
			if bestj == -1 {
				panicf("How did we get here?! %d -> %d == %d", i, bestj, bestd)
			}

			// Otherwise we have a new winner!
			bestni = i
			bestnj = bestj
			bestnd = bestd
		}

		// Ensure one was found
		if bestni == -1 {
			panicf("No match found! it=%d bestni=%d bestnj=%d, bestnd=%d", it, bestni, bestnj, bestnd)
			continue
		}

		// Now that we've found the closest two
		// (not already grouped) points, we can
		// merge those two groups
		//
		// fmt.Printf("Merging group %d and group %d\n", ps[bestni].g, ps[bestnj].g)
		ps = groupPoints(ps, ps[bestni].g, ps[bestnj].g)

		// And add a wire
		wires[makeWire(ps[bestni].i, ps[bestnj].i)] = true
	}

	// Get the top 3 group sizes
	top := getTop3GroupSizes(ps)
	fmt.Printf("Top 3 group sizes: %+v\n", top)

	// Multiply them together
	total := 1
	for _, s := range top {
		total *= s
	}
	fmt.Printf("Total: %d\n", total)
}

type wire struct{ a, b int }

func makeWire(a, b int) wire {
	if a > b {
		return wire{b, a}
	}
	return wire{a, b}
}

func getGroupSizes(ps []point) map[int]int {
	sizes := make(map[int]int)
	for _, p := range ps {
		sizes[p.g] = sizes[p.g] + 1
	}
	return sizes
}

func getSortedGroupSizes(ps []point) []int {
	groups := getGroupSizes(ps)
	sizes := make([]int, 0, len(groups))
	for _, n := range groups {
		sizes = append(sizes, n)
	}
	slices.Sort(sizes)
	slices.Reverse(sizes)
	return sizes
}

func getTop3GroupSizes(ps []point) []int {
	sortedSizes := getSortedGroupSizes(ps)
	return sortedSizes[:3]
}

func groupPoints(ps []point, g1, g2 int) []point {
	g3 := min(g1, g2)
	for i, p := range ps {
		if p.g == g1 || p.g == g2 {
			ps[i] = p.withGroup(g3)
		}
	}
	return ps
}

type point struct{ i, g, x, y, z int }

func (p point) String() string {
	return fmt.Sprintf(
		"[id=%d|g=%d|p=(%d,%d,%d)]",
		p.i, p.g,
		p.x, p.y, p.z,
	)
}

func (p point) withGroup(g int) point {
	return point{
		i: p.i,
		g: g,
		x: p.x,
		y: p.y,
		z: p.z,
	}
}

func (p point) sub(o point) point {
	return point{
		i: -1,
		g: -1,
		x: p.x - o.x,
		y: p.y - o.y,
		z: p.z - o.z,
	}
}

func (p point) mag() int {
	return (p.x * p.x) + (p.y * p.y) + (p.z * p.z)
}

func (p point) dist(o point) int {
	return o.sub(p).mag()
}

func readInput() []point {
	s := bufio.NewScanner(os.Stdin)

	var ps []point

	for s.Scan() {
		// Read the next line
		line := s.Text()
		sns := strings.Split(line, ",")
		if len(sns) != 3 {
			panicf("expected line %q to have three nums", line)
		}
		var p point
		for i, s := range sns {
			n, err := strconv.Atoi(s)
			if err != nil {
				panicf("failed")
			}
			switch i {
			case 0:
				p.x = n
			case 1:
				p.y = n
			case 2:
				p.z = n
			}
		}
		p.i = len(ps) // ID is row #
		p.g = p.i     // Point starts in solo group
		ps = append(ps, p)

	}
	if err := s.Err(); err != nil {
		panic(err)
	}
	return ps
}

func panicf(e string, a ...any) {
	panic(fmt.Errorf(e, a...))
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
