package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

const (
	Start = "svr"
	DAC   = "dac"
	FFT   = "fft"
	End   = "out"
)

var pathCache = map[string]int{}

func main() {
	// Read in the edges
	es := readInput()

	// Sort the edges
	slices.SortFunc(es, ordEdge)

	// Get the count
	// count := countPathsOut(es, Start, End, nil)
	startToDAC := countPathsOut(es, Start, DAC, nil)
	startToFFT := countPathsOut(es, Start, FFT, nil)

	dacToFFT := countPathsOut(es, DAC, FFT, []string{DAC})
	fftToDAC := countPathsOut(es, FFT, DAC, []string{FFT})

	dacToEnd := countPathsOut(es, DAC, End, []string{DAC, FFT})
	fftToEnd := countPathsOut(es, FFT, End, []string{DAC, FFT})

	var count int
	count += startToDAC * dacToFFT * fftToEnd
	count += startToFFT * fftToDAC * dacToEnd

	fmt.Printf("Found %d paths from %q to %q\n", count, Start, End)
	// fmt.Printf("pathCache: %+v\n", pathCache)
}

func countPathsOut(es []edge, pos, goal string, seen []string) int {
	k := pos + ":" + goal
	if n, ok := pathCache[k]; ok {
		return n
	}

	n := _countPathsOut(es, pos, goal, seen)
	pathCache[k] = n
	return n
}

func _countPathsOut(es []edge, pos, goal string, seen []string) int {
	// Are we there yet?
	if pos == goal {
		return 1
	}

	// This "end" is not the goal but it *is*
	// a dead end, so stop here...
	if pos == End {
		return 0
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
		count += countPathsOut(es, e.to, goal, seen2)
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
