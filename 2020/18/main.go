package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reduce(exp []rune) int {
	fmt.Printf("%s\n", string(exp))
	var a, b int
	var op rune

	if len(exp) == 1 {
		return int(exp[0] - '0')
	} else if len(exp) <= 3 {
		a = int(exp[0] - '0')
		b = int(exp[2] - '0')
		op = exp[1]
	} else {
		if exp[len(exp)-1] == ')' {
			count := 1
			for i := len(exp) - 2; i >= 0; i-- {
				if exp[i] == ')' {
					count++
				} else if exp[i] == '(' {
					count--
					if count == 0 {
						if i > 0 {
							fmt.Println(string(exp[:i-1]), string(exp[i-1]), string(exp[i+1:len(exp)-1]))
							a = reduce(exp[:i-1])
							op = exp[i-1]
						}
						b = reduce(exp[i+1 : len(exp)-1])
					}
				}
			}
		} else {
			b = int(exp[len(exp)-1] - '0')
			op = exp[len(exp)-2]
			a = reduce(exp[:len(exp)-2])
		}
	}

	fmt.Printf("%d %c %d\n", a, op, b)

	switch op {
	case '*':
		return a * b
	case '+':
		return a + b
	case ')':
		return a
	case '(':
		return b
	}

	return b
}

// Task1 ...
func Task1(input []string) int {
	var result int
	for i := range input {
		result += reduce([]rune(input[i]))
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
