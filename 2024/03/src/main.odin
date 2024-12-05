package main

import "core:log"
import "core:os"

main :: proc() {
	context.logger = log.create_console_logger(ident = "MAIN")
	defer log.destroy_console_logger(context.logger)

	debug1: bool
	debug2: bool
	if len(os.args) > 1 {
		for arg in os.args {
			switch arg {
			case "-d1":
				debug1 = true
				log.infof("Enabled Task 1 debug")
			case "-d2":
				debug2 = true
				log.infof("Enabled Task 2 debug")
			}
		}
	}

	input := parse_input_file("input.txt")
	defer free_input(&input)

	{
		level: log.Level = .Debug if debug1 else .Info
		context.logger = log.create_console_logger(lowest = level, ident = "TASK 1")
		defer log.destroy_console_logger(context.logger)

		result1 := run_task1(&input, debug1)
		print_result1(&result1)
	}

	{
		level: log.Level = .Debug if debug2 else .Info
		context.logger = log.create_console_logger(lowest = level, ident = "TASK 2")
		defer log.destroy_console_logger(context.logger)

		result2 := run_task2(&input, debug2)
		print_result2(&result2)
	}
}
