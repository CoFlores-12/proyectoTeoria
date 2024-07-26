package main

import (
	"fmt"
	"priorityqueue/priorityqueue"
	"testing"
)

func TestFunctions(t *testing.T) {
	pq := priorityqueue.NewPriorityQueue(2)
	pq.Push("TR1", 1)
	pq.Push("R1", 2)
	pq.Push("R2", 2)
	pq.Push("R3", 2)
	pq.Push("TR2", 1)
	pq.Push("R4", 2)
	pq.Push("TR3", 1)
	pq.Push("R5", 2)
	val := pq.Pop()
	for val != nil {
		fmt.Println(val)
		val = pq.Pop()
	}
}

func TestLen(t *testing.T) {
	pq := priorityqueue.NewPriorityQueue(3)

	ln := pq.GetLenElements()
	if ln != 8 {
		t.Errorf("Esperado: %v / Obtenido: %v", 8, ln)
	}
}
