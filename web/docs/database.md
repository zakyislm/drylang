# Database

dryLang has built-in SQL database support via the `db()` function. Supports SQLite, MySQL, and PostgreSQL out of the box.

## `db()` — Execute SQL

```rust
db(driver, dsn, query, ...params)
```

| Param | Type | Required | Description |
|-------|------|:--------:|-------------|
| `driver` | string | ✅ | `"sqlite"`, `"mysql"`, or `"postgres"` |
| `dsn` | string | ✅ | Connection string / file path |
| `query` | string | ✅ | SQL query |
| `...params` | any | ❌ | Parameterized query values (prevents SQL injection) |

## Drivers & Connection Strings

```rust
// SQLite — file-based, no setup needed
db("sqlite", "app.db", "...")

// MySQL
db("mysql", "user:password@tcp(localhost:3306)/dbname", "...")

// PostgreSQL
db("postgres", "host=localhost port=5432 user=admin password=secret dbname=app sslmode=disable", "...")
```

## Creating Tables

```rust
db("sqlite", "app.db", "CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, age INTEGER)")
```

## Inserting Data

Use `?` for parameterized queries (safe from SQL injection):

```rust
db("sqlite", "app.db", "INSERT INTO users (name, age) VALUES (?, ?)", "Zaky", 25)
db("sqlite", "app.db", "INSERT INTO users (name, age) VALUES (?, ?)", "Alice", 30)
```

**Returns** a map with:
- `last_insert_id` — ID of the inserted row
- `rows_affected` — number of affected rows

```rust
res db("sqlite", "app.db", "INSERT INTO users (name, age) VALUES (?, ?)", "Bob", 22)
pt res["last_insert_id"]    // e.g., 3
pt res["rows_affected"]     // 1
```

## Querying Data

SELECT queries return an **array of maps**, where each map represents a row:

```rust
users db("sqlite", "app.db", "SELECT * FROM users")
pt users
// [{id: 1, name: Zaky, age: 25}, {id: 2, name: Alice, age: 30}]

// Access individual rows and fields
pt users[0]["name"]     // "Zaky"
pt users[1]["age"]      // 30

// Loop through results
lp len(users) {
    pt users[i]["name"] + " is " + str(users[i]["age"])
}
```

With WHERE clause:

```rust
older db("sqlite", "app.db", "SELECT name, age FROM users WHERE age > ?", 24)
pt older
```

## Updating Data

```rust
res db("sqlite", "app.db", "UPDATE users SET age = ? WHERE name = ?", 23, "Bob")
pt res["rows_affected"]    // 1
```

## Deleting Data

```rust
res db("sqlite", "app.db", "DELETE FROM users WHERE name = ?", "Alice")
pt res["rows_affected"]    // 1
```

## Also Supports

- `PRAGMA` queries (SQLite)
- `SHOW` queries (MySQL)
- `DESCRIBE` queries (MySQL)

These are treated as SELECT-like queries and return arrays of maps.

## Complete Example

```rust
DB "app.db"

// Setup
db("sqlite", DB, "CREATE TABLE IF NOT EXISTS tasks (id INTEGER PRIMARY KEY, title TEXT, done INTEGER DEFAULT 0)")

// Insert
db("sqlite", DB, "INSERT INTO tasks (title) VALUES (?)", "Learn dryLang")
db("sqlite", DB, "INSERT INTO tasks (title) VALUES (?)", "Build an API")

// Query
tasks db("sqlite", DB, "SELECT * FROM tasks")
lp len(tasks) {
    status "[ ]"
    if tasks[i]["done"] = 1 { status "[x]" }
    pt "${status} ${tasks[i][\"title\"]}"
}

// Update
db("sqlite", DB, "UPDATE tasks SET done = 1 WHERE title = ?", "Learn dryLang")

// Cleanup
del(DB)
```

---

Prev : [HTTP Server](http-server.md) | Next : [Comments](comments.md)
<!-- W3Schools-like web docs integration placeholder -->
