# Hello, World (Go Fundamentals 01)

Functions are defined using the `func` keyword.
Packages are imported using the `import "package-name"` syntax, e.g. `import "fmt"`
Note: the `fmt` package includes the `Println` function which lets us print stuff.

## Testing

Note: It's good practice to separate "domain" code from the "side-effects" code (and testing the "side-effects").

Testing is built into Go. To test a Go file named `filename.go`, create a file (in the same directory) named `filename_test.go`.
Import the "testing" library, and you're good to go.
To test `SomeFunction`, create a function `TestSomeFunction` which receives a variable `t` of type `*testing.T` in your test file.
Inside the function, two variables save the expected output and the expected output, and are compared to see if they match. 
e.g.
```go
func TestSomeFunction(t *testing.T) {
	got := SomeFunction() // this will evalutate to the actual value that the function returns
	want := "This is the expected value that SomeFunction should return"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
```
Note: `%q` is a formatting string to print strings.

Also, remember to initialize a module in your working directory before doing anything.
Your tests won't work otherwise. And... just do it.

### This is how the testing cycle goes:
- Write a test.
- Make the compiler pass.
- Run the test, see that it fails and check the error message is meaningful
- Write enough code to make the test pass.
-	Refactor.

## Basic Constructs

### Conditional statements
`if` statements are like the usual if statements.
```go
if condition {
	do this
}
```

When you have lots of `if` statements checking a particular value it is common to use a `switch` statement instead.
```go
switch valueToCheck {
		case thisValue:
			// do this
		case anotherValue:
			// or do this
		default:
			// if all else fails, do this
}
```

### Functions
Functions are defined like this 
```go
func SomeFunc(parameter parameter-type) return-type {
	// function body
}
```

Variables are defined like so-
```go
var normalVariable = "This is how you define variable the normal way, using the var keyword. The type if inferred here."

var anotherNormalVariable int // variables declared without a corresponding initialization are zero-valued. e.g. 0 for an int.

shorthandVariable := "This is short hand for declaring and initializing a variable." // This syntax is only available inside functions

const someConstant = "This is a constant value."
```

