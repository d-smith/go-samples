package main

import "fmt"

func main() {
	var x = make(map[string] int)
	x["foo"] = 10
	x["bar"] = 11
	fmt.Println(x)

	fmt.Println(len(x))

	delete(x, "bar")

	fmt.Println(x)
	delete(x, "flibby")
}