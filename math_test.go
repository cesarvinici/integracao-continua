package main

import "testing"

func TestSoma(t *testing.T) {
	total := soma(15, 15)

	if total != 30 {
		t.Errorf("Soma incorreta, resultado: %d, esperado: %d.", total, 30)
	}
}

func TestSubtracao(t *testing.T) {
	total := subtracao(15, 15)

	if total != 0 {
		t.Errorf("Subtracao incorreta, resultado: %d, esperado: %d.", total, 0)
	}
}
