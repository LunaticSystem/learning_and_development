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
