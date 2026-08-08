# Errors

Format: `line:col message`

## Compile-Time
| error | meaning |
|-------|---------|
| `close }` | extra `}` |
| `want }` | missing `}` |
| `want )` | missing `)` |
| `want ]` | missing `]` |
| `want fn` | `asn` needs `fn` |
| `want err` | `try` needs `err(e){}` |
| `illegal X` | unexpected token |
| `locked X` | reassigning constant |
| `stray rev` | `rev` outside `fn` |
| `stray done` | `done` outside `lp` |
| `stray con` | `con` outside `lp` |
| `stray awt` | `awt` outside `asn fn` |
| `unused X` | declared, never used |

## Runtime
| error | meaning |
|-------|---------|
| `want number` | arithmetic needs number |
| `want string` | string fn needs string |
| `want array` | array fn needs array |
| `want map` | dot access needs map |
| `want fn` | called non-function |
| `want N args` | wrong arg count |
| `div by 0` | division by zero |
| `bounds N` | index out of range |
| `empty array` | `pop()` on `[]` |
| `unknown X` | undefined variable |
| `read fail` | `r()` file not found |
| `write fail` | `w()` permission denied |
