//go:build !windows

package runsc

func initState(l *Launcher, _ []byte) (*state, error) {
	return &state{l: l}, nil
}
