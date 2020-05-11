package autoupdate

import "strings"

// Updater represents a type capable of automatically updating a Go binary.
type Updater interface {
	Check(currentTagName string) (string, error)
	Update() error
}

func normalizeSemver(ver string) string {
	return strings.Trim(ver, "v")
}
