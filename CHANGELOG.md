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
- Comprehensive documentation overhaul (`docs/` and `drydocs/`).
- Standardized open-source files (`LICENSE`, `CHANGELOG.md`, `CONTRIBUTING.md`).
- GitHub Actions CI/CD workflows for testing and release binaries.

### Changed
- Converted comment syntax from custom dot syntax (`. comment .`) to standard C-style (`//` and `/* */`).
- Re-architected documentation to follow the Writeless, get more philosophy.
- Refactored `dryLang` website to be fully responsive for mobile devices, including hamburger menu navigation.
- Moved website layout configurations from inline React styles to robust CSS media queries.

### Fixed
- Fixed 10 builtins that were registered in the compiler but missing implementation in the VM.
- Fixed mobile rendering issues including overflowing text, incorrect footer padding, and dropdown menu clipping in the documentation sidebar.
