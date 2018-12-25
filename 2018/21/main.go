package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
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

	LastResult [][6]int
}

func (p *Puzzle) RunProgram() {
	p.LastResult = make([][6]int, len(p.ROM)/4)
	ticks := 0
	p.PrintState()
	waitEnter()
	for {
		if p.IP < 0 || p.IP*4 >= len(p.ROM) {
			fmt.Println("PANIC! Access OOB @", p.IP, "after", ticks, "iterations")
			return
		}

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
		copy(p.LastResult[p.IP][:], p.Regs[:])

		p.IP++
		ticks++

		//fmt.Printf("%3c %3c %3c %3c %3c %3c <- %d\n", regs[0], regs[1], regs[2], regs[3], regs[4], regs[5], ticks)
		p.PrintState()
		waitEnter()
	}
}

func (p *Puzzle) PrintState() {
	fmt.Printf("\033c")

	color.New(color.FgWhite).Printf("IP: ")
	color.New(color.FgHiWhite).Printf("%2d", p.IP)
	color.New(color.FgWhite).Printf("\tIPreg: ")
	color.New(color.FgHiCyan).Printf("r%d\n", p.IPreg)
	color.New(color.FgWhite).Printf("Registers: ")
	fmt.Printf("[%d, %d, %d, %d, %d, %d]\n", p.Regs[0], p.Regs[1], p.Regs[2], p.Regs[3], p.Regs[4], p.Regs[5])
	color.New(color.FgWhite).Printf("%s\n", strings.Repeat("-", 93))
	for t := 0; t < len(p.ROM); t += 4 {
		o := opcodes[p.ROM[t]]

		gutterColor := color.New(color.FgWhite)
		normalColor := color.New(color.FgHiWhite)
		registerColor := color.New(color.FgHiCyan)
		if t/4 == p.IP {
			gutterColor = color.New(color.FgHiWhite).Add(color.BgBlue)
			normalColor.Add(color.BgBlue)
			registerColor = color.New(color.FgHiCyan).Add(color.BgBlue)
		}

		gutterColor.Printf("%02d|", t/4)
		normalColor.Printf(" %s", o.Name)
		if o.A == ArgumentTypeRegister {
			registerColor.Printf(" % 6sr%d", " ", p.ROM[t+1])
		} else {
			normalColor.Printf(" %8d", p.ROM[t+1])
		}
		if o.B == ArgumentTypeRegister {
			registerColor.Printf(" % 6sr%d", " ", p.ROM[t+2])
		} else {
			normalColor.Printf(" %8d", p.ROM[t+2])
		}
		registerColor.Printf(" % 6sr%d", " ", p.ROM[t+3])

		regs := make([]string, len(p.Regs))
		for t, r := range p.LastResult[t/4] {
			if r > 100000000 {
				regs[t] += ">9999999"
			} else {
				regs[t] += fmt.Sprintf("%8d", r)
			}
		}
		gutterColor.Printf(" = [%s]", strings.Join(regs, ","))

		fmt.Printf("\n")
	}
}

func main() {
	// Solution 1
	if true {
		fmt.Println("Solution 1\n---")
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

			//			fmt.Println(op, a, b, c)
		}
		if err := scanner.Err(); err != nil {
			panic(err.Error())
		}

		fmt.Printf("\n")
		p.RunProgram()
	}
}
