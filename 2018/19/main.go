package main

import (
	"bufio"
	"fmt"
	"os"
	"time"
)

type ArgumentType int

const (
	ArgumentTypeNone ArgumentType = iota
	ArgumentTypeRegister
	ArgumentTypeValue
)

type Opcode struct {
	Name string
	A    ArgumentType
	B    ArgumentType
	Exe  func(a, b int) int
}

var opcodes = []Opcode{
	{
		Name: "banr",
		A:    ArgumentTypeRegister,
		B:    ArgumentTypeRegister,
		Exe:  ban,
	},
	{
		Name: "eqrr",
		A:    ArgumentTypeRegister,
		B:    ArgumentTypeRegister,
		Exe:  eq,
	},
	{
		Name: "setr",
		A:    ArgumentTypeRegister,
		B:    ArgumentTypeNone,
		Exe:  set,
	},
	{
		Name: "eqir",
		A:    ArgumentTypeValue,
		B:    ArgumentTypeRegister,
		Exe:  eq,
	},
	{
		Name: "bori",
		A:    ArgumentTypeRegister,
		B:    ArgumentTypeValue,
		Exe:  bor,
	},
	{
		Name: "muli",
		A:    ArgumentTypeRegister,
		B:    ArgumentTypeValue,
		Exe:  mul,
	},
	{
		Name: "bani",
		A:    ArgumentTypeRegister,
		B:    ArgumentTypeValue,
		Exe:  ban,
	},
	{
		Name: "borr",
		A:    ArgumentTypeRegister,
		B:    ArgumentTypeRegister,
		Exe:  bor,
	},
	{
		Name: "gtir",
		A:    ArgumentTypeValue,
		B:    ArgumentTypeRegister,
		Exe:  gt,
	},
	{
		Name: "gtrr",
		A:    ArgumentTypeRegister,
		B:    ArgumentTypeRegister,
		Exe:  gt,
	},
	{
		Name: "addi",
		A:    ArgumentTypeRegister,
		B:    ArgumentTypeValue,
		Exe:  add,
	},
	{
		Name: "gtri",
		A:    ArgumentTypeRegister,
		B:    ArgumentTypeValue,
		Exe:  gt,
	},
	{
		Name: "eqri",
		A:    ArgumentTypeRegister,
		B:    ArgumentTypeValue,
		Exe:  eq,
	},
	{
		Name: "addr",
		A:    ArgumentTypeRegister,
		B:    ArgumentTypeRegister,
		Exe:  add,
	},
	{
		Name: "mulr",
		A:    ArgumentTypeRegister,
		B:    ArgumentTypeRegister,
		Exe:  mul,
	},
	{
		Name: "seti",
		A:    ArgumentTypeValue,
		B:    ArgumentTypeNone,
		Exe:  set,
	},
}

func add(a, b int) int {
	return a + b
}

func mul(a, b int) int {
	return a * b
}

func ban(a, b int) int {
	return a & b
}

func bor(a, b int) int {
	return a | b
}

func set(a, b int) int {
	return a
}

func gt(a, b int) int {
	if a > b {
		return 1
	}
	return 0
}

func eq(a, b int) int {
	if a == b {
		return 1
	}
	return 0
}

func waitEnter() {
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}

type Test struct {
	Before [4]int
	Op     [4]int
	After  [4]int
}

func sliceEqual(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for t := range a {
		if a[t] != b[t] {
			return false
		}
	}
	return true
}

type Puzzle struct {
	Regs  [6]int
	ROM   []int
	IP    int
	IPreg int
}

func (p *Puzzle) RunProgram() {
	oldRegs := [6]int{}
	regs := [6]rune{}
	ticks := 0
	start := time.Now()
	for {
		if p.IP < 0 || p.IP*4 >= len(p.ROM) {
			fmt.Println("PANIC! Access OOB @", p.IP, "after", ticks, "iterations")
			return
		}

		copy(oldRegs[:], p.Regs[:])

		p.Regs[p.IPreg] = p.IP
		op := p.ROM[p.IP*4 : (p.IP*4)+4]
		o := opcodes[op[0]]
		a := op[1]
		b := op[2]
		c := op[3]
		if o.A == ArgumentTypeRegister {
			a = p.Regs[a]
		}
		if o.B == ArgumentTypeRegister {
			b = p.Regs[b]
		}

		p.Regs[c] = o.Exe(a, b)
		p.IP = p.Regs[p.IPreg]

		p.IP++
		ticks++

		for t := range regs {
			if oldRegs[t] == p.Regs[t] {
				regs[t] = '.'
				continue
			}
			regs[t] = '*'
		}
		if ticks%10000000 == 0 {
			fmt.Println(time.Now().Sub(start), ticks)
		}
		//p.PrintRegs()
		fmt.Printf("%3c %3c %3c %3c %3c %3c <- %d\n", regs[0], regs[1], regs[2], regs[3], regs[4], regs[5], ticks)
	}
}

func (p *Puzzle) PrintRegs() {
	fmt.Println("Regs:", p.Regs)
}

func main() {
	// Program 1
	if true {
		fmt.Println("Program 1\n---")
		p := Puzzle{}

		input, err := os.Open("input.txt")
		if err != nil {
			panic(err.Error())
		}
		defer input.Close()
		scanner := bufio.NewScanner(input)

		for scanner.Scan() {
			var op string
			var a, b, c int
			fmt.Sscanf(scanner.Text(), "%s %d %d %d", &op, &a, &b, &c)

			if op == "#ip" {
				p.IPreg = a
				continue
			}

			for k, o := range opcodes {
				if o.Name == op {
					p.ROM = append(p.ROM, k, a, b, c)
					break
				}
			}

			fmt.Println(op, a, b, c)
		}
		if err := scanner.Err(); err != nil {
			panic(err.Error())
		}

		fmt.Printf("\n")
		p.RunProgram()
		p.PrintRegs()
	}

	// Program 2
	// NOTE: This runs for a looooong time, time to reverseengineer
	if false {
		fmt.Println("\nProgram 2\n---")
		p := Puzzle{}
		p.Regs[0] = 1

		// Bruteforcing after lengthy debugging and trying to understanding what the assembly is doing
		//p.Regs = [6]int{0, 0, 1, 9, 10551326, 10551325}
		//p.Regs = [6]int{1, 0, 10551326, 9, 10551326, 10551326}
		//p.Regs = [6]int{1, 0, 10551326, 9, 10551326, 10551325}
		//p.IP = 9

		input, err := os.Open("input.txt")
		if err != nil {
			panic(err.Error())
		}
		defer input.Close()
		scanner := bufio.NewScanner(input)

		for scanner.Scan() {
			var op string
			var a, b, c int
			fmt.Sscanf(scanner.Text(), "%s %d %d %d", &op, &a, &b, &c)

			if op == "#ip" {
				p.IPreg = a
				continue
			}

			for k, o := range opcodes {
				if o.Name == op {
					p.ROM = append(p.ROM, k, a, b, c)
					break
				}
			}

			fmt.Println(op, a, b, c)
		}
		if err := scanner.Err(); err != nil {
			panic(err.Error())
		}

		fmt.Printf("\n")
		p.RunProgram()
		p.PrintRegs()
	}
	return

	// Attempt to re-impl algo
	var r0, r1, r2, r3, r4, r5 int

	r0 = 1
	r4 += 2
	r4 *= r4
	r4 *= 19
	r4 *= 11
	r1 += 4
	r1 += 22
	r1 += 2
	r4 += r1

	r1 = 27
	r1 *= 28
	r1 += 29
	r1 *= 30
	r1 *= 14
	r1 *= 32
	r4 += r1
	r0 = 0

	//start
	r2 = 1
	//start1
	r5 = 1
	//start3
	tick := 0
	for {
		r1 = r2 * r5
		if r1 == r4 {
			//check
			if r5 > r4 {
				//skip
				r2++
				if r2 > r4 {
					fmt.Println("DONE", r0, r1, r2, r3, r4, r5)
					return
				} else {
					r5 = 1
				}
			}
		} else {
			r5++
		}

		if tick%10000 == 0 {
			fmt.Println(tick, "-", r0, r1, r2, r3, r4, r5)
		}
		tick++
	}
}
