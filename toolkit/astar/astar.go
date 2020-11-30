package astar

import (
	"container/heap"
	"fmt"

	"github.com/perlw/advent_of_code/toolkit/grid"
	"github.com/perlw/advent_of_code/toolkit/pq"
)

type Node struct {
	x  int
	y  int
	px int
	py int
}

func getNeighbours(node Node, width, height int) []Node {
	result := make([]Node, 0, 4)
	if node.x > 0 {
		result = append(result, Node{
			x: node.x - 1,
			y: node.y,
		})
	}
	if node.x < width-1 {
		result = append(result, Node{
			x: node.x + 1,
			y: node.y,
		})
	}
	if node.y > 0 {
		result = append(result, Node{
			x: node.x,
			y: node.y - 1,
		})
	}
	if node.y < height-1 {
		result = append(result, Node{
			x: node.x,
			y: node.y + 1,
		})
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func heuristic(a, b Node) int {
	return (abs(a.x-b.x) + abs(a.y-a.y))
}

// FindPath finds the path between start and goal (inclusively).
func FindPath(
	g grid.Grid, startx, starty, goalx, goaly int, genPathGrid bool,
) ([]Node, *grid.Grid) {
	open := make(pq.PriorityQueue, 0, g.Width*g.Height)
	heap.Init(&open)
	heap.Push(&open, &pq.Item{
		Value: Node{
			x:  startx,
			y:  starty,
			px: startx,
			py: starty,
		},
		Priority: 0,
	})

	closed := make(pq.PriorityQueue, 0, g.Width*g.Height)
	heap.Init(&closed)

	var outGrid *grid.Grid
	if genPathGrid {
		outGrid = &grid.Grid{
			Cells:  make([]grid.Cell, len(g.Cells)),
			Width:  g.Width,
			Height: g.Height,
			Label:  g.Label,
		}
		copy(outGrid.Cells, g.Cells)

		outGrid.Set(startx, starty, grid.Cell{
			R: 'S',
		})
		outGrid.Set(goalx, goaly, grid.Cell{
			R: 'G',
		})
	}

	goal := Node{
		x: goalx,
		y: goaly,
	}
	for len(open) > 0 {
		item := heap.Pop(&open).(*pq.Item)
		current := item.Value.(Node)
		if current.x == goal.x && current.y == goal.y {
			break
		}
		heap.Push(&closed, &pq.Item{
			Value:    current,
			Priority: item.Priority,
		})

		neighbours := getNeighbours(current, g.Width, g.Height)
		for _, next := range neighbours {
			if closed.Has(next, func(a, b interface{}) bool {
				aa := a.(Node)
				bb := b.(Node)
				return aa.x == bb.x && aa.y == bb.y
			}) {
				continue
			}
			if open.Has(next, func(a, b interface{}) bool {
				aa := a.(Node)
				bb := b.(Node)
				return aa.x == bb.x && aa.y == bb.y
			}) {
				continue
			}

			heap.Push(&open, &pq.Item{
				Value: Node{
					x:  next.x,
					y:  next.y,
					px: current.x,
					py: current.y,
				},
				Priority: item.Priority + heuristic(next, goal),
			})
		}
	}

	fmt.Printf("open: [\n")
	for _, item := range open {
		fmt.Printf("%d %+v,\n", item.Priority, item.Value.(Node))
	}
	fmt.Printf("]\n")
	fmt.Printf("closed: [\n")
	for _, item := range closed {
		fmt.Printf("%d %+v,\n", item.Priority, item.Value.(Node))
	}
	fmt.Printf("]\n")

	if genPathGrid {
		for _, item := range open {
			node := item.Value.(Node)

			dir := Node{
				x: node.x - node.px,
				y: node.y - node.py,
			}
			var r rune
			if dir.x > 0 {
				r = '←'
			} else if dir.x < 0 {
				r = '→'
			} else if dir.y > 0 {
				r = '↑'
			} else if dir.y < 0 {
				r = '↓'
			}
			if outGrid.Get(node.x, node.y).R == ' ' {
				outGrid.Set(node.x, node.y, grid.Cell{
					R: r,
				})
			}
		}
	}

	return []Node{}, outGrid
}
