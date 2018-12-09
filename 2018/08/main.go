package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/davecgh/go-spew/spew"
)

type Puzzle struct {
	Root *Node
}

func (p *Puzzle) Solution1() int {
	return p.Root.SumMetadata()
}

func (p *Puzzle) Solution2() int {
	return p.Root.Value()
}

/*
2 3 0 3 10 11 12 1 1 0 1 99 2 1 1 2
A----------------------------------
    B----------- C-----------
                     D-----
*/
type Node struct {
	Children []*Node
	Metadata []int
}

func (n *Node) SumMetadata() int {
	total := 0
	for _, child := range n.Children {
		total += child.SumMetadata()
	}
	for _, meta := range n.Metadata {
		total += meta
	}
	return total
}

func (n *Node) Value() int {
	if len(n.Children) == 0 {
		return n.SumMetadata()
	}

	total := 0
	for _, meta := range n.Metadata {
		if meta > 0 && meta <= len(n.Children) {
			total += n.Children[meta-1].Value()
		}
	}
	return total
}

func readNode(scanner *bufio.Scanner) *Node {
	if !scanner.Scan() {
		return nil
	}
	if err := scanner.Err(); err != nil {
		panic(err.Error())
	}
	children, _ := strconv.Atoi(scanner.Text())

	if !scanner.Scan() {
		return nil
	}
	if err := scanner.Err(); err != nil {
		panic(err.Error())
	}
	metadata, _ := strconv.Atoi(scanner.Text())

	node := Node{
		Children: make([]*Node, children),
		Metadata: make([]int, metadata),
	}
	for t := 0; t < children; t++ {
		node.Children[t] = readNode(scanner)
	}
	for t := 0; t < metadata; t++ {
		if !scanner.Scan() {
			return nil
		}
		if err := scanner.Err(); err != nil {
			panic(err.Error())
		}
		meta, _ := strconv.Atoi(scanner.Text())
		node.Metadata[t] = meta
	}

	return &node
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords)

	p := Puzzle{}
	p.Root = readNode(scanner)
	spew.Dump(p.Root)

	fmt.Println(p.Solution1())
	fmt.Println(p.Solution2())
}
