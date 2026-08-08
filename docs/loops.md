# Loops

dryLang has a single, universal loop keyword: `lp`.

## Counted Loop

```
lp count {
  .body — executes 'count' times.
}
```

```
lp 5 {
  pt "Hello!"
}
.Prints "Hello!" 5 times.
```

### Loop Counter `i`

Inside every `lp` block, the variable `i` is automatically available as the loop counter (0-indexed):

```
lp 5 {
  pt i
}
.Output: 0, 1, 2, 3, 4.
```

Use `i` in expressions:

```
lp 5 {
  pt "Item ${i}: ${i * 10}"
}
.Output:
  Item 0: 0
  Item 1: 10
  Item 2: 20
  Item 3: 30
  Item 4: 40
.
```

The counter can also be a variable:

```
n 10
lp n {
  pt i
}
```

## Infinite Loop

```
lp {
  .body — runs forever until 'done'.
}
```

```
count 0
lp {
  pt count
  count = count + 1
  if count >= 5 {
    done
  }
}
```

## Break: `done`

`done` immediately exits the innermost loop:

```
lp 100 {
  if i = 10 {
    done    .exit loop when i reaches 10.
  }
  pt i
}
.Prints 0 through 9.
```

## Continue: `con`

`con` skips the rest of the current iteration and moves to the next:

```
lp 10 {
  if i % 2 = 0 {
    con    .skip even numbers.
  }
  pt i
}
.Prints: 1, 3, 5, 7, 9.
```

## Nested Loops

`done` and `con` apply to the **innermost** loop:

```
lp 3 {
  outer i
  lp 3 {
    if i = 1 {
      done    .breaks inner loop only.
    }
    pt "${outer},${i}"
  }
}
```

> **Note:** The automatic `i` counter applies to each loop independently. In nested loops, the inner loop's `i` shadows the outer loop's `i`. Save the outer counter to a separate variable if needed.

## No For-Each

dryLang intentionally does not have a for-each loop. To iterate over collections, use a counted loop with indexing:

```
colors ["red", "green", "blue"]

lp len(colors) {
  pt colors[i]
}
```

## While-Style Loop

Use an infinite loop with a condition:

```
x 100
lp {
  if x <= 0 {
    done
  }
  pt x
  x = x - 10
}
```
