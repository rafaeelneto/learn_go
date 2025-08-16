package main

import "fmt"

func main() {
	var s = []int{10, 20, 30, 40, 50, 60}

	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(s), cap(s), s)
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(s[:0]), cap(s[:0]), s[:0])
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(s[:4]), cap(s[:4]), s[:4])

	s = append(s, 70)
	fmt.Printf("len=%d, cap=%d, slice=%v\n", len(s), cap(s), s)
}
