package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

// Food ...
type Food struct {
	Ingredients []string
	Allergens   []string
}

func inSlice(s []string, a string) bool {
	for i := range s {
		if s[i] == a {
			return true
		}
	}
	return false
}

func delFromSlice(s []string, a string) []string {
	result := make([]string, 0, 10)
	for i := range s {
		if s[i] != a {
			result = append(result, s[i])
		}
	}
	return result
}

func intersectSlice(a, b []string) []string {
	result := make([]string, 0, 10)
	for i := range a {
		if inSlice(b, a[i]) && !inSlice(result, a[i]) {
			result = append(result, a[i])
		}
	}
	for i := range b {
		if inSlice(a, b[i]) && !inSlice(result, b[i]) {
			result = append(result, b[i])
		}
	}
	return result
}

func compileIngredients(input []Food) map[string]int {
	result := make(map[string]int)
	for _, in := range input {
		for _, i := range in.Ingredients {
			result[i]++
		}
	}
	return result
}

func findAllergens(input []Food) map[string][]string {
	result := make(map[string][]string)
	for {
		stop := true
		for _, in := range input {
			for _, a := range in.Allergens {
				if _, ok := result[a]; !ok {
					result[a] = append(result[a], in.Ingredients...)
				} else {
					result[a] = intersectSlice(result[a], in.Ingredients)
				}
			}
		}

		for a, all := range result {
			if len(all) == 1 {
				for i := range result {
					if i != a {
						result[i] = delFromSlice(result[i], all[0])
					}
				}
			} else {
				stop = false
			}
		}

		if stop {
			break
		}
	}
	return result
}

// Task1 ...
func Task1(input []Food, debug bool) int {
	ingredients := compileIngredients(input)
	allergens := findAllergens(input)

	var result int
	for i, c := range ingredients {
		var found bool
		for _, a := range allergens {
			if i == a[0] {
				found = true
				break
			}
		}
		if !found {
			result += c
		}
	}

	if debug {
		fmt.Printf("%+v\n", ingredients)
		fmt.Printf("%+v\n", allergens)
	}

	return result
}

// Task2 ...
func Task2(input []Food, debug bool) string {
	allergens := findAllergens(input)
	sorted := make([]string, 0, len(allergens))
	for a := range allergens {
		sorted = append(sorted, a)
	}
	sort.Strings(sorted)
	ingredients := make([]string, 0, len(sorted))
	for _, a := range sorted {
		ingredients = append(ingredients, allergens[a][0])
	}
	return strings.Join(ingredients, ",")
}

func readInput(reader io.Reader) []Food {
	input := make([]Food, 0, 10)

	scanner := bufio.NewScanner(reader)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "(contains ")
		ingredients := strings.Split(strings.TrimSpace(parts[0]), " ")
		allergens := strings.Split(parts[1][:len(parts[1])-1], ", ")
		input = append(input, Food{
			Ingredients: ingredients,
			Allergens:   allergens,
		})
	}

	return input
}

func main() {
	var debug bool
	flag.BoolVar(&debug, "debug", false, "debug")
	flag.Parse()

	var input []Food

	file, _ := os.Open("input.txt")
	input = readInput(file)
	file.Close()

	result1 := Task1(input, debug)
	fmt.Printf("Task 1: %d\n", result1)

	result2 := Task2(input, debug)
	fmt.Printf("Task 2: %s\n", result2)
}
