# Arrays and slices (Go Fundamentals 05)

Arrays allow you to store multiple elements of the same type in a variable in a particular order.
They have a fixed capacity, which you define when you declare the variable.
We can initialize arrays in two ways:
- [N]type{value1, value2, ..., valueN} e.g. `numbers := [5]int{1, 2, 3, 4, 5}`
- [...]type{value1, value2, ..., valueN} e.g. `numbers := [...]int{1, 2, 3, 4, 5}`

> Side note: Package main only contains integration of other packages and not unit-testable code.

We can access elements in the arrays the usual way, using square brackets.
Range is a helpful construct when iterating through an array. It returns two values - the index and the value
```go
	for i, el := range arr {
	}

	// the index can be ignored like so,
	for _, el := range arr {
	}
```

The size is of an array is encoded into it's type, so,
if you try to pass an [4]int array into a function that expects [5]int, it won't compile.
They have different types.

Slices are similar to arrays, but more powerful and flexible.
Slices, just like arrays, are used to store multiple elements of the same time in a single variable.
But the length of the slice can grow and shrink as you see fit.

`reflect.DeepEqual` is a nice function to use when comparing two slices
> It's not type safe though, so the code will compile if you compare to different typed values (like comparing a string to a slice)
