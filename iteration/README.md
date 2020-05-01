Lesson 3 | Iteration
===

In Go there are no `while`, `do`, `until` keywords, you can only use `for`.

`+=` the Add AND assignment operator, adds the right operand to the left operand and assigns the result to left operand. 
It works with other types like integers.

#### Benchmarks
```go
func BenchmarkRepeat(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Repeat("a")
    }
}
```
The `testing.B` gives you access to the cryptically named `b.N`. 
When the benchmark code is executed, it runs `b.N` times and measures how long it takes.
To run the benchmarks do `go test -bench=.`
```shell script
goos: darwin
goarch: amd64
pkg: github.com/jasstkn/learn-go-with-tests/iteration
BenchmarkRepeat-8        8350526               136 ns/op
PASS
ok      github.com/jasstkn/learn-go-with-tests/iteration        1.364s
```