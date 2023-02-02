package main

import "fmt"

func printGradeResult (grade float64){
	if grade >= 7 {
		fmt.Println("Aprovado com nota", grade)
		} else {
		fmt.Println("Reprovado com nota", grade)
		
	}
}

func main() {
	printGradeResult(6.6)
	printGradeResult(8.7)

}