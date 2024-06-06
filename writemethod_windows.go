//go:build windows

package runsc

// Consts for supported copy methods.
const (
	NtWriteVirtualMemory WriteMethod = iota
	InvalidWrite
)

var writeMethods map[WriteMethod]wFunc = map[WriteMethod]wFunc{
	NtWriteVirtualMemory: writeMem,
}
