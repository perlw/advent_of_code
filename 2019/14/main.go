package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

type Resource struct {
	Type  string
	Count int
}

type Recipe struct {
	Materials []Resource
	Count     int
}

func produce(recipes map[string]Recipe, resources map[string]int, target string, debug bool) int {
	recipe, ok := recipes[target]
	if !ok {
		log.Fatalf("%s is not part of resources", target)
	}

	if debug {
		fmt.Printf("producing %d %s\n", recipe.Count, target)
	}
	for _, m := range recipe.Materials {
		if m.Type == "ORE" {
			resources[m.Type] += m.Count
			continue
		}
		if debug {
			fmt.Printf("need %d %s (%d)\n", m.Count, m.Type, resources[m.Type])
		}
		if resources[m.Type] < m.Count {
			var produced int
			for produced < m.Count-resources[m.Type] {
				produced += produce(recipes, resources, m.Type, debug)
			}
			resources[m.Type] += produced
			if debug {
				fmt.Printf("produced %+v %s\n", produced, m.Type)
				fmt.Printf("%+v\n", resources)
			}
		}
		resources[m.Type] -= m.Count
		if debug {
			fmt.Printf("took %+v %s\n", m.Count, m.Type)
			fmt.Printf("%+v\n", resources)
		}
	}

	return recipe.Count
}

func CalculateOreRequirement(recipes map[string]Recipe, debug bool) int {
	if debug {
		fmt.Println()
	}
	resources := make(map[string]int)
	produce(recipes, resources, "FUEL", debug)
	if debug {
		fmt.Printf("%+v\n", resources)
	}
	return resources["ORE"]
}

func CalculateMaxFuel(recipes map[string]Recipe, target int, debug bool) int {
	if debug {
		fmt.Println()
	}

	step := 100000
	var count, used int
	resources := make(map[string]int)
	for i := range recipes {
		recipe := recipes[i]
		recipe.Count *= step
		recipes[i] = recipe
		for j := range recipes[i].Materials {
			recipes[i].Materials[j].Count *= step
		}
	}
	it := 0
	for used < target {
		it++
		if it%100 == 0 {
			fmt.Println(count)
		}
		rec := make(map[string]int)
		for i := range resources {
			rec[i] = resources[i]
		}
		produce(recipes, rec, "FUEL", debug)

		ore := rec["ORE"]
		if used+ore > target {
			div := 10
			if step == 100000 {
				div = 100000
			}
			step /= div
			if step == 0 {
				break
			}
			for i := range recipes {
				recipe := recipes[i]
				recipe.Count /= div
				recipes[i] = recipe
				for j := range recipes[i].Materials {
					recipes[i].Materials[j].Count /= div
				}
			}
			for i := range resources {
				rec[i] = resources[i]
			}
			resources["ORE"] = 0
			continue
		}
		used += ore
		rec["ORE"] = 0
		count += step
		for i := range rec {
			resources[i] = rec[i]
		}
	}
	return count
}

func LoadRecipe(line string) (string, Recipe) {
	var in, out string
	parts := strings.Split(line, "=>")
	in = parts[0]
	out = parts[1]

	var count int
	var name string
	fmt.Sscanf(out, "%d %s", &count, &name)

	recipe := Recipe{
		Count: count,
	}
	for _, i := range strings.Split(in, ",") {
		var count int
		var name string
		fmt.Sscanf(i, "%d %s", &count, &name)
		recipe.Materials = append(recipe.Materials, Resource{
			Count: count,
			Type:  name,
		})
	}
	return name, recipe
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(fmt.Errorf("could not open input file: %w", err))
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	recipes := make(map[string]Recipe)
	for scanner.Scan() {
		name, recipe := LoadRecipe(strings.TrimSpace(scanner.Text()))
		recipes[name] = recipe
	}

	fmt.Printf("Part 1: %d\n", CalculateOreRequirement(recipes, false))
	fmt.Printf("Part 2: %d\n", CalculateMaxFuel(recipes, 1000000000000, false))
}
