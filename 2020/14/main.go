package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Op constants.
const (
	OpMask string = "mask"
	OpMem         = "mem"
)

// Op ...
type Op struct {
	Type string
	Mask string
	Addr int
	Val  int
}

// PortComputer ...
type PortComputer struct {
	Mask string
	Mem  map[int]int
}

// Set ...
func (p *PortComputer) Set(addr int, val int) {
	bit := 35
	for _, r := range p.Mask {
		if r != 'X' {
			if r == '1' {
				val |= 1 << bit
			} else {
				val &= ^(1 << bit)
			}
		}
		bit--
	}
	p.Mem[addr] = val
}

// Task1 ...
func Task1(input []Op) int {
	p := PortComputer{
		Mem: make(map[int]int),
	}

	for _, op := range input {
		switch op.Type {
		case OpMask:
			p.Mask = op.Mask
		case OpMem:
			p.Set(op.Addr, op.Val)
		}
	}

	var result int
	for _, v := range p.Mem {
		result += int(v)
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

	var input []Op
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), "=")
		if strings.HasPrefix(parts[0], "mask") {
			input = append(input, Op{
				Type: OpMask,
				Mask: parts[1][1:],
			})
		} else {
			var addr, val int
			fmt.Sscanf(parts[0], "mem[%d]", &addr)
			val, _ = strconv.Atoi(parts[1][1:])
			input = append(input, Op{
				Type: OpMem,
				Addr: addr,
				Val:  val,
			})
		}
	}
	file.Close()

	result := Task1(input)
	fmt.Printf("Task 1: %d\n", result)

	result = Task2()
	fmt.Printf("Task 2: %d\n", result)
}
