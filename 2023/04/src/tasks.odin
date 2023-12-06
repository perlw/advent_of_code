package main

import "core:bytes"
import "core:fmt"
import "core:io"
import "core:os"
import "core:strconv"
import "core:strings"

Card :: struct {
	winners: []int,
	numbers: []int,
}

Input :: struct {
	cards: [dynamic]Card,
}

Result1 :: distinct int
Result2 :: distinct int


// --- Input --- //
print_input :: proc(input: ^Input) {
	for card, i in input.cards {
		fmt.printf("Card % 3d: ", i + 1)

		for n in card.winners {
			fmt.printf("% 3d ", n)
		}
		fmt.printf(" |")
		for n in card.numbers {
			fmt.printf("% 3d ", n)
		}

		fmt.printf("\n")
	}
}

parse_numbers :: proc(numbers: string) -> []int {
	split_numbers := strings.fields(numbers)
	result := make([]int, len(split_numbers))
	for n, i in split_numbers {
		result[i] = strconv.atoi(n)
	}
	return result
}

parse_input_file :: proc(filepath: string) -> Input {
	input: Input

	raw_data, ok := os.read_entire_file_from_filename(filepath)
	if !ok {
		panic("oh no, could not read file")
	}

	buf: bytes.Buffer
	bytes.buffer_init(&buf, raw_data)
	defer bytes.buffer_destroy(&buf)

	input.cards = make([dynamic]Card, 0, 10)

	line: string
	err: io.Error
	for ; err != .EOF; line, err = bytes.buffer_read_string(&buf, '\n') {
		if line == "" || line == "\n" {
			continue
		}

		line = line[:len(line) - 1]

		card: Card
		start := strings.index(line, ":")
		data := strings.split(line[start + 2:], " | ")
		card.winners = parse_numbers(data[0])
		card.numbers = parse_numbers(data[1])

		append(&input.cards, card)
	}

	return input
}

free_input :: proc(input: ^Input) {
	for card in input.cards {
		delete(card.winners)
		delete(card.numbers)
	}
	delete(input.cards)
}
// --- Input --- //


// --- Task 1 --- //
run_task1 :: proc(input: ^Input, debug: bool) -> Result1 {
	result: Result1

	if debug {
		print_input(input)
	}

	for card in input.cards {
		count: int
		for n in card.numbers {
			for w in card.winners {
				if n == w {
					count = count * 2 if count > 0 else 1
					break
				}
			}
		}
		result += Result1(count)
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
		print_input(input)
	}

	won_cards := make([]int, len(input.cards))
	defer delete(won_cards)
	for _, i in won_cards {
		won_cards[i] = 1
	}

	for card, i in input.cards {
		count: int
		for n in card.numbers {
			for w in card.winners {
				if n == w {
					count += 1
					break
				}
			}
		}

		if count > 0 {
			for j := i + 1; j <= i + count && j < len(input.cards); j += 1 {
				won_cards[j] += won_cards[i]
			}
		}
	}

	for w in won_cards {
		result += Result2(w)
	}

	return result
}

print_result2 :: proc(result: ^Result2) {
	fmt.printf("Task 2: %d\n", result^)
}
// --- Task 2 --- //
