package main

func TestAsIWouldDoIt(t *testing.T) {
	positiveSize:=10
	arr:=[]int{}
	for i := 0; i <= 10; i++ {
		arr = append(arr, i-positiveSize)
	}
	for i := 1; i <= 10; i++ {
		arr = append(arr, i)
	}
}