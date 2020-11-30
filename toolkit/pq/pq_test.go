package pq

import (
	"container/heap"
	"testing"
)

func verify(t *testing.T, pq *PriorityQueue, expects []string) {
	i := 0
	for pq.Len() > 0 {
		item := heap.Pop(pq).(*Item)
		value := item.Value.(string)
		if value != expects[i] {
			t.Errorf("expected %s, got %s", expects[i], value)
		}
		i++
	}
}

func TestShouldContainItems(t *testing.T) {
	input := map[string]int{
		"cat":     0,
		"hamster": 2,
		"rat":     10,
		"dog":     25,
	}
	expects := []string{
		"cat",
		"hamster",
		"rat",
		"dog",
	}

	pq := make(PriorityQueue, 0, len(input))
	for value, priority := range input {
		pq = append(pq, &Item{
			Value:    value,
			Priority: priority,
		})
	}
	heap.Init(&pq)

	verify(t, &pq, expects)
}

func TestShouldSortItems(t *testing.T) {
	input := map[string]int{
		"rat":     10,
		"cat":     0,
		"dog":     25,
		"hamster": 2,
	}
	expects := []string{
		"cat",
		"hamster",
		"rat",
		"dog",
	}

	pq := make(PriorityQueue, 0, len(input))
	for value, priority := range input {
		pq = append(pq, &Item{
			Value:    value,
			Priority: priority,
		})
	}
	heap.Init(&pq)

	verify(t, &pq, expects)
}

func TestShouldPushItems(t *testing.T) {
	input := map[string]int{
		"rat":     10,
		"cat":     0,
		"dog":     25,
		"hamster": 2,
	}
	expects := []string{
		"cat",
		"hamster",
		"rat",
		"dog",
		"dangernoodle",
	}

	pq := make(PriorityQueue, 0, len(input))
	for value, priority := range input {
		pq = append(pq, &Item{
			Value:    value,
			Priority: priority,
		})
	}
	heap.Init(&pq)

	heap.Push(&pq, &Item{
		Value:    "dangernoodle",
		Priority: 99,
	})

	verify(t, &pq, expects)
}

func TestShouldHaveItem(t *testing.T) {
	input := map[string]int{
		"rat":     10,
		"cat":     0,
		"dog":     25,
		"hamster": 2,
	}

	pq := make(PriorityQueue, 0, len(input))
	for value, priority := range input {
		pq = append(pq, &Item{
			Value:    value,
			Priority: priority,
		})
	}
	heap.Init(&pq)

	if !pq.Has("cat", func(a, b interface{}) bool {
		return a.(string) == b.(string)
	}) {
		t.Errorf("expected to find cat")
	}
}
