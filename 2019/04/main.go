package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
)

func CheckPassword1(password int) bool {
	digits := strconv.Itoa(password)
	last := rune(0)
	dup := false
	for _, c := range digits {
		if c == last {
			dup = true
		}
		if c < last {
			return false
		}
		last = c
	}
	return dup
}

func CheckPassword2(password int) bool {
	digits := strconv.Itoa(password)
	last := rune(0)
	dups := make(map[rune]int)
	for _, c := range digits {
		dups[c]++
		if c < last {
			return false
		}
		last = c
	}
	for _, d := range dups {
		if d == 2 {
			return true
		}
	}
	return false
}

func main() {
	var start, end int
	flag.IntVar(&start, "start", 0, "start of range")
	flag.IntVar(&end, "end", 0, "end of range")
	flag.Parse()

	if start == 0 || end == 0 {
		flag.PrintDefaults()
		os.Exit(-1)
	}

	c1 := 0
	c2 := 0
	for i := start; i <= end; i++ {
		if CheckPassword1(i) {
			c1++
		}
		if CheckPassword2(i) {
			c2++
		}
	}

	fmt.Printf("c1: %d, c2: %d\n", c1, c2)
}
