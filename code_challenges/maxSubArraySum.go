package main

import (
	"fmt"
)
type set []int


// me explica sua dúvida
// qual função eu chamo pra inputar o array?
// subsets? mas ela não chama as outras e retorna os subarrays
// eu fragmentei o processo em 3:
// - separa os subsets
// - soma os valores dos subsets
// - verifica qual o maior
// Essa função sample faz o processo inteiro

// aaaaaaa acho que entendi

func sampleV1() {
	// creates a mirrored array with size 2n + 1
	positiveSize := 2 // will have 5 elements
	arr := []int{}
	for i := 0; i <= positiveSize; i++ {
		arr = append(arr, i-positiveSize)
	}
	for i := 1; i <= positiveSize; i++ {
		arr = append(arr, i)
	}

	// map pode dar ruim se mais de um subset tiver a mesma soma
	// preferi o map pq vc pode usar um só loop
	// a solução é usar *int ao invés de int <- isso daqui é legal vc entender
	// estas usando ponteiro pq mesmo se tiver valores de soma igual eles não vão se sobreescrever no map certo?
	// pera que vou precisar estudar com cuidado esse algoritmo
	sumList := map[*int]set{}
	i:=0
	max:=&i
	// split subsets
	for _, s := range subsets(arr) {
		sum:=s.sum()
		sumList[&sum] = s
		if sum > *max { 
			*max = sum
		}
	}
	
	// aqui deveria retornar o max subset certo?
	// pode ser, basta renomear a função e colocar um return na assinatura
	fmt.Println("max subset", sumList[max])
}

func sampleV2(arr set) {


	// map pode dar ruim se mais de um subset tiver a mesma soma <- isso aqui pode dar ruim
	// inclusive, quase todas as somas tem 2 subsets
	// preferi o map pq vc pode usar um só loop
	// a solução é usar *int ao invés de int <- isso daqui é legal vc entender
	// estas usando ponteiro pq mesmo se tiver valores de soma igual eles não vão se sobreescrever no map certo?
	// pera que vou precisar estudar com cuidado esse algoritmo
	sumList := map[int]set{}
	max:=0
	// split subsets
	for _, s := range subsets(arr) {
		sum:=s.sum()
		sumList[sum] = s
		if sum > max { 
			max = sum
		}
	}//agora rodou como esperado
	// tendi
	// mas tem aquele risco de dar ruim ali né?
	
	// aqui deveria retornar o max subset certo?
	// pode ser, basta renomear a função e colocar um return na assinatura
	fmt.Println("max subset", sumList[max],"sum",max)
}

func subsets(nums []int) []set {
	subarr := []set{}
	for i := 0; i < len(nums); i++ {
		for j := range nums {
			if i+j > len(nums) {
				// this is the part that makes it below O(n2)
				continue
			}
			subset := nums[i : i+j]
			if len(subset) > 0 {
				subarr = append(subarr, subset)
			}
		}
	}
	return append(subarr, nums)
}

// this is the part that makes it avoe O(2n)
func (s set) sum() int {
	sum := 0
	for i := range s {
		sum += s[i]
	}
	return sum
}


func main() {
	// vou rodar um dos testes
	sampleV2([]int{5,4,-1,7,8}) // expect 23
	sampleV2([]int{-2,1,-3,4,-1,2,1,-5,4}) // expect 6

	sampleV1() // expect 23
	sampleV1() // expect 6
	// funcionou bacana
	// fmt.Println(subsets([]int{5,4,-1,7,8})) // expect 23

	// // essa aqui retornou 0, mas não deveria. na vdd deveria kkk
	// fmt.Println(subsets([]int{-2,-1})) // expect -1
	// fmt.Println(subsets([]int{-2,-3,-1})) // expect -1
}