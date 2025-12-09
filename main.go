package main

import (
	"bufio"
	"fmt"
	"os"
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
	fmt.Printf("Found %d input points\n", len(ps))

	var besti int
	var bestj int
	var bestArea int
	for i, a := range ps {
		for j, b := range ps {
			ar := area(a, b)
			if ar > bestArea {
				besti = i
				bestj = j
				bestArea = ar
			}
		}
	}
	fmt.Printf("%s -> %s == %d\n", ps[besti], ps[bestj], bestArea)
}

type point struct{ x, y int }

func (p point) String() string {
	return fmt.Sprintf("(%d,%d)", p.x, p.y)
}

func area(a, b point) int {
	x := abs(a.x-b.x) + 1
	y := abs(a.y-b.y) + 1
	return x * y
}

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func readInput() []point {
	var ps []point
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		line := s.Text()
		sns := strings.Split(line, ",")
		if len(sns) != 2 {
			panic(fmt.Errorf("expected line to have 2 parts %q", line))
		}
		ns := make([]int, 2)
		for i, s := range sns {
			n, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			ns[i] = n
		}
		ps = append(ps, point{
			x: ns[0],
			y: ns[1],
		})
	}
	if err := s.Err(); err != nil {
		panic(err)
	}
	return ps
}
