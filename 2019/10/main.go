package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"sort"
	"strings"
)

type Point struct {
	x, y int
}

func distance(a, b Point) float64 {
	return math.Sqrt(math.Pow(math.Abs(float64(b.x-a.x)), 2) + math.Pow(math.Abs(float64(b.y-a.y)), 2))
}

func distanceToLine(a, b, c Point) float64 {
	px := b.x - a.x
	py := b.y - a.y
	norm := (px * px) + (py * py)
	u := float64(((c.x-a.x)*px)+((c.y-a.y)*py)) / float64(norm)

	if u > 1 {
		u = 1
	} else if u < 0 {
		u = 0
	}

	x := float64(a.x) + u*float64(px)
	y := float64(a.y) + u*float64(py)
	dx := x - float64(c.x)
	dy := y - float64(c.y)

	return math.Sqrt((dx * dx) + (dy * dy))
}

func FindBestPosition(grid [][]rune, debug bool) (Point, int) {
	var result Point
	var max int

	points := make([]Point, 0)
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == '#' {
				points = append(points, Point{x: x, y: y})
			}
		}
	}

	for i := range points {
		var count int
		for j := range points {
			if i == j {
				continue
			}

			p1, p2 := points[i], points[j]
			if p2.x < p1.x {
				p1, p2 = p2, p1
			}

			var hit bool
			for k := range points {
				if k == i || k == j {
					continue
				}

				if distanceToLine(p1, p2, points[k]) == 0 {
					hit = true
					break
				}
			}
			if hit {
				continue
			}

			count++
		}

		if debug {
			grid[points[i].y][points[i].x] = rune(count)
		}
		if count > max {
			max = count
			result = points[i]
		}
	}

	return result, max
}

type Asteroid struct {
	pos      Point
	distance float64
}

func VaporizeAsteroids(grid [][]rune, station Point, debug bool) Point {
	radDeg := 180.0 / math.Pi
	angles := make([]float64, 0)
	angleDistances := make(map[float64][]Asteroid)
	for y := range grid {
		for x := range grid[y] {
			if x == station.x && y == station.y {
				continue
			}

			if grid[y][x] == '#' {
				angle := math.Atan2(float64(x-station.x), -float64(y-station.y)) * radDeg
				if angle < 0 {
					angle += 360.0
				}
				if _, ok := angleDistances[angle]; !ok {
					angleDistances[angle] = make([]Asteroid, 0)
					angles = append(angles, angle)
				}
				p := Point{
					x: x,
					y: y,
				}
				angleDistances[angle] = append(angleDistances[angle], Asteroid{
					pos:      p,
					distance: distance(station, p),
				})
			}
		}
	}

	sort.Slice(angles, func(a, b int) bool {
		return angles[a] < angles[b]
	})

	for i := range angleDistances {
		sort.Slice(angleDistances[i], func(a, b int) bool {
			return angleDistances[i][a].distance < angleDistances[i][b].distance
		})
	}

	if debug {
		for _, a := range angles {
			fmt.Printf("%f -> [\n", a)
			for _, as := range angleDistances[a] {
				fmt.Printf("\t%+v\n", as)
			}
			fmt.Println("]")
		}
		fmt.Println("Vaporizing...")
	}

	var count int
	for count < 200 {
		for _, a := range angles {
			if len(angleDistances[a]) == 0 {
				continue
			}

			as := angleDistances[a][0]
			angleDistances[a] = angleDistances[a][1:]

			if debug {
				fmt.Printf("Asteroid #%d, %+v\n", count+1, as)
			}
			count++
			if count == 200 {
				return as.pos
			}
		}
	}

	return Point{}
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(fmt.Errorf("could not open input file: %w", err))
	}
	defer input.Close()

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)

	grid := make([][]rune, 0)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		x := make([]rune, len(line))
		for i := range line {
			x[i] = rune(line[i])
		}
		grid = append(grid, x)
	}

	pos, count := FindBestPosition(grid, false)
	fmt.Printf("Part 1: %d\n", count)

	pos = VaporizeAsteroids(grid, pos, false)
	fmt.Printf("Part 2: %d\n", (pos.x*100)+pos.y)
}
