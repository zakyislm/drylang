// test_01_vars.y
a "hello"
b 10
c t
d f
? e

// Reassignment
a = "world"
b = 20

// Constants
x = 100
// x = 200 // uncommenting should error

// Print all
pt a
pt b
pt c
pt d
pt e

// Shadowing and scoping test
a "global"
fn test_scope() {
    a = "local_new" // wait, this sets local because it's not a global unless global tracked
    pt a
}
test_scope()
pt a
