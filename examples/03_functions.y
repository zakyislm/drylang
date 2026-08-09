// 03_functions.y - Functions and Returns

fn greet(name) {
    rev "Hello, ${name}"
}

fn add(a, b) {
    rev a + b
}

// Early return
fn check_age(age) {
    if age < 18 {
        rev "Too young"
    }
    rev "Welcome"
}

pt greet("Zaky")
pt "5 + 7 = ${add(5, 7)}"
pt check_age(17)
