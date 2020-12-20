// Note: Do NOT use this solution as an example. I hate it. It's a horrible mess
// of brute force, workarounds, and shortcuts, and is generally really not a
// good idea. A syntax tree would be a better idea and what I should have done
// from the beginning but I tried to save time by not thinking|I mean using
// recursion instead.  I've fought overflows, indexing errors, recursion loops,
// general stubborn stupidity, and I'm even two days behind.
// But I guess I've learnt what _not_ to do as well.
// Next year AoC, next year..

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func reduce(exp []rune, debug bool) int {
	if debug {
		fmt.Printf("iter ")
		for _, r := range exp {
			if r < ' ' {
				fmt.Printf("%d", r)
			} else {
				fmt.Printf("%c", r)
			}
		}
		fmt.Println()
	}
	var a, b int
	var op rune

	if len(exp) == 1 {
		return int(exp[0])
	}
	if len(exp) <= 3 {
		a = int(exp[0])
		b = int(exp[2])
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
						}

						a = reduce(exp[:i-1], debug)
						b = reduce(exp[i+1:len(exp)-1], debug)
						op = exp[i-1]
						break
					}
				}
			}
		} else {
			a = reduce(exp[:len(exp)-2], debug)
			b = int(exp[len(exp)-1])
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

func reduceAdditions(exp []int, debug bool) []int {
	newExp := make([]int, 0, len(exp))

	for i := 0; i < len(exp); i++ {
		switch exp[i] {
		case '(':
			var count int
			for j := i; j < len(exp); j++ {
				if exp[j] == '(' {
					count++
				} else if exp[j] == ')' {
					count--
					if count == 0 {
						if debug {
							fmt.Printf("mm ")
							for _, r := range exp[i+1 : j] {
								if r == '*' || r == '+' || r == '(' || r == ')' {
									fmt.Printf("%c", r)
								} else {
									fmt.Printf("%d", r-255)
								}
							}
							fmt.Println()
						}
						e := reduceAdditions(exp[i+1:j], debug)
						for {
							prev := e
							e = reduceAdditions(e, debug)
							if len(prev) == len(e) {
								break
							}
						}

						if debug {
							fmt.Printf("++ ")
							for _, r := range e {
								if r == '*' || r == '+' || r == '(' || r == ')' {
									fmt.Printf("%c", r)
								} else {
									fmt.Printf("%d", r-255)
								}
							}
							fmt.Println()
						}

						e = reduceMultiplications(e, debug)
						for {
							prev := e
							if len(prev) == len(e) {
								break
							}
						}
						if debug {
							fmt.Printf("** ")
							for _, r := range e {
								if r == '*' || r == '+' || r == '(' || r == ')' {
									fmt.Printf("%c", r)
								} else {
									fmt.Printf("%d", r-255)
								}
							}
							fmt.Println()
						}
						if len(e) == 1 {
							newExp = append(newExp, e...)
						} else {
							newExp = append(newExp, '(')
							newExp = append(newExp, e...)
							newExp = append(newExp, ')')
						}

						i = j
						break
					}
				}
			}
		case '*':
			newExp = append(newExp, exp[i])
		case '+':
			if exp[i-1] >= 255 && exp[i+1] >= 255 {
				if debug {
					fmt.Printf("add %d %d\n", newExp[len(newExp)-1]-255, exp[i+1]-255)
				}
				newExp[len(newExp)-1] = ((newExp[len(newExp)-1] - 255) + (exp[i+1] - 255)) + 255
				i++
			} else {
				newExp = append(newExp, exp[i])
			}
		default:
			newExp = append(newExp, exp[i])
		}
	}

	return newExp
}

func reduceMultiplications(exp []int, debug bool) []int {
	newExp := make([]int, 0, len(exp))

	for i := 0; i < len(exp); i++ {
		switch exp[i] {
		case '*':
			if exp[i-1] >= 255 && exp[i+1] >= 255 {
				newExp[len(newExp)-1] = ((newExp[len(newExp)-1] - 255) * (exp[i+1] - 255)) + 255
				i++
			} else {
				newExp = append(newExp, exp[i])
			}
		default:
			newExp = append(newExp, exp[i])
		}
	}

	return newExp
}

func reduce2(exp []int, debug bool) int {
	exp = reduceAdditions(exp, debug)
	for {
		prev := exp
		exp = reduceAdditions(exp, debug)
		if debug {
			fmt.Printf("additions ")
			for _, r := range exp {
				if r == '*' || r == '+' || r == '(' || r == ')' {
					fmt.Printf("%c", r)
				} else {
					fmt.Printf("%d", r-255)
				}
			}
			fmt.Println()
		}

		if len(prev) == len(exp) {
			break
		}
	}

	exp = reduceMultiplications(exp, debug)
	for {
		prev := exp
		exp = reduceMultiplications(exp, debug)
		if debug {
			fmt.Printf("multiplications ")
			for _, r := range exp {
				if r == '*' || r == '+' || r == '(' || r == ')' {
					fmt.Printf("%c", r)
				} else {
					fmt.Printf("%d", r-255)
				}
			}
			fmt.Println()
		}

		if len(prev) == len(exp) {
			break
		}
	}

	return exp[0] - 255
}

func prepareInput(str string, fudge bool) []rune {
	line := []rune(strings.ReplaceAll(str, " ", ""))
	for i, r := range line {
		if r != '(' && r != ')' && r != '*' && r != '+' {
			if fudge {
				line[i] = 255 + r - '0'
			} else {
				line[i] = r - '0'
			}
		}
	}
	return line
}

// Task1 ...
func Task1(input [][]rune) int {
	var result int
	for i := range input {
		result += reduce(input[i], false)
	}
	return result
}

// Task2 ...
func Task2(input [][]rune) int {
	var result int
	for i := range input {
		in := make([]int, len(input[i]))
		for i, j := range input[i] {
			in[i] = int(j)
		}
		r := reduce2(in, false)
		result += r
	}
	return result
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	input := make([]string, 0, 100)
	for scanner.Scan() {
		input = append(input, scanner.Text())
	}
	file.Close()

	in := make([][]rune, len(input))
	for i, line := range input {
		in[i] = prepareInput(line, false)
	}
	result := Task1(in)
	fmt.Printf("Task 1: %d\n", result)

	in = make([][]rune, len(input))
	for i, line := range input {
		in[i] = prepareInput(line, true)
	}
	result = Task2(in)
	fmt.Printf("Task 2: %d\n", result)
}
