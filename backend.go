package gol

import "fmt"

type backend string

const (
	BACKEND_CONSOLE backend = "console"
	BACKEND_FILE    backend = "file"
	BACKEND_SYSLOG  backend = "syslog"
)

func backendByName(name string) (backend, error) {
	switch (name) {
	case string(BACKEND_CONSOLE):
		return BACKEND_CONSOLE, nil
	case string(BACKEND_FILE):
		return BACKEND_FILE, nil
	case string(BACKEND_SYSLOG):
		return BACKEND_SYSLOG, nil
	}
	return "", fmt.Errorf("unknown backend '%s', try 'console', 'file' or 'syslog'", name)
}
