package main

import (
	"bufio"
	"fmt"
	"os"
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
		Name: "addr",
		A:    ArgumentTypeRegister,
		B:    ArgumentTypeRegister,
		Exe:  add,
	},
	{
		Name: "addi",
		A:    ArgumentTypeRegister,
		B:    ArgumentTypeValue,
		Exe:  add,
	},
	{
		Name: "mulr",
		A:    ArgumentTypeRegister,
		B:    ArgumentTypeRegister,
		Exe:  mul,
	},
	{
		Name: "muli",
		A:    ArgumentTypeRegister,
		B:    ArgumentTypeValue,
		Exe:  mul,
	},
	{
		Name: "banr",
		A:    ArgumentTypeRegister,
		B:    ArgumentTypeRegister,
		Exe:  ban,
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
		Name: "bori",
		A:    ArgumentTypeRegister,
		B:    ArgumentTypeValue,
		Exe:  bor,
	},
	{
		Name: "setr",
		A:    ArgumentTypeRegister,
		B:    ArgumentTypeNone,
		Exe:  set,
	},
	{
		Name: "seti",
		A:    ArgumentTypeValue,
		B:    ArgumentTypeNone,
		Exe:  set,
	},
	{
		Name: "gtir",
		A:    ArgumentTypeValue,
		B:    ArgumentTypeRegister,
		Exe:  gt,
	},
	{
		Name: "gtri",
		A:    ArgumentTypeRegister,
		B:    ArgumentTypeValue,
		Exe:  gt,
	},
	{
		Name: "gtrr",
		A:    ArgumentTypeRegister,
		B:    ArgumentTypeRegister,
		Exe:  gt,
	},
	{
		Name: "eqir",
		A:    ArgumentTypeValue,
		B:    ArgumentTypeRegister,
		Exe:  eq,
	},
	{
		Name: "eqri",
		A:    ArgumentTypeRegister,
		B:    ArgumentTypeValue,
		Exe:  eq,
	},
	{
		Name: "eqrr",
		A:    ArgumentTypeRegister,
		B:    ArgumentTypeRegister,
		Exe:  eq,
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
	Tests      []Test
	Potentials [16][]string

	Instructions [16]*Opcode
	Registers    [4]int
	Program      []int
	PC           int
}

func (p *Puzzle) IdentifyPotentials() {
	total := 0
	for t, test := range p.Tests {
		potentials := 0
		fmt.Printf("Test #%d\n\tI: %+v\n\tOP:%+v\n\tT: %+v\n", t, test.Before, test.Op, test.After)
		for _, op := range opcodes {
			registers := [4]int{}
			copy(registers[:], test.Before[:])

			a := test.Op[1]
			b := test.Op[2]
			c := test.Op[3]
			if op.A == ArgumentTypeRegister {
				a = registers[a]
			}
			if op.B == ArgumentTypeRegister {
				b = registers[b]
			}

			registers[c] = op.Exe(a, b)
			fmt.Printf("\t%s -> A:%d B:%d C:%d R:%+v", op.Name, a, b, c, registers)
			if sliceEqual(registers[:], test.After[:]) {
				fmt.Printf(" OK")
				potentials++

				o := test.Op[0]
				if p.Potentials[o] == nil {
					p.Potentials[o] = make([]string, 0, 10)
				}
				found := false
				for _, i := range p.Potentials[o] {
					if i == op.Name {
						found = true
						break
					}
				}
				if !found {
					p.Potentials[o] = append(p.Potentials[test.Op[0]], op.Name)
				}
			}
			fmt.Printf("\n")
		}
		fmt.Println("Potentials:", potentials)
		if potentials >= 3 {
			total++
		}
	}
	fmt.Println("\nTotal samples with more >=3 potentials:", total)

	for t, i := range p.Potentials {
		fmt.Printf("%d -> %+v\n", t, i)
	}
}

func (p *Puzzle) IdentifyInstructions() {
	for {
		found := false
		for j, i := range p.Potentials {
			if len(i) != 1 {
				continue
			}

			op := i[0]
			for u := range opcodes {
				if opcodes[u].Name == op {
					p.Instructions[j] = &opcodes[u]
					break
				}
			}

			for h, u := range p.Potentials {
				for t, v := range u {
					if v == op {
						p.Potentials[h] = append(p.Potentials[h][:t], p.Potentials[h][t+1:]...)
						break
					}
				}
			}

			found = true
			break
		}
		if !found {
			break
		}
	}

	fmt.Println("\nIdentified:")
	for t, i := range p.Instructions {
		fmt.Printf("%02d -> %s\n", t, i.Name)
	}
}

func (p *Puzzle) RunProgram() {
	for p.PC = 0; p.PC < len(p.Program); p.PC += 4 {
		op := p.Program[p.PC : p.PC+4]
		o := p.Instructions[op[0]]
		a := op[1]
		b := op[2]
		c := op[3]
		if o.A == ArgumentTypeRegister {
			a = p.Registers[a]
		}
		if o.B == ArgumentTypeRegister {
			b = p.Registers[b]
		}
		p.Registers[c] = o.Exe(a, b)
	}
}

func (p *Puzzle) PrintRegisters() {
	fmt.Println("Registers:", p.Registers)
}

func main() {
	p := Puzzle{}

	{
		input, err := os.Open("input1.txt")
		if err != nil {
			panic(err.Error())
		}
		defer input.Close()
		scanner := bufio.NewScanner(input)

		for scanner.Scan() {
			test := Test{}
			if err := scanner.Err(); err != nil {
				panic(err.Error())
			}
			fmt.Sscanf(scanner.Text(), "Before: [%d, %d, %d, %d]", &test.Before[0], &test.Before[1], &test.Before[2], &test.Before[3])
			scanner.Scan()
			if err := scanner.Err(); err != nil {
				panic(err.Error())
			}
			fmt.Sscanf(scanner.Text(), "%d %d %d %d", &test.Op[0], &test.Op[1], &test.Op[2], &test.Op[3])
			scanner.Scan()
			if err := scanner.Err(); err != nil {
				panic(err.Error())
			}
			fmt.Sscanf(scanner.Text(), "After: [%d, %d, %d, %d]", &test.After[0], &test.After[1], &test.After[2], &test.After[3])
			p.Tests = append(p.Tests, test)

			scanner.Scan()
			if err := scanner.Err(); err != nil {
				panic(err.Error())
			}
		}
		if err := scanner.Err(); err != nil {
			panic(err.Error())
		}
	}
	{
		input, err := os.Open("input2.txt")
		if err != nil {
			panic(err.Error())
		}
		defer input.Close()
		scanner := bufio.NewScanner(input)

		for scanner.Scan() {
			op := [4]int{}
			fmt.Sscanf(scanner.Text(), "%d %d %d %d", &op[0], &op[1], &op[2], &op[3])
			p.Program = append(p.Program, op[:]...)
		}
		if err := scanner.Err(); err != nil {
			panic(err.Error())
		}
	}

	p.IdentifyPotentials()
	p.IdentifyInstructions()
	p.RunProgram()
	p.PrintRegisters()
}
