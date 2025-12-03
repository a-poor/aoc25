package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var total int

	// Scan the input
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		ln := s.Text()
		cs := []rune(ln)

		// Find the best first character
		var i1 int
		for i := 1; i < len(ln)-1; i++ {
			if cs[i] > cs[i1] {
				i1 = i
			}
		}

		// Find the best second character
		i2 := i1 + 1
		for i := i2 + 1; i < len(ln); i++ {
			if cs[i] > cs[i2] {
				i2 = i
			}
		}

		// Get the number
		n, err := strconv.Atoi(string(cs[i1]) + string(cs[i2]))
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
