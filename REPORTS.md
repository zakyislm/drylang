# dryLang Syntax Analysis Report

This report outlines edge cases, parser limitations, and unsupported syntaxes discovered during massive-scale test generation. Re-evaluated against the official `FEATURES.md` and `DICTIONARY.md`.

## Part A: Confirmed Parser Limitations / Bugs
These are features that appear to break despite being validly within standard parsing expectations.

### 1. Multi-line Map Literal Parsing (FIXED)
- **Error:** `[E109] illegal token \n in prefix`
- **Context:** Defining maps spanning multiple lines (e.g., `cns ERROR_CODES = {\n "400": ... }`).
- **Resolution:** The parser chokes on newlines immediately following `{`. Maps must be written on a single line.
- **Fix:** Added `SkipSemicolons()` calls in `ParseMapLit` (`parser/expr/map.go`) to consume newlines after braces and commas.

### 2. String Interpolation Property Access (FIXED)
- **Error:** `[E109] illegal .`
- **Context:** Accessing object properties within string interpolation: `pt("ID: ${this.id}")`.
- **Resolution:** Interpolation evaluates standard math (`${1+2}`) but fails on DotExpr `.` parsing. Use string concatenation: `"ID: " + str(this.id)`.
- **Fix:** Added `TOKEN_DOT` handling to `nextInterpToken` (`lexer/read_string.go`).

### 3. String Interpolation Nested Quote Lexing (FIXED)
- **Error:** `[E109] illegal \` or `[E109] illegal \``
- **Context:** Trying to use quotes (`\"`) or raw strings (`` ` ``) inside an interpolation block, e.g. `pt("Hello ${ \"MR \" + name }")`.
- **Resolution:** The lexer completely breaks when encountering nested quotes or backticks inside `${ ... }`. 
- **Recommendation:** Fix the lexer state machine to track quote contexts properly when inside string interpolation blocks.
- **Fix:** Registered string literals (`"`, `'`, ``` ` ```) inside `nextInterpToken` (`lexer/read_string.go`).

### 4. Dynamic Method Reassignment (FIXED)
- **Error:** None (Silent failure)
- **Context:** Assigning a closure to a class instance method: `d.bark = () -> { pt("Meow") }`.
- **Resolution:** The compiler evaluates this, but because methods are bound to prototypes (not the specific instance fields map initially), calling `d.bark()` still prints `"Woof"`. The closure is discarded.
- **Fix:** Tracked method names in `a.classMethods` and blocked method reassignment in `analyzeDotAssign` (`analyzer/stmt_assign.go`).

### 5. Builtin Function Shadowing (FIXED)
- **Error:** None (Silent shadowing)
- **Context:** `len = 10`
- **Resolution:** Can overwrite core builtins globally.
- **Fix:** Added check in `analyzeAssign` (`analyzer/stmt_vars.go`) against `core.BuiltinNames`. Throws `E109 cannot assign to builtin`.

### 6. Constants (`cns`) Mutable (FIXED)
- **Error:** None (Silent overwrite)
- **Context:** `cns X = 1; X = 2`
- **Resolution:** Constants are parsed, but the compiler does not emit immutability checks, allowing them to act exactly like `let`.
- **Fix:** Checked `info.isConst` in `analyzeAssign` (`analyzer/stmt_vars.go`). Throws `E109 cannot assign to cns`.

### 7. `pv cl` Unusable [FIXED]
- **Context:** `pv cl Secret { ... }`.
- **Resolution:** Removed the overly restrictive cross-file block in `analyzer/expr.go` since dryLang merges AST globally, making instantiation public but internal fields/methods safe.

### 8. Private Class (`pv cl`) Inheritance Leak (FIXED)
- **Error:** None (Bypass of `pv cl`)
- **Context:** `pv cl Base { } cl Public <- Base { }`
- **Resolution:** While the compiler blocks direct instantiation of `Base()`, it fails to prevent `Public` from inheriting from `Base`. The developer can simply instantiate `Public()`, entirely bypassing the `pv` lock and exposing all of `Base`'s logic.
- **Fix:** Checked parent visibility in `analyzeClass` (`analyzer/stmt_class.go`). Throws `E109 cannot inherit from pv cl` if a public class extends a `pv cl`.

### 9. Silent Import Failure in Block Scopes (FIXED)
- **Error:** None (Silent failure leading to runtime `undefined` errors)
- **Context:** Placing a `use "file.dry"` statement inside a block scope, such as a function or `if` statement.
- **Resolution:** The parser successfully accepts the `use` statement anywhere, but the AST evaluator completely ignores it if it is inside a block scope. The file is never loaded, but no compilation error is thrown, leading to extremely confusing `undefined` identifier errors later in the block.
- **Recommendation:** The compiler should either throw a parse error (`[E109] use must be at top level`) or properly evaluate imports dynamically within the block scope.
- **Fix:** Added block depth tracking to `ParserCore`. Throws `E109` if `use` is parsed inside a block (`parser/stmt/use.go`).

### 10. Duplicate Function Parameters (FIXED)
- **Error:** None (Silent shadowing)
- **Context:** Declaring a function with duplicate parameter names: `fn test(x, x) { rev x }`.
- **Resolution:** The compiler allows duplicate parameter names. The later parameter silently shadows the earlier one.
- **Fix:** Added duplicate parameter identifier check in `analyzeFnDecl` (`analyzer/stmt_fn.go`). Throws `E109 duplicate argument name`.

### 11. Arrow Functions Do Not Capture `this` [FIXED]
- **Context:** Using an arrow function inside a method, e.g., `arr.map(el -> this.doSomething(el))`.
- **Resolution:** Arrow functions were missing `this` bindings. Fixed by modifying `OpClosure` and the compiler to properly pass `LocalNames` from `ClassMethod` to the bound `CompiledFn`, allowing the arrow function's `OpClosure` to dynamically inherit `this` from the parent environment.

### 12. Circular Class Inheritance Permitted (FIXED)
- **Error:** `Fatal VM Crash (Stack Overflow)` or silent hangs.
- **Context:** `cl A <- B {}; cl B <- A {};` or `cl A <- A {}`
- **Resolution:** The analyzer does not build an inheritance graph or detect cycles, which causes an infinite loop during instance initialization and property resolution.
- **Fix:** Added circular inheritance detection inside `analyzeClass` (`analyzer/stmt_class.go`). Throws `E109 circular inheritance`.

### 13. Class Definitions are Fully Mutable (FIXED)
- **Error:** None (Silent overwrite)
- **Context:** `cl App {} ; App = 5;`
- **Resolution:** The global definition of `App` is simply reassigned to a number.
- **Fix:** Tracked class definitions and builtins in `analyzeAssign` (`analyzer/stmt_vars.go`). Throws `E109 cannot assign to class`.

### 14. Fatal VM Crash: `pop()` on Empty Array [FIXED]
- **Context:** Calling the built-in `pop()` function on an empty array: `pop([])`.
- **Resolution:** Confirmed that `handler/core/pop.go` already correctly implements the `len(arr) == 0` check and avoids slicing issues. No native panic occurs.

### 15. Stack Overflow on Recursive Print / JSON Stringify [FIXED]
- **Context:** Cyclic references, e.g., `a = {}; a.self = a; pt(a)` or `json(a)`.
- **Resolution:** Added a `visited map[uintptr]bool` tracker to `core/value.go` stringifier and `handler/data/json.go` JSON converter. Circular arrays and maps now safely format as `[Circular]` or `{Circular}` instead of blowing up the call stack.

---

## Part B: Unsupported Syntaxes (Working as Intended)
These were initially flagged as bugs but are verified against `FEATURES.md` as **intentional language design choices** or strict unsupported syntaxes.

### 3. Inline Anonymous `fn` (Unsupported)
- **Context:** Trying to pass `fn() { ... }` as an argument.
- **Status:** **Not a bug.** The docs explicitly dictate Arrow Functions (`->`) for expressions and closures. Standard `fn` is exclusively for statements.

### 4. Implicit `i` Index in `mul` (Unsupported)
- **Context:** Trying to use `i` within a `mul` parallel execution.
- **Status:** **Not a bug.** Unlike `lp` which injects `i`, `mul` does not inject context.

### 5. Strict Unused Analyzer Checks (Intentional)
- **Context:** `[E300] unused var`.
- **Status:** **Not a bug.** Core characteristic of dryLang to enforce excellent code hygiene.

### 6. Trailing Commas & Destructuring (Unsupported)
- **Context:** `arr = [1, 2,]` and `a, b = [1, 2]`.
- **Status:** **Not a bug.** dryLang intentionally opts for a minimal grammar. Trailing commas and multiple assignment are explicitly not supported.

### 7. Ternary Conditional Operator (Unsupported)
- **Context:** `val = cond ? t : f`.
- **Status:** **Not a bug.** The language does not have `? :`. It uses `??` for nullish coalescing and the `unknown` literal for unresolvable boolean states.

### 8. Single Letter `t` and `f` Collisions (Intentional)
- **Context:** `f = obj.method` throws `invalid assignment target: *ast.BoolLit = *ast.DotExpr`.
- **Status:** **Not a bug.** `FEATURES.md` dictates that `t` and `f` are strictly reserved as boolean true/false literals. Developers cannot use them as variable names.
