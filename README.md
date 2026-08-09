# dryLang

> **Writeless, get more**

dryLang is a minimalist, dynamically-typed programming language compiled to bytecode and executed on a fast, stack-based virtual machine.

## Why dryLang? (Origin Story)

I used to write long, verbose messages. I once wrote **20 paragraphs** just to apologize to someone. It was exhausting, repetitive, and — looking back — it didn't actually convey my emotions better. It just added noise.

Then I realized: **text is just text**. It doesn't change the reality beneath it. Adding more words doesn't make you more sincere. It just wastes time.

I started texting differently — shorter, more direct. "sry" instead of a paragraph. "y?" instead of "why are you doing this??". No fluff. Just signal, no noise.

Around the same time, I got frustrated with programming languages. They all felt bloated — you need a `package.json`, a virtual environment, an npm ecosystem, just to write a simple HTTP server or script. It felt like writing 20 paragraphs when you only need one sentence.

**So I built dryLang:** A language where you say exactly what you mean, nothing more.

- No boilerplate ceremony
- No dependency management theater
- No keywords longer than 4 characters
- Everything compressed to its essence

**Inspirations:**
- **Go**: Single binary output, fast compilation, `net/http` simplicity.
- **Python**: Clean variable naming, dynamic typing.
- **JavaScript**: String interpolation (`${}`).

But distilled. Stripped of everything unnecessary.

The result is a language compressed to its absolute essence: **No keyword or built-in function exceeds 4 characters.**

## Proof of Concept: The 4-Character Limit
dryLang enforces a strict limit on all built-in syntax to save keystrokes.

| Category | Keywords / Functions |
|----------|----------------------|
| **Core** | `cns`, `t`, `f`, `fn`, `rev`, `if`, `el`, `elif`, `on`, `lp`, `done`, `con`, `try`, `err`, `pv`, `use` |
| **I/O & Sys** | `pt`, `in`, `r`, `w`, `dir`, `del`, `arg`, `env`, `cmd`, `die`, `q`, `now`, `date` |
| **Data** | `num`, `str`, `get`, `len`, `add`, `rm`, `pop`, `sort`, `key`, `val`, `json` |
| **String** | `cap`, `low`, `trm`, `spl`, `j`, `mod`, `has` |
| **Math** | `abs`, `min`, `max`, `rnd`, `ran`, `math` |
| **Net & DB** | `req`, `op`, `db` |

## Size & Performance
- **Tiny Toolchain**: The entire compiler, VM, and runtime is a single ~5MB binary. Compare this to Node.js (~50MB) or Python (~30MB).
- **Fast Execution**: Code compiles to bytecode instantly and runs on a highly optimized stack-based VM written in Go.
- **Zero Dependencies**: HTTP servers, SQLite/PostgreSQL drivers, JSON parsing, and advanced math are all baked into the single binary.

### Pros & Cons
**Pros (+):**
- Instantly start a web server or database connection with 1 line of code.
- No `package.json`, no `pip install`, no `go mod init`.
- Incredibly small syntax footprint.
- Great for quick scripts, automations, and simple APIs.

**Cons (-):**
- Small ecosystem (no external package manager yet).
- Strict (unused variables will stop compilation).
- Single-letter keywords require a slight learning curve.

## Real-World Examples

### 1. HTTP Server in 5 Lines
```rust
fn handler(req) {
    rev { "status": 200, "body": "{\"hello\": \"world\"}" }
}
op(8080, handler, "mul", 100)
```

### 2. SQLite Database
```rust
db("sqlite", "app.db", "CREATE TABLE users (id INTEGER PRIMARY KEY, name TEXT)")
db("sqlite", "app.db", "INSERT INTO users (name) VALUES (?)", "Zaky")
users = db("sqlite", "app.db", "SELECT * FROM users")
pt users
```

### 3. Fetch JSON API
```rust
body req("https://api.github.com/users/octocat")
data json(body)
pt "Followers: " + str(data["followers"])
```

## Installation

```bash
git clone https://github.com/zakyislm/drylang.git
cd drylang
go build -o dry .
```

(You can alias the executable to `y`, `dry`, or `drylang`).

## Running Scripts

dryLang supports `.y` and `.dry` file extensions.

```bash
# Single file
dry main.y

# Multiple files (concatenated in order)
dry file1.y,file2.y

# Entire directory
dry myfolder
```

## Documentation

- Full Documentation (docs/) — comprehensive language reference
- DryDocs (drydocs/) — ultra-concise reference (dryLang style)
- Changelog (CHANGELOG.md) — release history
- Contributing (CONTRIBUTING.md) — how to help out

## License
MIT