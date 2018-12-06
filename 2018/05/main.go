package main

import (
	"bufio"
	"fmt"
	"os"
)

func findNextNonSpace(buf []byte) int {
	for t := 0; t < len(buf); t++ {
		if buf[t] != ' ' {
			return t
		}
	}
	return -1
}

func reactChain(chain []byte) []byte {
	t := findNextNonSpace(chain)
	modded := false
	for {
		i := t
		u := findNextNonSpace(chain[t+1:])
		if u == -1 {
			break
		}
		t += u + 1
		j := t

		a := chain[i]
		b := chain[j]

		if a-b == 32 || b-a == 32 {
			chain[i] = ' '
			chain[j] = ' '
			modded = true
		}
	}

	if !modded {
		return chain
	}

	return reactChain(chain)
}

type Puzzle struct {
	Chain []byte
}

func countBufLength(buf []byte) int {
	count := 0
	t := findNextNonSpace(buf)
	for {
		b := findNextNonSpace(buf[t+1:])
		if b == -1 {
			if buf[t] != ' ' {
				count++
			}
			break
		}
		t += b + 1
		count++
	}
	return count
}

func (p *Puzzle) Solution1() int {
	buf := make([]byte, len(p.Chain))
	copy(buf, p.Chain)
	return countBufLength(reactChain(buf))
}

func (p *Puzzle) Solution2() int {
	done := make(chan int)
	var a, b byte
	for a = 'A'; a <= 'Z'; a++ {
		b = a + 32

		go (func(a, b byte) {
			buf := make([]byte, len(p.Chain))
			copy(buf, p.Chain)
			t := 0
			for {
				u := findNextNonSpace(buf[t+1:])
				if u == -1 {
					break
				}

				if buf[t] == a || buf[t] == b {
					buf[t] = ' '
				}

				t += u + 1
			}

			length := countBufLength(reactChain(buf))
			done <- length
		})(a, b)
	}

	count := 26
	shortest := -1
	for count > 0 {
		length := <-done
		count--
		if shortest == -1 || length < shortest {
			shortest = length
		}
	}

	return shortest
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)

	p := Puzzle{}
	for scanner.Scan() {
		p.Chain = scanner.Bytes()
	}
	if err := scanner.Err(); err != nil {
		panic(err.Error())
	}

	fmt.Println(p.Solution1())
	fmt.Println(p.Solution2())
}
