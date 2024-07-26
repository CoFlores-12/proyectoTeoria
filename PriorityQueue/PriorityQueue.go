package priorityqueue

import "fmt"

type Item struct {
	Data     any
	Priority int
	Next     *Item
}

func (it Item) String() string {
	return fmt.Sprintf("Data: %v, Priority: %d", it.Data, it.Priority)
}

type PriorityQueue struct {
	priorities    []*Item
	maxPriority   int
	elements      int
	countThirdAge int
	countRegular  int
}

func NewPriorityQueue(max int) *PriorityQueue {
	return &PriorityQueue{
		priorities:  make([]*Item, max+1),
		maxPriority: max,
	}
}

func (pq *PriorityQueue) Push(data any, priority int) {
	if priority < 0 || priority > pq.maxPriority {
		fmt.Println("Fuera de rango")
		return
	}

	if pq.priorities[priority] == nil {
		pq.priorities[priority] = &Item{
			Data:     data,
			Priority: priority,
		}
		pq.elements++
		return
	}

	tmp := pq.priorities[priority]
	for tmp.Next != nil {
		tmp = tmp.Next
	}
	tmp.Next = &Item{Data: data, Priority: priority}
	pq.elements++
}

func (pq PriorityQueue) GetMaxPriority() int {
	return pq.maxPriority
}

func (pq PriorityQueue) GetLenElements() int {
	return pq.elements
}

func (pq *PriorityQueue) Pop() *Item {
	if pq.elements == 0 {
		return nil
	}
	max := pq.GetMaxPriority()
	for max > 0 {
		tmp := pq.priorities[max]
		if tmp != nil {
			break
		}
		max--
	}
	tmp := pq.priorities[max]
	pq.priorities[max] = tmp.Next
	pq.elements--
	return tmp
}
