package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func CalculateChecksum(img []int, width, height int) int {
	// Find layer
	least0 := 999999
	var leastLayer int
	layers := len(img) / (width * height)
	for l := 0; l < layers; l++ {
		num0 := 0
		offset := l * (width * height)
		for y := 0; y < height; y++ {
			i := offset + (y * width)
			for x := 0; x < width; x++ {
				if img[i+x] == 0 {
					num0++
				}
			}
		}
		if num0 < least0 {
			least0 = num0
			leastLayer = l
		}
	}

	// Count 1s and 2s
	var ones, twos int
	offset := leastLayer * (width * height)
	for y := 0; y < height; y++ {
		i := offset + (y * width)
		for x := 0; x < width; x++ {
			if img[i+x] == 1 {
				ones++
			} else if img[i+x] == 2 {
				twos++
			}
		}
	}

	return ones * twos
}

func FlattenImage(img []int, width, height int) []int {
	result := make([]int, width*height)

	layers := len(img) / (width * height)
	for l := layers - 1; l >= 0; l-- {
		offset := l * (width * height)
		for y := 0; y < height; y++ {
			i := offset + (y * width)
			for x := 0; x < width; x++ {
				if img[i+x] != 2 {
					result[i+x-offset] = img[i+x]
				}
			}
		}
	}

	return result
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(fmt.Errorf("could not open input file: %w", err))
	}
	defer input.Close()

	raw, err := ioutil.ReadAll(input)
	if err != nil {
		log.Fatal(fmt.Errorf("error while reading input: %w", err))
	}
	raw = []byte(strings.TrimSpace(string(raw)))

	img := make([]int, len(raw))
	for i := range raw {
		img[i] = int(raw[i] - 48)
	}

	result := CalculateChecksum(img, 25, 6)
	fmt.Printf("Part 1: %d\n", result)

	flattened := FlattenImage(img, 25, 6)
	fmt.Println("Part 2:")
	for y := 0; y < 6; y++ {
		i := (y * 25)
		for x := 0; x < 25; x++ {
			var r rune
			if flattened[i+x] == 0 {
				r = ' '
			} else {
				r = '*'
			}
			fmt.Printf("%c", r)
		}
		fmt.Println()
	}
}
