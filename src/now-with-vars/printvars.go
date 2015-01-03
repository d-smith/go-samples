package main

import "fmt"

func main() {
	var x string = "Hell, o World"

	fmt.Println(x)

	x = "foo"
	fmt.Println(x)

	fmt.Println(x == "bar")

	const y = "x"
	fmt.Println(y)

	var (a = 2 
		c = 4)

	fmt.Println(a)
	fmt.Println(c)
}