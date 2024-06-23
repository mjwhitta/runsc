//go:build windows

package runsc

import (
	"golang.org/x/sys/windows"

	"github.com/mjwhitta/errors"
	w32 "github.com/mjwhitta/win/api"
)

func getProcessHandle(s *state) (*state, error) {
	var e error

	// Get process handle
	if s.l.pid == 0 {
		s.proc = windows.CurrentProcess()
	} else {
		s.proc, e = w32.NtOpenProcess(
			s.l.pid,
			w32.Winnt.ProcessAllAccess,
		)
		if e != nil {
			e = errors.Newf("failed to get process handle: %w", e)
			return nil, e
		}
	}

	return s, nil
}

func initState(l *Launcher, sc []byte) (*state, error) {
	var e error
	var s *state = &state{l: l, sz: len(sc)}

	if s, e = getProcessHandle(s); e != nil {
		return nil, e
	}

	return s, nil
}
