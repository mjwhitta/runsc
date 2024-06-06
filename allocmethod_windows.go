//go:build windows

package runsc

// Consts for supported allocation methods.
const (
	HeapAlloc AllocMethod = iota + 1
	NtAllocateVirtualMemory
	NtCreateSection
)

var allocMethods map[AllocMethod]aFunc = map[AllocMethod]aFunc{
	HeapAlloc:               allocHeap,
	NtAllocateVirtualMemory: allocStack,
	NtCreateSection:         allocSection,
}
