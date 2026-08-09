# Modules

## Importing Modules

Use the `use` keyword to import a dryLang module:

```rust
use "helpers"
```

This loads `helpers.y` from the same directory.

### Path Resolution

| Statement | File Loaded |
|-----------|-------------|
| `use "helpers"` | `./helpers.y` |
| `use "utils/math"` | `./utils/math.y` |
| `use "lib/string"` | `./lib/string.y` |

The `.y` extension is added automatically.

## Exporting

By default, **all** top-level variables and functions in a module are exported (accessible to importers).

Use `pv` to mark items as private:

```rust
// math.y

// Exported — accessible by importers
fn add(a, b) { rev a + b }
fn sub(a, b) { rev a - b }
PI 3.14

// Private — NOT accessible by importers
pv fn internal_calc(x) { rev x * x }
pv secret "abc123"
```

## Multi-File Programs

### Running Multiple Files

You can run multiple files together:

```bash
y file1.y,file2.y
```

Or run all `.y` files in a directory:

```bash
y myfolder
y all
```

Files are concatenated in order, so declarations in earlier files are available in later files.

## Importing via URL

You can import scripts directly from the internet by providing a full URL:

```rust
use "https://example.com/math.y"

pt add(2, 3)
```

The compiler fetches the file, parses its AST, and prepends its statements into your program.

## GitHub Shorthand

dryLang natively supports importing scripts directly from GitHub repositories:

```rust
use "github.com/user/repo"
```

The compiler automatically translates this to:
`https://raw.githubusercontent.com/user/repo/main/idx.y`

It looks for `idx.y` as the default entry point of the remote module.

You can also specify a path within the repo:

```rust
use "github.com/user/repo/lib/utils.y"
```

## Cycle Prevention

`use` implements strict cycle prevention by caching the absolute paths/URLs of all visited files. If a module has already been loaded, it is skipped — preventing infinite import loops.

## Duplicate Prevention

If two modules import the same file, dryLang only loads it once. The resolved path is cached globally across the entire compilation.

---

Prev : [Error Handling](error-handling.md) | Next : [Built-ins](builtins.md)
<!-- W3Schools-like web docs integration placeholder -->
