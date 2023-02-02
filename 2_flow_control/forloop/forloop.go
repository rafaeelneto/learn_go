package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0
	for i < 10 { // like a while stament
		fmt.Print(i, " ")
		i++
	}

	for i := 0; i<20; i+=2 {
		fmt.Println(i)
	}

	// infinite loop
	for {
		fmt.Println("Running...")
		time.Sleep(time.Second*2)
	}
}