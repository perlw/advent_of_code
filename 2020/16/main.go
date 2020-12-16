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

	rulePositions := make(map[string]int)
	/*
		for i := 0; i < len(validTickets[0].Values); i++ {
			for _, t := range validTickets {
				possible := make([]int, 0, 2)
				for _, r := range rules {
					if r.IsValid(t.Values[i]) {
						possible = append(possible, i)
					}
					fmt.Printf("checking %d for rule %s: %v\n", t.Values[i], r.Name, r.IsValid(t.Values[i]))
				}
				fmt.Printf("\tpossible: %+v\n", possible)
			}
		}
	*/
	fmt.Printf("%+v\n", rulePositions)

	return -1
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
