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
	rs, _ := readInput()
	// fmt.Printf("Read in %d ranges\n", len(rs))
	compacted := true
	for compacted {
		// fmt.Println("Compacting...")
		rs, compacted = compact(rs)
	}

	// fmt.Printf("%+v\n", rs)

	// Count up the id size
	var count int
	for _, r := range rs {
		count += r.size()
	}

	// Done!
	fmt.Printf("There were %d unique ranges with a total id count of %d\n", len(rs), count)
}

func compact(rs []idrange) ([]idrange, bool) {
	var compacted bool
	var rs2 []idrange
	for _, ra := range rs {
		// Hit will track if we see one
		// in the new list that matches
		hit := -1

		// Look for overlap
		for j, rb := range rs2 {
			// if ra.isbefore(rb) {
			// 	continue
			// }
			// if ra.isafter(rb) {
			// 	break
			// }
			if ra.overlaps(rb) {
				hit = j
				break
			}
		}

		// Did it match one?
		switch hit {
		case -1:
			rs2 = append(rs2, ra)
		default:
			rs2[hit] = rs2[hit].merge(ra)
			compacted = true
		}
	}
	return rs2, compacted
}

type idrange struct{ start, end int }

func (r idrange) isbefore(o idrange) bool {
	return r.end < o.start
}

func (r idrange) isafter(o idrange) bool {
	return r.start > o.end
}

func (r idrange) contains(n int) bool {
	return n >= r.start && n <= r.end
}

func (r idrange) size() int {
	return r.end - r.start + 1
}

func (r idrange) overlaps(o idrange) bool {
	return r.contains(o.start) ||
		r.contains(o.end) ||
		o.contains(r.start) ||
		o.contains(r.end)

}

func (r idrange) merge(o idrange) idrange {
	if !r.overlaps(o) {
		panic(fmt.Sprintf("%+v and %+v don't overlap", r, o))
	}
	start := r.start
	if o.start < start {
		start = o.start
	}
	end := r.end
	if o.end > end {
		end = o.end
	}
	return idrange{
		start: start,
		end:   end,
	}
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
