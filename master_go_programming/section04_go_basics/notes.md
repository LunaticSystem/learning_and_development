## Variables in Go
* A variable is a name for a memory location where a value of a specific type is stored.
* In Go a variable belongs and is created at runtime.
* A Declared variable *<b>MUST</b>* be used or we get an error
* "_" is the blank identifier and mutes the compile-time error returned by unused variables.
* Examples:
  * Using the var keyword `var x int = 7` or `var s1 string` or `s1 = "Learning Go!"
  * Using the Short Declaration Operator `:=` like `age := 30`
* No variable into another variable if using `:=` like `age := name`

## Multiple Declerations

**EXAMPLES**
```
var (
    salary float64
    firstName string
    gendar bool
)
```

or 

```
var a, b, c int
fmt.Println(a, b, c)
```

## Types and Zero Values
* Does type checking at compile time.
* Must provide type at compile time or it can infer type (not always accurate)
* variables of one type cannot be stored in variables of another type until converted.
* Example:
  ```
  package main

  import "fmt"

  func main() {
    var a = 4 // int
    var b = 5.2 // float64

    a = b // will give an error at compile time.
    a = int(b) // Converts b to int type and stores it in variable a
    fmt.Println(a, b)

    var x int // Variable x is of type int
    x = "5" // String literal cannot be added to a variable of type int
  }
  ```

* Will fail if variables are uninitialized.
* Example of value zero:
  ```
  var value int
  var price float64
  var name string
  var done bool
  fmt.Println(value, price, name, done)
  ```
* Result:
  ```
  0 0  false
  ```
* Go Zero Values:
  * numeric types: 0
  * bool types: false
  * string type: "" (empty string)
  * point type: nil
 
## Comments in Go

* Comments is text that tells a user about what something does.
* Comment format:
  ```
  // This is a comment
  ```
* Comments are not executed.
* Multiple Line Comments:
  ```
  /* Starts Comment
  ...
  */ Ends Comment
  ```
* Idiotmatic to use "//" only use "/*" for debugging.
* Inline comments should not be used to often.

## Naming Conventions In Go

* Names start with a letter or an underscore (_)
* Case matters: quickSort and QuickSort are differnt variables.
* Go keywords (25) can not be used as names
* Use the first letters of the words
  ```
  var mv int // mv -> max value
  ```
* Use fewer letters in smaller scopes and the complete word in larger scopes
  ```
  var packetsReceived int // NOT OK, to verbose
  var n int // OK -> no. of packets received
  var taskDone bool //ok in larger scopes
  ```

## Package "fmt"

### Docs
* [FMT Package Docs](https://pkg.go.dev/fmt)

### Example
```go
/////////////////////////////////
// Package fmt
// Go Playground: https://play.golang.org/p/JGb4akovl8W
/////////////////////////////////
 
package main
 
// Package fmt implements formatted I/O with functions analogous to C's printf and scanf.
// It's used mainly to print out to stdout
import "fmt"
 
func main() {
 
    // fmt.Println() writes to standard output.
    // spaces are always added between operands and a newline is appended.
    fmt.Println("Hello Go World!") // => Hello Go World!
 
    var name, age = "Andrei", 35
    fmt.Println(name, "is", age, "years old.") // => Andrei is 35 years old.
 
    //** fmt.Printf() **//
 
    // fmt.Printf() prints out to stdout according to a format specifier called verb.
    // It doesn't add a newline (\n)
 
    // VERBS:
    // %d -> decimal
    // %f -> float
    // %s -> string
    // %q -> double-quoted string
    // %v -> value (any)
    // %#v -> a Go-syntax representation of the value
    // %T -> value Type
    // %t -> bool (true or false)
    // %p -> pointer (address in base 16, with leading 0x)
    // %c -> char (rune) represented by the corresponding Unicode code point
 
    a, b, c := 10, 15.5, "Gophers"
    grades := []int{10, 20, 30}
 
    fmt.Printf("a is %d, b is %f, c is %s \n", a, b, c)    // => a is 10, b is 15.500000, c is Gophers
    fmt.Printf("%q\n", c)                      // => "Gophers"
    fmt.Printf("%v\n", grades)                 // => [10 20 30]
    fmt.Printf("%#v\n", grades)                // => b is of type float64 and grades is of type []int
    fmt.Printf("b is of type %T and grades is of type %T\n", b, grades) 
    // => b is of type float64 and grades is of type []int
    fmt.Printf("The address of a: %p\n", &a)    // => The address of a: 0xc000016128
    fmt.Printf("%c and %c\n", 100, 51011)       // =>  d and 읃  (runes for code points 101 and 51011)
 
    const pi float64 = 3.14159265359
    fmt.Printf("pi is %.4f\n", pi) // => formatting with 4 decimal points
 
    // %b -> base 2
    // %x -> base 16
    fmt.Printf("255 in base 2 is %b\n", 255)  //  => 255 in base 2 is 11111111
    fmt.Printf("101 in base 16 is %x\n", 101) // => 101 in base 16 is 65
 
    // fmt.Sprintf() returns a string. Uses the same verbs as fmt.Printf()
    s := fmt.Sprintf("a is %d, b is %f, c is %s \n", a, b, c)
    fmt.Println(s) // => a is 10, b is 15.500000, c is Gophers
}
```

## Constants in Go

* Constants are used to represent fixed (unchanging) values
* We use constants to avoid possible errors (variables that change when they shouldn't) or to replace a value only in one place instead of in many places
* All basic literals (1, 3.4, "hello", true) are in fact unnamed constants
* A constant belongs to the compile time and its created at compile time. It's value cannot be changed while the program is running.
* An advantage of using constants is that go can not detect runtime errors at compile-time but constants belong to compile time so errors can be detected earlier.
* You can declare constants that store numbers, strings or booleans.

### Examples

```go
package main

func main(){
  const days int = 7
  
  var i int
  fmt.Println(i)

  const pi float64 // if not declared with a variable this will fail
  const secondsInHour = 3600

  duration := 234 //in hours
  fmt.Printf("Duration in seconds: %v\n", duration*secondsInHour)

  x, y := 5, 0 // not a constant
  fmt.Println(x / y) // Evaluation of this will not happen at compile time but during runtime.

  const a = 5
  const b = 0

  fmt.Println(a / b) // This will error at compile time

  // Declaring multiple constants
  const n, m int = 4, 5
  const n1, m1 = 6, 7

  // Declaring multiple constants at the same time (Grouped Constants)
  const (
    min1 = -500
    min2 = -300
    min3 = 100
  )

  fmt.Println(min1, min2, min3)

  // Constants undeclared in a grouped constant block recieve their value from the previous one with a value.
    const (
    min1 = -500
    min2
    min3
  )

  fmt.Println(min1, min2, min3)
}
```

## Constant Rules

1. You cannot change a constant
  ```go
    const temp int = 100
    temp = 50 // This will give an error
  ```
2. You can not initiate a constant at runtime
```go
  const power = math..Pow(2, 3)
```
3. You can not use a variable to initialize a constant
```go
  t := 5
  const tc = t
```
4. You can use a function like len (len returns the number of elements in a string) to ainitialize a constant if it has as an argument a constant string literal to a variable. String literally is an unnnamed constant.
```go
  const l1 = len("hello")
```

## Constant Expressions: Typed vs. Untyped Constants

* Typed constants behave differently than untyped constants
```go
package main

func main(){
  const a float64 = 5.1 // typed constant

  const b = 6.7 // untyped constant

  const c float64 = a * b
  const str = "Hello " + "Go!"

  const d = 5 > 10
  fmt.Println(d)

  const x int = 5
  const y float64 = 2.2 * x // fails as you cannot multple typed and untyped constants

  const x = 5
  const y = 2.2 * 5 // This is allowed because we are not explicitly setting a type so go assumes the type and doesn't fail with the normal strict guidlines of typed and untyped values being multiplied together.
  fmt.Printf("%T\n", y)

  var i int = x // x changes to int
  var j float64 = x // var j float64 = float64(x)
  var p byte = x // var p byte = byte(x)
}
```
* Untyped constant is implicitely given a type only when used in an expression
```go
 const r = 5 // implicitly an int
 var rr = r
 fmt.Printf("%T", rr) // Shows that var rr which is const r is an int
```
* Untyped constants need to be assigned to variables in order to set the default type.

## IOTA
* Within a constant declaration, the predeclared identifier iota represents successive untyped integer constants. Its value is the index of the respective ConstSpec in that constant declaration, starting at zero. It can be used to construct a set of related constants.
```go
package main

func main() {
  const (
    c1 = iota
    c2 = iota
    c3 = iota
  )

  fmt.Println(c1, c2, c3) // prints out 0, 1, 2

  // Simplified
  const (
    c11 = iota
    c22
    c33
  )
  fmt.Println(c11, c22, c33) // prints out 0, 1, 2

  const (
    North = iota //default 0
    East
    South
    West
  )
  fmt.Println(North, East, South, West) // prints 0, 1, 2, 3

  const (
    a = iota *2 // 0
    b          // 2
    c          // 4
  )
  fmt.Println(a, b, c) // 0, 2, 4

  const (
    a = (iota *2) + 1
    b
    c
  )
  fmt.Println(a, b, c) // 1, 3, 5

  //X = -2, y = -4, z = -5
  const (
    x = -(iota +2) // -2
    _              // -3
    y              // -4
    z              // -5
  )
  fmt.Println(x, y, z) // -2, -4, -5
}
```

## Go Data Types - Part 1

