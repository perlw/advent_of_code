package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reduce(exp []rune, debug bool) int {
	if debug {
		fmt.Printf("iter %s\n", string(exp))
	}
	var a, b int
	var op rune

	if len(exp) == 1 {
		return int(exp[0] - '0')
	}
	if len(exp) <= 3 {
		a = int(exp[0] - '0')
		b = int(exp[2] - '0')
		op = exp[1]

		if debug {
			fmt.Printf("\t? %d %c %d\n", a, op, b)
		}
	} else {
		if exp[len(exp)-1] == ')' {
			count := 1
			for i := len(exp) - 2; i >= 0; i-- {
				if exp[i] == ')' {
					count++
				} else if exp[i] == '(' {
					count--
					if count == 0 {
						if i == 0 {
							return reduce(exp[i+1:len(exp)-1], debug)
						} else {
							a = reduce(exp[:i-1], debug)
							b = reduce(exp[i+1:len(exp)-1], debug)
							op = exp[i-1]
						}
						break
					}
				}
			}
		} else {
			a = reduce(exp[:len(exp)-2], debug)
			b = int(exp[len(exp)-1] - '0')
			op = exp[len(exp)-2]
		}

		if debug {
			fmt.Printf("\t# %d %c %d\n", a, op, b)
		}
	}

	var result int
	switch op {
	case '*':
		result = a * b
	case '+':
		result = a + b
	}
	if debug {
		fmt.Printf("%d %c %d => %d\n", a, op, b, result)
	}

	return result
}

// Task1 ...
func Task1(input []string) int {
	var result int
	for i := range input {
		result += reduce([]rune(input[i]), false)
	}
	return result
}

// Task2 ...
func Task2() int {
	return -1
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	input := make([]string, 0, 100)
	for scanner.Scan() {
		line := scanner.Text()
		input = append(input, strings.ReplaceAll(line, " ", ""))
	}
	file.Close()

	result := Task1(input)
	fmt.Printf("Task 1: %d\n", result)

	result = Task2()
	fmt.Printf("Task 2: %d\n", result)
}
