# Error Messages

dryLang errors are intentionally minimal. Format:

```
line:col message
```

## Compile-Time Errors

These errors occur during parsing or compilation, before the program runs.

| Error | Meaning | Fix |
|-------|---------|-----|
| `close }` | Unmatched closing brace | Check for missing opening `{` |
| `want }` | Expected closing brace | Check for missing `}` at end of block |
| `want )` | Expected closing paren | Check for missing `)` at end of expression |
| `want ]` | Expected closing bracket | Check for missing `]` at end of array |
| `want fn` | Expected `fn` after `asn` | `asn` must be followed by `fn` |
| `want err` | Expected `err` after `try` block | Every `try {}` needs `err(e) {}` |
| `want IDENT` | Expected identifier | Check for missing variable or function name |
| `want STRING` | Expected string literal | `use` requires a string argument |
| `illegal X` | Unexpected token | Check syntax around the indicated position |
| `locked X` | Constant reassignment | Cannot reassign a `cns` or ALL CAPS variable |
| `stray rev` | `rev` outside function | `rev` can only be used inside `fn` blocks |
| `stray done` | `done` outside loop | `done` can only be used inside `lp` blocks |
| `stray con` | `con` outside loop | `con` can only be used inside `lp` blocks |
| `stray awt` | `awt` outside async fn | `awt` can only be used inside `asn fn` blocks |
| `unused X` | Declared but never used | Remove or use the variable |
| `bad number` | Invalid number literal | Check number format |

## Runtime Errors

These errors occur during program execution.

| Error | Meaning | Fix |
|-------|---------|-----|
| `want number` | Expected numeric operand | Arithmetic ops require numbers |
| `want string` | Expected string argument | String functions require strings |
| `want array` | Expected array argument | Array functions require arrays |
| `want map` | Expected map argument | Map/dot functions require maps |
| `want fn` | Called a non-function | Variable is not a function |
| `want N args` | Wrong argument count | Function expects N parameters |
| `want array\|map` | Expected indexable value | `[]` only works on arrays and maps |
| `want string\|array\|map` | Expected lengthed value | `len()` requires string, array, or map |
| `div by 0` | Division by zero | Check divisor before dividing |
| `bounds N` | Array index out of range | Index N is beyond array length |
| `empty array` | Operation on empty array | `pop()` on empty array |
| `unknown X` | Undefined variable | Variable X was never assigned |
| `read fail` | File read error | File doesn't exist or no permission |
| `write fail` | File write error | Path invalid or no permission |

## Reading Error Messages

```
3:12 close }
│ │  └── message
│ └── column (character position)
└── line number
```

Example program with error:

```
fn greet(name) {
  pt "Hello, " + name
}
```

If you add a stray `}`:

```
fn greet(name) {
  pt "Hello, " + name
}
}
```

Error: `4:1 illegal }`

## Philosophy

Error messages in dryLang are deliberately short:

- **No suggestions** — read the docs
- **No verbose explanations** — the error name + position is enough
- **No colors** — plain text, pipe-friendly
- **Always includes position** — `line:col` for every error
