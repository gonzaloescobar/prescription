package main

import (
	"testing"
)


func TestSuma(t *testing.T) {
	valor := Suma(7, 23)
	if valor != 30 {
		t.Error("Se esperaba 30 y se obtuvo", valor)
	}
}