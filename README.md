# dryLang

> **WDYNIWDYG** — *What Do You Need Is What You Get.*

dryLang is a minimalist, dynamically-typed programming language designed for maximum expressiveness with minimum syntax. Inspired by Go's simplicity, Python's readability, and JavaScript's string interpolation — then compressed to its absolute essence.

## Quick Example

```rs
.Hello World in dryLang.
name "Zaky"
age 17

pt "Hello, ${name}! You are ${age} years old."

fn greet(who) {
  rev "Hi, " + who + "!"
}

pt greet("World")
```

## Features

- **Ultra-short keywords** — `pt` (print), `fn` (function), `rev` (return), `lp` (loop), `el` (else)
- **No variable keyword** — `name "Zaky"` or `name = "Zaky"`
- **Comma decimal** — `pi 3,14` (comma = decimal, dot = comment)
- **Dot comments** — `.this is a comment.`
- **27 built-in functions** — zero imports needed
- **Structs without keywords** — `user { name age }`
- **String interpolation** — `"Hello ${name}"`
- **Raw strings** — `` `\d+\.txt` ``
- **Stack-based VM** — compiled to bytecode, fast execution
- **Strict unused check** — no dead code allowed

## Installation

### Build from Source

```bash
git clone https://github.com/user/drylang.git
cd drylang
go build -o y.exe .
```

### Run

```bash
# Single file
y main.y

# Multiple files
y file1.y,file2.y

# Entire directory
y myfolder

# All .y files in current directory
y all
```

## Language at a Glance

| Feature | dryLang | Other Languages |
|---------|---------|-----------------|
| Print | `pt "hi"` | `print("hi")` / `fmt.Println("hi")` |
| Variable | `name "Zaky"` | `name = "Zaky"` / `var name = "Zaky"` |
| Constant | `MAXLIFE 5` | `const MAXLIFE = 5` |
| Function | `fn add(a, b) { rev a + b }` | `def add(a, b): return a + b` |
| Loop | `lp 5 { pt i }` | `for i in range(5): print(i)` |
| If | `if x = 5 { pt "yes" }` | `if x == 5: print("yes")` |
| Comment | `.this is a comment.` | `// this is a comment` |
| Boolean | `t` / `f` | `true` / `false` |
| Float | `3,14` | `3.14` |

## Documentation

- 📖 [Full Documentation](docs/) — comprehensive language reference
- ⚡ [DryDocs](drydocs/) — ultra-concise reference (dryLang style)

## File Extension

dryLang source files use the `.y` extension.

## Keywords

```
cns  t  f  fn  rev  if  elif  el  on  lp  done  con
asn  awt  try  err  pv  use  unknown  pt  in
```

## Built-in Functions

| Function | Description | Example |
|----------|-------------|---------|
| `pt` | Print | `pt "hello"` |
| `in` | Input | `in("name?")` |
| `r` | Read file | `r("f.txt")` |
| `w` | Write file | `w("f.txt", data)` |
| `len` | Length | `len(arr)` |
| `get` | Type check | `get(x)` |
| `add` | Add to array | `add(arr, 5)` |
| `num` | To number | `num("42")` |
| `str` | To string | `str(42)` |
| `abs` | Absolute | `abs(-5)` |
| `min` | Minimum | `min(a, b)` |
| `max` | Maximum | `max(a, b)` |
| `rnd` | Round | `rnd(3,7)` |
| `cap` | Uppercase | `cap("hi")` |
| `low` | Lowercase | `low("HI")` |
| `trm` | Trim | `trm(" hi ")` |
| `spl` | Split | `spl("a,b", ",")` |
| `j` | Join | `j(arr, ",")` |
| `mod` | Replace | `mod(s, "a", "b")` |
| `has` | Contains | `has(s, "sub")` |
| `sort` | Sort | `sort(arr)` |
| `pop` | Pop last | `pop(arr)` |
| `rm` | Remove | `rm(arr, 2)` |
| `key` | Map keys | `key(map)` |
| `val` | Map values | `val(map)` |
| `ran` | Random | `ran()` |
| `q` | Wait/sleep | `q(1000)` |

## Philosophy

dryLang follows a strict design philosophy:

1. **DRY** — Don't Repeat Yourself. Every keyword and function name is as short as possible.
2. **WDYNIWDYG** — What Do You Need Is What You Get. No bloat, no extras.
3. **Zero import basics** — File I/O, math, string ops — all built-in.
4. **Strict** — Unused variables are compile errors, not warnings.
5. **Minimal error messages** — `3:12 close }` — short, with location. Read the docs.

## License

MIT
