package main

import (
	"bufio"
	"fmt"
	"os"
)

func reactChain(chain []byte) []byte {
	result := make([]byte, 0, len(chain))
	for t := 0; t < len(chain)-1; t++ {
		a, b := chain[t], chain[t+1]

		if a-b == 32 || b-a == 32 {
			result = append(result, chain[:t]...)
			result = append(result, chain[t+2:]...)
			break
		}
	}

	if len(result) == 0 {
		return chain
	}

	return reactChain(result)
}

type Puzzle struct {
	Chain []byte
}

func (p *Puzzle) Solution1() int {
	return len(reactChain(p.Chain))
}

func (p *Puzzle) Solution2() int {
	return 0
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
