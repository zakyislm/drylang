# Structs

Structs are user-defined data types with named fields.

## Declaration

Declare a struct by naming it followed by field names in `{}`:

```rust
user {
  name
  age
  email
}
```

Or on one line:

```rust
user { name age email }
```

> **No types.** dryLang is dynamically typed — fields accept any value.

## Instantiation

Create an instance by providing the variable name, struct type, and field values:

```rust
user { name age email }

u user {
  name "Zaky"
  age 17
  email "zaky@example.com"
}
```

## Accessing Fields

Use dot notation:

```rust
pt u.name     // Zaky
pt u.age      // 17
pt u.email    // zaky@example.com
```

Or bracket notation:

```rust
pt u["name"]    // Zaky
```

## Modifying Fields

```rust
u.age = 18
u["email"] = "new@example.com"
```

## Structs in Arrays

```rust
player { name score }

players [
  player { name "Zaky" score 100 },
  player { name "Andi" score 85 }
]

// Note: struct instances inside arrays work like maps
// Access via regular array indexing + dot notation:

.lp len(players) {
  pt players[i].name
.}
```

## Under the Hood

Structs are internally represented as maps. A struct instance is a map with field names as keys. This means all map functions work on struct instances:

```rust
user { name age }
u user { name "Zaky" age 17 }

pt key(u)     // shows fields
pt len(u)     // 2 (plus __struct__ meta key)
```

## Struct vs Map

| Feature | Struct | Map |
|---------|--------|-----|
| Declaration | `user { name age }` | — |
| Creation | `u user { name "Zaky" }` | `u {"name": "Zaky"}` |
| Type identity | Yes (`__struct__` meta) | No |
| Dot access | ✅ | ✅ |
| Bracket access | ✅ | ✅ |
| Dynamic keys | ❌ | ✅ |

Use structs when you want a **named type** with a **fixed set of fields**. Use maps when you need **dynamic keys**.

---

Prev : [Collections](collections.md) | Next : [Error Handling](error-handling.md)
<!-- W3Schools-like web docs integration placeholder -->
