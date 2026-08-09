# Comments

dryLang uses C-style comments with `//` for single-line and `/* */` for multi-line.

## Single-Line Comments

```rust
// This is a comment
pt "Hello"  // inline comment
```

Everything after `//` until the end of the line is ignored.

## Multi-Line Comments

```rust
/*
This is a multi-line comment.
It can span as many lines as you need.
*/
```

## Inline Comments

```rust
name "Zaky"    // assign name
age 17         // user age
PI 3.14        // constant
```

## Nested Content in Comments

Comments can contain code-like text, URLs, and any punctuation:

```rust
// This function calculates the sum of two numbers
fn add(a, b) {
    rev a + b
}

/*
TODO: Add support for http://example.com
This will be added in v2.0
See: https://github.com/user/repo
*/
```

## Comment Best Practices

```rust
// Good: explains WHY
// Skip negative values because they indicate errors
lp len(data) {
    if data[i] < 0 { con }
    pt data[i]
}

// Bad: explains WHAT (obvious from code)
// Loop through data and print
lp len(data) {
    pt data[i]
}
```

---

Prev : [Database](database.md) | Next : [CLI](cli.md)
<!-- W3Schools-like web docs integration placeholder -->
