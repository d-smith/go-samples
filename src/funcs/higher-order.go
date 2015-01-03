package main

import "fmt"

func printIt(s string) {
	fmt.Println(s)
}

func doSomething(p func(string)) {
	x := "Some string..."
	p(x)
}

func repeatThing(i int, f func(string), farg string) {
	for  idx := 0; idx < i; idx++ {
		f(farg)
	}
}

func main() {
	f := printIt
	doSomething(f)

	repeatThing(3, printIt, "Repeat this")
}