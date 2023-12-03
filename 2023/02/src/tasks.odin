package main

import "core:bytes"
import "core:fmt"
import "core:io"
import "core:os"
import "core:strconv"
import "core:strings"

Color :: enum {
	Red,
	Green,
	Blue,
}

Combination :: struct {
	num: int,
	col: Color,
}

Set :: [dynamic]Combination

Game :: struct {
	id:   int,
	sets: [dynamic]Set,
}

Input :: struct {
	games: [dynamic]Game,
}

Result1 :: distinct int
Result2 :: distinct int


// --- Input --- //
parse_input_file :: proc(filepath: string) -> Input {
	input: Input

	raw_data, ok := os.read_entire_file_from_filename(filepath)
	if !ok {
		panic("oh no, could not read file")
	}

	buf: bytes.Buffer
	bytes.buffer_init(&buf, raw_data)
	defer bytes.buffer_destroy(&buf)

	input.games = make([dynamic]Game, 0, 10)

	line: string
	err: io.Error
	for ; err != .EOF; line, err = bytes.buffer_read_string(&buf, '\n') {
		if line == "" || line == "\n" {
			continue
		}

		line = line[:len(line) - 1]

		game_and_sets := strings.split(line, ": ")
		current_game := strings.split(game_and_sets[0], " ")
		sets := strings.split(game_and_sets[1], "; ")

		game: Game
		game.id = strconv.atoi(current_game[1])
		game.sets = make([dynamic]Set, 0, 10)

		for set in sets {
			parsed_set := make([dynamic]Combination, 0, 10)

			combinations := strings.split(set, ", ")
			for combination in combinations {
				number_and_color := strings.split(combination, " ")

				col: Color
				switch number_and_color[1] {
				case "red":
					col = .Red
				case "green":
					col = .Green
				case "blue":
					col = .Blue
				}

				append(&parsed_set, Combination{num = strconv.atoi(number_and_color[0]), col = col})
			}

			append(&game.sets, parsed_set)
		}

		append(&input.games, game)
	}

	return input
}

free_input :: proc(input: ^Input) {
	for game in input.games {
		for set in game.sets {
			delete(set)
		}
		delete(game.sets)
	}
	delete(input.games)
}
// --- Input --- //


// --- Task 1 --- //
task1_rules := map[Color]int {
	.Red   = 12,
	.Green = 13,
	.Blue  = 14,
}
run_task1 :: proc(input: ^Input, debug: bool) -> Result1 {
	result: Result1

	if debug {
		fmt.printf("input: %#v\n", input)
	}

	for game in input.games {
		ok := true
		for set in game.sets {
			for combination in set {
				if task1_rules[combination.col] < combination.num {
					ok = false
					break
				}
			}
		}
		if ok {
			if debug {
				fmt.printf("Possible game: %d\n", game.id)
			}
			result += Result1(game.id)
		}
	}

	return result
}

print_result1 :: proc(result: ^Result1) {
	fmt.printf("Task 1: %d\n", result^)
}
// --- Task 1 --- //


// --- Task 2 --- //
run_task2 :: proc(input: ^Input, debug: bool) -> Result2 {
	result: Result2

	if debug {
		fmt.printf("input: %#v\n", input)
	}

	for game in input.games {
		fewest_cubes_necessary := map[Color]int {
			.Red   = 0,
			.Green = 0,
			.Blue  = 0,
		}
		for set in game.sets {
			for combination in set {
				if num, _ := fewest_cubes_necessary[combination.col]; combination.num > num {
					fewest_cubes_necessary[combination.col] = combination.num
				}
			}
		}
		power := fewest_cubes_necessary[.Red] * fewest_cubes_necessary[.Green] * fewest_cubes_necessary[.Blue]
		if debug {
			fmt.printf("Possible game: %d\n", game.id)
			fmt.printf("Fewest cubes necessary: %#v -> %d\n", fewest_cubes_necessary, power)
		}
		result += Result2(power)
	}

	return result
}

print_result2 :: proc(result: ^Result2) {
	fmt.printf("Task 2: %d\n", result^)
}
// --- Task 2 --- //
