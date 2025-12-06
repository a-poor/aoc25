package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"unicode"
)

type operation rune

const (
	Mul operation = '*'
	Add operation = '+'
	Div operation = '/'
	Sub operation = '-'
)

func main() {
	var total int
	eqs := readInput()
	for _, eq := range eqs {
		total += eq.solve()
	}
	fmt.Printf("total = %d\n", total)
}

type equation struct {
	ns []int
	op operation
}

func (e equation) solve() int {
	n := e.ns[0]
	for i := 1; i < len(e.ns); i++ {
		switch e.op {
		case Mul:
			n *= e.ns[i]
		case Add:
			n += e.ns[i]
		case Div:
			n /= e.ns[i]
		case Sub:
			n -= e.ns[i]
		default:
			panic(fmt.Sprintf("Unknown rune %q", string(e.op)))
		}
	}
	return n
}

func readInput() []equation {
	var eqs []equation
	s := bufio.NewScanner(os.Stdin)
	for s.Scan() {
		// Read the next line
		line := s.Text()

		// Handle different line cases
		switch {
		case len(eqs) == 0:
			ns := parseNumberLine(line)
			for _, n := range ns {
				eqs = append(eqs, equation{
					ns: []int{n},
					op: operation(0),
				})
			}

		case isNumberLine(line):
			ns := parseNumberLine(line)
			for i, n := range ns {
				eqs[i].ns = append(eqs[i].ns, n)
			}

		default:
			os := parseOpLine(line)
			for i, o := range os {
				eqs[i].op = o
			}
		}
	}
	if err := s.Err(); err != nil {
		panic(err)
	}
	return eqs
}

func isNumberLine(ln string) bool {
	for _, r := range []rune(ln) {
		if unicode.IsDigit(r) {
			return true
		}
	}
	return false
}

func parseNumberLine(ln string) []int {
	ln = strings.TrimSpace(ln)
	var ns []int
	re := regexp.MustCompile(`^\d+`)
	for len(ln) > 0 {
		m := re.FindString(ln)
		n, err := strconv.Atoi(m)
		if err != nil {
			panic(err)
		}
		ns = append(ns, n)
		ln = ln[len(m):]
		ln = strings.TrimSpace(ln)
	}
	return ns
}

func parseOpLine(ln string) []operation {
	ln = strings.TrimSpace(ln)
	var os []operation
	re := regexp.MustCompile(`^[+\-*/]`)
	for len(ln) > 0 {
		m := re.FindString(ln)
		os = append(os, operation(m[0]))
		ln = ln[len(m):]
		ln = strings.TrimSpace(ln)
	}
	return os
}
