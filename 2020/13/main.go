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
	for _, i := range input {
		next := i.NextDeparture(now)
		if next < soonestDeparture {
			soonestDeparture = next
			soonestBusID = int(i)
		}
	}

	return (soonestDeparture - now) * soonestBusID
}

// Task2 ...
func Task2() int {
	return -1
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
		if bus != "x" {
			t, _ := strconv.Atoi(bus)
			input = append(input, Bus(t))
		}
	}
	file.Close()

	result := Task1(now, input)
	fmt.Printf("Task 1: %d\n", result)

	result = Task2()
	fmt.Printf("Task 2: %d\n", result)
}
