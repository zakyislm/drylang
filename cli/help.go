package cli

import "fmt"

// Version is stamped at build time via:
//   go build -ldflags "-X drylang/cli.Version=v1.2.3"
var Version = "1.0.0"

func printHelp() {
	fmt.Print(`dryLang - Writeless, get more.

Usage:
  dry [file|folder|url|github_repo]
  dry init [dir] [template]

Commands:
  init        Scaffold a new project in the specified directory.
              Templates: api, web, crud, fetch, cli, automation, scraper, hello.
              (If no template is specified, creates an empty idx.dry)

Options:
  -h, --help     Show this help message.
  -v, --version  Show version information.

Examples:
  dry main.dry                   # Run a local file
  dry .                        # Run all files in current directory
  dry github.com/user/repo     # Run from a GitHub repo
  dry init . api               # Create a REST API project in current dir
`)
}
