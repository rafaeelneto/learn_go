package main

import "fmt"

func weekendBuys(job1, job2 bool) (bool, bool, bool){
	buyTv50 := job1 && job2; 
	buyTv32 := job1 != job2; // ou exclusivo 
	buyIceCream := job1 || job2; 

	return buyTv50, buyTv32, buyIceCream
}

func main(){
	tv50, tv32, iceCream := weekendBuys(true, true)
	fmt.Printf("%v %v %v %v", tv50, tv32, iceCream, !iceCream)
}