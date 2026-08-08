# Collections

## Arrays

Arrays are ordered lists of values. They can hold mixed types.

### Creating Arrays

```
numbers [1, 2, 3, 4, 5]
names ["Zaky", "Andi", "Budi"]
mixed [42, "hello", t, [1, 2]]
empty []
```

### Accessing Elements

Use bracket notation with a 0-based index:

```
colors ["red", "green", "blue"]

pt colors[0]    .red.
pt colors[1]    .green.
pt colors[2]    .blue.
```

### Modifying Elements

```
colors[0] = "crimson"
pt colors[0]    .crimson.
```

### Array Functions

| Function | Description | Example |
|----------|-------------|---------|
| `len(arr)` | Get length | `len([1,2,3])` → `3` |
| `add(arr, item)` | Append item | `add(arr, 4)` |
| `sort(arr)` | Sort (returns new) | `sort([3,1,2])` → `[1,2,3]` |
| `pop(arr)` | Get last element | `pop([1,2,3])` → `3` |
| `rm(arr, idx)` | Remove by index | `rm([1,2,3], 1)` → `[1,3]` |

```
nums [3, 1, 4, 1, 5]

pt len(nums)       .5.
pt sort(nums)      .[1, 1, 3, 4, 5].
pt pop(nums)       .5.
pt rm(nums, 0)     .[1, 4, 1, 5].
```

### Iterating Arrays

```
fruits ["apple", "banana", "cherry"]

lp len(fruits) {
  pt "${i}: ${fruits[i]}"
}
.Output:
  0: apple
  1: banana
  2: cherry
.
```

## Maps

Maps are key-value collections. Keys are strings.

### Creating Maps

```
person {"name": "Zaky", "age": 17}
config {"host": "localhost", "port": 8080}
empty_map {}
```

### Accessing Values

Two syntaxes are supported:

```
person {"name": "Zaky", "age": 17}

.Dot access.
pt person.name     .Zaky.

.Bracket access.
pt person["name"]  .Zaky.
```

Dot access is shorter and preferred for known keys. Bracket access is useful for dynamic keys:

```
field "name"
pt person[field]    .Zaky.
```

### Modifying Values

```
person.age = 18
person["email"] = "zaky@email.com"
```

### Map Functions

| Function | Description | Example |
|----------|-------------|---------|
| `len(map)` | Count entries | `len({"a": 1})` → `1` |
| `key(map)` | Get all keys | `key({"a": 1})` → `["a"]` |
| `val(map)` | Get all values | `val({"a": 1})` → `[1]` |
| `has(map, key)` | Check key exists | — |

```
config {"host": "localhost", "port": 8080}

pt len(config)     .2.
pt key(config)     .["host", "port"].
pt val(config)     .["localhost", 8080].
```

### Iterating Maps

```
data {"a": 1, "b": 2, "c": 3}
keys key(data)

lp len(keys) {
  k keys[i]
  pt "${k}: ${data[k]}"
}
```

## Nested Collections

Arrays and maps can be nested:

```
users [
  {"name": "Zaky", "age": 17},
  {"name": "Andi", "age": 20}
]

pt users[0].name       .Zaky.
pt users[1]["age"]     .20.

matrix [[1, 2], [3, 4], [5, 6]]
pt matrix[0][1]        .2.
pt matrix[2][0]        .5.
```
