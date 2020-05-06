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

### nil
`nil` is synonymous with `null` from other programming languages. Errors can be `nil` because the return type of 
`Withdraw` will be `error`, which is an interface. If you see a function that takes arguments or returns values that are
interfaces, they can be nilable.

Like `null` if you try to access a value that is `nil` it will throw a **runtime panic**.

### errors
```go
errors.New("oh no")
```

## Summary
**Pointers:**
- Go copies values when you pass them to functions/methods so if you're writing a function that needs to mutate state 
you'll need it to take a pointer to the thing you want to change.
- The fact that Go takes a copy of values is useful a lot of the time but sometimes you won't want your system to make 
a copy of something, in which case you need to pass a reference. 

**nil:**
- Pointers can be nil
- When a function returns a pointer to something, you need to make sure you check if it's nil or you might raise a runtime 
exception, the compiler won't help you here
- Useful for when you want to describe a value that could be missing

**Errors:**
- Errors are they way to signify failure when calling a function/method
- By listening to our tests we concluded that checking for a string in an error would result in a flaky test.
So we refactored to use a meaningful value instead and this resulted an easier to test code and concluded this would be 
easier for users of our API too.

**Create new types from existing one:**
- Useful for adding more domain specific meaning to values
- Can let you implement interfaces