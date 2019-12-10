package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/perlw/advent_of_code/2019/intcode"
)

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(fmt.Errorf("could not open input file: %w", err))
	}
	defer input.Close()

	raw, err := ioutil.ReadAll(input)
	if err != nil {
		log.Fatal(fmt.Errorf("error while reading input: %w", err))
	}
	ops, err := func(val []byte) ([]int, error) {
		vals := strings.Split(strings.TrimSpace(string(val)), ",")
		result := make([]int, len(vals))
		for i := range vals {
			result[i], err = strconv.Atoi(vals[i])
			if err != nil {
				return nil, fmt.Errorf("could not convert to int: %w", err)
			}
		}
		return result, nil
	}(raw)
	if err != nil {
		log.Fatal(fmt.Errorf("could not read ops: %w", err))
	}

	// Part 1
	com := intcode.NewComputer(ops)
	fmt.Printf("Part 1:\n")
	for !com.Halted {
		out, halt, err := com.Iterate(func() int {
			return 1
		}, false, false)
		if err != nil {
			log.Fatal(fmt.Errorf("failure while executing computer: %w", err))
		}
		if !halt {
			fmt.Printf("%d\n", out)
		}
	}

	// Part 2
	com = intcode.NewComputer(ops)
	fmt.Printf("Part 2:\n")
	for !com.Halted {
		out, halt, err := com.Iterate(func() int {
			return 2
		}, false, false)
		if err != nil {
			log.Fatal(fmt.Errorf("failure while executing computer: %w", err))
		}
		if !halt {
			fmt.Printf("%d\n", out)
		}
	}
}
