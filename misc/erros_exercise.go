package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e *ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", e)
}
	

func Sqrt(x float64) (float64, error) {
	if x < 0 {
		return &ErrNegativeSqrt{x}
	}
	
	var z, zPrevious float64 = 1, 0;
	counter := 1
	
	for math.Abs(zPrevious - z) > 0.0000000000001 {
		zPrevious = z;
		z -= (z*z-x)/(2*z);
		counter += 1
	}
	return z;
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}
