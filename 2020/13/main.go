package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Bus ...
type Bus int

// NextDeparture ...
func (b Bus) NextDeparture(now int) int {
	return int(now) - (now % int(b)) + int(b)
}

// Task1 ...
func Task1(now int, input []Bus) int {
	soonestDeparture := math.MaxInt32
	var soonestBusID int
	for _, bus := range input {
		if bus == -1 {
			continue
		}

		next := bus.NextDeparture(now)
		if next < soonestDeparture {
			soonestDeparture = next
			soonestBusID = int(bus)
		}
	}

	return (soonestDeparture - now) * soonestBusID
}

// Task2 ...
func Task2(input []Bus) int {
	timestamp := 1
	increment := 1

	for i, bus := range input {
		if bus == -1 {
			continue
		}

		for {
			if (timestamp+i)%int(bus) == 0 {
				increment *= int(bus)
				break
			}
			timestamp += increment
		}
	}

	return timestamp
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	input := make([]Bus, 0, 100)
	scanner.Scan()
	now, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	buses := strings.Split(scanner.Text(), ",")
	for _, bus := range buses {
		var val int
		if bus == "x" {
			val = -1
		} else {
			val, _ = strconv.Atoi(bus)
		}
		input = append(input, Bus(val))
	}
	file.Close()

	result := Task1(now, input)
	fmt.Printf("Task 1: %d\n", result)

	result = Task2(input)
	fmt.Printf("Task 2: %d\n", result)
}
