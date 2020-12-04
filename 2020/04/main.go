package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var validFields = []string{
	"byr", // Birth Year
	"iyr", // Issue Year
	"eyr", // Expiration Year
	"hgt", // Height
	"hcl", // Hair Color
	"ecl", // Eye Color
	"pid", // Passport ID
	"cid", // Country ID
}

var requiredFields = []string{"byr", "iyr", "eyr", "hgt", "hcl", "ecl", "pid"}
var validEyeColors = []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

// Passport ...
type Passport map[string]string

func stringInSlice(slice []string, str string) bool {
	for _, s := range slice {
		if s == str {
			return true
		}
	}
	return false
}

// HasAllFields checks that all requried fields are present.
func (p Passport) HasAllFields() bool {
	for _, s := range requiredFields {
		if _, ok := p[s]; !ok {
			return false
		}
	}
	return true
}

// IsValid validates a passport.
func (p Passport) IsValid() bool {
	for k, v := range p {
		switch k {
		case "byr":
			val, _ := strconv.Atoi(v)
			if !(val >= 1920 && val <= 2002) {
				return false
			}
		case "iyr":
			val, _ := strconv.Atoi(v)
			if !(val >= 2010 && val <= 2020) {
				return false
			}
		case "eyr":
			val, _ := strconv.Atoi(v)
			if !(val >= 2020 && val <= 2030) {
				return false
			}
		case "hgt":
			if len(v) < 3 {
				return false
			}

			unit := v[len(v)-2:]
			measurement, _ := strconv.Atoi(v[:len(v)-2])
			switch unit {
			case "cm":
				if !(measurement >= 150 && measurement <= 193) {
					return false
				}
			case "in":
				if !(measurement >= 59 && measurement <= 76) {
					return false
				}
			default:
				return false
			}
		case "hcl":
			if len(v) != 7 {
				return false
			}
			matched, _ := regexp.MatchString("[a-fA-F0-9]+", v[1:])
			if !matched {
				return false
			}
		case "ecl":
			if !stringInSlice(validEyeColors, v) {
				return false
			}
		case "pid":
			if len(v) != 9 {
				return false
			}
			matched, _ := regexp.MatchString(`^\d+$`, v[1:])
			if !matched {
				return false
			}
		default:
		}
	}

	return true
}

func (p Passport) String() string {
	var result strings.Builder

	fmt.Fprintf(
		&result, `╭─────────────────────────────────────────╮
│ Passport Card                           │
│ +-------+ Issuing Country   Passport ID │
│ |  /-\  |  %03s               %09s  │
│ |  \-/  |                               │
│ | /-u-\ | Issued            Expires     │
│ +-------+  %04s              %04s       │
│                                         │
│ Height       Hair Color       Eye Color │
│  %05s        %07s          %03s      │
│ Birth Year                              │
│  %04s                                   │
╰─────────────────────────────────────────╯`, p["cid"], p["pid"],
		p["iyr"], p["eyr"], p["hgt"], p["hcl"], p["ecl"], p["byr"],
	)

	return result.String()
}

// Task1 ...
func Task1(input []Passport) int {
	var result int
	for _, i := range input {
		if i.HasAllFields() {
			result++
		}
	}
	return result
}

// Task2 ...
func Task2(input []Passport) int {
	var result int
	for _, i := range input {
		if i.HasAllFields() && i.IsValid() {
			result++
		}
	}
	return result
}

func main() {
	var input []Passport
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	p := Passport{}
	for scanner.Scan() {
		data := strings.Split(scanner.Text(), " ")
		if len(data) == 1 && data[0] == "" {
			input = append(input, p)
			p = Passport{}
			continue
		}

		for _, d := range data {
			kv := strings.Split(d, ":")
			if stringInSlice(validFields, kv[0]) {
				p[kv[0]] = kv[1]
			}
		}
	}
	input = append(input, p)
	file.Close()

	result := Task1(input)
	fmt.Printf("Task 1: %d\n", result)

	result = Task2(input)
	fmt.Printf("Task 2: %d\n", result)
}
