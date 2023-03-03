package main

import (
	"fmt"
)
type set []int

func maxSubArray(arr set) int {
	
	sumList := map[*int]set{}

	subArrays := subsets(arr)

	//! i := 0
	i := subArrays[0][0]
	max:=&i
	// split subsets
	for _, s := range subArrays {
		sum:=s.sum()
		fmt.Println(sum)
		sumList[&sum] = s
		if sum > *max {
			//! *max = sum
			max = &sum
		}
	}

	fmt.Println("max subset", sumList[max], "sum",*max)

	return *max
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

	// fmt.Println(maxSubArray([]int{5,4,-1,7,8}) == 23) // expect 23
	// fmt.Println(maxSubArray([]int{-2,1,-3,4,-1,2,1,-5,4}) == 6) // expect 6
	// fmt.Println(maxSubArray([]int{-2,-1,-3}) == -1) // expect -1
	// fmt.Println(maxSubArray([]int{3, -2, -3, -3, 1, 3, 0}) == 4) // expect 4
	// fmt.Println(maxSubArray([]int{})) // expect 4

	
}