// Package GPG implements Go bindings to the GPG command line tool.
package gpg

import (
	"git.missionfocus.com/ours/code/libraries/go/multierror"
	"io"
	"os/exec"
)

// Imports the bytes from r as a PGP key with a passphrase.
func ImportWithPassphrase(r io.Reader, passphrase string) ([]byte, error) {
	errCount := 3
	hasPassphrase := passphrase != ""
	args := []string{"--import", "--batch"}

	if hasPassphrase {
		args = append(args, "--passphrase-fd", "0")
		errCount++
	}

	cmd := exec.Command("gpg", args...)
	pipe, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	errs := make(chan error, errCount)
	go func() {
		if hasPassphrase {
			_, err := pipe.Write([]byte(passphrase + "\n"))
			errs <- err
		}
		_, err := io.Copy(pipe, r)
		errs <- err
		errs <- pipe.Close()
	}()

	output, err := cmd.CombinedOutput()
	errs <- err
	return output, multierror.Collect(errCount, errs)
}

// Imports the bytes from r as a PGP key.
func Import(r io.Reader) ([]byte, error) {
	return ImportWithPassphrase(r, "")
}
