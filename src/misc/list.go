package main

import (
	"container/list"
	"fmt"
)

func printList(l *list.List) {
	fmt.Println("-----")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value.(int))
	}
}

func main() {
	var x list.List
	x.PushBack(1)
	x.PushBack(2)
	x.PushBack(3)

	printList(&x)

	x.PushFront(0)

	var y list.List
	y.PushBack((4))
	y.PushBack((5))

	x.PushBackList(&y)

	printList(&x)

}
