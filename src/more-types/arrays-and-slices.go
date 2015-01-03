package main

import "fmt"

func main() {
	var x [5]int
	x[4] = 100

	fmt.Println(x)

	total := 0
	for _, value := range x {
		total += value
	}

	fmt.Println(total)

	arr := [5] float64{ 1,2,3,4,5}
	y := arr[ 0: 5]
	y2 := append(y, 6,7,8)
	fmt.Println(y2)

	slice2 := make([]float64, 2)
	copy(slice2, y2)
	fmt.Println(slice2)
}