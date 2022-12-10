package main

import "core:fmt"
import "core:os"

parse_input_file :: proc(filepath: string) -> [dynamic]Game {
	data, ok := os.read_entire_file_from_filename(filepath)
	if !ok {
		panic("oh no, could not read file")
	}

	games := make([dynamic]Game, 0, 100)

	for i := 0; i < len(data); i += 4 {
		current := data[i:i + 4]
		append(&games, Game{rune(current[0]), rune(current[2])})
	}

	return games
}

Game :: [2]rune

// rock 0, paper 1, scissors 2
losing_moves := []int{1, 2, 0}
winning_moves := []int{2, 0, 1}

get_score :: proc(game: Game) -> int {
	score: int
	elf_shape := int(game[0]) - 'A'
	player_shape := int(game[1]) - 'X'

	if elf_shape == player_shape {
		score = 3
	} else {
		if losing_moves[elf_shape] == player_shape {
			score = 6
		}
	}
	score += player_shape + 1

	return score
}

task1 :: proc(games: []Game) -> int {
	total_score: int
	for game in games {
		score := get_score(game)
		total_score += score
	}
	return total_score
}

task2 :: proc(games: []Game) -> int {
	total_score: int
	for game in games {
		move: rune
		elf_shape := int(game[0]) - 'A'
		switch game[1] {
		// Lose
		case 'X':
			move = rune(winning_moves[elf_shape] + 'X')
		// Draw
		case 'Y':
			move = rune(elf_shape + 'X')
		// Win
		case 'Z':
			move = rune(losing_moves[elf_shape] + 'X')
		}
		score := get_score({game[0], move})
		total_score += score
	}
	return total_score
}

main :: proc() {
	games := parse_input_file("input.txt")
	defer delete(games)

	result1 := task1(games[:])
	fmt.printf("Task 1 result: %v\n", result1)

	result2 := task2(games[:])
	fmt.printf("Task 2 result: %v\n", result2)
}
