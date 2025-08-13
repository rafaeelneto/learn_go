package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Tags on structs
type Conta struct {
	Numero int `json:"numero"`
	// Tag - is optional or "omit"
	Saldo int `json:"-"`
}

func main() {
	conta := Conta{
		Numero: 1232,
		Saldo:  2309,
	}

	fmt.Println(conta)

	// returns in bytes
	res, err := json.Marshal(conta)
	if err != nil {
		panic(err)
	}

	println(string(res))

	enconder := json.NewEncoder(os.Stdout)

	err = enconder.Encode(conta)
	if err != nil {
		panic(err)
	}

	jsonPuro := []byte(`{"numero": 231, "saldo": 90}`)
	var contaReceived Conta

	err = json.Unmarshal(jsonPuro, &contaReceived)
	if err != nil {
		panic(err)
	}

	fmt.Println(contaReceived)
}
