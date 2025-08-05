package main

import "fmt"

type Address struct {
	street   string
	number   string
	zip_code string
	state    string
	country  string
}

type Client struct {
	name   string
	age    int
	active bool
	// Composition
	Address

	// Creating a variable with the type of Address
	endereco Address
}

func main() {
	client := Client{
		name:   "Rafael",
		age:    30,
		active: true,

		endereco: Address{
			street: "aaaaa",
		},
	}

	fmt.Println(client)

	// Can be called like this
	client.Address.street = "Rua miau"
	// Or this
	client.country = "BRASIL"
	fmt.Println(client.name)
}
