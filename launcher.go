package runsc

import "github.com/mjwhitta/errors"

// Launcher is a struct containing configuration options for running
// shellcode.
type Launcher struct {
	alloc   AllocMethod
	pid     uint32
	run     RunMethod
	suspend bool
	write   WriteMethod
}

// New will return a pointer to a new Launcher instance.
func New() *Launcher {
	return &Launcher{}
}

// AllocVia will set the Windows API to be used for allocating memory.
func (l *Launcher) AllocVia(a AllocMethod) *Launcher {
	l.alloc = a
	return l
}

// Exe will execute the shellcode using the Launcher as it's
// configured.
func (l *Launcher) Exe(sc []byte) error {
	var e error
	var s *state

	if e = l.validate(sc); e != nil {
		return e
	}

	if s, e = initState(l, sc); e != nil {
		return e
	}

	if s, e = allocMethods[l.alloc](s); e != nil {
		return e
	}

	if s, e = writeMethods[l.write](s, sc); e != nil {
		return e
	}

	if _, e = runMethods[l.run](s); e != nil {
		return e
	}

	return nil
}

// InPID will set the PID, causing the Launcher to execute shellcode
// within the specified process.
func (l *Launcher) InPID(pid uint32) *Launcher {
	l.pid = pid
	return l
}

// InSelf will set the PID to 0, causing the Launcher to execute
// shellcode within its own process.
func (l *Launcher) InSelf() *Launcher {
	l.pid = 0
	return l
}

// RunVia will set the Windows API to be used for running shellcode.
func (l *Launcher) RunVia(r RunMethod) *Launcher {
	l.run = r
	return l
}

// Suspend will set new threads to be suspended when running
// shellcode.
func (l *Launcher) Suspend(s ...bool) *Launcher {
	if len(s) == 0 {
		l.suspend = true
	} else {
		l.suspend = s[0]
	}

	return l
}

func (l *Launcher) validate(sc []byte) error {
	if len(sc) == 0 {
		return errors.New("no shellcode provided")
	}

	if _, ok := allocMethods[l.alloc]; !ok {
		return errors.New("invalid allocation method provided")
	}

	if _, ok := writeMethods[l.write]; !ok {
		return errors.New("invalid write method provided")
	}

	if _, ok := runMethods[l.run]; !ok {
		return errors.New("invalid run method provided")
	}

	return nil
}

// WriteVia will set the Windows API to be used for writing shellcode
// to memory.
func (l *Launcher) WriteVia(w WriteMethod) *Launcher {
	l.write = w
	return l
}
