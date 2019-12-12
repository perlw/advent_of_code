package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func abs(a int) int {
	if a < 0 {
		a = -a
	}
	return a
}

type Moon struct {
	x, y, z    int
	vx, vy, vz int
}

func (m *Moon) CalcVelocity(self int, moons []Moon) {
	for i := range moons {
		if i == self {
			continue
		}

		if m.x > moons[i].x {
			m.vx--
		} else if m.x < moons[i].x {
			m.vx++
		}
		if m.y > moons[i].y {
			m.vy--
		} else if m.y < moons[i].y {
			m.vy++
		}
		if m.z > moons[i].z {
			m.vz--
		} else if m.z < moons[i].z {
			m.vz++
		}
	}
}

func (m *Moon) Step() {
	m.x += m.vx
	m.y += m.vy
	m.z += m.vz
}

func (m *Moon) Energy() int {
	pot := abs(m.x) + abs(m.y) + abs(m.z)
	kin := abs(m.vx) + abs(m.vy) + abs(m.vz)
	return pot * kin
}

func IterateMoons(moons []Moon, steps int, debug bool) int {
	for i := 0; i < steps; i++ {
		if debug {
			fmt.Printf("Step %d\n", i+1)
		}
		for j := range moons {
			moons[j].CalcVelocity(j, moons)
		}
		for j := range moons {
			moons[j].Step()
			if debug {
				fmt.Printf("%+v\n", moons[j])
			}
		}
		if debug {
			fmt.Println()
		}
	}

	var energy int
	for i := range moons {
		energy += moons[i].Energy()
	}
	return energy
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
		/*t := b
		  b = a % b
		  a = t*/
	}
	return a
}

func lcm(a, b, c int) int {
	result := a * b / gcd(a, b)
	result = result * c / gcd(result, c)
	return result
}

func FindSteps(moons []Moon, debug bool) int {
	initial := make([]Moon, len(moons))

	var xFound, yFound, zFound bool
	xSteps, ySteps, zSteps := 1, 1, 1

	copy(initial, moons)
	for {
		for i := range moons {
			moons[i].CalcVelocity(i, moons)
		}

		xLooped := true
		yLooped := true
		zLooped := true
		for i := range moons {
			moons[i].Step()

			if moons[i].x != initial[i].x || moons[i].vx != initial[i].vx {
				xLooped = false
			}
			if moons[i].y != initial[i].y || moons[i].vy != initial[i].vy {
				yLooped = false
			}
			if moons[i].z != initial[i].z || moons[i].vz != initial[i].vz {
				zLooped = false
			}
		}
		if xLooped {
			xFound = true
		}
		if yLooped {
			yFound = true
		}
		if zLooped {
			zFound = true
		}

		if !xFound {
			xSteps++
		}
		if !yFound {
			ySteps++
		}
		if !zFound {
			zSteps++
		}

		if xFound && yFound && zFound {
			break
		}
	}

	return lcm(xSteps, ySteps, zSteps)
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(fmt.Errorf("could not open input file: %w", err))
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	moons := make([]Moon, 0)
	for scanner.Scan() {
		var moon Moon
		line := strings.TrimSpace(scanner.Text())
		fmt.Sscanf(line, "<x=%d, y=%d, z=%d>", &moon.x, &moon.y, &moon.z)
		moons = append(moons, moon)
	}

	fmt.Printf("Part 1: %d\n", IterateMoons(moons, 1000, false))
	fmt.Printf("Part 2: %d\n", FindSteps(moons, false))
}
