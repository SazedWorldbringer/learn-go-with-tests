# Pointers & errors (Go Fundamentals 06)

Symbols starting with lowercase are private outside of the package they're defined in.

When calling a function or method, the arguments are copied,
which means they're passed by value.
Place an `&` before the symbol to get its pointer(memory address).
The formatting string for a pointer is`v`.
Adding an `*` before a symbol, you can pass it's pointer,
and then change the underlying values the address points to.

```go
func (w *Wallet) Deposit(amount int) {
    w.balance += amount
}
```

Types are defined like this
```go
// type NewType OriginalType
type Bitcoin int

// to make Bitcoins you use this syntax
Bitcoin(999)
```

A Stringer helps you format how your types are printed when used with the `%s` format strings.
```go
func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
```
The syntax for creating a method on a type declaration is the same as it is on a struct.

The `var` keyword allows us to define values global to the package.
