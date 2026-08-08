# CLI

```
y <target>
```

| target | what |
|--------|------|
| `y main.y` | single file |
| `y a.y,b.y` | multiple files (comma, no space) |
| `y myfolder` | all `.y` in folder |
| `y all` | all `.y` in cwd |

Pipeline: `Source → Lexer → Parser → Compiler → VM`

Exit 0 = ok. Exit 1 = error (stderr).
