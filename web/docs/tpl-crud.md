# Crud

```javascript
// template: crud
// desc: Full CRUD app — SQLite database + REST API

DB "app.db"

 // Initialize database
db("sqlite", DB, "CREATE TABLE IF NOT EXISTS items (id INTEGER PRIMARY KEY, name TEXT, done INTEGER DEFAULT 0)")

fn handler(req) {
     // GET / → list all items
    if req["method"] = "G" & req["path"] = "/" {
        items = db("sqlite", DB, "SELECT * FROM items")
        rev {
            "status": 200,
            "body": str(items)
        }
    }

     // POST /add?name=xxx → create item
    elif req["method"] = "PO" & req["path"] = "/add" {
        name = req["body"]
        if len(name) = 0 {
            name = "Untitled"
        }
        res = db("sqlite", DB, "INSERT INTO items (name) VALUES (?)", name)
        rev {
            "status": 201,
            "body": "{\"created\": true, \"id\": " + str(res["last_insert_id"]) + "}"
        }
    }

     // PUT /done → mark item as done (body = id)
    elif req["method"] = "PUT" & req["path"] = "/done" {
        id = num(req["body"])
        db("sqlite", DB, "UPDATE items SET done = 1 WHERE id = ?", id)
        rev {
            "status": 200,
            "body": "{\"updated\": true}"
        }
    }

     // DELETE /del → delete item (body = id)
    elif req["method"] = "DEL" & req["path"] = "/del" {
        id = num(req["body"])
        db("sqlite", DB, "DELETE FROM items WHERE id = ?", id)
        rev {
            "status": 200,
            "body": "{\"deleted\": true}"
        }
    }

    el {
        rev {
            "status": 404,
            "body": "{\"error\": \"not found\"}"
        }
    }
}

pt "CRUD API on http://localhost:3000"
pt "  GET  /       → list items"
pt "  POST /add    → create (body = name)"
pt "  PUT  /done   → mark done (body = id)"
pt "  DEL  /del    → delete (body = id)"
op(3000, handler, "mul", 100)

```
