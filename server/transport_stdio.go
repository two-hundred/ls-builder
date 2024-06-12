package server

import (
	"errors"
	"os"
)

type Stdio struct{}

// Read reads from os.Stdin.
// Fulfils the io.Reader interface.
func (Stdio) Read(p []byte) (int, error) {
	return os.Stdin.Read(p)
}

// Write writes to os.Stdout.
// Fulfils the io.Writer interface.
func (Stdio) Write(p []byte) (int, error) {
	return os.Stdout.Write(p)
}

// Close closes os.Stdin and os.Stdout.
// Fulfils the io.Closer interface.
func (Stdio) Close() error {
	return errors.Join(os.Stdin.Close(), os.Stdout.Close())
}
