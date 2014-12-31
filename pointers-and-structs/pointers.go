package main

import "fmt"

func setToEleven(x *int) {
	*x = 11
}

func main() {
	var x int = 0
	setToEleven(&x)
	fmt.Println(x)

	var x2 *int = new(int)
	setToEleven(x2)
	fmt.Println(x)
}