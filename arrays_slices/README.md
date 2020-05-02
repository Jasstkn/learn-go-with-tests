Lesson 4 | Arrays and slices
===

#### Array initialization
- [N]type{value1, value2, ..., valueN} e.g. `numbers := [5]int{1, 2, 3, 4, 5}`
- [...]type{value1, value2, ..., valueN} e.g. `numbers := [...]int{1, 2, 3, 4, 5}`

the `%v` placeholder which is the "default" format, which works well for arrays.

### Package `"fmt"`
#### Printing
The verbs:
- General:
```shell script
%v	the value in a default format
	when printing structs, the plus flag (%+v) adds field names
%#v	a Go-syntax representation of the value
%T	a Go-syntax representation of the type of the value
%%	a literal percent sign; consumes no value
```
- Boolean:
```shell script
%t	the word true or false
```
- Integer:
```shell script
%b	base 2
%c	the character represented by the corresponding Unicode code point
%d	base 10
%o	base 8
%O	base 8 with 0o prefix
%q	a single-quoted character literal safely escaped with Go syntax.
%x	base 16, with lower-case letters for a-f
%X	base 16, with upper-case letters for A-F
%U	Unicode format: U+1234; same as "U+%04X"
```
- Floating-point and complex constituents:
```shell script
%b	decimalless scientific notation with exponent a power of two,
	in the manner of strconv.FormatFloat with the 'b' format,
	e.g. -123456p-78
%e	scientific notation, e.g. -1.234456e+78
%E	scientific notation, e.g. -1.234456E+78
%f	decimal point but no exponent, e.g. 123.456
%F	synonym for %f
%g	%e for large exponents, %f otherwise. Precision is discussed below.
%G	%E for large exponents, %F otherwise
%x	hexadecimal notation (with decimal power of two exponent), e.g. -0x1.23abcp+20
%X	upper-case hexadecimal notation, e.g. -0X1.23ABCP+20
```
- String and slice of bytes (treated equivalently with these verbs):
```shell script
%s	the uninterpreted bytes of the string or slice
%q	a double-quoted string safely escaped with Go syntax
%x	base 16, lower-case, two characters per byte
%X	base 16, upper-case, two characters per byte
```
- Slice:
```shell script
%p	address of 0th element in base 16 notation, with leading 0x
```
- Pointer:
```shell script
%p	base 16 notation, with leading 0x
The %b, %d, %o, %x and %X verbs also work with pointers,
formatting the value exactly as if it were an integer.
```
- The default format for `%v` is:
```shell script
bool:                    %t
int, int8 etc.:          %d
uint, uint8 etc.:        %d, %#x if printed with %#v
float32, complex64, etc: %g
string:                  %s
chan:                    %p
pointer:                 %p
```

### Range
`range` lets you iterate over an array. 
Every time it is called it returns two values, the index and the value. 
We are choosing to ignore the index value by using `_` blank identifier.
```go
for _, number := range numbers {
    sum += number
}
```
### Arrays
An interesting property of arrays is that the size is encoded in its type. 
If you try to pass an `[4]int` into a function that expects `[5]int`, it won't compile. 
They are different types so it's just the same as trying to pass a string into a function that wants an `int`.

### Slices
`make` allows you to create a slice with a starting capacity of the len of the numbersToSum we need to work through.
```go
sums := make([]int, lengthOfNumbers)
```
`append` is a function which takes a slice and a new value, returning a new slice with all the items in it.

### Test coverage
```shell script
go test -cover
PASS
coverage: 100.0% of statements
ok      github.com/jasstkn/learn-go-with-tests/arrays_slices    0.344s
```