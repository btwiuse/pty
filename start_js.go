//go:build js
// +build js

package pty

import (
	"os"
	"os/exec"
)

// StartWithSize assigns a pseudo-terminal tty os.File to c.Stdin, c.Stdout,
// and c.Stderr, calls c.Start, and returns the File of the tty's
// corresponding pty.
func StartWithSize(cmd *exec.Cmd, ws *Winsize) (*os.File, error) {
	_ = cmd
	_ = ws
	return nil, ErrUnsupported
}
