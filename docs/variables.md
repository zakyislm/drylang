# Variables & Constants

## Variables

dryLang uses **no keyword** for variable declaration. The first assignment creates the variable.

### With `=` (explicit)

```
name = "Zaky"
age = 17
height = 175,5
```

### Without `=` (space-separated)

```
name "Zaky"
age 17
height 175,5
```

Both forms are identical. Use whichever you prefer.

### Reassignment

Variables can be reassigned freely:

```
score 100
pt score    .prints 100.

score 200
pt score    .prints 200.
```

### Naming Rules

- Variable names must start with a letter or underscore
- Can contain letters, digits, and underscores
- Convention: **lowercase with underscores** (`my_variable`)
- **ALL CAPS** names are automatically treated as constants (see below)

```
.Valid names.
name "ok"
my_var 42
_private "yes"
x2 100

.These are auto-constants (ALL CAPS).
MAXLIFE 5
API_KEY "abc123"
```

## Constants

Constants are values that cannot be reassigned after declaration. There are two ways to create them.

### Explicit: `cns` Keyword

```
cns pi 3,14
cns max_speed 299792458
```

With `=`:

```
cns pi = 3,14
```

### Implicit: ALL CAPS

Any variable name in ALL UPPERCASE is automatically constant:

```
MAXLIFE 5
API_URL "https://api.example.com"
MAX_RETRIES 3
```

### Constant Reassignment Error

Attempting to reassign a constant produces a compile error:

```
cns pi 3,14
pi 3,15          .Error: locked pi.

MAXLIFE 5
MAXLIFE 10       .Error: locked MAXLIFE.
```

## Private Variables

Use the `pv` prefix to mark a variable as private (not exported when the module is imported):

```
pv secret_key "abc123"
pv count 0
```

Private also works with functions:

```
pv fn internal_helper() {
  rev "internal"
}
```

## Unknown Bool (`?`)

The `?` prefix declares a boolean variable with an initial value of `unknown`:

```
? status

pt status          .prints "unknown".
pt get(status)     .prints "unknown".

. Can be assigned t or f later .
status = t
pt status          .prints "t".

. Check for unknown state .
if status = unknown {
  pt "Status not yet determined"
}
```

This is useful for tri-state booleans where the value hasn't been determined yet.
