package main

import "fmt"
import "math"

type Circle struct {
	x float64
	y float64
	r float64
}


func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

type Rectangle struct {
	l,w float64
}

func (r *Rectangle) area() float64 {
	return r.l*r.w
}

type Shape interface {
	area() float64
}

func printArea(s Shape) {
	fmt.Println(s.area())
}

func main() {
	var c Circle
	c2 := new(Circle)
	c3 := Circle{x: 0, y:0, r: 10}
	c4 := Circle{0,0,5}

	fmt.Println(c.x)
	fmt.Println(c2.y)
	fmt.Println(c3.r)
	fmt.Println(c4)

	r1 := Rectangle{l:4, w:4}

	printArea(&c3)
	printArea(&r1)
}