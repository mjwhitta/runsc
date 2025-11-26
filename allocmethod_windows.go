//go:build windows

package runsc

import (
	"golang.org/x/sys/windows"

	"github.com/mjwhitta/errors"
	w32 "github.com/mjwhitta/win/api"
)

// Consts for supported allocation methods.
const (
	HeapAlloc AllocMethod = iota + 1
	NtAllocateVirtualMemory
	NtCreateSection
)

var allocMethods map[AllocMethod]aFunc = map[AllocMethod]aFunc{
	HeapAlloc:               allocHeap,
	NtAllocateVirtualMemory: allocStack,
	NtCreateSection:         allocSection,
}

func allocHeap(s *state) (*state, error) {
	var e error

	if s.l.pid != 0 {
		e = errors.New("cannot allocate via Heap in remote process")
		return nil, e
	}

	s.heap, e = w32.HeapCreate(
		w32.Winnt.HeapCreateEnableExecute,
		0,
		0, // Can grow as needed
	)
	if e != nil {
		e = errors.Newf("failed to create memory: %w", e)
		return nil, e
	}

	s.addr, e = w32.HeapAlloc(
		s.heap,
		w32.Winnt.HeapZeroMemory,
		uintptr(s.sz),
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
		s.sz,
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
		s.sz,
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
		s.sz,
		w32.Winnt.MemCommit|w32.Winnt.MemReserve,
		w32.Winnt.PageExecuteReadwrite,
	)
	if e != nil {
		e = errors.Newf("failed to allocate memory: %w", e)
		return nil, e
	}

	return s, nil
}
