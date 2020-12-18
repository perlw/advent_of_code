package main

import (
	"bufio"
	"fmt"
	"os"
)

func printWorld(world []rune, stride int) {
	for z := 0; z < stride; z++ {
		i := z * (stride * stride)
		fmt.Printf("Z: %d\n", z)
		for y := 0; y < stride; y++ {
			j := i + (y * stride)
			for x := 0; x < stride; x++ {
				k := j + x
				r := world[k]
				if r == 0 {
					fmt.Printf("%c", '+')
				} else {
					fmt.Printf("%c", r)
				}
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func printWorld4D(world []rune, stride int) {
	for w := 0; w < stride; w++ {
		t := w * (stride * stride * stride)
		fmt.Printf("W: %d\n", w)
		for z := 0; z < stride; z++ {
			i := t + (z * (stride * stride))
			fmt.Printf("Z: %d\n", z)
			for y := 0; y < stride; y++ {
				j := i + (y * stride)
				for x := 0; x < stride; x++ {
					k := j + x
					r := world[k]
					if r == 0 {
						fmt.Printf("%c", '+')
					} else {
						fmt.Printf("%c", r)
					}
				}
				fmt.Println()
			}
			fmt.Println()
		}
		fmt.Println()
	}
}

func getNumNeighbors(world []rune, stride int, x, y, z int) int {
	var result int
	for zz := z - 1; zz < z+2; zz++ {
		if zz < 0 || zz >= stride {
			continue
		}

		i := zz * (stride * stride)
		for yy := y - 1; yy < y+2; yy++ {
			if yy < 0 || yy >= stride {
				continue
			}

			j := i + (yy * stride)
			for xx := x - 1; xx < x+2; xx++ {
				if xx == x && yy == y && zz == z {
					continue
				}
				if xx < 0 || xx >= stride {
					continue
				}

				k := j + xx
				if world[k] == '#' {
					result++
				}
			}
		}
	}
	return result
}

func getNumNeighbors4D(world []rune, stride int, x, y, z, w int) int {
	var result int
	for ww := w - 1; ww < w+2; ww++ {
		if ww < 0 || ww >= stride {
			continue
		}
		t := ww * (stride * stride * stride)
		for zz := z - 1; zz < z+2; zz++ {
			if zz < 0 || zz >= stride {
				continue
			}

			i := t + (zz * (stride * stride))
			for yy := y - 1; yy < y+2; yy++ {
				if yy < 0 || yy >= stride {
					continue
				}

				j := i + (yy * stride)
				for xx := x - 1; xx < x+2; xx++ {
					if xx == x && yy == y && zz == z && ww == w {
						continue
					}
					if xx < 0 || xx >= stride {
						continue
					}

					k := j + xx
					if world[k] == '#' {
						result++
					}
				}
			}
		}
	}
	return result
}

// Task1 ...
func Task1(input []rune, stride int) int {
	offset := 1
	worldStride := stride + (offset * 2)
	size := worldStride * worldStride * worldStride
	world := make([]rune, size)
	for i := range world {
		world[i] = '.'
	}

	for y := 0; y < stride; y++ {
		for x := 0; x < stride; x++ {
			zz := offset * (worldStride * worldStride)
			yy := (y + offset) * worldStride
			xx := x + offset
			world[zz+yy+xx] = input[(y*stride)+x]
		}
	}

	printWorld(world, worldStride)
	prevWorld := make([]rune, size)
	for step := 0; step < 6; step++ {
		//fmt.Printf("Step %d\n==========\n", step+1)
		copy(prevWorld, world)

		var grow bool
		for z := 0; z < worldStride; z++ {
			i := z * (worldStride * worldStride)
			for y := 0; y < worldStride; y++ {
				j := i + (y * worldStride)
				for x := 0; x < worldStride; x++ {
					k := j + x
					num := getNumNeighbors(prevWorld, worldStride, x, y, z)
					if prevWorld[k] == '#' {
						if num != 2 && num != 3 {
							world[k] = '.'
						}
					} else if prevWorld[i] == '.' {
						if num == 3 {
							world[k] = '#'

							if x == 0 || x == stride-1 ||
								y == 0 || y == stride-1 ||
								z == 0 || z == stride-1 {
								grow = true
							}
						}
					}
				}
			}
		}

		if grow {
			offset++
			oldStride := worldStride
			worldStride = stride + (offset * 2)
			size = worldStride * worldStride * worldStride

			prevWorld = make([]rune, size)
			for i := range prevWorld {
				prevWorld[i] = '.'
			}
			for z := 0; z < oldStride; z++ {
				for y := 0; y < oldStride; y++ {
					for x := 0; x < oldStride; x++ {
						zz := (z + 1) * (worldStride * worldStride)
						yy := (y + 1) * worldStride
						xx := x + 1
						prevWorld[zz+yy+xx] = world[(z*(oldStride*oldStride))+(y*oldStride)+x]
					}
				}
			}
			world = make([]rune, size)
			copy(world, prevWorld)
		}

		//printWorld(world, worldStride)
	}

	var result int
	for i := range world {
		if world[i] == '#' {
			result++
		}
	}
	return result
}

// Task2 ...
func Task2(input []rune, stride int) int {
	offset := 1
	worldStride := stride + (offset * 2)
	size := worldStride * worldStride * worldStride * worldStride
	world := make([]rune, size)
	for i := range world {
		world[i] = '.'
	}

	for y := 0; y < stride; y++ {
		for x := 0; x < stride; x++ {
			ww := offset * (worldStride * worldStride * worldStride)
			zz := offset * (worldStride * worldStride)
			yy := (y + offset) * worldStride
			xx := x + offset
			world[ww+zz+yy+xx] = input[(y*stride)+x]
		}
	}

	//printWorld4D(world, worldStride)
	prevWorld := make([]rune, size)
	for step := 0; step < 6; step++ {
		//fmt.Printf("Step %d\n==========\n", step+1)
		copy(prevWorld, world)

		var grow bool
		for w := 0; w < worldStride; w++ {
			t := w * (worldStride * worldStride * worldStride)
			for z := 0; z < worldStride; z++ {
				i := t + (z * (worldStride * worldStride))
				for y := 0; y < worldStride; y++ {
					j := i + (y * worldStride)
					for x := 0; x < worldStride; x++ {
						k := j + x
						num := getNumNeighbors4D(prevWorld, worldStride, x, y, z, w)
						if prevWorld[k] == '#' {
							if num != 2 && num != 3 {
								world[k] = '.'
							}
						} else if prevWorld[i] == '.' {
							if num == 3 {
								world[k] = '#'

								if x == 0 || x == stride-1 ||
									y == 0 || y == stride-1 ||
									z == 0 || z == stride-1 ||
									w == 0 || w == stride-1 {
									grow = true
								}
							}
						}
					}
				}
			}
		}

		if grow {
			offset++
			oldStride := worldStride
			worldStride = stride + (offset * 2)
			size = worldStride * worldStride * worldStride * worldStride

			prevWorld = make([]rune, size)
			for i := range prevWorld {
				prevWorld[i] = '.'
			}
			for w := 0; w < oldStride; w++ {
				for z := 0; z < oldStride; z++ {
					for y := 0; y < oldStride; y++ {
						for x := 0; x < oldStride; x++ {
							ww := (w + 1) * (worldStride * worldStride * worldStride)
							zz := (z + 1) * (worldStride * worldStride)
							yy := (y + 1) * worldStride
							xx := x + 1
							prevWorld[ww+zz+yy+xx] = world[(w*(oldStride*oldStride*oldStride))+(z*(oldStride*oldStride))+(y*oldStride)+x]
						}
					}
				}
			}
			world = make([]rune, size)
			copy(world, prevWorld)
		}

		//printWorld4D(world, worldStride)
	}

	var result int
	for i := range world {
		if world[i] == '#' {
			result++
		}
	}
	return result
}

func main() {
	file, _ := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	var stride int
	input := make([]rune, 0, 100)
	for scanner.Scan() {
		line := scanner.Text()
		if stride == 0 {
			stride = len(line)
		}
		for _, r := range line {
			input = append(input, rune(r))
		}
	}
	file.Close()

	result := Task1(input, stride)
	fmt.Printf("Task 1: %d\n", result)

	result = Task2(input, stride)
	fmt.Printf("Task 2: %d\n", result)
}
