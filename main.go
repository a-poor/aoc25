package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

const (
	Start = "you"
	End   = "out"
)

func main() {
	// Read in the edges
	es := readInput()

	// Sort the edges
	slices.SortFunc(es, ordEdge)

	// Get the count
	count := countPathsOut(es, Start, nil)
	fmt.Printf("Found %d paths from %q to %q\n", count, Start, End)
}

func countPathsOut(es []edge, pos string, seen []string) int {
	// Are we there yet?
	if pos == End {
		return 1
	}

	// Otherwise look for next steps...
	//
	// Find the start
	start, _ := slices.BinarySearchFunc(es, edge{from: pos}, func(a, b edge) int {
		return strings.Compare(a.from, b.from)
	})

	// Doesn't exist?
	if es[start].from != pos {
		panic(fmt.Sprintf("point %q not in from list? (start=%d)", pos, start))
	}

	// Create a new "seen" slice that includes "pos"
	seen2 := binaryInsert(seen, pos)

	// Look through the list for all edges from pos
	var count int
	for i := start; i < len(es) && es[i].from == pos; i++ {
		// Get the edge
		e := es[i]

		// Have we already been there?
		//
		// If so, it's a dead end
		if binaryContains(seen, e.to) {
			continue
		}

		// Otherwise, recurse
		n := countPathsOut(es, e.to, seen2)
		// (Optionally cache?)
		count += n
	}
	return count
}

func binaryContains(ps []string, p string) bool {
	_, ok := slices.BinarySearch(ps, p)
	return ok
}

func binaryInsert(ps []string, p string) []string {
	// Create a new slice with 1 extra capacity
	ps2 := make([]string, len(ps), len(ps)+1)

	// And copy over the existing values
	copy(ps2, ps)

	// Find the position of the new item
	i, _ := slices.BinarySearch(ps2, p)

	// And insert it
	return slices.Insert(ps2, i, p)
}

type edge struct{ from, to string }

func (e edge) String() string {
	return fmt.Sprintf("(%s->%s)", e.from, e.to)
}

func ordEdge(a, b edge) int {
	if c := strings.Compare(a.from, b.from); c != 0 {
		return c
	}
	return strings.Compare(a.to, b.to)
}

func readInput() []edge {
	var es []edge
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		line := s.Text()
		ps := strings.Split(line, ": ")
		if len(ps) != 2 {
			panic(fmt.Errorf("expected line to have 2 parts %q", line))
		}

		f := ps[0]
		for t := range strings.SplitSeq(ps[1], " ") {
			es = append(es, edge{
				from: f,
				to:   t,
			})
		}
	}
	if err := s.Err(); err != nil {
		panic(err)
	}
	return es
}
