# Maps (Go Fundamentals 07)

We declare Maps just like we do arrays, except that it requires two types.
First is the key type(written inside []) and second is the value type(after the []).
```go
testMap := map[string]string("key": "value")
```
You can access a value, _mapped_ to a particular `key` using syntax pretty similar to Arrays.
```go
testMap["key"]
// returns "value"
```

When looking up values from a map, it can return two values.
First is the actual `value` corresponding the `key` passed,
and second is a boolean which indicates if the key was found successfully.
```go
value, ok := testMap["key"] 
// returns value = "value" and ok = true
```

Adding to a map is just like arrays. Just specify the key and set it to a value.
```go
testMap["newKey"] = "Value associated to newKey in testMap"
```

You don't have to pass maps as an address in order to modify them.
So, when you pass a map to a function/method, you are copying it, but just the pointer part.

A nil map behaves like an empty map when reading, but an attempt to write to it will cause a runtime panic.
Never initialize an empty map variable
```go
var m map[string]string
```
Instead, initialize an empty map like this-
```go
var dictionary = map[string]string{}
// OR with the `make` keyword
var dictionary = make(map[string]string)
```

Make sure to return precise error messages for each individual cases.
```go
ErrNotFound         = "could not find the word you were looking for"
ErrWordExists       = "cannot add word because it already exists"
ErrWordDoesNotExist = "cannot update word because it does not exist"
```
Having specific errors gives you more information about what went wrong. eg -
> You can redirect the user when ErrNotFound is encountered,
but display an error message when ErrWordDoesNotExist is encountered.

Go has a `delete` function to delete keys from maps.
It takes two arguments, first a map, and second a key.
```go
delete(testMap, "key")
```
