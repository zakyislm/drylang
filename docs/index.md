# dryLang Documentation

Welcome to the official dryLang documentation. dryLang is a minimalist, dynamically-typed programming language compiled to bytecode and executed on a stack-based virtual machine.

## Table of Contents

1. [Getting Started](getting-started.md)
2. [Variables & Constants](variables.md)
3. [Data Types](types.md)
4. [Operators](operators.md)
5. [Strings](strings.md)
6. [Control Flow](control-flow.md)
7. [Loops](loops.md)
8. [Functions](functions.md)
9. [Collections](collections.md)
10. [Structs](structs.md)
11. [Error Handling](error-handling.md)
12. [Modules](modules.md)
13. [Built-in Functions](builtins.md)
14. [HTTP Server](http-server.md)
15. [Database](database.md)
16. [Comments](comments.md)
17. [CLI Reference](cli.md)
18. [Error Messages](errors.md)

## Design Philosophy

dryLang follows the **Writeless, get more** principle — Writeless, get more.

- **Go** inspired: simplicity, single-binary output, fast compilation
- **Python** inspired: clean variable naming, dynamic typing
- **JavaScript** inspired: `${}` string interpolation
- All compressed to the absolute minimum syntax

Every keyword is 1–4 characters. Every built-in function is 1–4 characters. No exceptions.

## What Can dryLang Do?

- **HTTP Server** — Built-in `op()` starts a web server (powered by Go's `net/http`)
- **SQL Database** — Built-in `db()` connects to SQLite, MySQL, or PostgreSQL
- **HTTP Client** — Built-in `req()` fetches data from any URL
- **JSON** — Built-in `json()` parses JSON into native maps/arrays
- **Module System** — Import local files, URLs, or directly from GitHub repos
- **File I/O** — Read, write, list, and delete files
- **System Commands** — Execute shell commands, read env variables
- **Advanced Math** — sqrt, pow, trig, log via `math()`
- **38 Built-in Functions** — all with 1–4 character names

---

Next : [Getting Started](getting-started.md)
<!-- W3Schools-like web docs integration placeholder -->
