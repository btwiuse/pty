//go:build !linux && !darwin && !freebsd && !dragonfly && !netbsd && !openbsd && !solaris && !zos && !js
// +build !linux,!darwin,!freebsd,!dragonfly,!netbsd,!openbsd,!solaris,!zos,!js

package pty

import (
	"os"
)

func open() (pty, tty *os.File, err error) {
	return nil, nil, ErrUnsupported
}
