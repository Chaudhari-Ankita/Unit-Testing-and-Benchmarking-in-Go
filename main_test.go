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

func BenchmarkDivide(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Divide(100, 5)
	}
}
