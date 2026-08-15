package core

import (
	"os"
	"strings"
)

// Sandbox gates for privileged builtins. Configuration is read once from the
// environment. Default = locked (everything denied) so a script can never
// run shell commands, touch databases, or fetch remote URLs unless the host
// explicitly opts in.
type Sandbox struct {
	allowCmd []string          // command names allowed by `cmd`/`sys.run`
	allowDB  bool              // `db` allowed
	allowURL []string          // URL host prefixes allowed by `req` and `use`
}

var sandboxCfg = loadSandbox()

// GetSandbox returns the process sandbox configuration.
func GetSandbox() Sandbox {
	return sandboxCfg
}

func loadSandbox() Sandbox {
	return Sandbox{
		allowCmd: splitCSV(os.Getenv("DRY_ALLOW_CMD")),
		allowDB:  os.Getenv("DRY_ALLOW_DB") == "1" || strings.EqualFold(os.Getenv("DRY_ALLOW_DB"), "true"),
		allowURL: splitCSV(os.Getenv("DRY_ALLOW_URL")),
	}
}

func splitCSV(s string) []string {
	if strings.TrimSpace(s) == "" {
		return nil
	}
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		if p = strings.TrimSpace(p); p != "" {
			out = append(out, p)
		}
	}
	return out
}

// AllowCmd reports whether running cmd is permitted.
func (s Sandbox) AllowCmd(cmd string) bool {
	if len(s.allowCmd) == 0 {
		return false
	}
	for _, c := range s.allowCmd {
		if c == "*" || c == cmd {
			return true
		}
	}
	return false
}

// AllowDB reports whether database access is permitted.
func (s Sandbox) AllowDB() bool {
	return s.allowDB
}

// AllowURL reports whether fetching a URL with the given host is permitted.
func (s Sandbox) AllowURL(host string) bool {
	if len(s.allowURL) == 0 {
		return false
	}
	for _, h := range s.allowURL {
		if h == "*" || strings.HasPrefix(host, h) {
			return true
		}
	}
	return false
}
