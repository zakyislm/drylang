.dryLang — test program.

. Variables .
name "Zaky"
age 17
batteryhealth 89,5

. Print .
pt "=== dryLang ==="
pt name
pt age
pt batteryhealth

. Constant .
cns pi 3,14
MAXLIFE 5
pt pi
pt MAXLIFE

. Boolean .
active t
offline f
pt active
pt offline

. Array .
colors = ["red", "green", "blue"]
pt colors
pt len(colors)

. Map .
config = {"host": "localhost", "port": 8080}
pt config

. Arithmetic .
x = 10
y = 3
pt x + y
pt x - y
pt x * y
pt x / y
pt x % y

. String built-ins .
msg = "  Hello World  "
pt trm(msg)
pt cap("hello")
pt low("HELLO")
pt has("hello world", "world")

. Array built-ins .
nums = [3, 1, 2]
pt sort(nums)

. Math .
pt abs(-42)
pt min(10, 20)
pt max(10, 20)
pt rnd(3,7)

. Type check .
pt get(name)
pt get(age)
pt get(active)
pt get(colors)

. If-elif-el .
score = 85
if score >= 90 {
  pt "A"
} elif score >= 80 {
  pt "B"
} el {
  pt "C"
}

. Loop .
pt "--- loop ---"
lp 3 {
  pt i
}

. Function .
fn greet(who) {
  rev "Hello, " + who + "!"
}
result = greet("World")
pt result

. Try-err .
try {
  pt "trying..."
  err "boom!"
} err(e) {
  pt "caught: " + e
}

. Done! .
pt "=== all tests passed ==="
