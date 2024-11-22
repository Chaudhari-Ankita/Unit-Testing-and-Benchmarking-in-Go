package main

import (
	"errors"
	"fmt"
)

func Divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func Square(a int ) int {
		return a*a
	}

func main() {
	fmt.Println(Divide(5, 3))
	fmt.Println(Square(5))
}
