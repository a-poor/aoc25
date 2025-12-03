package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const ndigits = 12

func main() {
	var total int

	// Scan the input
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		// Get the next line
		ln := s.Text()
		cs := []rune(ln)

		// An array to store the digits
		ds := make([]rune, 0, ndigits)

		// Start the loop
		lasti := -1
		for len(ds) < ndigits {
			// Find the best character...

			// Start looking at the next character
			// with the base case being that the first
			// remaining character is best
			bi := lasti + 1

			// Define the loop bounds (to compare against
			// the base case) starting with the character
			// *after* next and ending with enough characters
			// so that we will still have enough to finish
			// our 12-digit (or n-digit) number
			start := bi + 1
			end := len(ln) - (ndigits - len(ds) - 1)

			// Look for better numbers?
			for i := start; i < end; i++ {
				if cs[i] > cs[bi] {
					bi = i
				}
			}

			// Push the digit
			ds = append(ds, cs[bi])
			lasti = bi
		}

		// Get the number
		n, err := strconv.Atoi(string(ds))
		if err != nil {
			panic(err)
		}
		// fmt.Printf("found one! %d\n", n)

		// Add to the total
		total += n
	}
	if err := s.Err(); err != nil {
		panic(err)
	}

	// Done!
	fmt.Printf("total: %d\n", total)
}
