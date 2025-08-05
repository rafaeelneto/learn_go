package main

import (
	"fmt"
	"time"
)

func task(taskText string) {
	for i := 0; i < 10; i++ {
		fmt.Println(taskText, ":", i+1)
		time.Sleep(time.Second)
	}
}

// T1
func main_() {

	// T2
	go task("Task 1") // New Thread

	// T3
	go task("Task 2") // New Thread

	// T1
	task("Task 3")
	// If you put go in the front of the above line the function will be putted on another thread and the main thread program
	// will not have anything to run and then will stop execution
}

// T1
func main() {

	// T2
	go task("Task 1") // New Thread

	// T3
	go task("Task 2") // New Thread

	// T1
	task("Task 3")
	// If you put go in the front of the above line the function will be putted on another thread and the main thread program
	// will not have anything to run and then will stop execution
}
