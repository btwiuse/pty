//go:build js
// +build js

package pty

import (
	"os"
	"sync"
)

var (
	defaultWinsize = Winsize{
		Rows: 24,
		Cols: 80,
	}
	jsWinsizes sync.Map
)

// Winsize describes the terminal size.
type Winsize struct {
	Rows uint16
	Cols uint16
	X    uint16
	Y    uint16
}

// Setsize resizes t to s.
func Setsize(t *os.File, ws *Winsize) error {
	if t == nil || ws == nil {
		return os.ErrInvalid
	}

	jsWinsizes.Store(t.Fd(), *ws)
	return nil
}

// GetsizeFull returns the full terminal size description.
func GetsizeFull(t *os.File) (*Winsize, error) {
	if t == nil {
		return nil, os.ErrInvalid
	}

	if ws, ok := jsWinsizes.Load(t.Fd()); ok {
		size := ws.(Winsize)
		return &size, nil
	}

	size := defaultWinsize
	return &size, nil
}
