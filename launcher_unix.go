//go:build !windows

package runsc

import "github.com/mjwhitta/errors"

func (l *Launcher) exe(sc []byte) error {
	return errors.New("unsupported OS")
}
