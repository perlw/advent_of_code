package pq

// CompareFunc defines a comparison function.
type CompareFunc func(a, b interface{}) bool

// Item is an item in the PriorityQueue.
type Item struct {
	Value    interface{}
	Priority int
}

// PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Priority < pq[j].Priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

// Push pushes an item into the queue.
func (pq *PriorityQueue) Push(x interface{}) {
	if item, ok := x.(*Item); ok {
		*pq = append(*pq, item)
	}
}

// Pop pops an item from the queue.
func (pq *PriorityQueue) Pop() interface{} {
	n := len(*pq)
	item := (*pq)[n-1]
	(*pq)[n-1] = nil
	*pq = (*pq)[0 : n-1]
	return item
}

// Has checks if queue has an item.
func (pq PriorityQueue) Has(x interface{}, cmp CompareFunc) bool {
	for _, item := range pq {
		if cmp(x, item.Value) {
			return true
		}
	}
	return false
}
