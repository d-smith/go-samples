#Concurrency

##Goroutines
* A goroutine is a lightweight thread managed by the Go runtime.
* go f(x, y, z) starts a new goroutine running f(x,y,z)
* The evaluation of f, x, y, and z happens in the current goroutine and the execution of f happens 
in the new goroutine.
* Goroutines run in the same address space, so access to shared memory must be synchronized. 
The sync package provides useful primitives, although you won't need them much in Go as there are 
other primitives. (See the next slide.)

##Channels
* Channels are a typed conduit through which you can send and receive values

			ch <- v    // Send v to channel ch.
			v := <-ch  // Receive from ch, and
			           // assign value to v.

* The data flows in the direction of the arrow.
* Like maps and slices, channels must be created before use:

			ch := make(chan int)

* By default, sends and receives block until the other side is ready. This allows goroutines
to synchronize without explicit locks or condition variables.

##Buffered Channels
* Channels can be buffered. Provide the buffer length as the second argument 
to make to initialize a buffered channel: ch := make(chan int, 100)
* Sends only block when the channel is full, recieves only block when the 
channel is empty

##Range and Close
* A sender can close a channel to indicate that no more values will be sent. Receivers c
an test whether a channel has been closed by assigning a second parameter 
to the receive expression: after

			v, ok := <-ch

* ok is false if there are no more values to receive and the channel is closed.
* The loop for i := range c receives values from the channel repeatedly until it is closed.
* Note: Only the sender should close a channel, never the receiver. Sending on a closed 
channel will cause a panic.
* Another note: Channels aren't like files; you don't usually need to close them. 
Closing is only necessary when the receiver must be told there are no more values coming, such 
as to terminate a range loop.

##Select
* The select statement lets a goroutine wait on multiple communication operations.
* A select blocks until one of its cases can run, then it executes that case. It 
chooses one at random if multiple are ready.
* The default case in a select is run if no other case is ready. se a default case to try 
a send or receive without blocking:

			select {
			case i := <-c:
			    // use i
			default:
			    // receiving from c would block
			}

Binary Tree Equivalence

			package main

			import (
				"code.google.com/p/go-tour/tree"
				"fmt"
			)

			// Walk walks the tree t sending all values
			// from the tree to the channel ch.
			func Walk(t *tree.Tree, ch chan int) {
			    var walker func(t *tree.Tree)
			    walker = func (t *tree.Tree) {
			        if (t == nil) {
			            return
			        }
			        walker(t.Left)
			        ch <- t.Value
			        walker(t.Right)
			    }
			    walker(t)
			    close(ch)
			}


			// Same determines whether the trees
			// t1 and t2 contain the same values.
			func Same(t1, t2 *tree.Tree) bool {
			ch1, ch2 := make(chan int), make(chan int)

			    go Walk(t1, ch1)
			    go Walk(t2, ch2)

			    for {
			        v1,ok1 := <- ch1
			        v2,ok2 := <- ch2

			        if v1 != v2 || ok1 != ok2 {
			            return false
			        }

			        if !ok1 {
			            break
			        }
			    }

			    return true
			}

			func main() {
				fmt.Println(Same(tree.New(1), tree.New(1)))
			    fmt.Println(Same(tree.New(1), tree.New(2)))
			}

Where to go from here: https://tour.golang.org/concurrency/10