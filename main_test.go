package main

import (
	"testing"
)

func BenchmarkDivide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Divide(100, 5)
	}
}
