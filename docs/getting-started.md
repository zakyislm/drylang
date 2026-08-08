# Getting Started

## Installation

### Prerequisites

- [Go](https://go.dev/dl/) 1.21 or later

### Build from Source

```bash
git clone https://github.com/user/drylang.git
cd drylang
go build -o y.exe .
```

This produces the `y` binary — the dryLang compiler and runtime.

### Verify Installation

```bash
y --help
```

## Your First Program

Create a file called `hello.y`:

```
pt "Hello, World!"
```

Run it:

```bash
y hello.y
```

Output:

```
Hello, World!
```

## A More Complete Example

Create `intro.y`:

```
.Welcome to dryLang!.

. Variables — no keyword needed .
name "Zaky"
age 17
height 175,5

. Print with string interpolation .
pt "Name: ${name}"
pt "Age: ${age}"
pt "Height: ${height} cm"

. Constants .
cns pi 3,14
MAXLIFE 5

. Array .
colors ["red", "green", "blue"]
pt colors
pt "Count: ${len(colors)}"

. Function .
fn greet(who) {
  rev "Hello, " + who + "!"
}

pt greet("World")

. Loop .
lp 3 {
  pt "Iteration ${i}"
}

. Conditional .
score 85

if score >= 90 {
  pt "Grade: A"
} elif score >= 80 {
  pt "Grade: B"
} el {
  pt "Grade: C"
}
```

Run:

```bash
y intro.y
```

## File Extension

All dryLang source files use the `.y` extension.

## Next Steps

- [Variables & Constants](variables.md) — how to store data
- [Data Types](types.md) — what types exist
- [Functions](functions.md) — how to write functions
