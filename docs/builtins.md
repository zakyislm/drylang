# Built-in Functions

dryLang provides 27 built-in functions available without any imports. All function names are 1–4 characters.

## I/O

### `pt` — Print

Outputs a value to stdout, followed by a newline.

```
pt "Hello, World!"
pt 42
pt [1, 2, 3]
pt {"key": "value"}
```

`pt` is a **statement**, not a function. It does not return a value.

### `in` — Input

Reads a line of text from stdin. Optionally displays a prompt.

```
.Without prompt.
value in()

.With prompt.
name in("Enter your name: ")
pt "Hello, " + name
```

`in()` always returns a string. Use `num()` to convert to a number:

```
age_str in("Age: ")
age num(age_str)
```

### `r` — Read File

Reads the entire contents of a file as a string.

```
content r("data.txt")
pt content
```

Returns the file content as a string. Throws an error if the file doesn't exist.

### `w` — Write File

Writes a string to a file. Creates the file if it doesn't exist, overwrites if it does.

```
w("output.txt", "Hello, File!")
w("data.json", `{"key": "value"}`)
```

Returns `t` on success. Throws an error on failure.

---

## Type Conversion

### `num` — To Number

Converts a string to a number.

```
x num("42")       .42.
y num("3,14")     .3.14. .note: follows dryLang comma decimal.
```

Throws an error if the string is not a valid number.

### `str` — To String

Converts any value to its string representation.

```
s str(42)          ."42".
s str(t)           ."t".
s str([1, 2])      ."[1, 2]".
```

### `get` — Type Check

Returns the type of a value as a string.

```
pt get(42)          ."number".
pt get("hello")     ."string".
pt get(t)           ."bool".
pt get([1, 2])      ."array".
pt get({"a": 1})    ."map".
pt get(greet)       ."fn".
pt get(unknown)     ."unknown".
```

---

## Math

### `abs` — Absolute Value

```
pt abs(-42)    .42.
pt abs(42)     .42.
```

### `min` — Minimum

Returns the smaller of two numbers.

```
pt min(10, 20)    .10.
pt min(-5, 5)     .-5.
```

### `max` — Maximum

Returns the larger of two numbers.

```
pt max(10, 20)    .20.
pt max(-5, 5)     .5.
```

### `rnd` — Round

Rounds a number to the nearest integer.

```
pt rnd(3,7)     .4.
pt rnd(3,2)     .3.
pt rnd(2,5)     .3. .rounds up at .5.
```

### `ran` — Random

Returns a random float between 0.0 (inclusive) and 1.0 (exclusive).

```
pt ran()           .e.g., 0.7291....
random_int rnd(ran() * 100)    .random 0-100.
```

---

## String Functions

### `cap` — Capitalize (Uppercase)

Converts a string to UPPERCASE.

```
pt cap("hello")    .HELLO.
pt cap("World")    .WORLD.
```

### `low` — Lowercase

Converts a string to lowercase.

```
pt low("HELLO")    .hello.
pt low("World")    .world.
```

### `trm` — Trim

Removes leading and trailing whitespace.

```
pt trm("  hello  ")    .hello.
pt trm("\t hi \n")      .hi.
```

### `spl` — Split

Splits a string by a delimiter, returns an array.

```
parts spl("a,b,c", ",")
pt parts    .["a", "b", "c"].

words spl("hello world", " ")
pt words    .["hello", "world"].
```

### `j` — Join

Joins an array of strings with a separator.

```
arr ["a", "b", "c"]

pt j(arr, ",")     .a,b,c.
pt j(arr, " - ")   .a - b - c.
pt j(arr, "")      .abc.
```

### `mod` — Modify (Replace)

Replaces all occurrences of a substring.

```
pt mod("hello world", "world", "earth")    .hello earth.
pt mod("aabaa", "a", "x")                  .xxbxx.
```

### `has` — Contains

Checks if a string contains a substring. Returns `t` or `f`.

```
pt has("hello world", "world")    .t.
pt has("hello world", "xyz")      .f.
```

---

## Array Functions

### `len` — Length

Returns the number of elements in an array, characters in a string, or entries in a map.

```
pt len([1, 2, 3])        .3.
pt len("hello")           .5.
pt len({"a": 1, "b": 2}) .2.
pt len([])                .0.
```

### `add` — Add (Append)

Appends an element to an array. Returns a new array.

```
arr [1, 2, 3]
arr add(arr, 4)
pt arr    .[1, 2, 3, 4].
```

### `sort` — Sort

Returns a sorted copy of an array. Numbers sort numerically, strings sort alphabetically.

```
pt sort([3, 1, 4, 1, 5])         .[1, 1, 3, 4, 5].
pt sort(["cherry", "apple"])      .["apple", "cherry"].
```

### `pop` — Pop Last

Returns the last element of an array.

```
pt pop([1, 2, 3])    .3.
```

Throws an error on empty arrays.

### `rm` — Remove

Removes an element at the given index. Returns a new array.

```
arr [10, 20, 30, 40]
arr rm(arr, 1)
pt arr    .[10, 30, 40].
```

---

## Map Functions

### `key` — Keys

Returns all keys of a map as an array.

```
m {"name": "Zaky", "age": 17}
pt key(m)    .["name", "age"].
```

### `val` — Values

Returns all values of a map as an array.

```
m {"name": "Zaky", "age": 17}
pt val(m)    .["Zaky", 17].
```

---

## Utility

### `q` — Quiet (Sleep)

Pauses execution for the specified number of milliseconds.

```
pt "Starting..."
q(1000)        .wait 1 second.
pt "Done!"

q(500)         .wait 500ms.
```
