package main

import "core:bytes"
import "core:fmt"
import "core:io"
import "core:os"
import "core:strconv"
import "core:strings"

Op :: enum {
	Add,
	Mul,
}

Monkey :: struct {
	items:          [dynamic]uint,
	new_op:         Op,
	new_arg:        uint,
	new_arg_old:    bool,
	test:           uint,
	if_true_index:  int,
	if_false_index: int,
	num_inspected:  int,
}

// NOTE: Have to free result.
parse_input_file :: proc(filepath: string) -> [dynamic]Monkey {
	data, ok := os.read_entire_file_from_filename(filepath)
	if !ok {
		panic("oh no, could not read file")
	}

	buf: bytes.Buffer
	bytes.buffer_init(&buf, data)
	defer bytes.buffer_destroy(&buf)

	monkeys := make([dynamic]Monkey, 0, 10)

	line: string
	err: io.Error
	for ; err != .EOF; line, err = bytes.buffer_read_string(&buf, '\n') {
		if line == "" || line == "\n" {
			continue
		}
		line = line[:len(line) - 1]

		monkey: Monkey
		monkey.items = make([dynamic]uint, 0, 2)

		line, err = bytes.buffer_read_string(&buf, '\n')
		line = line[:len(line) - 1]
		{
			item_index := strings.last_index_byte(line, ':') + 1
			items_split := strings.split(line[item_index:], ",")
			for id in items_split {
				append(&monkey.items, uint(strconv.atoi(id[1:])))
			}
			delete(items_split)
		}


		line, err = bytes.buffer_read_string(&buf, '\n')
		line = line[:len(line) - 1]
		{
			if strings.last_index_byte(line, '+') > 0 {
				monkey.new_op = .Add
			} else if strings.last_index_byte(line, '*') > 0 {
				monkey.new_op = .Mul
			}
			arg_index := strings.last_index_byte(line, ' ') + 1
			if line[arg_index:] != "old" {
				monkey.new_arg = uint(strconv.atoi(line[arg_index:]))
			} else {
				monkey.new_arg_old = true
			}
		}

		line, err = bytes.buffer_read_string(&buf, '\n')
		line = line[:len(line) - 1]
		{
			index := strings.last_index_byte(line, ' ') + 1
			monkey.test = uint(strconv.atoi(line[index:]))
		}

		line, err = bytes.buffer_read_string(&buf, '\n')
		line = line[:len(line) - 1]
		{
			index := strings.last_index_byte(line, ' ') + 1
			monkey.if_true_index = strconv.atoi(line[index:])
		}

		line, err = bytes.buffer_read_string(&buf, '\n')
		line = line[:len(line) - 1]
		{
			index := strings.last_index_byte(line, ' ') + 1
			monkey.if_false_index = strconv.atoi(line[index:])
		}

		_, err = bytes.buffer_read_string(&buf, '\n')

		append(&monkeys, monkey)
	}

	return monkeys
}

print_monkey_items :: proc(monkeys: []Monkey) {
	for monkey, i in monkeys {
		fmt.printf("Monkey %d: %v\n", i, monkey.items)
	}
}

task1 :: proc(monkeys: []Monkey, debug := false) -> int {
	if debug {
		print_monkey_items(monkeys)
	}

	for round in 1 ..= 20 {
		for monkey, monkey_index in monkeys {
			if debug {
				fmt.printf("Monkey %d:\n", monkey_index)
			}
			for item in monkey.items {
				if debug {
					fmt.printf("\tMonkey inspects an item with a worry level of %d.\n", item)
				}

				worry: uint
				if monkey.new_op == .Add {
					worry = item + (monkey.new_arg if !monkey.new_arg_old else item)
				} else if monkey.new_op == .Mul {
					worry = item * (monkey.new_arg if !monkey.new_arg_old else item)
				}

				if debug {
					fmt.printf(
						"\t\tWorry level is %s by %d to %d.\n",
						monkey.new_op,
						monkey.new_arg if !monkey.new_arg_old else item,
						worry,
					)
				}

				worry /= 3
				if debug {
					fmt.printf(
						"\t\tMonkey gets bored with item. Worry level is divided by 3 to %d.\n",
						worry,
					)
				}

				throw_to: int
				if worry % monkey.test == 0 {
					if debug {
						fmt.printf("\t\tCurrent worry level is divisible by %d.\n", monkey.test)
					}
					throw_to = monkey.if_true_index
				} else {
					if debug {
						fmt.printf(
							"\t\tCurrent worry level is not divisible by %d.\n",
							monkey.test,
						)
					}
					throw_to = monkey.if_false_index
				}
				if debug {
					fmt.printf(
						"\t\tItem with worry level %d is thrown to monkey %d.\n",
						worry,
						throw_to,
					)
				}
				append(&monkeys[throw_to].items, worry)
			}

			monkeys[monkey_index].num_inspected += len(monkeys[monkey_index].items)
			delete(monkeys[monkey_index].items)
			monkeys[monkey_index].items = make([dynamic]uint, 0, 2)
		}

		if debug {
			fmt.printf("\nRound %d\n", round)
			print_monkey_items(monkeys)
		}
	}

	if debug {
		fmt.printf("\n")
	}
	first, second: int
	for monkey, i in monkeys {
		if debug {
			fmt.printf("Monkey %d inspected items %d times\n", i, monkey.num_inspected)
		}
		if monkey.num_inspected > first {
			second, first = first, monkey.num_inspected
		} else if monkey.num_inspected > second {
			second = monkey.num_inspected
		}
	}

	return first * second
}

task2 :: proc(monkeys: []Monkey, debug := false) -> int {
	if debug {
		print_monkey_items(monkeys)
	}

	mod_constant: uint = 1
	for monkey, monkey_index in monkeys {
		mod_constant *= monkey.test
	}

	for round in 1 ..= 10000 {
		for monkey, monkey_index in monkeys {
			if debug {
				fmt.printf("Monkey %d:\n", monkey_index)
			}
			for item in monkey.items {
				if debug {
					fmt.printf("\tMonkey inspects an item with a worry level of %d.\n", item)
				}

				worry: uint
				if monkey.new_op == .Add {
					worry = item + (monkey.new_arg if !monkey.new_arg_old else item)
				} else if monkey.new_op == .Mul {
					worry = item * (monkey.new_arg if !monkey.new_arg_old else item)
				}

				if debug {
					fmt.printf(
						"\t\tWorry level is %s by %d to %d.\n",
						monkey.new_op,
						monkey.new_arg if !monkey.new_arg_old else item,
						worry,
					)
				}

				worry %= mod_constant

				throw_to: int
				if worry % monkey.test == 0 {
					if debug {
						fmt.printf("\t\tCurrent worry level is divisible by %d.\n", monkey.test)
					}
					throw_to = monkey.if_true_index
				} else {
					if debug {
						fmt.printf(
							"\t\tCurrent worry level is not divisible by %d.\n",
							monkey.test,
						)
					}
					throw_to = monkey.if_false_index
				}
				if debug {
					fmt.printf(
						"\t\tItem with worry level %d is thrown to monkey %d.\n",
						worry,
						throw_to,
					)
				}
				append(&monkeys[throw_to].items, worry)
			}

			monkeys[monkey_index].num_inspected += len(monkeys[monkey_index].items)
			delete(monkeys[monkey_index].items)
			monkeys[monkey_index].items = make([dynamic]uint, 0, 2)
		}

		if debug {
			fmt.printf("\nRound %d\n", round)
			print_monkey_items(monkeys)
		}
		if round == 1 || round == 20 || round % 1000 == 0 {
			fmt.printf("\nAfter round %d\n", round)
			for monkey, i in monkeys {
				fmt.printf("Monkey %d inspected items %d times\n", i, monkey.num_inspected)
			}
		}
	}

	if debug {
		fmt.printf("\n")
	}
	first, second: int
	for monkey, i in monkeys {
		if debug {
			fmt.printf("Monkey %d inspected items %d times\n", i, monkey.num_inspected)
		}
		if monkey.num_inspected > first {
			second, first = first, monkey.num_inspected
		} else if monkey.num_inspected > second {
			second = monkey.num_inspected
		}
	}

	return first * second
}

free_monkeys :: proc(monkeys: []Monkey) {
	for monkey in monkeys {
		delete(monkey.items)
	}
}

main :: proc() {
	{
		monkeys := parse_input_file("input.txt")
		defer delete(monkeys)
		defer free_monkeys(monkeys[:])

		result1 := task1(monkeys[:])
		fmt.printf("Task 1 result: %d\n", result1)
	}

	{
		monkeys := parse_input_file("input.txt")
		defer delete(monkeys)
		defer free_monkeys(monkeys[:])

		result2 := task2(monkeys[:])
		fmt.printf("Task 2 result: %d\n", result2)
	}
}
