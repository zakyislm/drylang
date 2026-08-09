# HTTP Server

dryLang has a built-in HTTP server powered by Go's `net/http`. Start one with a single function call.

## `op()` — Start Server

```rust
op(port, handler, mode, maxWorkers)
```

| Param | Type | Required | Description |
|-------|------|:--------:|-------------|
| `port` | number | ✅ | Port to listen on |
| `handler` | function | ✅ | Request handler function |
| `mode` | string | ❌ | `"uni"` (default) or `"mul"` |
| `maxWorkers` | number | ❌ | Max concurrent workers (default: 100) |

### Modes

- **`"uni"`** — Single-threaded with mutex locking. Requests are processed one at a time. Safe for shared state.
- **`"mul"`** — Multi-threaded. Each request gets a fresh VM clone with shared globals. Use for high-throughput APIs.

## Handler Function

The handler receives a **request map** and must `rev` (return) a response.

### Request Map

```rust
fn handler(req) {
    pt req["method"]    // "G", "PO", "PUT", "PAT", "DEL", "OPT", "H"
    pt req["path"]      // "/api/users"
    pt req["body"]      // request body string
    pt req["query"]     // map of query parameters
}
```

**HTTP Method Mappings** (shortened for dryLang):

| HTTP Method | dryLang |
|-------------|---------|
| GET | `"G"` |
| POST | `"PO"` |
| PUT | `"PUT"` |
| PATCH | `"PAT"` |
| DELETE | `"DEL"` |
| OPTIONS | `"OPT"` |
| HEAD | `"H"` |

### Response

Return a **string** for simple text response:

```rust
fn handler(req) {
    rev "Hello, World!"
}
```

Return a **map** for custom status and body:

```rust
fn handler(req) {
    rev {
        "status": 200,
        "body": "{\"message\": \"success\"}"
    }
}
```

## Complete Example

```rust
fn handler(req) {
    if req["method"] = "G" & req["path"] = "/" {
        rev {
            "status": 200,
            "body": "{\"message\": \"Welcome to dryLang API\"}"
        }
    } elif req["method"] = "PO" & req["path"] = "/echo" {
        rev {
            "status": 200,
            "body": req["body"]
        }
    } el {
        rev {
            "status": 404,
            "body": "{\"error\": \"not found\"}"
        }
    }
}

pt "API running on http://localhost:8080"
op(8080, handler, "mul", 100)
```

## Notes

- `op()` is **blocking** — it runs indefinitely until the process is killed.
- In `"mul"` mode, globals are **shared** across requests (useful for in-memory state).
- The server prints a startup message: `Starting dryLang server on port XXXX (mode: yyy)...`

---

Prev : [Built-ins](builtins.md) | Next : [Database](database.md)
<!-- W3Schools-like web docs integration placeholder -->
