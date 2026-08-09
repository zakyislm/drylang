pt "Testing assignment vs comparison"
a f
if a = f {
    pt "a = f is true!"
} el {
    pt "a = f is false!"
}

b "hello"
if b = "hello" {
    pt "b = hello is true!"
}

if a = t {
    pt "a = t is true! (This means it ASSIGNED!)"
} el {
    pt "a = t is false! (This means it COMPARED!)"
}
