//go:build !windows

package runsc

// Consts for supported write methods.
// const (
// 	InvalidWrite WriteMethod = iota + 1
// )

var writeMethods map[WriteMethod]wFunc = map[WriteMethod]wFunc{}
