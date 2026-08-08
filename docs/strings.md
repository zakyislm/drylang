# Strings

dryLang supports three types of string literals.

## Regular Strings

Enclosed in double quotes (`"`) or single quotes (`'`). Both are identical in behavior — both support escape sequences and interpolation.

```
greeting "Hello, World!"
name 'Zaky'
```

## Escape Sequences

| Escape | Character |
|--------|-----------|
| `\n` | Newline |
| `\t` | Tab |
| `\\` | Backslash |
| `\"` | Double quote |
| `\'` | Single quote |
| `\$` | Literal dollar sign |
| `\0` | Null character |

```
pt "Line 1\nLine 2"
.Output:
  Line 1
  Line 2
.

pt "Column 1\tColumn 2"
pt "She said \"hello\""
pt "Price: \$100"
```

## String Interpolation

Use `${}` to embed expressions inside strings. Any valid dryLang expression can be used.

### Variables

```
name "Zaky"
age 17

pt "My name is ${name}"
pt "${name} is ${age} years old"
```

### Expressions

```
pt "2 + 3 = ${2 + 3}"
pt "Half of 100 is ${100 / 2}"
```

### Array Elements

```
colors ["red", "green", "blue"]
pt "First: ${colors[0]}"
```

### Map Fields

```
config {"host": "localhost", "port": 8080}
pt "Server: ${config.host}"
```

### Nested Access

```
users [{"name": "Zaky"}, {"name": "Andi"}]
pt "First user: ${users[0].name}"
```

## Raw Strings

Enclosed in backticks (`` ` ``). No escape processing or interpolation occurs. Everything is literal.

```
regex_pattern `\d{3}-\d{4}`
file_path `C:\Users\zaky\Documents`
template `Hello ${name}`       .literal "${name}", NOT interpolated.
```

Raw strings can span concepts that would require heavy escaping in regular strings:

```
json_template `{"key": "value", "count": 0}`
sql `SELECT * FROM users WHERE name = 'Zaky'`
```

## String Operations

dryLang provides built-in functions for string manipulation:

```
msg "  Hello, World!  "

.Transform.
pt cap("hello")              .HELLO.
pt low("HELLO")              .hello.
pt trm(msg)                  .Hello, World!.

.Search.
pt has("hello world", "world")   .t.

.Split & Join.
parts spl("a,b,c", ",")     .["a", "b", "c"].
back j(parts, "-")           .a-b-c.

.Replace.
pt mod("hello world", "world", "earth")    .hello earth.

.Length.
pt len("hello")              .5.
```

## String Concatenation

Use the `+` operator:

```
full "Hello" + ", " + "World!"
pt full    .Hello, World!.
```

Non-string values are automatically converted:

```
pt "Score: " + 100    .Score: 100.
pt "Active: " + t     .Active: t.
```

## Converting To/From Strings

```
.Number to string.
s str(42)          ."42".

.String to number.
n num("42")        .42.
```
