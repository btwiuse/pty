//go:build js
// +build js

package pty

import "os"

func open() (pty, tty *os.File, err error) {
	return nil, nil, ErrUnsupported
}
