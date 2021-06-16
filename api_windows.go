package runsc

import (
	"unsafe"

	"golang.org/x/sys/windows"

	hl "gitlab.com/mjwhitta/hilighter"
)

var ntdll *windows.LazyDLL

func init() {
	// Load DLLs
	ntdll = windows.NewLazySystemDLL("ntdll")
}

// NtAllocateVirtualMemory from ntdll.
func NtAllocateVirtualMemory(
	pHndl windows.Handle,
	size uint64,
	allocType uintptr,
	protection uintptr,
) (addr uintptr, e error) {
	var err uintptr

	err, _, _ = ntdll.NewProc("NtAllocateVirtualMemory").Call(
		uintptr(pHndl),
		uintptr(unsafe.Pointer(&addr)),
		0,
		uintptr(unsafe.Pointer(&size)),
		allocType,
		protection,
	)
	if err != 0 {
		e = hl.Errorf(
			"NtAllocateVirtualMemory returned %0x",
			uint32(err),
		)
	} else if addr == 0 {
		e = hl.Errorf(
			"NtAllocateVirtualMemory failed for unknown reason",
		)
	} else {
		e = nil
	}

	// WTF?! Why is a Printf needed?! time.Sleep() doesn't work?
	// Oh well, print newline and escape sequence for "go up 1 line"
	hl.Printf("\n\x1b[1A")

	return
}

// NtCreateSection from ntdll.
func NtCreateSection(
	sHndl *windows.Handle,
	access uintptr,
	size uint64,
	pagePerms uintptr,
	secPerms uintptr,
) error {
	var err uintptr

	err, _, _ = ntdll.NewProc("NtCreateSection").Call(
		uintptr(unsafe.Pointer(sHndl)),
		access,
		0,
		uintptr(unsafe.Pointer(&size)),
		pagePerms,
		secPerms,
		0,
	)
	if err != 0 {
		return hl.Errorf("NtCreateSection returned %0x", uint32(err))
	} else if *sHndl == 0 {
		return hl.Errorf("NtCreateSection failed for unknown reason")
	}

	return nil
}

// NtMapViewOfSection from ntdll.
func NtMapViewOfSection(
	sHndl windows.Handle,
	pHndl windows.Handle,
	size uint64,
	inheritPerms uintptr,
	pagePerms uintptr,
) (scBase uintptr, e error) {
	var err uintptr
	var scOffset uintptr

	err, _, _ = ntdll.NewProc("NtMapViewOfSection").Call(
		uintptr(sHndl),
		uintptr(pHndl),
		uintptr(unsafe.Pointer(&scBase)),
		0,
		0,
		uintptr(unsafe.Pointer(&scOffset)),
		uintptr(unsafe.Pointer(&size)),
		inheritPerms,
		0,
		pagePerms,
	)
	if err != 0 {
		e = hl.Errorf("NtMapViewOfSection returned %0x", uint32(err))
	} else if scBase == 0 {
		e = hl.Errorf("NtMapViewOfSection failed for unknown reason")
	} else {
		e = nil
	}

	return
}

// NtOpenProcess from ntdll.
func NtOpenProcess(
	pid uint32,
	access uintptr,
) (pHndl windows.Handle, e error) {
	var err uintptr

	err, _, _ = ntdll.NewProc("NtOpenProcess").Call(
		uintptr(unsafe.Pointer(&pHndl)),
		access,
		uintptr(unsafe.Pointer(&objectAttrs{0, 0, 0, 0, 0, 0})),
		uintptr(unsafe.Pointer(&clientID{uintptr(pid), 0})),
	)
	if err != 0 {
		e = hl.Errorf("NtOpenProcess returned %0x", uint32(err))
	} else if pHndl == 0 {
		e = hl.Errorf("NtOpenProcess failed for unknown reason")
	} else {
		e = nil
	}

	return
}

// NtQueueApcThread from ntdll.
func NtQueueApcThread(
	tHndl windows.Handle,
	apcRoutine uintptr,
) error {
	var err uintptr

	err, _, _ = ntdll.NewProc("NtQueueApcThread").Call(
		uintptr(tHndl),
		apcRoutine,
		0, // arg1
		0, // arg2
		0, // arg3
	)
	if err != 0 {
		return hl.Errorf(
			"NtQueueApcThread returned: %0x",
			uint32(err),
		)
	}

	return nil
}

// NtQueueApcThreadEx from ntdll.
func NtQueueApcThreadEx(
	tHndl windows.Handle,
	apcRoutine uintptr,
) error {
	var err uintptr

	err, _, _ = ntdll.NewProc("NtQueueApcThreadEx").Call(
		uintptr(tHndl),
		0x1, // userApcReservedHandle
		apcRoutine,
		0, // arg1
		0, // arg2
		0, // arg3
	)
	if err != 0 {
		return hl.Errorf(
			"NtQueueApcThreadEx returned: %0x",
			uint32(err),
		)
	}

	return nil
}

// NtResumeThread from ntdll.
func NtResumeThread(tHndl windows.Handle) error {
	var err uintptr

	err, _, _ = ntdll.NewProc("NtResumeThread").Call(
		uintptr(tHndl),
		0, // previousSuspendCount
	)
	if err != 0 {
		return hl.Errorf(
			"NtResumeThread returned: %0x",
			uint32(err),
		)
	}

	return nil
}

// NtWriteVirtualMemory from ntdll.
func NtWriteVirtualMemory(
	pHndl windows.Handle,
	dst uintptr,
	b []byte,
) error {
	var err uintptr

	err, _, _ = ntdll.NewProc("NtWriteVirtualMemory").Call(
		uintptr(pHndl),
		dst,
		uintptr(unsafe.Pointer(&b[0])),
		uintptr(len(b)),
	)
	if err != 0 {
		return hl.Errorf(
			"NtWriteVirtualMemory returned %0x",
			uint32(err),
		)
	}

	return nil
}

// RtlCreateUserThread from ntdll.
func RtlCreateUserThread(
	pHndl windows.Handle,
	addr uintptr,
	sspnd bool,
) (tHndl windows.Handle, e error) {
	var err uintptr
	var suspend uintptr

	if sspnd {
		suspend = 1
	}

	err, _, _ = ntdll.NewProc("RtlCreateUserThread").Call(
		uintptr(pHndl),
		0,
		suspend,
		0,
		0,
		0,
		addr,
		0,
		uintptr(unsafe.Pointer(&tHndl)),
		0,
	)
	if err != 0 {
		e = hl.Errorf("RtlCreateUserThread returned %0x", uint32(err))
	} else if tHndl == 0 {
		e = hl.Errorf("RtlCreateUserThread failed for unknown reason")
	} else {
		e = nil
	}

	return
}
