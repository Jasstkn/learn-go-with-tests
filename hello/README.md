Lesson 1 | Hello world
===

#### Writing tests
Writing a test is just like writing a function, with a few rules
- It needs to be in a file with a name like `xxx_test.go`
- The test function must start with the word `Test`
- The test function takes one argument only `t *testing.T`

It's enough to know that your `t` of type `*testing.T` is your "hook" into the testing framework 
so you can do things like `t.Fail()` when you want to fail.

#### t.Errorf
We are calling the `Errorf` method on our `t` which will print out a message and fail the test. 
The `f` stands for format which allows us to build a string with values inserted into the placeholder values `%q`. 
When you made the test fail it should be clear how it works.

For tests `%q` is very useful as it wraps your values in double quotes.

#### Go doc
You can launch the docs locally by running `godoc -http :8000`

#### Constants
```go
const englishHelloPrefix = "Hello, "
```

> `t.Helper()` is needed to tell the test suite that this method is a helper. 
> By doing this when it fails the line number reported will be in our function call rather than inside our test helper. 
> This will help other developers track down problems easier. 

#### switch
When you have lots of `if` statements checking a particular value it is common to use a switch statement instead.
We can use `switch` to refactor the code to make it easier to read and more extensible.

```go
switch language {
case french:
    prefix = frenchHelloPrefix
case spanish:
    prefix = spanishHelloPrefix
default:
    prefix = englishHelloPrefix
}
```

- In our function signature we have made a named return value `(prefix string)`.
- This will create a variable called `prefix` in your function.
- It will be assigned the `"zero"` value. This depends on the type, for example ints are `0` and for strings it is `""`.
- You can return whatever it's set to by just calling return rather than return prefix.
- This will display in the Go Doc for your function so it can make the intent of your code clearer.
- `default` in the `switch case` will be branched to if none of the other case statements match.

> The function name starts with a lowercase letter. In Go public functions start with a capital letter and private ones 
> start with a lowercase. We don't want the internals of our algorithm to be exposed to the world, so we made this 
> function private.


