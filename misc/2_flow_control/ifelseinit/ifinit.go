package main

import (
	"fmt"
	"math/rand"
	"time"
)

func randomNumber () int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return r.Intn(10)
}

func main() {
	if i:=randomNumber(); i>4{
		fmt.Println(i)
	}else {
		fmt.Println("<=4")
	}
}