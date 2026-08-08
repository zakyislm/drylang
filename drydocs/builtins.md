# Built-in Functions

## I/O
| fn | sig | what |
|----|-----|------|
| `pt` | `pt val` | print + newline |
| `in` | `in("prompt")` | read stdin → string |
| `r` | `r("file.txt")` | read file → string |
| `w` | `w("file.txt", data)` | write file → `t` |

## Type
| fn | sig | what |
|----|-----|------|
| `get` | `get(val)` | → `"number"` / `"string"` / `"bool"` / `"array"` / `"map"` / `"fn"` / `"unknown"` |
| `num` | `num("42")` | string → number |
| `str` | `str(42)` | anything → string |

## Math
| fn | sig | what |
|----|-----|------|
| `abs` | `abs(-5)` | absolute value |
| `min` | `min(a, b)` | smaller of two |
| `max` | `max(a, b)` | larger of two |
| `rnd` | `rnd(3,7)` | round to int |
| `ran` | `ran()` | random 0.0–1.0 |

## String
| fn | sig | what |
|----|-----|------|
| `cap` | `cap("hi")` | → `"HI"` |
| `low` | `low("HI")` | → `"hi"` |
| `trm` | `trm("  hi  ")` | trim whitespace |
| `spl` | `spl("a,b", ",")` | split → array |
| `j` | `j(arr, ",")` | join → string |
| `mod` | `mod(s, "old", "new")` | replace all |
| `has` | `has(s, "sub")` | contains → `t`/`f` |

## Array
| fn | sig | what |
|----|-----|------|
| `len` | `len(x)` | length (string/array/map) |
| `add` | `add(arr, item)` | append → new array |
| `sort` | `sort(arr)` | sorted copy |
| `pop` | `pop(arr)` | last element |
| `rm` | `rm(arr, idx)` | remove at index → new array |

## Map
| fn | sig | what |
|----|-----|------|
| `key` | `key(map)` | all keys → array |
| `val` | `val(map)` | all values → array |

## Utility
| fn | sig | what |
|----|-----|------|
| `q` | `q(1000)` | sleep N ms |
