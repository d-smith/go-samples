# Go Tour Notes - Pointers, Structs, etc

## Pointers

* A pointer holdes the memory address of a variable
* The type *T is a pointer to a T value.
* The zero value of a pointer is nil.
* The & operator generates a pointer to its operand.
* The * operator denotes the pointers underlying value

      package main

      import "fmt"

      func main() {
      	i, j := 42, 2701

      	p := &i         // point to i
      	fmt.Println(*p) // read i through the pointer
      	*p = 21         // set i through the pointer
      	fmt.Println(i)  // see the new value of i

      	p = &j         // point to j
      	*p = *p / 37   // divide j through the pointer
      	fmt.Println(j) // see the new value of j
      }

## Structs

* A struct us a collection of fields

      type Vertex struct {
      	X int
      	Y int
      }

      func main() {
      	fmt.Println(Vertex{1, 2})
      }

* Struct fields are accessed with a dot, e.g. v.X
* Struct fields can be accessed through a struct pointer. The
indirection through the pointer is transparent.
* A struct literal denotes a newly allocated struct value by listing the values of its fields.
* The special prefix & returns a pointer to the struct value.

      package main

      import "fmt"

      type Vertex struct {
      	X, Y int
      }

      var (
      	v1 = Vertex{1, 2}  // has type Vertex
      	v2 = Vertex{X: 1}  // Y:0 is implicit
      	v3 = Vertex{}      // X:0 and Y:0
      	p  = &Vertex{1, 2} // has type *Vertex
      )

      func main() {
      	fmt.Println(v1, p, v2, v3)
      }

## Arrays

* The type [n]T is an array of n values of type T
* The array length is part of its type - arrays cannot be resized

## Slices
* A slice points to an array of values and also includes a length.
* []T is a slice with elements of type T
* Slices can be re-sliced, creating a new slice value that points to the same array.
* s[lo:hi] evaluates to a slice of the elements from lo through hi-1, inclusive
* s[lo:lo] is empty
* s[lo:lo+1] has one element

      package main

      import "fmt"

      func main() {
      	p := []int{2, 3, 5, 7, 11, 13}
      	fmt.Println("p ==", p)
      	fmt.Println("p[1:4] ==", p[1:4])

      	// missing low index implies 0
      	fmt.Println("p[:3] ==", p[:3])

      	// missing high index implies len(s)
      	fmt.Println("p[4:] ==", p[4:])
      }
* Slices are created with the make function. It works by allocating a zeroed array and returning a slice that refers to that array, e.g. a := make([]int, 5)
* To specify a capacity, pass a third argument to make: b := make([]int, 0, 5)
* The zero value of a slice is nil. A nil slice has
a length and capacity of 0.

Adding elements to a slice:
* func append(s []T, vs ...T) []T
* If the backing array is too small a bigger array will be allocated.

Range
* The range form of the for loop iterates over a slice or map.
* You can skip the index or value by assigning to _
* If you only want the index, drop the ", value" entirely.

      var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

      func main() {
      	for i, v := range pow {
      		fmt.Printf("2**%d = %d\n", i, v)
      	}
      }

      pow := make([]int, 10)
      for i := range pow {
      	pow[i] = 1 << uint(i)
      }
      for _, value := range pow {
      	fmt.Printf("%d\n", value)
      }

Slices Exercise

    package main

    import "code.google.com/p/go-tour/pic"

    func Pic(dx, dy int) [][]uint8 {
    	picdata := make([][]uint8, dy)
    	for i := range picdata {
    		picdata[i] = make([]uint8,dx)
    	}

    	for y := range picdata {
    		for x := range picdata[y] {
    			picdata[y][x] = uint8 (x*y)
    		}
    	}

    	return picdata
    }

    func main() {
    	pic.Show(Pic)
    }

## Maps

* A map maps keys to values
* Maps must be created with make before use, the nil map is empty and cannot
be assigned to

      type Vertex struct {
      	Lat, Long float64
      }

      var m map[string]Vertex

      func main() {
      	m = make(map[string]Vertex)
      	m["Bell Labs"] = Vertex{
      		40.68433, -74.39967,
      	}
      	fmt.Println(m["Bell Labs"])
      }

* Map literals are like struct literals, but the keys are required.

      var m = map[string]Vertex{
      	"Bell Labs": Vertex{
      		40.68433, -74.39967,
      	},
      	"Google": Vertex{
      		37.42202, -122.08408,
      	},
      }

* If the top-level type is just a type name, you can omit it from the elements of the literal.

      var m = map[string]Vertex{
      	"Bell Labs": {40.68433, -74.39967},
      	"Google":    {37.42202, -122.08408},
      }

Working with maps
* m[key] = elem
* elem = m[key]
* delete(m, key)
* elem, ok = m[key] - If key is in m, ok is true. If not, ok is false and elem is the zero value for the map's element type.
Similarly, when reading from a map if the key is not present the result is the zero value for the map's element type.


Map Exercise

      package main

      import (
      	"code.google.com/p/go-tour/wc"
      	"strings"
      )

      func WordCount(s string) map[string]int {

      	wcmap := make(map[string]int)

      	words := strings.Fields(s)
      	for i:= range words {
      		word := words[i]
      		count := wcmap[word]
      		wcmap[word] = count + 1
      	}
      	return wcmap
      }

      func main() {
      	wc.Test(WordCount)
      }

Function Values
* Functions are values too

      func main() {
      	hypot := func(x, y float64) float64 {
      		return math.Sqrt(x*x + y*y)
      	}

      	fmt.Println(hypot(3, 4))
      }


Closures
* Go functions may be closures. A closure is a function value that references
variables from outside its body. The function may access and assign to the
referenced variables; in this sense the function is "bound" to the variables.

      func adder() func(int) int {
      	sum := 0
      	return func(x int) int {
      		sum += x
      		return sum
      	}
      }

      func main() {
      	pos, neg := adder(), adder()
      	for i := 0; i < 10; i++ {
      		fmt.Println(
      			pos(i),
      			neg(-2*i),
      		)
      	}
      }


Exercise - Function returning a function that returns successive
Fibonacci numbers

      package main

      import "fmt"

      // fibonacci is a function that returns
      // a function that returns an int.
      func fibonacci() func() int {
      	a, b := 0, 1

      	return func() int {
      		a, b = b, a+b
      		return a
      	}
      }

      func main() {
      	f := fibonacci()
      	for i := 0; i < 10; i++ {
      		fmt.Println(f())
      	}
      }
