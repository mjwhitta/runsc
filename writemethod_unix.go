//go:build !windows

package runsc

// Consts for supported write methods.
const (
	InvalidWrite WriteMethod = iota
)

var writeMethods map[WriteMethod]wFunc = map[WriteMethod]wFunc{}
