// test_06_errors.y

fn risky() {
    err "something went wrong"
}

try {
    risky()
    pt "This should not print"
} err (e) {
    pt "Caught error: " + e
}

pt "Execution continues"
