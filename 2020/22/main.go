package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Deck ...
type Deck []int

func sameDeck(a, b Deck) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

// Task1 ...
func Task1(input []Deck, debug bool) int {
	p1, p2 := input[0], input[1]
	round := 1

	for len(p1) > 0 && len(p2) > 0 {
		if debug {
			fmt.Printf("\n-- Round %d --\n", round)
			fmt.Printf("Player 1's deck: %+v\n", p1)
			fmt.Printf("Player 2's deck: %+v\n", p2)
		}

		a, b := p1[0], p2[0]
		if debug {
			fmt.Printf("Player 1 plays: %d\n", a)
			fmt.Printf("Player 2 plays: %d\n", b)
		}
		if a > b {
			if debug {
				fmt.Printf("Player 1 wins the round!\n")
			}
			p1 = append(p1[1:], a, b)
			p2 = p2[1:]
		} else {
			if debug {
				fmt.Printf("Player 2 wins the round!\n")
			}
			p1 = p1[1:]
			p2 = append(p2[1:], b, a)
		}
		round++
	}

	if debug {
		fmt.Println("\n== Post-game results ==")
		fmt.Printf("Player 1's deck: %+v\n", p1)
		fmt.Printf("Player 2's deck: %+v\n", p2)
		fmt.Println()
	}

	var deck Deck
	if len(p1) > 0 {
		deck = p1
	} else {
		deck = p2
	}

	var result int
	for i := 0; i < len(deck); i++ {
		if debug {
			fmt.Printf("+ %2d * %2d\n", deck[i], len(deck)-i)
		}
		result += deck[i] * (len(deck) - i)
	}
	return result
}

func playRecursiveCombat(game int, p1, p2 Deck, debug bool) (int, Deck, Deck) {
	if debug {
		fmt.Printf("\n=== Game %d ===\n", game)
	}

	round := 1
	memory := make(map[int][]Deck)
	for len(p1) > 0 && len(p2) > 0 {
		if debug {
			fmt.Printf("\n-- Round %d (Game %d) --\n", round, game)
			fmt.Printf("Player 1's deck: %+v\n", p1)
			fmt.Printf("Player 2's deck: %+v\n", p2)
		}

		a, b := p1[0], p2[0]
		p1 = p1[1:]
		p2 = p2[1:]
		if debug {
			fmt.Printf("Player 1 plays: %d\n", a)
			fmt.Printf("Player 2 plays: %d\n", b)
		}
		var playerWon int
		if len(p1) >= a && len(p2) >= b {
			if debug {
				fmt.Println("Playing a sub-game to determine the winner...")
			}
			p1C := make([]int, a)
			p2C := make([]int, b)
			copy(p1C, p1[:a])
			copy(p2C, p2[:b])
			playerWon, _, _ = playRecursiveCombat(game+1, p1C, p2C, debug)
			if debug {
				fmt.Printf("The winner of game %d is player %d\n", game+1, playerWon)
			}
		} else {
			if a > b {
				playerWon = 1
			} else {
				playerWon = 2
			}
			if debug {
				fmt.Printf(
					"Player %d wins round %d of game %d!\n", playerWon, round, game,
				)
			}
		}
		if playerWon == 1 {
			p1 = append(p1, a, b)
		} else {
			p2 = append(p2, b, a)
		}

		for _, m := range memory {
			if sameDeck(m[0], p1) && sameDeck(m[1], p2) {
				if debug {
					fmt.Println("infinite loop, p1 won")
				}
				return 1, p1, p2
			}
		}

		memory[round] = []Deck{p1, p2}
		round++
	}

	if len(p1) > len(p2) {
		return 1, p1, p2
	}
	return 2, p1, p2
}

// Task2 ...
func Task2(input []Deck, debug bool) int {
	p1, p2 := input[0], input[1]
	game := 1
	player, p1, p2 := playRecursiveCombat(game, p1, p2, debug)

	if debug {
		fmt.Println("\n== Post-game results ==")
		fmt.Printf("Player 1's deck: %+v\n", p1)
		fmt.Printf("Player 2's deck: %+v\n", p2)
		fmt.Println()
	}

	var deck Deck
	if player == 1 {
		deck = p1
	} else {
		deck = p2
	}

	var result int
	for i := 0; i < len(deck); i++ {
		if debug {
			fmt.Printf("+ %2d * %2d\n", deck[i], len(deck)-i)
		}
		result += deck[i] * (len(deck) - i)
	}
	return result
}

func readInput(reader io.Reader) []Deck {
	input := make([]Deck, 2)
	input[0] = make([]int, 0, 10)
	input[1] = make([]int, 0, 10)

	var current int
	scanner := bufio.NewScanner(reader)
	scanner.Scan()
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			scanner.Scan()
			current++
			continue
		}
		val, _ := strconv.Atoi(line)
		input[current] = append(input[current], val)
	}

	return input
}

func main() {
	var debug bool
	flag.BoolVar(&debug, "debug", false, "debug")
	flag.Parse()

	var input []Deck

	file, _ := os.Open("input.txt")
	input = readInput(file)
	file.Close()

	result := Task1(input, debug)
	fmt.Printf("Task 1: %d\n", result)

	result = Task2(input, debug)
	fmt.Printf("Task 2: %d\n", result)
}
