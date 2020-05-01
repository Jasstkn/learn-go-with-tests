Lesson 2 | Integers
===

> Go source files can only have one package per directory, make sure that your files are organised separately.

`%d` - is for integers format inside `Errorf`, `Printf` and etc.

Also note that we are no longer using the main package, instead we've defined a package named integers, as the name 
suggests this will group functions for working with integers such as Add.

When you have more than one argument of the same type you can shorten it to `(x, y int)`.

#### Examples
Go examples are executed just like tests so you can be confident examples reflect what the code actually does.
Examples are compiled (and optionally executed) as part of a package's test suite.

```go
func ExampleAdd() {
    sum := Add(1, 5)
    fmt.Println(sum)
    // Output: 6
}
```
Verify that it works
```shell script
go test -v
=== RUN   TestAdder
--- PASS: TestAdder (0.00s)
=== RUN   ExampleAdd
--- PASS: ExampleAdd (0.00s)
```
> Please note that the example function will not be executed if you remove the comment "//Output: 6".
> Although the function will be compiled, it won't be executed.