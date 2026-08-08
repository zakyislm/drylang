# Control Flow

## If Statement

```
if condition {
  .body.
}
```

The condition can be any expression. It is evaluated for truthiness (see [Data Types — Truthiness](types.md#truthiness)).

```
age 18

if age >= 18 {
  pt "You are an adult"
}
```

### If-Else

Use `el` for the else block:

```
score 45

if score >= 60 {
  pt "Pass"
} el {
  pt "Fail"
}
```

### If-ElIf-Else

Use `elif` for additional conditions:

```
score 85

if score >= 90 {
  pt "A"
} elif score >= 80 {
  pt "B"
} elif score >= 70 {
  pt "C"
} elif score >= 60 {
  pt "D"
} el {
  pt "F"
}
```

You can chain as many `elif` blocks as needed.

### Single Line

Blocks can be written on a single line:

```
if x = 5 { pt "five" }
if active { pt "on" } el { pt "off" }
```

### No Indentation Required

dryLang has **free indentation**. Curly braces `{}` define blocks, not whitespace.

```
.These are all valid:.
if x = 5 {
pt "five"
}

if x = 5 {
  pt "five"
}

if x = 5 { pt "five" }
```

### Equality Comparison

In an expression context (inside `if`), the `=` operator is **equality**, not assignment:

```
name "Zaky"

if name = "Zaky" {
  pt "Hello, Zaky!"
}

if name != "Andi" {
  pt "You are not Andi"
}
```

## Switch (on)

The `on` statement matches a value against multiple cases:

```
on(value) {
  case1 { .body. }
  case2 { .body. }
}
```

### Example

```
day 3

on(day) {
  1 { pt "Monday" }
  2 { pt "Tuesday" }
  3 { pt "Wednesday" }
  4 { pt "Thursday" }
  5 { pt "Friday" }
}
```

### String Cases

```
cmd "start"

on(cmd) {
  "start" { pt "Starting..." }
  "stop" { pt "Stopping..." }
  "restart" {
    pt "Restarting..."
  }
}
```

### Behavior

- **No fall-through** — each case is independent. Only the matching case body executes.
- Cases are checked in order from top to bottom.
- If no case matches, nothing happens (no default/else case).
