package main

import "core:bytes"
import "core:fmt"
import "core:io"
import "core:os"
import "core:slice"
import "core:sort"
import "core:strconv"
import "core:strings"

HandType :: enum (int) {
	HighCard,
	OnePair,
	TwoPair,
	ThreeOfAKind,
	FullHouse,
	FourOfAKind,
	FiveOfAKind,
}

Card := map[rune]int {
	'A' = 12,
	'K' = 11,
	'Q' = 10,
	'J' = 9,
	'T' = 8,
	'9' = 7,
	'8' = 6,
	'7' = 5,
	'6' = 4,
	'5' = 3,
	'4' = 2,
	'3' = 1,
	'2' = 0,
}

CardJoker := map[rune]int {
	'A' = 12,
	'K' = 11,
	'Q' = 10,
	'T' = 9,
	'9' = 8,
	'8' = 7,
	'7' = 6,
	'6' = 5,
	'5' = 4,
	'4' = 3,
	'3' = 2,
	'2' = 1,
	'J' = 0,
}

Hand :: struct {
	cards:         [5]rune,
	bid:           int,
	type:          HandType,
	type_if_joker: HandType,
}

Input :: struct {
	hands: [dynamic]Hand,
}

Result1 :: distinct int
Result2 :: distinct int


// --- Input --- //
print_input :: proc(input: ^Input) {
	for h in input.hands {
		fmt.printf("%v(%s|%s) -> %d\n", h.cards, h.type, h.type_if_joker, h.bid)
	}
}

identify_hand_type :: proc(cards: []rune) -> HandType {
	card_map := make(map[rune]int)
	defer delete(card_map)

	for r in cards {
		card_map[r] += 1
	}
	high: int
	for _, n in card_map {
		if n > high {
			high = n
		}
	}

	switch len(card_map) {
	case 1:
		return .FiveOfAKind
	case 2:
		if high > 3 {
			return .FourOfAKind
		}
		return .FullHouse
	case 3:
		if high > 2 {
			return .ThreeOfAKind
		}
		return .TwoPair
	case 4:
		return .OnePair
	}

	return .HighCard
}

identify_hand_type_if_joker :: proc(cards: []rune) -> HandType {
	card_map := make(map[rune]int)
	defer delete(card_map)

	num_jokers: int
	for r in cards {
		if r == 'J' {
			num_jokers += 1
			continue
		}
		card_map[r] += 1
	}
	high: int
	for _, n in card_map {
		if n > high {
			high = n
		}
	}
	high += num_jokers

	if high == 5 {
		return .FiveOfAKind
	}

	switch len(card_map) {
	case 1:
		return .FiveOfAKind
	case 2:
		if high > 3 {
			return .FourOfAKind
		}
		return .FullHouse
	case 3:
		if high > 2 {
			return .ThreeOfAKind
		}
		return .TwoPair
	case 4:
		return .OnePair
	}

	return .HighCard
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

	input.hands = make([dynamic]Hand, 0, 10)

	line: string
	err: io.Error
	for ; err != .EOF; line, err = bytes.buffer_read_string(&buf, '\n') {
		if line == "" || line == "\n" {
			continue
		}

		line = line[:len(line) - 1]

		hand_line := strings.fields(line)
		defer delete(hand_line)

		hand: Hand
		for r, i in hand_line[0] {
			hand.cards[i] = r
		}
		hand.bid = strconv.atoi(hand_line[1])
		hand.type = identify_hand_type(hand.cards[:])
		hand.type_if_joker = identify_hand_type_if_joker(hand.cards[:])

		append(&input.hands, hand)
	}

	return input
}

free_input :: proc(input: ^Input) {
	delete(input.hands)
}
// --- Input --- //


// --- Task 1 --- //
sort_hands :: proc(hands: ^[]Hand) -> sort.Interface {
	it := sort.Interface {
		collection = rawptr(hands),
		len = proc(it: sort.Interface) -> int {
			hands := (^[]Hand)(it.collection)
			return len(hands^)
		},
		less = proc(it: sort.Interface, i, j: int) -> bool {
			hands := (^[]Hand)(it.collection)
			if hands[i].type == hands[j].type {
				for t in 0 ..< 5 {
					if hands[i].cards[t] != hands[j].cards[t] {
						return Card[hands[i].cards[t]] < Card[hands[j].cards[t]]
					}
				}
			}
			return hands[i].type < hands[j].type
		},
		swap = proc(it: sort.Interface, i, j: int) {
			hands := (^[]Hand)(it.collection)
			hands[i], hands[j] = hands[j], hands[i]
		},
	}
	return it
}

run_task1 :: proc(input: ^Input, debug: bool) -> Result1 {
	result: Result1

	if debug {
		print_input(input)
	}

	sorted_hands := slice.clone(input.hands[:])
	defer delete(sorted_hands)

	sort.sort(sort_hands(&sorted_hands))
	for hand, i in sorted_hands {
		result += Result1((i + 1) * hand.bid)
	}

	return result
}

print_result1 :: proc(result: ^Result1) {
	fmt.printf("Task 1: %d\n", result^)
}
// --- Task 1 --- //


// --- Task 2 --- //
sort_joker_hands :: proc(hands: ^[]Hand) -> sort.Interface {
	it := sort.Interface {
		collection = rawptr(hands),
		len = proc(it: sort.Interface) -> int {
			hands := (^[]Hand)(it.collection)
			return len(hands^)
		},
		less = proc(it: sort.Interface, i, j: int) -> bool {
			hands := (^[]Hand)(it.collection)
			if hands[i].type_if_joker == hands[j].type_if_joker {
				for t in 0 ..< 5 {
					if hands[i].cards[t] != hands[j].cards[t] {
						return CardJoker[hands[i].cards[t]] < CardJoker[hands[j].cards[t]]
					}
				}
			}
			return hands[i].type_if_joker < hands[j].type_if_joker
		},
		swap = proc(it: sort.Interface, i, j: int) {
			hands := (^[]Hand)(it.collection)
			hands[i], hands[j] = hands[j], hands[i]
		},
	}
	return it
}

run_task2 :: proc(input: ^Input, debug: bool) -> Result2 {
	result: Result2

	if debug {
		print_input(input)
	}

	sorted_hands := slice.clone(input.hands[:])
	defer delete(sorted_hands)

	sort.sort(sort_joker_hands(&sorted_hands))
	for hand, i in sorted_hands {
		result += Result2((i + 1) * hand.bid)
	}

	return result
}

print_result2 :: proc(result: ^Result2) {
	fmt.printf("Task 2: %d\n", result^)
}
// --- Task 2 --- //
