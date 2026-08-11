# Changelog

All notable changes to dryLang will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added
- Native HTTP server support via `op()` function (powered by Go `net/http`).
- Native SQL database support via `db()` function (SQLite, MySQL, PostgreSQL).
- Advanced math functions via `math()` (`sqrt`, `pow`, `ceil`, `floor`, `sin`, `cos`, `tan`, `log`, `log10`).
- System builtins: `now()`, `date()`, `arg()`, `env()`, `cmd()`, `dir()`, `del()`, `die()`.
- Networking & JSON builtins: `req()`, `json()`.
- Support for both `.y` and `.dry` file extensions.
- Comprehensive documentation overhaul (`docs/`).
- Standardized open-source files (`LICENSE`, `CHANGELOG.md`, `CONTRIBUTING.md`).
- GitHub Actions CI/CD workflows for testing and release binaries.
- Added `DICTIONARY.md`: a complete, 100% exhaustive reference mapping of keywords, built-ins, operators, delimiters, and error codes directly to the compiler pipeline.

### Changed
- Converted comment syntax from custom dot syntax (`. comment .`) to standard C-style (`//` and `/* */`).
- Re-architected documentation to follow the Write Less, Get More philosophy.
- Refactored `dryLang` website to be fully responsive for mobile devices, including hamburger menu navigation.
- Moved website layout configurations from inline React styles to robust CSS media queries.

### Fixed
- Fixed critical VM closure panics when resolving variables in `async` blocks escaping the stack.
- Fixed variable shadowing and resolution logic inside nested closures.
- Fixed polymorphic behavior for `rm()` built-in to correctly support `os.Remove` for file paths alongside array index removal.
- Fixed JSON stringify array output support.
- Test suite stabilized to 100% pass rate (75/75 test cases passing).
- Fixed 10 builtins that were registered in the compiler but missing implementation in the VM.
- Fixed mobile rendering issues including overflowing text, incorrect footer padding, and dropdown menu clipping in the documentation sidebar.
