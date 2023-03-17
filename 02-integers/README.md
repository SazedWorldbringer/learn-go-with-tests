# Integers (Go Fundamentals 02)

Integers work the usual way in Go. Addition, subtraction and stuff.

Format string for integers is `%d`.

A note on testing: Write the test first, then write just enough code to satisfy the compiler.
This is to see if the tests are failing for the correct reasons.

When a function has more than one argument of the same type, say integers-
```go
func Add(x int, y int) {} // this can be shortened to

func Add(x, y int) // this
```

Comments added before function declarations show up in the Go Doc.

You can add `Examples` to your code. Go examples are executed just like tests, so they reflect what the code actually does.
Examples are compiled and executed as a part of the package's test suite.
Go examples are written in the `_test.go` files. Example example-
```go
func ExampleAdd() {
		sum := Add(1, 5)
		fmt.Println(sum)
		// Output: 6
}
```
The example functions are executed because the comment saying `Output: x` is added.
The function is compiled without the comment, but it isn't executed.
These examples also show up in the documentation in `godoc`, making your code more accessible.
