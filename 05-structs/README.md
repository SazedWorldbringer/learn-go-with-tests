# Structs, methods & interfaces (Go Fundamentals 05)

`float64` is the type used for floating point numbers, like `123.45`.
The formatting string for floating point numbers is `f`,
and you can describe how many decimal places you want to print with it using `.n`,
where n is the desired decimal places. e.g.
```go
var decimalNumber float64 = 40.0

// this prints the decimalNumber with two decimal places
fmt.Println("decimalNumber is %.2f", decimalNumber)
```

You can define your own types using a `struct`.
A `struct` is just a named collection of fields where you can store data.
```go
type Rectangle struct {
	Width float64
	Height float64
}
```

Methods are functions with receivers. Methods are called by invoking them on an instance of a particular type.
Functions can be called anywhere you like, such as `Area(rectangle)`. Methods are called on 'things'.
Declaration for a method is almost similar to that of a function.
```go
func (receiverName ReceiverType) MethodName(args) {
	// method body
}
```
When the method is called on a variable of the `ReceiverType`,
a reference to it's data is received via the `receiverName` variable.
It's a convention in Go to have the receiver variable be the first letter of the type.

Interfaces, like structs, are a named collection of methods for objects.
```go
type Shape interface {
	Area() float64
}
```

`%g` is a format string to print out more precise floating point numbers.

Decoupling. The approach using interfaces to declare only what you need, is very important in software design.
In the code for this chapter, the helper is *decoupled* from the concrete types
and only has to deal with the methods it needs for it's job.

Table driven tests are a great fit for testing various implementations of an interface.
Each table is a complete test case with an input and expected output.
If you ever find yourself copy pasting when writing tests,
think about whether refactoring into a table-driven test or pulling the copied 
code out into a helper function might be a better option.

When using Table driven tests, it's helpful to write the test tables in `field: value` format for better readability.

