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
		"cat":     99,
		"hamster": 98,
		"rat":     75,
		"dog":     0,
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
		"hamster": 98,
		"dog":     0,
		"rat":     75,
		"cat":     99,
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
		"hamster": 98,
		"dog":     0,
		"rat":     75,
		"cat":     99,
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
		Priority: -1,
	})

	verify(t, &pq, expects)
}
