// dryexamples/cheatsheet.y - 100% Syntax in 1 file

// Vars & Consts
v "Z"
cns P 3.14

// Print
pt v

// Arrays & Maps
a = [1,2]
m = {"k":"v"}
add(a, 3)

// Struct
U { n }
u = U("Z")

// If & Loop
if t { pt "y" } el { pt "n" }
lp 3 { pt i }

// Functions
fn mul(x) { rev x*2 }
pt mul(5)

// Try
try { err "e" } err(e) { pt e }

// Use (local, URL, GitHub)
use "../examples/01_hello.y"
// use "https://example.com/lib.y"
// use "github.com/user/repo"

// Advanced Math
pt math("sqrt", 16)
pt math("pow", 2, 10)

// Database (SQLite)
db("sqlite", "app.db", "CREATE TABLE IF NOT EXISTS t (id INTEGER PRIMARY KEY, name TEXT)")
db("sqlite", "app.db", "INSERT INTO t (name) VALUES (?)", "Zaky")
users = db("sqlite", "app.db", "SELECT * FROM t")
pt users

// HTTP Client & JSON
// body = req("https://api.github.com/")
// data = json(body)

// System & File I/O
// pt now()
// w("test.txt", "data")
// r("test.txt")
// del("test.txt")

// HTTP Server
// fn handler(req) { rev "Hello " + req["method"] }
// op(8080, handler, "mul", 100)
