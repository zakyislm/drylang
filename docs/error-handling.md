# Error Handling

dryLang uses `try`/`err` for error handling.

## Try-Catch

```
try {
  .code that might fail.
} err(e) {
  .handle the error.
  .e = error value (string).
}
```

### Example

```
try {
  pt "Attempting..."
  err "Something went wrong!"
} err(e) {
  pt "Caught: " + e
}

.Output:
  Attempting...
  Caught: Something went wrong!
.
```

## Throwing Errors

Use the `err` keyword (outside of catch clause context) to throw an error:

```
fn divide(a, b) {
  if b = 0 {
    err "division by zero"
  }
  rev a / b
}

try {
  result divide(10, 0)
  pt result
} err(e) {
  pt "Error: " + e
}
.Output: Error: division by zero.
```

## Error Propagation

If an error is thrown and not caught by a `try`/`err`, the program terminates:

```
err "fatal error"
pt "This never prints"
```

Output:

```
1:0 fatal error
```

## Nesting

Try blocks can be nested:

```
try {
  try {
    err "inner error"
  } err(e) {
    pt "Inner caught: " + e
    err "re-thrown: " + e
  }
} err(e) {
  pt "Outer caught: " + e
}

.Output:
  Inner caught: inner error
  Outer caught: re-thrown: inner error
.
```

## Error Values

Error values in dryLang are always strings. When an `err` is thrown, the value becomes available as a string in the catch clause:

```
try {
  err 42    .42 is converted to "42".
} err(e) {
  pt get(e)    .prints "string".
  pt e         .prints "42".
}
```

## Pattern: Validate or Die

```
fn require(value, msg) {
  if value = unknown {
    err msg
  }
  rev value
}

try {
  config require(unknown, "config missing")
} err(e) {
  pt e
}
```

## Pattern: Safe Call

```
fn safe_divide(a, b) {
  try {
    if b = 0 {
      err "div by 0"
    }
    rev a / b
  } err(e) {
    rev 0
  }
}

pt safe_divide(10, 2)    .5.
pt safe_divide(10, 0)    .0.
```
