package main

import (
	"testing"
)

func TestFunctions(t *testing.T) {

}

func TestLen(t *testing.T) {

	ln := 0
	if ln != 8 {
		t.Errorf("Esperado: %v / Obtenido: %v", 8, ln)
	}
}
