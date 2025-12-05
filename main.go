package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"slices"
	"strconv"
	"strings"
)

func main() {
	// Read in the input ranges and
	// ids to check
	rs, ids := readInput()

	// Track the count of good ids
	var count int

	// Loop through the ids
	for _, id := range ids {
		// TODO: Add binary search to find start
		for _, r := range rs {
			if r.start > id {
				break
			}
			if r.end < id {
				continue
			}
			if r.isin(id) {
				count++
				break
			}
		}
	}

	// Done!
	fmt.Printf("Final count: %d\n", count)
}

type idrange struct{ start, end int }

func (r idrange) isin(n int) bool {
	return n >= r.start && n <= r.end
}

func readInput() ([]idrange, []int) {
	// Pre-define regex for range
	rre := regexp.MustCompile(`^(\d+)-(\d+)$`)

	// Pre-define arrays to be returned and a
	// flag to check which list we're on
	var rs []idrange
	var ids []int
	var inids bool

	// Scan through the input lines
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		// Get the next line
		ln := s.Text()

		switch {
		case strings.TrimSpace(ln) == "":
			// Empty line means switch list type
			inids = true

		case !inids:
			// Line is a range
			gs := rre.FindStringSubmatch(ln)
			if gs == nil {
				panic("error finding regex match with '" + ln + "'")
			}
			sl, sr := gs[1], gs[2]

			// Parse as numbers
			nl, err := strconv.Atoi(sl)
			if err != nil {
				panic("failed to parse number '" + sl + "': " + err.Error())
			}
			nr, err := strconv.Atoi(sr)
			if err != nil {
				panic("failed to parse number '" + sr + "': " + err.Error())
			}

			// Add it to the list
			rs = append(rs, idrange{start: nl, end: nr})

		case inids:
			// Line is an id candidate
			//
			// Just parse it!
			n, err := strconv.Atoi(ln)
			if err != nil {
				panic("Failed to parse id '" + ln + "':" + err.Error())
			}
			ids = append(ids, n)
		}
	}
	if err := s.Err(); err != nil {
		panic(err)
	}

	// Sort the ranges
	slices.SortFunc(rs, func(ra, rb idrange) int {
		if ra.start < rb.start {
			return -1
		}
		if ra.start > rb.start {
			return 1
		}
		if ra.end < rb.end {
			return -1
		}
		if ra.end > rb.end {
			return 1
		}
		return 0
	})

	// Return the input data!
	return rs, ids
}
