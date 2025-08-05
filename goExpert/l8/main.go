package main

import (
	"errors"
	"fmt"
)

// Entry point
func main() {
	fmt.Println(sum(1, 2))
	fmt.Println(sumMultipleReturns(1, 2))
	fmt.Println(sumWithError(50, 60))
}

func sum(a int, b int) int {
	return a + b
}

func sumMultipleReturns(a, b int) (int, bool) {
	tempSum := a + b
	if tempSum >= 50 {
		return tempSum, true
	}

	return tempSum, false
}

func sumWithError(a, b int) (int, error) {
	tempSum := a + b
	if tempSum >= 50 {
		return tempSum, errors.New("A soma é maior que 50")
	}

	return tempSum, nil
}
