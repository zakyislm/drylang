# Functions

## Declaration

```rust
fn name(parameters) {
  // body
}
```

```rust
fn greet(name) {
  pt "Hello, " + name + "!"
}

greet("World")
```

## Return Values

Use `rev` (revert) to return a value:

```rust
fn add(a, b) {
  rev a + b
}

result add(2, 3)
pt result    // 5
```

### Default Return

If a function has no `rev` statement, it returns `unknown`:

```rust
fn do_stuff() {
  pt "Working..."
}

x do_stuff()
pt x           // prints "unknown"
```

### Early Return

`rev` can be used for early exit:

```rust
fn divide(a, b) {
  if b = 0 {
    rev "Error: division by zero"
  }
  rev a / b
}
```

## Parameters

Parameters are separated by commas:

```rust
fn calculate(x, y, op) {
  if op = "add" { rev x + y }
  if op = "sub" { rev x - y }
  rev 0
}

pt calculate(10, 5, "add")    // 15
pt calculate(10, 5, "sub")    // 5
```

### No Default Parameters

dryLang does not support default parameter values. All parameters must be provided:

```rust
fn greet(name) {
  rev "Hello, " + name
}

greet()         // Error: want 1 args
greet("Zaky")   // OK
```

## Arrow Functions

Anonymous functions use `->`:

```rust
double -> (x) {
  rev x * 2
}

pt double(5)    // 10
```

Arrow functions can be passed as arguments:

```rust
fn apply(value, transformer) {
  rev transformer(value)
}

result apply(5, -> (x) { rev x * 3 })
pt result    // 15
```

## Function Calls

Function calls **always require parentheses** `()`. This is how dryLang distinguishes calls from assignments:

```rust
greet("World")     // function call
name "World"       // variable assignment
greet()            // call with no args
```

## Async Functions

Use `asn fn` to declare an async function (goroutine-based):

```rust
asn fn fetch_data(url) {
  // async operation
  rev "data from " + url
}
```

### Await

Use `awt` to wait for an async function's result:

```rust
result awt fetch_data("https://api.example.com")
pt result
```

> **Note:** Async is simplified in the current version. Full goroutine support is planned for a future release.

## Recursion

Functions can call themselves:

```rust
fn factorial(n) {
  if n <= 1 {
    rev 1
  }
  rev n * factorial(n - 1)
}

pt factorial(5)    // 120
```

## Private Functions

Use `pv fn` to mark a function as private (not exported from a module):

```rust
pv fn internal_calc(x) {
  rev x * x
}
```

## Top-Level `rev`

Using `rev` outside any function exits the program:

```rust
pt "This prints"
rev
pt "This never prints"
```

---

Prev : [Loops](loops.md) | Next : [Collections](collections.md)
<!-- W3Schools-like web docs integration placeholder -->
