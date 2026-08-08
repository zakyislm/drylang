# Modules

## Importing Modules

Use the `use` keyword to import a dryLang module:

```
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

```
.math.y.

.Exported — accessible by importers.
fn add(a, b) { rev a + b }
fn sub(a, b) { rev a - b }
PI 3,14

.Private — NOT accessible by importers.
pv fn internal_calc(x) { rev x * x }
pv secret "abc123"
```

## Multi-File Programs

### Running Multiple Files

You can run multiple files together:

```bash
.file1.y defines functions.
.file2.y uses them.

y file1.y,file2.y
```

Or run all `.y` files in a directory:

```bash
y myfolder
y all
```

Files are concatenated in order, so declarations in earlier files are available in later files.

## Future: Namespaced Modules

The following module patterns are planned for post-MVP:

```
.Standard library modules (planned).
use "regex"
use "http"
use "json"

result regex.match(`\d+`, input)
resp http.open("GET", "https://api.example.com")
data json.parse(resp)
```

These will provide namespaced access to extended functionality without polluting the global scope.
