package main

import "fmt"

func main() {
	var salarios = map[string]int{"rafael": 1, "renato": 2, "lucas": 3}

	fmt.Println(salarios)
	fmt.Println(salarios["rafael"])

	delete(salarios, "rafael")

	fmt.Println(salarios)

	sal := make(map[string]int)

	fmt.Println(sal)

	for nome, salario := range salarios { // blank identifier _
		fmt.Printf("%s é %d \n", nome, salario)
	}
}
