package main

import "fmt"
import "time"
import "math/rand"

func randomSleep() {
	sleepTime := time.Duration(rand.Intn(250))
	time.Sleep(sleepTime * time.Millisecond)
}

func f(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n , ": ", i)
		randomSleep()
	}
}

func main() {
	for i:= 0; i < 10; i++ {
		go f(i)	
	}
	
	var input string
	fmt.Println("Enter some test and press return")
	fmt.Scanln(&input)
	fmt.Println("Your text: ", input)
}