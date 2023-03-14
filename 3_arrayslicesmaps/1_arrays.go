package main

import "fmt"

func main() {
	// homogenea e estática (tamanho fixo)
	var grades [3]float64

	fmt.Println(grades)

	grades[0], grades[1], grades[2] = 7.8, 8.7, 9.1;

	// 👇 out bound error 
	// grades[3] = 9.4;

	fmt.Println(grades)

	total := 0.0;

	for i:=0; i<len(grades); i++ {
		total += grades[i]
	}

	fmt.Println(total/float64(len(grades)))
}