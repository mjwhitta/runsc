package runsc

func (l *Launcher) exe(sc []byte) error {
	var e error
	var s *state = &state{l: l, sz: len(sc)}

	if s, e = getProcessHandle(s); e != nil {
		return e
	}

	if s, e = allocate(s); e != nil {
		return e
	}

	if s, e = write(s, sc); e != nil {
		return e
	}

	if _, e = run(s); e != nil {
		return e
	}

	return nil
}
