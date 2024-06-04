//go:build windows

package runsc

// Consts for supported allocation methods.
const (
	HeapAlloc AllocMethod = iota
	NtAllocateVirtualMemory
	NtCreateSection
	InvalidAlloc
)

func init() {
	allocMethods[HeapAlloc] = allocHeap
	allocMethods[NtAllocateVirtualMemory] = allocStack
	allocMethods[NtCreateSection] = allocSection
}
