# File Server

```javascript
// template: file-server
// desc: Basic static file server serving files from the current directory

fn handler(req) {
    path = req["path"]
    
     // Default to index.html for root path
    if path = "/" {
        path = "/index.html"
    }

     // Remove leading slash for local file path
    local_path = mod(path, "/", "")

    try {
        content = r(local_path)
        
         // Basic content type guessing
        ctype = "text/plain"
        if has(local_path, ".html") { ctype = "text/html" }
        elif has(local_path, ".json") { ctype = "application/json" }
        elif has(local_path, ".css") { ctype = "text/css" }
        elif has(local_path, ".js") { ctype = "application/javascript" }

        rev {
            "status": 200,
            "headers": {"Content-Type": ctype},
            "body": content
        }
    } err(e) {
        rev {
            "status": 404,
            "body": "File not found: " + local_path
        }
    }
}

PORT 8080
pt "Serving files on http://localhost:${PORT}"
op(PORT, handler, "mul", 100)

```
