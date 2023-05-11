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
