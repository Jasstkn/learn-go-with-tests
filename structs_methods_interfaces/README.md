Lesson 5 | Structs, methods & interfaces
===

> The `f` is for our float64 and the `.2` means print 2 decimal places.

### Struct
A **struct** is just a named collection of fields where you can store data.

Declare a struct like this:
```go
type Rectangle struct {
    Width float64
    Height float64
}
```
You can access the fields of a struct with the syntax of `myStruct.field`.

> With `'g'` we get a complete decimal number in the error message

### What are methods?
When we call `t.Errorf` we are calling the method `Errorf` on the instance of our `t (testing.T)`.
A method is a function with a receiver. 
A method declaration binds an identifier, the method name, to a method, 
and associates the method with the receiver's base type.

Methods are very similar to functions, but they are called by invoking them on an instance of a particular type. 
Where you can just call functions wherever you like, such as Area(rectangle) you can only call methods on "things".

The syntax for declaring methods is almost the same as functions and that's because they're so similar. 
The only difference is the syntax of the method receiver `func (receiverName ReceiverType) MethodName(args)`.

When your method is called on a variable of that type, you get your reference to its data via the `receiverName` variable. 
In many other programming languages this is done implicitly and you access the receiver via `this`.

It is a convention in Go to have the receiver variable be the first letter of the type.
```go
r Rectangle
```

### Interfaces
**Interfaces** are a very powerful concept in statically typed languages like Go because they allow you to make 
functions that can be used with different types and create highly-decoupled code whilst still maintaining type-safety.
```go
type Shape interface {
    Area() float64
}
```
We're creating a new `type` just like we did with `Rectangle` and `Circle` but this time it is an `interface` rather than a struct.

In our case:
- `Rectangle` has a method called `Area` that returns a `float64` so it satisfies the `Shape` interface
- `Circle` has a method called `Area` that returns a `float64` so it satisfies the `Shape` interface
- `string` does not have such a method, so it doesn't satisfy the interface
- etc.

In Go interface resolution is implicit. If the type you pass in matches what the interface is asking for, it will compile.

