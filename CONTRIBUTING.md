# Contributing to dryLang

Thank you for your interest in dryLang! We welcome contributions from the community.

## Design Philosophy (Writeless, get more)
Before contributing, please understand the core philosophy: **Writeless, get more.**
1. **Ultra-Short Names**: All keywords and builtins must be 1 to 4 characters.
2. **Minimalism**: No syntax bloat. Only add features if they solve a common, practical problem.
3. **No External Dependencies in Scripts**: If a feature requires a complex 3rd party package, it should probably be built-in using Go standard library or high-quality Go modules inside the VM.

## Setting Up the Development Environment
1. Clone the repository: `git clone https://github.com/zakyislm/drylang.git`
2. Ensure you have Go 1.22+ installed.
3. Run tests (if applicable) or build: `go build -o dry .`

## How to Contribute

### 1. Reporting Bugs
Please open an issue providing:
- The dryLang script that caused the error.
- The expected vs actual behavior.
- The operating system and version.

### 2. Adding a New Built-in Function
1. Add the constant to `BuiltinNames` in `core/types.go` (ensure it's ≤ 4 chars).
2. Implement the logic in the `executeBuiltin` switch statement inside `vm/vm.go`.
3. Add tests for the new function in the `tests/` directory.
4. Document the function in `docs/builtins.md`.

### 3. Submitting Pull Requests
- Keep PRs focused on a single feature or bug fix.
- Ensure your code follows standard Go formatting (`gofmt`).
- Do not submit code that breaks the 1-4 character keyword rule.

## Code of Conduct
Please treat all contributors with respect. Constructive feedback is welcome; toxicity is not.
