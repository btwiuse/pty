//go:build js
// +build js

package pty

import (
	"errors"
	"os"
	"os/exec"
	"testing"
)

func TestOpenUnsupportedJS(t *testing.T) {
	t.Parallel()

	ptmx, tty, err := Open()
	if !errors.Is(err, ErrUnsupported) {
		t.Fatalf("Open() error = %v, want %v", err, ErrUnsupported)
	}
	if ptmx != nil || tty != nil {
		t.Fatalf("Open() = (%v, %v), want (nil, nil)", ptmx, tty)
	}
}

func TestStartUnsupportedJS(t *testing.T) {
	t.Parallel()

	ptmx, err := Start(exec.Command("echo"))
	if !errors.Is(err, ErrUnsupported) {
		t.Fatalf("Start() error = %v, want %v", err, ErrUnsupported)
	}
	if ptmx != nil {
		t.Fatalf("Start() returned %v, want nil", ptmx)
	}
}

func TestWinsizeJS(t *testing.T) {
	t.Parallel()

	f, err := os.CreateTemp("", "pty-js-winsize")
	if err != nil {
		t.Fatalf("CreateTemp() error = %v", err)
	}
	t.Cleanup(func() { _ = os.Remove(f.Name()) })
	t.Cleanup(func() { _ = f.Close() })

	size, err := GetsizeFull(f)
	if err != nil {
		t.Fatalf("GetsizeFull() error = %v", err)
	}
	if size.Rows != 24 || size.Cols != 80 || size.X != 0 || size.Y != 0 {
		t.Fatalf("GetsizeFull() = %+v, want Rows=24 Cols=80 X=0 Y=0", *size)
	}

	want := &Winsize{Rows: 40, Cols: 120, X: 640, Y: 480}
	if err := Setsize(f, want); err != nil {
		t.Fatalf("Setsize() error = %v", err)
	}

	got, err := GetsizeFull(f)
	if err != nil {
		t.Fatalf("GetsizeFull() after Setsize() error = %v", err)
	}
	if *got != *want {
		t.Fatalf("GetsizeFull() after Setsize() = %+v, want %+v", *got, *want)
	}
}
