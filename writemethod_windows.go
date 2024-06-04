//go:build windows

package runsc

// Consts for supported copy methods.
const (
	NtWriteVirtualMemory WriteMethod = iota
	InvalidWrite
)

func init() {
	writeMethods[NtWriteVirtualMemory] = writeMem
}
