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

// Any type that has a desativar method will implement the Pessoa interface
type Pessoa interface {
	desativar()
}

func (c Client) desativar() {
	c.active = false

	fmt.Printf("O client %s foi desativado \n", c.name)
}

func desativacao(pessoa Pessoa) {
	pessoa.desativar()
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

	fmt.Println(client.active)

	client.desativar()

	fmt.Println(client.active)

	desativacao(client)
}
