package main

import (
	"testing"
)

func TestSquare(t *testing.T) {
	result := Square(5)
	expected := 25
	if result != expected {
		t.Fatalf("Square(5) will result %d, but we got %d", result, expected)
	}
}

func BenchmarkDivide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Divide(100, 5)
	}
}
