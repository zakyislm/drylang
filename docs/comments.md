# Comments

dryLang uses **dot-delimited comments**. The dot (`.`) starts a comment, and the next dot ends it.

## Syntax

```
.This is a comment.
```

The content between two dots is completely ignored by the compiler.

## Inline Comments

```
name "Zaky"    .assign name.
age 17         .user age.
```

## Multi-Line Comments

Comments can span multiple lines:

```
.
This is a multi-line comment.
It can contain anything except a lone dot
that would end it.
.
```

## Nested Content in Comments

Comments can contain code-like text, URLs, and most punctuation. The only restriction is that a standalone dot ends the comment:

```
.This function calculates the sum of two numbers.
fn add(a, b) {
  rev a + b
}

.
TODO: Add support for http://example.com
This will be added in v2.0
.
```

## Why Dots?

- The `.` character is rarely used at the start/end of meaningful text
- Frees up `//`, `#`, `/* */` for other potential uses
- Keeps the syntax unique and minimal
- The comma (`,`) handles decimals, so dots were available for comments

## Comment Best Practices

```
.Good: explains WHY.
.Skip negative values because they indicate errors.
lp len(data) {
  if data[i] < 0 { con }
  pt data[i]
}

.Bad: explains WHAT (obvious from code).
.Loop through data and print.
lp len(data) {
  pt data[i]
}
```
