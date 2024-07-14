package main

import (
	"fmt"
	"priorityqueue/priorityqueue"
	"testing"
)

func TestFunctions(t *testing.T) {
	pq := priorityqueue.NewPriorityQueue(3)
	pq.Push("TR1", true)
	pq.Push("R1", false)
	pq.Push("R2", false)
	pq.Push("R3", false)
	pq.Push("TR2", true)
	pq.Push("R4", false)
	pq.Push("TR3", true)
	pq.Push("R5", false)
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
