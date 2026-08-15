# dryLang Features

Semua fitur dryLang, sintaks, contoh pemakaian. Terakhir diverifikasi: 2026-08-13.

---

## 1. Variabel & Konstanta

### Assignment (`=`)
Variabel dideklarasi otomatis saat pertama di-assign. `=` juga dipakai sebagai operator perbandingan (single `=` untuk assign DAN compare).

```rust
name = "Zaky"
age = 17
x = 5
```

### Konstanta (`cns`)
```rust
cns MAX = 100
cns API_URL = "https://api.example.com"
```

### Strict Unused Check
Variabel/fungsi yang dideklarasi tapi tidak dipakai → error compile:
```rust
x = 5        // ERROR: unused x
```

### Unknown / Null (`unknown`)
```rust
a = unknown
pt(a ?? "fallback")   // ?? = null coalescing
```

---

## 2. Tipe Data

| Tipe | Contoh | Ket |
| --- | --- | --- |
| Number | `42`, `3.14`, `1e400`, `2.5e-3` | float64, support exponent |
| String | `"hello"`, `'hello'` | kutip ganda/tunggal |
| Raw String | `` `raw text` `` | backtick, tanpa escape |
| Bool | `t`, `f` | true/false |
| Array | `[1, 2, 3]` | |
| Map | `{"a": 1, "b": 2}` | |
| Struct | `Point { x }` | data container |
| Class | `cl Dog { ... }` | OOP penuh |
| Unknown | `unknown` | null/undefined |

```rust
arr = [1, 2, 3]
m = {"name": "zaky", "age": 17}
pt(arr[0])      // 1
pt(m["name"])   // zaky
```

---

## 3. String Interpolation (`${}`)

```rust
name = "Zaky"
pt("Hello, ${name}!")     // Hello, Zaky!
pt("sum=${1 + 2}")        // sum=3
```

---

## 4. Struktur Kontrol

### if / elif / el
```rust
x = 10
if (x > 5) {
    pt("big")
} elif (x > 0) {
    pt("small")
} el {
    pt("zero")
}
```

### on (switch)
```rust
v = 2
on (v) {
    0 { pt("zero") }
    1 { pt("one") }
    2 { pt("two") }
}
```

### lp (loop universal)
```rust
// Counted loop — variabel i otomatis
lp 5 {
    pt(i)
}

// Infinite loop
lp {
    // ...
    done
}
```

### done (break) / con (continue)
```rust
lp 10 {
    if (i > 5) { done }
    if (i = 2) { con }
}
```

---

## 5. Fungsi

### fn
```rust
fn add(a, b) {
    rev a + b
}
pt(add(2, 3))     // 5
```

### Arrow function (`->`)
```rust
// single param
f = x -> x * 2
pt(f(4))          // 8

// multi param
add2 = (a, b) -> a + b
pt(add2(2, 3))    // 5

// block body
f2 = x -> {
    y = x * 2
    rev y + 1
}
pt(f2(5))         // 11

// sebagai argument
fn apply(f, v) { rev f(v) }
pt(apply(x -> x + 1, 10))   // 11
```

### rev (return)
```rust
fn isEven(n) {
    if (n % 2 = 0) { rev t }
    rev f
}
```

### async (asn / mul / uni / awt)
```rust
asn fn worker(n) {
    pt("working on", n)
}

// parallel (mul)
mul 4 worker(23)

// single (uni)
uni worker(1)

// tunggu semua selesai
awt
```

---

## 6. Struct (data container)

```rust
Point {
    x
    y
}

p = Point(1, 2)
pt(p.x)         // 1
p.x = 99        // dot assign
```

Struct murni data — tanpa method, tanpa inheritance. OOP pakai class.

---

## 7. Class (OOP penuh)

### Deklarasi + init + method
```rust
cl Dog {
    name
    fn init(n) {
        this.name = n
    }
    fn bark() {
        rev this.name + "!"
    }
}

d = Dog("rex")
pt(d.bark())    // rex!
```

### Inheritance (`<-`) — multi-parent
```rust
cl Animal {
    fn speak() { rev "..." }
}
cl Dog <- Animal {
    fn speak() { rev "woof" }
}
cl Cat <- Animal {
}

d = Dog()
pt(d.speak())   // woof (override)
c = Cat()
pt(c.speak())   // ... (inherit)
```

```rust
cl A { fn fa() { rev "A" } }
cl B { fn fb() { rev "B" } }
cl C <- A, B {}       // multi-parent
c = C()
pt(c.fa())      // A
pt(c.fb())      // B
```

### Private (`pv`)
```rust
cl Box {
    pv secret          // private field
    pv fn helper() {}  // private method
    fn init() { this.secret = 1 }
}

b = Box()
pt(b.secret)    // COMPILE ERROR: pv access b.secret
```

### Private class (`pv cl`)
```rust
pv cl Secret {
    x
    fn init() { this.x = 5 }
}
s = Secret()    // COMPILE ERROR: pv class access Secret
```

---

## 8. Error Handling (try / err)

```rust
try {
    err("boom")     // throw
} err(e) {
    pt("caught: " + e)
}
```

```rust
// nested try
try {
    err("inner")
} err(e) {
    try {
        err("outer: " + e)
    } err(e2) {
        pt(e2)
    }
}
```

---

## 9. Module (`use`)

```rust
// file import
use "helpers.dry"
// folder import — semua file di-load, sorted, .dry > .y > lainnya
use "lib"
```

```rust
// lib/idx.dry
fn greet() { rev "hi" }

// main.dry
use "lib"
pt(greet())     // hi
```

---

## 10. Built-in Functions

### String
```rust
str(42)          // "42"
num("42")        // 42
low("ABC")       // "abc"
cap("hello")     // "HELLO"
trm("  x  ")     // "x"
spl("a,b,c", ",") // ["a","b","c"]
j(["x","y"], "-") // "x-y"
mod("aaa","a","b") // "bbb"
fmt("%s-%d","x",5) // "x-5"
len("abc")       // 3
has("hello","ell") // t
```

### Math
```rust
abs(-5)        // 5
min(3, 9)      // 3
max(3, 9)      // 9
rnd(3.7)       // 4
math.sqrt(16)  // 4
math.pow(2, 10) // 1024
```

### Array/Map
```rust
len([1,2,3])   // 3
add([1,2], 3)  // [1,2,3]
pop([1,2,3])   // 3
rm([1,2,3], 1) // [1,3]
sort([3,1,2])  // [1,2,3]
key({"a":1})   // ["a"]
val({"a":1})   // [1]
has({"a":1},"a") // t
get("x")       // "string"
```

### Crypto
```rust
hash.md5("abc")      // 900150983cd24fb0d6963f7d28e17f72
enc.b64("hello")     // aGVsbG8=
enc.hex("ab")        // 6162
jwt.gen({"a":1}, "secret")  // token
jwt.fit(token, "secret")    // {"a":1}
```

### Regex
```rust
rgx.find("abc123", "[0-9]+") // result
rgx.rep("a1b2", "[0-9]", "X") // "aXbX"
```

### I/O
```rust
pt("print")           // stdout
in("prompt: ")        // stdin
r("file.txt")         // read file
w("file.txt", "data") // write file
```

### HTTP
```rust
// request
res = req("https://api.example.com", {"method": "G", "body": "..."})
data = json(res)

// server
fn handler(req) {
    rev {"status": 200, "body": "hello"}
}
op(8080, handler, "mul", 100)
```

### Database (SQLite/Postgres/MySQL)
```rust
rows = db("sqlite", "data.db", "SELECT * FROM users")
dbpool("sqlite", "data.db", 10, 2, 30)
```

### State
```rust
mem.set("key", 42, 60)     // TTL cache
mem.get("key")
sess.set("sid", "user", "zaky", 3600)
sess.get("sid", "user")
flag.set("beta", t)
flag.check("beta")
flow.create("id", "new")
flow.transition("id", "paid")
flow.get("id")
```

### Rate limit / Jobs / Cron / Pipe
```rust
rate.check("1.1.1.1", 100, 60)   // t/f
job.push(fn2, 1, 2)
cron.add("* * * * *", fn2)
ch = pipe.make(10)
pipe.send(ch, 5)
pipe.recv(ch)
```

### Validation / Template / Misc
```rust
valid.mail("a@b.com")  // t
valid.num("42")        // t
tmpl.render("Hi {{.name}}", {"name":"zaky"})
sys.os()               // {"name":"windows",...}
now()                  // unix ms
date()                 // map year/month/day/...
```

---

## 11. Operator

| Operator | Fungsi |
| --- | --- |
| `=` | assign DAN perbandingan (single =) |
| `+ - * / %` | aritmatika |
| `< > <= >=` | perbandingan (number + string) |
| `!=` | tidak sama |
| `& \| !` | logika and/or/not |
| `?.` | safe navigation |
| `??` | null coalescing |
| `.` | property access |
| `->` | arrow function |
| `<-` | class inheritance |

```rust
x = 5
if (x = 5) { pt("same") }    // single = bandingkan
"a" < "b"                     // t (string lexicographic)
```

---

## 12. Angka

```rust
42        // integer
3.14      // float
1e400     // exponent (overflow literal -> E110 bad number)
2.5e-3    // 0.0025
99999999999999999999   // 1e20 (no garbage)
```

---

## 13. Sandbox (security)

Builtin berbahaya **terkunci default**. Buka via env:

| Env | Efek |
| --- | --- |
| `DRY_ALLOW_CMD=echo,cat` | izinkan `cmd`/`sys.run` untuk perintah tersebut |
| `DRY_ALLOW_DB=1` | izinkan `db` |
| `DRY_ALLOW_URL=api.example.com` | izinkan `req` + `use` untuk host tersebut |
| `DRY_CORS_ORIGIN=https://app.example.com` | CORS origin server (default `*`) |
| `DRY_LOG_JSON=1` | log format JSON |

```bash
# dev open
DRY_ALLOW_CMD=* DRY_ALLOW_DB=1 DRY_ALLOW_URL=* dry script.dry

# prod locked (default) — cmd/db/url ditolak
dry script.dry
```

Terkunci: `cmd denied: echo`, `sys.run denied: echo`, `db denied`, `url denied`.

---

## 14. HTTP Server (prod hardening)

- Timeout: read 10s, header 5s, write 10s, idle 60s, max header 1MB
- **Graceful shutdown**: SIGINT/SIGTERM → `Shutdown(ctx)` (10s timeout)
- **Metrics**: `/debug/vars` (expvar) — `dry_requests`, `dry_errors`, `dry_builtins`, `dry_async_jobs`

```rust
fn handler(req) {
    rev {"status": 200, "body": "hello"}
}
op(8080, handler, "mul", 100)
// curl localhost:8080/debug/vars  -> metrics JSON
```

---

## 15. Error handling (prod)

- `errfmt.Format` return `DryError{Code,Line,Col,Message}` (unified)
- Format: `line:col [CODE] message` + source pointer

---

## 16. Build & version

```bash
make build        # go build -ldflags "-X drylang/cli.Version=..."
make test         # go test ./tools/tests/...
make race         # go test -race ./tools/tests/...
make vet          # go vet ./...
make release      # build + version stamp

# version stamp
go build -ldflags "-X drylang/cli.Version=v1.2.3" -o dry .
dry --version     # dryLang version v1.2.3
```

CI (`.github/workflows/ci.yml`): Go 1.26, build, vet, `go test ./tools/tests/...`, script tests, **secret scan**.

---

## TODO — yang perlu ditambah/diperbaiki

### Bug / fitur teknis
- [ ] **Closure mutation (upvalues)** — closure yang menangkap variabel local fungsi tidak bisa mutate secara persisten. `fn counter(){ n=0; fn inc(){ n=n+1 }; rev inc }` → `1,1,1` (harus `1,2,3`). Butuh upvalue (heap-shared captured locals) di compiler + VM.
- [ ] **pv inheritance enforcement penuh** — private field parent lewat child instance sudah terdeteksi; parent-private diakses lewat class lain perlu audit.
- [ ] **Exponent overflow** — `1e400` → E110 bad number (ditolak). Kalau mau render `Inf`, ubah ParseFloat handling.

### Sandbox / ops (partial)
- [ ] **Command timeout** — `cmd`/`sys.run` belum punya timeout (hanya allowlist gate). Tambah context timeout.
- [ ] **Graceful shutdown test** — wiring SIGTERM sudah ada, belum teruji di CI (Windows pakai SIGKILL).

### Prod (dari PROD_PROPOSE)
- [ ] **DryError full migration** — errfmt sudah return `DryError`; handler lain masih string-concat `"E300 at %d:%d:"`. Migrasi penuh.
- [ ] **/metrics** — expvar `/debug/vars` aktif; counters Requests/Errors sudah wired, Builtins/AsyncJobs belum (hot path).
- [ ] **Structured logging default** — `DRY_LOG_JSON=1` tersedia; default masih text.
- [ ] **P6 ops sisa** — lint golangci, repo cleanup (binary tracked), release checksum di release.yml.
- [ ] **`ns` namespace** — dibatalkan (collision builtin). Jika butuh module-scoped pv, desain ulang.

### Selesai (verifikasi terakhir 2026-08-13)
- [x] **Arrow function** — `x -> expr`, `(a,b) -> expr`, `x -> {body}`, sebagai call arg
- [x] **Exponent notation** — `1e2`, `2.5e-3`, `1e309` (overflow ditolak)
- [x] **String compare** — `< > <= >=` lexicographic
- [x] **Number overflow** — `99999999999999999999` → 1e20 (tanpa garbage MinInt64)
- [x] **Analyzer strict** — unused source-order, class method analysis, undefined var, pv enforce
- [x] **pv inheritance** — child instance akses parent private terdeteksi
- [x] **pv cl** — private class akses luar error
- [x] **Parent init inherit** — `cl B <- A`, `B(args)` pakai init A
- [x] **Keyword-named methods** — `rt.on`, `pipe.close`, `mem.set` parse OK
- [x] **use "dir"** — load semua file sorted, .dry > .y > lainnya
- [x] **Struct pv parse** — `Point { x; pv hidden }`
- [x] **Block scope** — var baru dalam block = local (tidak bocor)
- [x] **Sandbox** — cmd/db/url terkunci default (env-driven)
- [x] **HTTP hardening** — timeouts, CORS config, graceful shutdown
- [x] **DryError** — errfmt unified type
- [x] **Metrics** — expvar `/debug/vars` + counters
- [x] **CI** — secret scan, Go 1.26, vet, test suite
- [x] **Version stamping** — `-ldflags -X drylang/cli.Version=`
- [x] **Makefile**
