# dryLang Reference

## Variables
```
name "Zaky"          .no keyword, space-separated.
name = "Zaky"        .same thing, with =.
cns pi 3,14          .explicit constant.
MAXLIFE 5            .ALL CAPS = auto constant.
pv secret "abc"      .private (not exported).
? status             .unknown bool.
```

## Types
```
42        .number (all float64 internally).
3,14      .float — comma is decimal point.
"hi"      .string (double or single quotes).
`raw`     .raw string — no escape/interpolation.
t         .true.
f         .false.
unknown   .uninitialized/undetermined.
[1,2,3]   .array.
{"k":"v"} .map.
```

## Operators
```
+  -  *  /  %     .arithmetic.
=  !=  <  >  <=  >=  .comparison (= is equality in expressions).
&  |  !            .logic (AND, OR, NOT).
```

`=` at statement start → assignment. `=` inside expression → equality.

## Strings
```
"Hello ${name}"     .interpolation with ${}.
"Line\n2"           .escape sequences: \n \t \\ \" \' \$ \0.
`raw\nstring`       .no escapes, no interpolation.
```

## Comments
```
.single dot-comment.

.
multi-line
comment
.
```

## If
```
if x = 5 { pt "five" }
elif x = 3 { pt "three" }
el { pt "other" }
```

## Switch
```
on(val) {
  1 { pt "one" }
  2 { pt "two" }
}
```

## Loop
```
lp 5 { pt i }       .counted loop, i = 0..4.
lp { pt "forever" } .infinite loop.
done                 .break.
con                  .continue.
```

## Functions
```
fn add(a, b) { rev a + b }   .fn declaration + return.
asn fn fetch() { ... }        .async function.
awt fetch()                    .await.
double -> (x) { rev x * 2 }   .arrow function.
```

No `rev` → returns `unknown`.

## Collections
```
arr[0]         .array index.
map.key        .map dot access.
map["key"]     .map bracket access.
arr[0] = 5     .index assign.
map.key = 5    .dot assign.
```

## Structs
```
user { name age }                    .declare.
u user { name "Zaky" age 17 }       .instantiate.
pt u.name                            .access.
```

## Error Handling
```
try { err "boom" } err(e) { pt e }  .try-catch.
err "fatal"                          .throw (outside try = crash).
```

## Modules
```
use "helpers"   .loads helpers.y.
pv fn x() {}    .private — not exported.
```

## Truthiness
Falsy: `f`, `0`, `""`, `[]`, `{}`, `unknown`. Everything else truthy.
