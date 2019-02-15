// Package multierror implements utility functions for handling processes that may produce multiple errors concurrently.
package multierror

import "strings"

// A single Error representing multiple errors.
type Error []error

// Formats a MultiError as a list of errors.
func (err Error) Error() string {
	var sb strings.Builder

	sb.WriteString("errors:\n")
	for _, e := range err {
		if err != nil {
			sb.WriteString("\t* ")
			sb.WriteString(e.Error())
			sb.WriteRune('\n')
		}
	}

	return sb.String()
}

// Collects `count` errors from channel `errs` then creates a single MultiError from them.
func Collect(count int, errs <-chan error) error {
	var multi Error = make([]error, 0)

	for i := 0; i < count; i++ {
		if err := <-errs; err != nil {
			multi = append(multi, err)
		}
	}

	nErrors := len(multi)
	if nErrors == 0 {
		return nil
	} else if nErrors == 1 {
		return multi[0]
	} else {
		return multi
	}
}
