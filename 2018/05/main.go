package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

func findPrevNonSpace(buf []byte) int {
	for t := len(buf) - 1; t >= 0; t-- {
		if buf[t] != ' ' {
			return t
		}
	}
	return -1
}

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
			t = findPrevNonSpace(chain[:i])
			if t == -1 {
				t = findNextNonSpace(chain)
				if t == -1 {
					break
				}
			}
		}
	}

	return chain
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
	var a, b byte
	shortest := math.MaxInt32
	for a = 'A'; a <= 'Z'; a++ {
		b = a + 32

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
		if length < shortest {
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
