package main

import "fmt"

func main() {
	arr := [6]int{3, 14, 64,1, 54,90} // array

	slc := []int{3, 41,21} // slice

	fmt.Println(arr)
	fmt.Println(slc)

	// a slice a part of a array. It defines a continuous piece a a array
	// a slice is not a array
	sliceEx := arr[1:3] //! 1 is inclusive and 3 is exclusive

	fmt.Println(sliceEx)

	sliceEx2 := arr[:2]

	fmt.Println(sliceEx2)

}