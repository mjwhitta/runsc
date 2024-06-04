//go:build windows

package runsc

import "golang.org/x/sys/windows"

type state struct {
	addr    uintptr
	heap    uintptr
	l       *Launcher
	proc    windows.Handle
	section windows.Handle
	sz      int
	thread  windows.Handle
}
