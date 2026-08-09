# Try Err

```javascript
// 07_try_err.y - Error Handling

fn risky_operation(fail) {
    if fail = t {
        err "Operation failed!"
    }
    rev "Success!"
}

try {
    pt "Running risky operation..."
    result = risky_operation(t)
    pt result
} err(e) {
    pt "Caught an error: " + e
}

```
