//go:build !windows

package runsc

func initState(l *Launcher, sc []byte) (*state, error) {
	return &state{l: l}, nil
}
