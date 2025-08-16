package main

type MyNumber int

type Number interface {
	~int | float64 | float32
}

func sum[T Number | float32](m map[string]T) T {
	var total T = 0
	for _, v := range m {
		total += v
	}
	return total
}

// Only to equallity
func compare[T comparable](a, b T) bool {
	if a == b {
		return true
	}

	return false
}

func main() {
	myMap := map[string]int{
		"oi":    2,
		"miau":  38,
		"aviao": 90,
	}
	myMapCustomType := map[string]MyNumber{
		"oi":    2,
		"miau":  38,
		"aviao": 90,
	}

	mapFloat := map[string]float64{
		"oi":    2.3,
		"miau":  38.34,
		"aviao": 90.34,
	}

	println(sum(myMap))

	println(sum(mapFloat))
	println(sum(myMapCustomType))

	println(compare(10, 10))
}
