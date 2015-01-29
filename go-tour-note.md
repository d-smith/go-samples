#Go Tour - Notes

Note taken while going through the [Go Tour](http://tour.golang.org/list)

## Basics - Packages, Variables, Functions

Packages
* Every go program consists of a set of packages
* Runnable programs are in the main package
* Convention - package name is the same as the last element of the import path, e.g. "math/rand" has files that begin with `package rand`


Import
* Packages are imported via import, after you can refer to the names exported by the package, which start with a capital letter
* Imported packages can be imported with an import per package per line, or
import with parens can be used (known as the factored import statement, which is the
preferred style)

Functions
* Can have zero or more arguments
* When two or more consecutive named function params have the same type, you can omit the type from all but the last
* A function can return multiple results; return values can be named and used like input parameters
* A return statement that omits the returned params (know as the naked return) returns the current values
of the return parameters

Variables
* Named using the var keyword
* Variables can be defined at the function or package level
* Variable declarations can have initializers, one per var
* Inside a function a variable can be declared and initialized with the `:=` assignment
instead of using the var keyword

Basic Types
* bool
* string
* int, int8, int16, int32, int64
* uint, uint8, uint16, uint32, uint64, uintptr
* byte (alias for uint8)
* rune (alias for int32, represents unicode code point)
* float32, float64
* complex64, complex128

Zero Values
* Variables not explicitly initialized are given their zero values
* 0 for numeric types, false for bool, "" for strings

Type Conversions
* T(v) converts value v to type T
* Example: var f float64 = float64(i)

Type Inference
* When declaring a variable without specifying its type, the type is inferred
from the right hand side of the assignment
* Precision of untyped numeric literals used to infer the type - int,
float64, or complex128

Constants
* Declared like variable but with the const keyword
* Cannot be declared using :=
* Numeric constants are high precision values


##Basic - Flow Control

For
* The only looping construct in go
* Looks like C or Java, exception no parans and brackets are required

      for i := 0; i < 10; i++ {
      		sum += i
      }

* Pre and post can be left empty

      for ; sum < 1000; {
          sum += sum
      }

* For the above you can drop the semi-colons, now it's a C while

  	for sum < 1000 {
  		sum += sum
  	}

* Loop forever

      for {
      }

If
* Like C and Java, except no parens and {} is required

      if x < 0 {
      	return sqrt(-x) + "i"
      }

* Can include a short statement to execute before the statement
* Vars declared in the short statement are accessible in any else blocks

      if v := math.Pow(x, n); v < lim {
        return v
      }


Exercise: Newton's Method

    package main

    import (
    	"fmt"
    )

    func Sqrt(x float64) float64 {
    	z := float64(1)
    	for i := 0; i < 10; i++ {
    		z = z - ((z*z) - x)/(2*z)
    	}

    	return z
    }

    func main() {
    	fmt.Println(Sqrt(2))
    }

Switch
* Case bodies break automatically unless ended with a fallthrough statement
* Evaluated top to bottomm stopping when a case succeeds

    	switch os := runtime.GOOS; os {
    	case "darwin":
    		fmt.Println("OS X.")
    	case "linux":
    		fmt.Println("Linux.")
    	default:
    		// freebsd, openbsd,
    		// plan9, windows...
    		fmt.Printf("%s.", os)
    	}


Switch with No Condition
* Switch without a condition is the same as switch true.
* This construct can be a clean way to write long if-then-else chains.

Defer
* A defer statement defers the execution of a function until the surrounding function returns.
* The deferred call's arguments are evaluated immediately, but the function call is not executed until the surrounding function returns.
* Deferred function calls are pushed onto a stack.
* When a function returns, its deferred calls are executed in last-in-first-out order.

##Go Tour Lessons
http://tour.golang.org/list

##More Information
* [Article on Go's declaration syntax](http://golang.org/doc/articles/gos_declaration_syntax.html)
* [More on defer statements](http://blog.golang.org/defer-panic-and-recover)
