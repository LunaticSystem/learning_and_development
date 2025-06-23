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

