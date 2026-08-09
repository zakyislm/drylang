# Operators

## Assignment

| Operator | Usage | Description |
|----------|-------|-------------|
| `=` | `x = 5` | Assign value (also used for equality in expressions) |
| *(space)* | `x 5` | Assign value without `=` |

In a **statement context** (start of line, after identifier), `=` is assignment. In an **expression context** (inside `if`, etc.), `=` is equality comparison.

```rust
x = 5              // assignment
if x = 5 { ... }   // equality comparison
```

## Arithmetic

| Operator | Example | Result |
|----------|---------|--------|
| `+` | `5 + 3` | `8` |
| `-` | `5 - 3` | `2` |
| `*` | `5 * 3` | `15` |
| `/` | `10 / 3` | `3.3333...` |
| `%` | `10 % 3` | `1` |

### String Concatenation

The `+` operator concatenates strings:

```rust
greeting "Hello" + ", " + "World!"
pt greeting    // prints "Hello, World!"
```

If either operand is a string, the other is converted to string automatically:

```rust
pt "Age: " + 17    // prints "Age: 17"
pt "Score: " + 99,5    // prints "Score: 99.5"
```

### Unary Minus

```rust
x -5
y -(3 + 2)
```

## Comparison

| Operator | Meaning | Example |
|----------|---------|---------|
| `=` | Equal | `x = 5` |
| `!=` | Not equal | `x != 5` |
| `<` | Less than | `x < 10` |
| `>` | Greater than | `x > 10` |
| `<=` | Less or equal | `x <= 10` |
| `>=` | Greater or equal | `x >= 10` |

All comparison operators return `t` or `f`.

```rust
pt 5 = 5       // prints "t"
pt 5 != 3      // prints "t"
pt 10 < 20     // prints "t"
pt 10 >= 10    // prints "t"
```

## Logical

| Operator | Meaning | Example |
|----------|---------|---------|
| `&` | AND | `a & b` |
| `\|` | OR | `a \| b` |
| `!` | NOT | `!a` |

```rust
online t
admin t

if online & admin {
  pt "Welcome, admin!"
}

if !online {
  pt "You are offline"
}

active t
premium f
if active | premium {
  pt "Access granted"
}
```

## Operator Precedence

From lowest to highest:

| Precedence | Operators |
|------------|-----------|
| 1 (lowest) | `\|` |
| 2 | `&` |
| 3 | `=` `!=` |
| 4 | `<` `>` `<=` `>=` |
| 5 | `+` `-` |
| 6 | `*` `/` `%` |
| 7 | `!` `-` (unary) |
| 8 (highest) | `()` `[]` `.` (call, index, access) |

Use parentheses to override precedence:

```rust
result (a + b) * c
```

---

Prev : [Types](types.md) | Next : [Strings](strings.md)
<!-- W3Schools-like web docs integration placeholder -->
