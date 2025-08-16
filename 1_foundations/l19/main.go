package main

func main() {
	var x interface{} = "Nome"

	println(x)
	println(x.(string))

	res, ok := x.(int)

	println(ok)
	println(res)
}
