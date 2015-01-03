
package main

import "fmt"

func main() {
	i := 1
	for i <= 10 {
		fmt.Println(i)
		i = i + 1
	}

	for i := 0; i < 10; i++ {
		if i % 2 == 0 {
			fmt.Println(i)
			} else {
				fmt.Println("yeah")
			}
	}

	x := "quux"

	switch x {
		case "foo": fmt.Println("foo")
		case "bar": fmt.Println("bar")
		case "baz": fmt.Println("baz")
		default: fmt.Println("some other meta syntactic variable")
	}
}