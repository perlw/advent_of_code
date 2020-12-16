package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Rule ...
type Rule struct {
	Name   string
	Ranges [][2]int
	Not    []int
	Column int
}

// IsValid ...
func (r Rule) IsValid(val int) bool {
	for _, i := range r.Ranges {
		if val >= i[0] && val <= i[1] {
			return true
		}
	}
	return false
}

// Ticket ...
type Ticket struct {
	Values []int
}

// Task1 ...
func Task1(rules []Rule, tickets []Ticket) int {
	var result int
	for _, t := range tickets {
		for _, i := range t.Values {
			var valid bool
			for _, r := range rules {
				if r.IsValid(i) {
					valid = true
					break
				}
			}

			if !valid {
				result += i
			}
		}
	}
	return result
}

// Task2 ...
func Task2(rules []Rule, tickets []Ticket, targetTicket Ticket) int {
	// Find valid tickets.
	validTickets := make([]Ticket, 0, len(tickets)/2)
	for _, t := range tickets {
		var invalid bool
		for _, i := range t.Values {
			var valid bool
			for _, r := range rules {
				if r.IsValid(i) {
					valid = true
					break
				}
			}
			if !valid {
				invalid = true
				break
			}
		}
		if !invalid {
			validTickets = append(validTickets, t)
		}
	}

	columns := make(map[int]int)
	for i := 0; i < len(validTickets[0].Values); i++ {
		columns[i] = -1
	}
	numColumns := len(columns)

	// Whittle down columns one by one by identifying which rule is _not_ valid in
	// a given column's values, then removing that column from the iteration, and
	// running over it again.
	for numColumns > 0 {
		for i := range rules {
			if rules[i].Column > -1 {
				continue
			}
			rules[i].Not = make([]int, 0, 2)
		}

		for i := 0; i < len(columns); i++ {
			if columns[i] > -1 {
				continue
			}

			for _, t := range validTickets {
				for ri, r := range rules {
					if r.Column > -1 {
						continue
					}

					if !r.IsValid(t.Values[i]) {
						rules[ri].Not = append(r.Not, i)
					}
				}
			}
		}

		for ri, r := range rules {
			if r.Column > -1 {
				continue
			}

			if len(r.Not) == numColumns-1 {
				for c, n := range columns {
					if n > -1 {
						continue
					}

					var found bool
					for _, n := range r.Not {
						if n == c {
							found = true
							break
						}
					}
					if !found {
						columns[c] = ri
						rules[ri].Column = c
						numColumns--
						break
					}
				}
			}
		}
	}

	// Calculate result by multiplying the correct column-values from the target
	// ticket.
	result := 1
	for _, r := range rules {
		if strings.HasPrefix(r.Name, "departure") {
			result *= targetTicket.Values[r.Column]
		}
	}
	return result
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	rules := make([]Rule, 0, 10)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}

		parts := strings.Split(line, ": ")
		rangeParts := strings.Split(parts[1], " or ")

		ranges := make([][2]int, len(rangeParts))
		for i := range rangeParts {
			fmt.Sscanf(rangeParts[i], "%d-%d", &ranges[i][0], &ranges[i][1])
		}
		rules = append(rules, Rule{
			Name:   parts[0],
			Ranges: ranges,
			Not:    make([]int, 0, 2),
			Column: -1,
		})
	}
	scanner.Scan()
	scanner.Scan()
	// My ticket
	parts := strings.Split(scanner.Text(), ",")
	values := make([]int, len(parts))
	for i := range parts {
		values[i], _ = strconv.Atoi(parts[i])
	}
	myTicket := Ticket{
		Values: values,
	}

	scanner.Scan()
	scanner.Scan()
	tickets := make([]Ticket, 0, 10)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ",")
		values := make([]int, len(parts))
		for i := range parts {
			values[i], _ = strconv.Atoi(parts[i])
		}
		tickets = append(tickets, Ticket{
			Values: values,
		})
	}
	file.Close()

	result := Task1(rules, tickets)
	fmt.Printf("Task 1: %d\n", result)

	result = Task2(rules, tickets, myTicket)
	fmt.Printf("Task 2: %d\n", result)
}
