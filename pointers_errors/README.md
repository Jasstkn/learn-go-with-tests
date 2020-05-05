Lesson 6 | Pointers & errors
===

> In Go if a symbol (so variables, types, functions et al) starts with a lowercase symbol then it is private outside 
>the package it's defined in.

In Go, when you **call a function or a method the arguments are copied.**

### Pointers
Pointers let us point to some values and then let us change them. 
So rather than taking a copy of the Wallet, we take a pointer to the wallet so we can change it.
```go
// The difference is the receiver type is *Wallet rather than Wallet 
// which you can read as "a pointer to a wallet"
func (w *Wallet) Deposit(amount int) {
    w.balance += amount
}
```

#### Stringer
Stringer is implemented by any value that has a String method, which defines the “native” format for that value. 
The String method is used to print values passed as an operand to any format that accepts a string or to an unformatted 
printer such as Print.
```go
type Stringer interface {
    String() string
}

// Example
package main

import (
	"fmt"
)

// Animal has a Name and an Age to represent an animal.
type Animal struct {
	Name string
	Age  uint
}

// String makes Animal satisfy the Stringer interface.
func (a Animal) String() string {
	return fmt.Sprintf("%v (%d)", a.Name, a.Age)
}

func main() {
	a := Animal{
		Name: "Gopher",
		Age:  2,
	}
	fmt.Println(a)
}
```