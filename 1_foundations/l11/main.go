package main

import "fmt"

type Client struct {
	name   string
	age    int
	active bool
}

func main() {
	client := Client{
		name:   "Rafael",
		age:    30,
		active: true,
	}

	fmt.Println(client)
	fmt.Println(client.name)
}
