# CLI Reference

The dryLang compiler provides a command-line interface to execute and manage your scripts.

## Running Scripts

You can run scripts using the `dry` (or `y`) command:

```bash
# Run a single file
dry script.y
dry script.dry

# Run multiple files (concatenated in order)
dry lib.y,main.y

# Run all .y and .dry files in a directory
dry myfolder
dry .

# Run all .y and .dry files in the current directory (alternative)
dry all
```

## Remote Execution

dryLang can execute scripts directly from the internet or GitHub:

```bash
# Run a script from a raw URL
dry https://example.com/script.y

# Run from a GitHub repository (looks for idx.y by default)
dry github.com/user/repo

# Run a specific file from a GitHub repository
dry github.com/user/repo/path/to/script.y
```

## Creating Projects (Scaffolding)

Use the `init` command to scaffold a new project from built-in templates.

```bash
dry init [directory] [template]
```

### Examples:
```bash
# Create an empty project in the current directory (creates idx.y)
dry init .

# Create a project in a new folder called "myapp"
dry init myapp

# Create a REST API project in the current directory
dry init . api

# Create a Web File Server in a new folder
dry init website web
```

### Available Templates:
- `api` — REST API Server
- `web` — Static File Server
- `crud` — Database CRUD application
- `fetch` — HTTP JSON fetch example
- `cli` — Basic CLI tool
- `automation` — Automation script
- `scraper` — Web Scraper skeleton
- `hello` — Hello World

## REPL (Interactive Mode)

If you run `dry` without any arguments, it starts the interactive REPL:

```bash
$ dry
dryLang REPL v1.0.0
> pt "Hello"
Hello
```

## Options

- **`--help` or `-h`**: Displays the help menu and usage instructions.
- **`--version` or `-v`**: Displays the current version of dryLang.

---

[< Prev (Comments)](comments.md) | [Home](index.md) | [Next (Errors) >](errors.md)
<!-- W3Schools-like web docs integration placeholder -->
