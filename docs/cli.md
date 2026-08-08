# CLI Reference

The dryLang CLI is a single binary: `y`.

## Usage

```
y <target>
```

## Targets

### Single File

```bash
y main.y
y script.y
y path/to/file.y
```

### Multiple Files

Comma-separated (no spaces):

```bash
y file1.y,file2.y,file3.y
```

Files are loaded and concatenated in order. Declarations in earlier files are available in later files.

### Directory

Run all `.y` files in a directory:

```bash
y myfolder
y src/scripts
```

### All Files

Run all `.y` files in the current directory:

```bash
y all
```

## Execution Pipeline

When you run `y`, the following pipeline executes:

```
Source (.y) → Lexer → Tokens → Parser → AST → Compiler → Bytecode → VM → Output
```

1. **Lexer** — tokenizes the source into tokens
2. **Parser** — builds an Abstract Syntax Tree (AST) using Pratt parsing
3. **Compiler** — compiles AST to bytecode instructions
4. **VM** — executes bytecode on a stack-based virtual machine

## Exit Codes

| Code | Meaning |
|------|---------|
| `0` | Success |
| `1` | Error (compile-time or runtime) |

## Error Output

Errors are printed to stderr in the format:

```
line:col message
```

Examples:

```
3:12 close }
5:1 unknown x
1:0 read main.y
```

Errors are intentionally minimal — read the docs for details.
