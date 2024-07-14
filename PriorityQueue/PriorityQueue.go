package priorityqueue

import "fmt"

type Item struct {
	data     any
	priority int
	next     *Item
}

func (it Item) String() string {
	return fmt.Sprintf("Data %v, Priority %d", it.data, it.priority)
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

const (
	ThirdAgePriority = 1
	RegularPriority  = 2
)

func (pq *PriorityQueue) Push(data any, isThirdAge bool) {
	var priority int
	if isThirdAge {
		priority = ThirdAgePriority
	} else {
		priority = RegularPriority
	}

	if priority < 0 || priority > pq.maxPriority {
		fmt.Println("Fuera de rango")
		return
	}

	if pq.priorities[priority] == nil {
		pq.priorities[priority] = &Item{
			data:     data,
			priority: priority,
		}
		pq.elements++
		return
	}

	tmp := pq.priorities[priority]
	for tmp.next != nil {
		tmp = tmp.next
	}
	tmp.next = &Item{data: data, priority: priority}
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

	if pq.countThirdAge < 1 && pq.priorities[ThirdAgePriority] != nil {
		pq.countThirdAge++
		pq.countRegular = 0
		tmp := pq.priorities[ThirdAgePriority]
		pq.priorities[ThirdAgePriority] = tmp.next
		pq.elements--
		return tmp
	}

	if pq.countRegular < 2 && pq.priorities[RegularPriority] != nil {
		pq.countRegular++
		if pq.countRegular == 2 {
			pq.countThirdAge = 0
		}
		tmp := pq.priorities[RegularPriority]
		pq.priorities[RegularPriority] = tmp.next
		pq.elements--
		return tmp
	}

	if pq.priorities[ThirdAgePriority] != nil {
		pq.countThirdAge = 1
		pq.countRegular = 0
		tmp := pq.priorities[ThirdAgePriority]
		pq.priorities[ThirdAgePriority] = tmp.next
		pq.elements--
		return tmp
	}

	if pq.priorities[RegularPriority] != nil {
		pq.countRegular++
		tmp := pq.priorities[RegularPriority]
		pq.priorities[RegularPriority] = tmp.next
		pq.elements--
		return tmp
	}

	return nil
}
