package main

import "fmt"

func main() {
	fmt.Printf("Lesson 5 - Array")

	var myArray [3]int

	myArray[0] = 1
	myArray[1] = 2

	fmt.Println(myArray)

	fmt.Println(myArray[1])

	var mySecondArray [3]int

	mySecondArray[0], mySecondArray[1], mySecondArray[2] = 10, 20, 30

	fmt.Println(mySecondArray)
	fmt.Println(len(mySecondArray))

	for i, v := range mySecondArray {
		println(i, v)
	}
}
