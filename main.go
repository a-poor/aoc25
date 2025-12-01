package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var re = regexp.MustCompile(`([LR])(\d+)`)

func main() {
	// Initialize the state
	p, n := 50, 0

	// Scan the input
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		// Get the line
		l := s.Text()
		// fmt.Printf("LINE: %s\n", l)

		// Parse it
		gs := re.FindStringSubmatch(l)
		if len(gs) != 3 {
			panic(fmt.Sprintf("Error with line %q -> %+v", l, gs))
		}

		// Parse the number
		c, err := strconv.Atoi(gs[2])
		if err != nil {
			panic(err)
		}

		// Move left or right
		if strings.ToLower(gs[1]) == "l" {
			p = moveLeft(p, c)
		} else {
			p = moveRight(p, c)
		}

		// Landed on zero?
		if p == 0 {
			n++
		}
	}
	if err := s.Err(); err != nil {
		panic(err)
	}

	// Done!
	fmt.Printf("Count: %d\n", n)
}

func moveRight(start, count int) int {
	n := start
	for range count {
		n++
		if n == 100 {
			n = 0
		}

	}
	return n
}

func moveLeft(start, count int) int {
	n := start
	for range count {
		n--
		if n == -1 {
			n = 99
		}

	}
	return n
}
