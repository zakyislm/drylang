# Http Server

```javascript
// 11_http_server.y
// Demonstrates the built-in HTTP server using op()

// The handler takes a request map containing method, path, body, and query.
fn handler(req) {
    pt "Received " + req["method"] + " request for " + req["path"]

    // Simple routing
    if req["method"] = "G" & req["path"] = "/" {
        // Return a map to control status code and body
        rev {
            "status": 200,
            "body": "{\"message\": \"Welcome to dryLang API\"}"
        }
    } elif req["method"] = "PO" & req["path"] = "/echo" {
        // Echo the body back
        rev {
            "status": 200,
            "body": req["body"]
        }
    } el {
        // Return a simple string (defaults to status 200)
        // Or return a map for a 404
        rev {
            "status": 404,
            "body": "{\"error\": \"Not found\"}"
        }
    }
}

pt "Starting API on http://localhost:8080"
pt "Try: curl http://localhost:8080"
pt "Try: curl -X POST -d 'test data' http://localhost:8080/echo"

// Start the server on port 8080
// Using "mul" (multi-threaded) mode with 100 max workers.
op(8080, handler, "mul", 100)

```
