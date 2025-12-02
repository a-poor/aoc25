package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

var ErrDone = errors.New("done")

func main() {
	var ninv int // Number of invalid IDs
	var tinv int

	// Scan the input
	ir := InputReader{
		r:    bufio.NewReader(os.Stdin),
		done: false,
	}
	for !ir.done {
		// Get the next span
		nl, nr, err := ir.NextNum()
		if err != nil {
			panic(err)
		}

		// Loop through the span...
		for i := nl; i <= nr; i++ {
			nd := numberOfDigits(i)

			// Try different window sizes
		window:
			for w := nd / 2; w > 0; w-- {
				// Does w divide s evenly?
				if nd%w != 0 {
					continue
				}

				// Make it a string
				s := strconv.Itoa(i)

				// Get the first window
				w1 := s[0:w]

				// Look through the rest
				for j := 1; j < nd/w; j++ {
					wn := s[w*j : w*(j+1)]
					if w1 != wn {
						continue window
					}
				}

				// If we get here it must be valid
				fmt.Printf("Found invlaid id! %d\n", i)
				ninv++
				tinv += i

				// No need to keep looking
				break
			}
		}
	}

	// Done!
	fmt.Printf("    n invalid: %d\n", ninv)
	fmt.Printf("total invalid: %d\n", tinv)
}

func numberOfDigits(x int) int {
	var n int
	for x > 0 {
		x /= 10
		n++
	}
	return n
}

type InputReader struct {
	r    *bufio.Reader
	done bool
}

func (ir *InputReader) Next() (string, string, error) {
	// Get the starting point...
	s1, err := ir.r.ReadString('-')
	if err != nil {
		return "", "", fmt.Errorf("failed to get first string: %w", err)
	}
	s1 = strings.ReplaceAll(s1, "\n", "")
	s1 = strings.ReplaceAll(s1, " ", "")
	s1 = strings.ReplaceAll(s1, "-", "")

	// Get the ending point...
	s2, err := ir.r.ReadString(',')
	if err != nil && !errors.Is(err, io.EOF) {
		return "", "", fmt.Errorf("failed to get second string: %w", err)
	}
	s2 = strings.ReplaceAll(s2, "\n", "")
	s2 = strings.ReplaceAll(s2, " ", "")
	s2 = strings.ReplaceAll(s2, ",", "")

	// Done?
	if errors.Is(err, io.EOF) {
		ir.done = true
	}

	// Return the results
	return s1, s2, nil
}

func (ir *InputReader) NextNum() (int, int, error) {
	s1, s2, err := ir.Next()
	if err != nil {
		return 0, 0, err
	}
	n1, err := strconv.Atoi(s1)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse first num as int: %w", err)
	}
	n2, err := strconv.Atoi(s2)
	if err != nil {
		return 0, 0, fmt.Errorf("failed to parse second num as int: %w", err)
	}
	return n1, n2, nil
}
