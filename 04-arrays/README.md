# Arrays and slices (Go Fundamentals 04)

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
> It's not type safe though, so the code will compile even if you compare to different typed values (like comparing a string to a slice)

To assign a value to an index inside a slice the `=` operator is used.

`make` lets you create a slice with a starting capacity you pass in.
This implies that slices have a set capacity.
Adding a 10th element to a slice that has a capacity of 2 will result in a runtime error.
Using the `append` function, which takes a slice and new value,
and returns a new slice with all the items in it, is better.

The "tail" a collection is all items in the collection except the first one (the "head").

To slice a slice (yes), the syntax is `slice[low:high]`. Omitting the value on one of the sides of the `:` captures everything on that side of it.
e.g.
```go
numbers[1:] // this says 'take from 1 to the end'
```

The `len` function returns the length of the array or slice passed into it.
