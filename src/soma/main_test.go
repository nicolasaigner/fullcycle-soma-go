package main

import "testing"

func TestSum(t *testing.T) {
	got := soma(5, 5)
	if got != 10 {
		t.Errorf("A soma retornou %d. O valor esperado é 10", got)
	}
}
