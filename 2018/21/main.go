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
	Ticks int

	LastResult        [][6]int
	Isljmp            []int // Instructions since last jump
	PrevIsljmp        []int
	LoopCount         int
	LastJumpFrom      int
	LastJumpTo        int
	PotentialLoopFrom int
	PotentialLoopTo   int
}

func (p *Puzzle) RunProgram() {
	p.PrevIsljmp = make([]int, 0, 10)
	p.Isljmp = make([]int, 0, 10)
	p.LastResult = make([][6]int, len(p.ROM)/4)
	p.LastJumpFrom = -1

	p.PrintState()
	waitEnter()
	for {
		if p.IP < 0 || p.IP*4 >= len(p.ROM) {
			color.New(color.BgRed).Add(color.FgHiWhite).Println("PANIC! Access OOB @", p.IP, "after", p.Ticks, "iterations")
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

		// Detect inf loops
		if c == p.IPreg {
			p.LastJumpFrom = p.IP
		} else {
			p.Isljmp = append(p.Isljmp, op[0])
		}

		p.Regs[c] = o.Exe(a, b)
		p.IP = p.Regs[p.IPreg]
		copy(p.LastResult[p.IP][:], p.Regs[:])

		p.IP++
		p.Ticks++

		if c == p.IPreg {
			p.LastJumpTo = p.IP

			if p.IP >= p.PotentialLoopTo && p.IP <= p.PotentialLoopFrom {
				if p.IP == p.PotentialLoopTo {
					if sliceEqual(p.Isljmp, p.PrevIsljmp) {
						p.LoopCount++
					} else {
						p.LoopCount = 0
					}
					p.PrevIsljmp = p.Isljmp
					p.Isljmp = make([]int, 0, 10)
				}
			} else {
				p.PotentialLoopFrom = p.LastJumpFrom
				p.PotentialLoopTo = p.LastJumpTo
				p.PrevIsljmp = p.Isljmp
				p.Isljmp = make([]int, 0, 10)
			}
		}

		p.PrintState()
		waitEnter()
	}
}

func (p *Puzzle) PrintState() {
	fmt.Printf("\033c")

	color.New(color.FgWhite).Printf("IP: ")
	color.New(color.FgHiWhite).Printf("%2d", p.IP)
	color.New(color.FgWhite).Printf("\tIPreg: ")
	color.New(color.FgHiCyan).Printf("r%d", p.IPreg)
	color.New(color.FgWhite).Printf("\tTicks: ")
	color.New(color.FgHiWhite).Printf("%d\n", p.Ticks)
	color.New(color.FgWhite).Printf("Registers: ")
	fmt.Printf("[%d, %d, %d, %d, %d, %d]\n", p.Regs[0], p.Regs[1], p.Regs[2], p.Regs[3], p.Regs[4], p.Regs[5])
	color.New(color.FgWhite).Printf("%s\n", strings.Repeat("-", 93))
	for t := 0; t < len(p.ROM); t += 4 {
		i := t / 4
		o := opcodes[p.ROM[t]]

		gutterColor := color.New(color.FgWhite)
		normalColor := color.New(color.FgHiWhite)
		registerColor := color.New(color.FgHiCyan)
		if i == p.IP {
			gutterColor = color.New(color.FgHiWhite).Add(color.BgBlue)
			normalColor.Add(color.BgBlue)
			registerColor = color.New(color.FgHiCyan).Add(color.BgBlue)
		}

		gutterColor.Printf("%02d|", i)
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
		for t, r := range p.LastResult[i] {
			if r > 100000000 {
				regs[t] += ">9999999"
			} else {
				regs[t] += fmt.Sprintf("%8d", r)
			}
		}
		gutterColor.Printf(" = [%s]", strings.Join(regs, ","))

		if p.LastJumpFrom > -1 {
			if i >= p.PotentialLoopTo && i <= p.PotentialLoopFrom {
				if i == p.PotentialLoopTo {
					registerColor.Printf(" *")
				}
				if i > p.PotentialLoopTo && i < p.PotentialLoopFrom {
					registerColor.Printf(" |")
				}
				if i == p.PotentialLoopFrom {
					registerColor.Printf(" ^")
				}
			}

			if i == p.LastJumpTo {
				gutterColor.Printf(" *")
			}
			if (i > p.LastJumpTo && i < p.LastJumpFrom) || (i < p.LastJumpTo && i > p.LastJumpFrom) {
				gutterColor.Printf(" |")
			}
			if i == p.LastJumpFrom {
				if p.LastJumpFrom < p.LastJumpTo {
					gutterColor.Printf(" v")
				} else {
					gutterColor.Printf(" ^")
				}
				gutterColor.Printf(" jump")
			}

			if i == p.PotentialLoopTo {
				if p.LoopCount > 0 {
					registerColor.Printf(" loop")
				}
			}

			if i == p.PotentialLoopFrom {
				if p.LoopCount > 0 {
					registerColor.Printf(" iteration %d", p.LoopCount)
				}
			}

			/*
				if i == p.IP {
					if (p.IP > p.LastJumpTo && p.IP < p.LastJumpFrom) || (p.IP < p.LastJumpTo && p.IP > p.LastJumpFrom) {
						inst := make([]string, len(p.Isljmp))
						for t, o := range p.Isljmp {
							inst[t] = opcodes[o].Name
						}
						gutterColor.Printf(" [%s]", strings.Join(inst, ","))
					}
				}
			*/
		}

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
