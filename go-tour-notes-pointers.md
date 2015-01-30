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
