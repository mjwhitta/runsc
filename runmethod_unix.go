//go:build !windows

package runsc

// Consts for supported run methods.
const (
	InvalidRun RunMethod = iota
)

var runMethods map[RunMethod]rFunc = map[RunMethod]rFunc{}
