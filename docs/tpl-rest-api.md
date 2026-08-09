# Rest Api

```javascript
// template: rest-api
// desc: REST API server with JSON responses and routing

fn handler(req) {
    if req["method"] = "G" & req["path"] = "/" {
        rev {
            "status": 200,
            "body": "{\"message\": \"Welcome to the API\", \"version\": \"1.0\"}"
        }
    } elif req["method"] = "G" & req["path"] = "/health" {
        rev {
            "status": 200,
            "body": "{\"status\": \"ok\"}"
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

PORT 8080
pt "API running on http://localhost:${PORT}"
op(PORT, handler, "mul", 100)

```
