# Methods, etc

Methods
* Go does not have classes, but structures can have methods.
* The method receiver appears in its own argument list between the func keyword and the method name.

      type Vertex struct {
      	X, Y float64
      }

      func (v *Vertex) Abs() float64 {
      	return math.Sqrt(v.X*v.X + v.Y*v.Y)
      }

      func main() {
      	v := &Vertex{3, 4}
      	fmt.Println(v.Abs())
      }

* You can declare a method on any type that is declared in your package, not just struct types.
* However, you cannot define a method on a type from another package (including built in types).

      func (f MyFloat) Abs() float64 {
      	if f < 0 {
      		return float64(-f)
      	}
      	return float64(f)
      }

      func main() {
      	f := MyFloat(-math.Sqrt2)
      	fmt.Println(f.Abs())
      }

Methods with Pointer Receivers
* Methods can be associated with a named type or a pointer to a named type.
* There are two reasons to use a pointer receiver. First, to avoid copying
the value on each method call (more efficient if the value type is a large
struct). Second, so that the method can modify the value that its receiver
points to.


      type Vertex struct {
      	X, Y float64
      }

      func (v *Vertex) Scale(f float64) {
      	v.X = v.X * f
      	v.Y = v.Y * f
      }

      func (v *Vertex) Abs() float64 {
      	return math.Sqrt(v.X*v.X + v.Y*v.Y)
      }

      func main() {
      	v := &Vertex{3, 4}
      	v.Scale(5)
      	fmt.Println(v, v.Abs())
      }

Interfaces
* An interface type is defined by a set of methods.
* A value of interface type can hold any value that implements those methods.
* A type implements an interface by implementing the methods. There is no
explicit declaration of intent; no "implements" keyword.
* Implicit interfaces decouple implementation packages from the packages that
define the interfaces: neither depends on the other.
* It also encourages the definition of precise interfaces, because you don't
have to find every implementation and tag it with the new interface name.

Stringers
* One of the most ubiquitous interfaces is Stringer defined by the fmt package.

      type Stringer interface {
          String() string
      }

* A Stringer is a type that can describe itself as a string. The fmt package
(and many others) look for this interface to print values.

      type Person struct {
      	Name string
      	Age  int
      }

      func (p Person) String() string {
      	return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
      }

      func main() {
      	a := Person{"Arthur Dent", 42}
      	z := Person{"Zaphod Beeblebrox", 9001}
      	fmt.Println(a, z)
      }

Stringer Exercise

      package main

      import "fmt"

      type IPAddr [4]byte

      func (ipaddr IPAddr) String() string {
      	return fmt.Sprintf("%d.%d.%d.%d",
      		ipaddr[0], ipaddr[1],ipaddr[2],ipaddr[3])
      }

      func main() {
      	addrs := map[string]IPAddr{
      		"loopback":  {127, 0, 0, 1},
      		"googleDNS": {8, 8, 8, 8},
      	}
      	for n, a := range addrs {
      		fmt.Printf("%v: %v\n", n, a)
      	}
      }

Errors
* Go programs express error state with error values.
* The error type is a built-in interface simliar to fmt.Stringer

      type error interface {
          Error() string
      }

* Functions often return an error value, and calling code should handle
errors by testing whether the error equals nil.


Error Exercise

      type ErrNegativeSqrt float64


      func (e ErrNegativeSqrt) Error() string {
      	return fmt.Sprintf("cannot Sqrt negative number: %f", float64(e))
      }



      func Sqrt(x float64) (float64, error) {
      	if(x < 0) {
      		return 0, ErrNegativeSqrt(x)
      	}

      	return 0, nil
      }

      func main() {
      	fmt.Println(Sqrt(2))
      	fmt.Println(Sqrt(-2))
      }

Readers
* The io package specifies the io.Reader interface, which represents the
read end of a stream of data.
* The Go standard library contains many implementations of these interfaces,
including files, network connections, compressors, ciphers, and others.
* The io.Reader interface has a Read method:

      func (T) Read(b []byte) (n int, err error)

*Read populates the given byte slice with data and returns the number of bytes
populated and an error value. It returns an io.EOF error when the stream ends.


      func main() {
      	r := strings.NewReader("Hello, Reader!")

      	b := make([]byte, 8)
      	for {
      		n, err := r.Read(b)
      		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
      		fmt.Printf("b[:n] = %q\n", b[:n])
      		if err == io.EOF {
      			break
      		}
      	}
      }


Reader Exercise

      package main

      import "code.google.com/p/go-tour/reader"

      type MyReader struct{}

      // TODO: Add a Read([]byte) (int, error) method to MyReader.
      func (mr MyReader)Read(b []byte) (int, error) {
      	for i := range b {
      	b[i] = 'A'
      	}

      	return len(b),nil
      }

      func main() {
      	reader.Validate(MyReader{})
      }


Rot 13 Reader Exercise


      package main

      import (
      	"io"
      	"os"
      	"strings"
      )

      type rot13Reader struct {
      	r io.Reader
      }

      func (rot13 *rot13Reader)Read(b []byte) (int, error) {
      	_,err := rot13.r.Read(b)
      	if(err != nil) {
      		return 0,err
      	}


      	for i := range b {
      		c := b[i]
      		switch {
      			case c >= 'a' && c <= 'm':
      				b[i] = c + 13
      			case c >= 'A' && c <= 'M':
      				b[i] = c + 13
      			case c >= 'n' && c <= 'z':
      				b[i] = c - 13
      			case c >= 'N' && c <= 'Z':
      				b[i] = c - 13
      		}
      	}

      	return len(b),nil
      }


      func main() {
      	s := strings.NewReader("Lbh penpxrq gur pbqr!")
      	r := rot13Reader{s}
      	io.Copy(os.Stdout, &r)
      }

Web Servers
* Package http serves HTTP requests using any value that implements http.Handler:

      package http

      type Handler interface {
          ServeHTTP(w ResponseWriter, r *Request)
      }
