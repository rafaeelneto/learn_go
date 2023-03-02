package main

import "fmt"

func gradeRangeClassification( grade float64) string {

	a := "Aprovado"
	r := "Reprovado"

	switch grade {
	case 10:
		fallthrough 
	case 9:
		return a

	case 4: 
		return r
	default:
		return "Invalido"
	}
}

func main() {
	classification := gradeRangeClassification(5)
	fmt.Println(classification)
}