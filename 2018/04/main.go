package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
	"time"
)

type TimeAction struct {
	DateTime time.Time
	Action   string
}

type ByTime []TimeAction

func (t ByTime) Len() int {
	return len(t)
}

func (t ByTime) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t ByTime) Less(i, j int) bool {
	return t[i].DateTime.Before(t[j].DateTime)
}

type Puzzle struct {
	RawLog     []TimeAction
	GuardSleep map[int][60]int
}

func (p *Puzzle) Add(dt time.Time, action string) {
	if p.RawLog == nil {
		p.RawLog = make([]TimeAction, 0, 100)
	}
	p.RawLog = append(p.RawLog, TimeAction{
		DateTime: dt,
		Action:   action,
	})
}

func (p *Puzzle) Finish() {
	sort.Sort(ByTime(p.RawLog))

	p.GuardSleep = make(map[int][60]int)

	var m, d int
	var guard int
	var sleep int
	for _, item := range p.RawLog {
		if int(item.DateTime.Month()) > m {
			m = int(item.DateTime.Month())
		}
		if item.DateTime.Day() > d {
			d = item.DateTime.Day()
		}

		switch item.Action {
		case "falls asleep":
			sleep = item.DateTime.Minute()
		case "wakes up":
			g, _ := p.GuardSleep[guard]
			for t := sleep; t < item.DateTime.Minute(); t++ {
				g[t]++
			}
			p.GuardSleep[guard] = g
		default:
			fmt.Sscanf(item.Action, "Guard #%d", &guard)
			sleep = 0
		}
	}
}

func (p *Puzzle) Solution1() int {
	most := 0
	guard := 0
	favMinute := 0

	for gid, g := range p.GuardSleep {
		totalSleep := 0
		highestTimes := 0
		minute := 0
		for min, times := range g {
			if times > 0 {
				totalSleep += times
			}
			if times > highestTimes {
				highestTimes = times
				minute = min
			}
		}
		if totalSleep > most {
			most = totalSleep
			guard = gid
			favMinute = minute
		}
	}

	return guard * favMinute
}

func (p *Puzzle) Solution2() int {
	highestMinute := 0
	highestMinuteVal := 0
	guard := 0

	for gid, g := range p.GuardSleep {
		highMin := 0
		highMinVal := 0
		for min, times := range g {
			if times > highMinVal {
				highMinVal = times
				highMin = min
			}
		}
		if highMinVal > highestMinuteVal {
			highestMinuteVal = highMinVal
			highestMinute = highMin
			guard = gid
		}
	}

	return guard * highestMinute
}

func main() {
	input, err := os.Open("input.txt")
	if err != nil {
		panic(err.Error())
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)

	p := Puzzle{}
	for scanner.Scan() {
		split := strings.SplitAfter(scanner.Text(), "] ")
		dt, _ := time.Parse("2006-01-02 15:04", strings.Trim(split[0], "[] "))
		action := split[1]
		p.Add(dt, action)
	}
	if err := scanner.Err(); err != nil {
		panic(err.Error())
	}
	p.Finish()

	fmt.Println(p.Solution1())
	fmt.Println(p.Solution2())
}
