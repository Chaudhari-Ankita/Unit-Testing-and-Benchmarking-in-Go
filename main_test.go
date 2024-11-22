package main

import (
	"testing"
)

func TestDivide(t *testing.T) {
	//error handling for wrong answer
	resultOne, _ := Divide(4, 2)
	expected := 2

	if resultOne != 2 {
		t.Errorf("Expected %d, got %d", expected, resultOne)
	}

	//error handling for divisible by 0 answer
	_, err := Divide(3, 0)
	if err == nil {
		t.Errorf("Expected an error when dividing by zero, got nil")
	}

	expectedError := "division by zero"

	if err.Error() != expectedError {
		t.Errorf("Expected error message '%s', got '%s'", expectedError, err.Error())
	}

}
func TestSquare(t *testing.T) {
	result := Square(5)
	expected := 25
	if result != expected {
		t.Fatalf("Square(5) will result %d, but we got %d", result, expected)
	}
}

func BenchmarkSquare(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Square(100)
	}
}
