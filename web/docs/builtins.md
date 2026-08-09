# Built-in Functions

dryLang provides 38 built-in functions available without any imports. All function names are 1–4 characters.

## I/O

### `pt` — Print

Outputs a value to stdout, followed by a newline.

```rust
pt "Hello, World!"
pt 42
pt [1, 2, 3]
pt {"key": "value"}
```

`pt` is a **statement**, not a function. It does not return a value.

### `in` — Input

Reads a line of text from stdin. Optionally displays a prompt.

```rust
// Without prompt
value in()

// With prompt
name in("Enter your name: ")
pt "Hello, ${name}"
```

`in()` always returns a string. Use `num()` to convert to a number:

```rust
age_str in("Age: ")
age num(age_str)
```

### `r` — Read File

Reads the entire contents of a file as a string.

```rust
content r("data.txt")
pt content
```

Returns the file content as a string. Throws an error if the file doesn't exist.

### `w` — Write File

Writes a string to a file. Creates the file if it doesn't exist, overwrites if it does.

```rust
w("output.txt", "Hello, File!")
w("data.json", `{"key": "value"}`)
```

Returns `t` on success. Throws an error on failure.

### `dir` — List Directory

Lists all entries (files and directories) in a directory. Returns an array of filenames.

```rust
files dir(".")
pt files    // ["main.y", "lib.y", "data"]

files dir("./src")
lp len(files) {
    pt files[i]
}
```

### `del` — Delete File

Deletes a file from the filesystem.

```rust
w("temp.txt", "temporary data")
del("temp.txt")    // file removed
```

Returns `t` on success. Throws an error if the file doesn't exist.

---

## Type Conversion

### `num` — To Number

Converts a string to a number.

```rust
x num("42")       // 42
y num("3.14")     // 3.14
```

Throws an error if the string is not a valid number.

### `str` — To String

Converts any value to its string representation.

```rust
s str(42)          // "42"
s str(t)           // "t"
s str([1, 2])      // "[1, 2]"
```

### `get` — Type Check

Returns the type of a value as a string.

```rust
pt get(42)          // "number"
pt get("hello")     // "string"
pt get(t)           // "bool"
pt get([1, 2])      // "array"
pt get({"a": 1})    // "map"
pt get(greet)       // "fn"
pt get(unknown)     // "unknown"
```

---

## Math

### `abs` — Absolute Value

```rust
pt abs(-42)    // 42
pt abs(42)     // 42
```

### `min` — Minimum

Returns the smaller of two numbers.

```rust
pt min(10, 20)    // 10
pt min(-5, 5)     // -5
```

### `max` — Maximum

Returns the larger of two numbers.

```rust
pt max(10, 20)    // 20
pt max(-5, 5)     // 5
```

### `rnd` — Round

Rounds a number to the nearest integer.

```rust
pt rnd(3.7)     // 4
pt rnd(3.2)     // 3
```

### `ran` — Random

Returns a random float between 0.0 (inclusive) and 1.0 (exclusive).

```rust
pt ran()           // e.g., 0.7291...
random_int rnd(ran() * 100)    // random 0-100
```

### `math` — Advanced Math

Provides advanced mathematical functions. First argument is the operation name.

```rust
// Square root
pt math("sqrt", 16)     // 4

// Power
pt math("pow", 2, 10)   // 1024

// Ceiling and floor
pt math("ceil", 4.1)    // 5
pt math("floor", 4.9)   // 4

// Trigonometry (radians)
pt math("sin", 0)       // 0
pt math("cos", 0)       // 1
pt math("tan", 0)       // 0

// Logarithms
pt math("log", 1)       // 0 (natural log)
pt math("log10", 100)   // 2
```

**Available operations:** `sqrt`, `pow`, `ceil`, `floor`, `sin`, `cos`, `tan`, `log`, `log10`

---

## String Functions

### `cap` — Capitalize (Uppercase)

Converts a string to UPPERCASE.

```rust
pt cap("hello")    // HELLO
pt cap("World")    // WORLD
```

### `low` — Lowercase

Converts a string to lowercase.

```rust
pt low("HELLO")    // hello
pt low("World")    // world
```

### `trm` — Trim

Removes leading and trailing whitespace.

```rust
pt trm("  hello  ")    // hello
pt trm("\t hi \n")      // hi
```

### `spl` — Split

Splits a string by a delimiter, returns an array.

```rust
parts spl("a,b,c", ",")
pt parts    // ["a", "b", "c"]

words spl("hello world", " ")
pt words    // ["hello", "world"]
```

### `j` — Join

Joins an array of strings with a separator.

```rust
arr ["a", "b", "c"]

pt j(arr, ",")     // a,b,c
pt j(arr, " - ")   // a - b - c
pt j(arr, "")      // abc
```

### `mod` — Modify (Replace)

Replaces all occurrences of a substring.

```rust
pt mod("hello world", "world", "earth")    // hello earth
pt mod("aabaa", "a", "x")                  // xxbxx
```

### `has` — Contains

Checks if a string contains a substring. Returns `t` or `f`.

```rust
pt has("hello world", "world")    // t
pt has("hello world", "xyz")      // f
```

---

## Array Functions

### `len` — Length

Returns the number of elements in an array, characters in a string, or entries in a map.

```rust
pt len([1, 2, 3])        // 3
pt len("hello")           // 5
pt len({"a": 1, "b": 2}) // 2
pt len([])                // 0
```

### `add` — Add (Append)

Appends an element to an array. Returns a new array.

```rust
arr [1, 2, 3]
arr add(arr, 4)
pt arr    // [1, 2, 3, 4]
```

### `sort` — Sort

Returns a sorted copy of an array. Numbers sort numerically, strings sort alphabetically.

```rust
pt sort([3, 1, 4, 1, 5])         // [1, 1, 3, 4, 5]
pt sort(["cherry", "apple"])      // ["apple", "cherry"]
```

### `pop` — Pop Last

Returns the last element of an array.

```rust
pt pop([1, 2, 3])    // 3
```

Throws an error on empty arrays.

### `rm` — Remove

Removes an element at the given index. Returns a new array.

```rust
arr [10, 20, 30, 40]
arr rm(arr, 1)
pt arr    // [10, 30, 40]
```

---

## Map Functions

### `key` — Keys

Returns all keys of a map as an array.

```rust
m {"name": "Zaky", "age": 17}
pt key(m)    // ["name", "age"]
```

### `val` — Values

Returns all values of a map as an array.

```rust
m {"name": "Zaky", "age": 17}
pt val(m)    // ["Zaky", 17]
```

---

## Time

### `now` — Unix Timestamp

Returns the current Unix timestamp in milliseconds.

```rust
pt now()    // e.g., 1786288352211
```

### `date` — Date Info

Returns a map with current date/time components.

```rust
d date()
pt d["year"]     // 2026
pt d["month"]    // 8
pt d["day"]      // 9
pt d["hour"]     // 22
pt d["min"]      // 12
pt d["sec"]      // 32
pt d["format"]   // "2026-08-09 22:12:32"
```

---

## HTTP & JSON

### `req` — HTTP Request

Makes an HTTP GET request and returns the response body as a string.

```rust
body req("https://api.example.com/data")
pt body
```

### `json` — Parse JSON

Parses a JSON string into a dryLang map or array.

```rust
data json("{\"name\":\"Zaky\",\"age\":17}")
pt data["name"]    // Zaky
pt data["age"]     // 17

// Works with arrays too
arr json("[1, 2, 3]")
pt arr[0]          // 1
```

Combined with `req`:

```rust
body req("https://api.example.com/users")
users json(body)
pt users[0]["name"]
```

---

## System

### `arg` — Command-Line Arguments

Returns an array of command-line arguments (excludes the binary and script name).

```rust
// Running: y script.y hello world
args arg()
pt args        // ["hello", "world"]
pt args[0]     // "hello"
```

### `env` — Environment Variable

Reads an environment variable. Returns empty string if not set.

```rust
pt env("USERNAME")    // "user"
pt env("HOME")        // "/home/user"
pt env("PATH")        // system PATH
```

### `cmd` — Execute Shell Command

Runs a shell command and returns the output as a string.

```rust
pt cmd("echo", "hello")     // hello
pt cmd("whoami")             // current user

// With multiple arguments
output cmd("git", "status")
pt output
```

### `die` — Exit with Error

Immediately terminates the program with exit code 1. Optionally prints an error message to stderr.

```rust
if len(arg()) < 1 {
    die("Usage: y script.y <name>")
}
```

---

## Utility

### `q` — Quiet (Sleep)

Pauses execution for the specified number of milliseconds.

```rust
pt "Starting..."
q(1000)        // wait 1 second
pt "Done!"

q(500)         // wait 500ms
```

---

## HTTP Server

### `op` — Start Server

Starts an HTTP server. See [HTTP Server](http-server.md) for full documentation.

```rust
fn handler(req) {
    rev "Hello from dryLang!"
}
op(8080, handler, "mul", 100)
```

---

## Database

### `db` — SQL Database

Executes SQL queries against a database. See [Database](database.md) for full documentation.

```rust
db("sqlite", "app.db", "CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT)")
db("sqlite", "app.db", "INSERT INTO users (name) VALUES (?)", "Zaky")
users db("sqlite", "app.db", "SELECT * FROM users")
pt users
```

---

Prev : [Modules](modules.md) | Next : [HTTP Server](http-server.md)
<!-- W3Schools-like web docs integration placeholder -->
