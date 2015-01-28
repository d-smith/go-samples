#Go Tour - Notes

Note taken while going through the [Go Tour](http://tour.golang.org/list)

## Basics

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

(Last was http://tour.golang.org/basics/15) 


More Information
* [Article on Go's declaration syntax](http://golang.org/doc/articles/gos_declaration_syntax.html)
