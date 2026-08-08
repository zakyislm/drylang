# Data Types

dryLang is dynamically typed. Variables can hold any type, and the type is determined at runtime.

## Number

Numbers represent both integers and floating-point values. Internally, all numbers are stored as 64-bit floats.

```
x 42          .integer.
y -17         .negative.
z 3,14        .float (comma = decimal point).
big 1000000   .large number.
```

### Float Decimal Rule

dryLang uses **comma** (`,`) as the decimal separator — the dot (`.`) is reserved for comments.

| Syntax | Meaning |
|--------|---------|
| `89,5` | Float `89.5` (no space after comma) |
| `89, 5` | Two separate values (space after comma = separator) |
| `fn(a, b)` | Two parameters (comma followed by space) |

```
pi 3,14          .float 3.14.
coords [45,5, 90,2]   .array of two floats: 45.5 and 90.2 — no! this won't work.
coords [45,5, 90,2]   .comma-space = separator, so this is [45.5, 90.2] ✓.
```

## String

Strings are sequences of characters enclosed in double or single quotes.

```
greeting "Hello, World!"
name 'Zaky'
```

Both `"` and `'` are equivalent — both support interpolation and escape sequences.

### Escape Sequences

| Escape | Character |
|--------|-----------|
| `\n` | Newline |
| `\t` | Tab |
| `\\` | Backslash |
| `\"` | Double quote |
| `\'` | Single quote |
| `\$` | Literal `$` (prevent interpolation) |
| `\0` | Null character |

```
pt "Line 1\nLine 2"
pt "Tab\there"
pt "Price: \$100"
```

### String Interpolation

Use `${}` inside strings to embed expressions:

```
name "Zaky"
age 17

pt "Hello, ${name}!"
pt "${name} is ${age} years old"
pt "2 + 3 = ${2 + 3}"
pt "First color: ${colors[0]}"
pt "Config host: ${config.host}"
```

### Raw Strings

Raw strings are enclosed in backticks (`` ` ``). No interpolation or escape processing occurs:

```
pattern `\d+\.\w+`
path `C:\Users\zaky\Desktop`
template `Hello ${name}`    .literal "${name}", not interpolated.
```

Raw strings are ideal for regex patterns and file paths.

## Boolean

dryLang uses single-character boolean values:

| Value | Meaning |
|-------|---------|
| `t` | true |
| `f` | false |

```
active t
deleted f

if active {
  pt "Active!"
}
```

## Unknown

`unknown` is a special value representing an uninitialized or undetermined state.

```
? status         .declares unknown bool.
pt status        .prints "unknown".

. Functions without rev return unknown .
fn do_stuff() {
  pt "doing stuff"
}
x = do_stuff()
pt x             .prints "unknown".
```

### Truthiness

| Value | Truthy? |
|-------|---------|
| `t` | ✅ |
| `f` | ❌ |
| `0` | ❌ |
| `""` (empty string) | ❌ |
| `[]` (empty array) | ❌ |
| `{}` (empty map) | ❌ |
| `unknown` | ❌ |
| Any other number | ✅ |
| Any non-empty string | ✅ |
| Any non-empty collection | ✅ |

## Type Checking

Use the `get()` built-in to check a value's type at runtime:

```
pt get(42)        .prints "number".
pt get("hello")   .prints "string".
pt get(t)         .prints "bool".
pt get([1,2,3])   .prints "array".
pt get({"a": 1})  .prints "map".
pt get(greet)     .prints "fn".
```

`get()` returns one of: `"number"`, `"string"`, `"bool"`, `"array"`, `"map"`, `"fn"`, `"unknown"`.
