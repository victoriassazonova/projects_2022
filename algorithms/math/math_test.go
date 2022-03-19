package math_test

import (
	"hsecode.com/stdlib/math"
	"testing"
)

func TestCompare(t *testing.T) {
	if math.NthPrime(3) != 5 {
		t.Fatal("wrong answer")
	}
	if math.NthPrime(24368) != 279121 {
		t.Fatal(math.NthPrime(24368))
	}
	if math.NthPrime(1) != 2 {
		t.Fatal("wrong answer")
	}
}
