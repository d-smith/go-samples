package main

import "fmt"

func average(xs []float64) float64 {
	total := 0.0
	for _,v := range xs {
		total += v
	}

	return total / float64(len(xs))
}


func main() {
	xs := []float64 { 98, 93, 77, 82, 84}

	fmt.Println(average(xs))
	fmt.Println(addOneAndTwo(1))
	fmt.Println(total(1,2,3))
}

func addOneAndTwo(x int) (int, int) {
	return x + 1, x + 2
}

func total(args ...int) int {
	total := 0
	for _,v := range args {
		total += v
	} 
	 return total
}