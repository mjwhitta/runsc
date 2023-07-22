package runsc

// Consts for supported allocation methods.
const (
	NtAllocateVirtualMemory AllocMethod = iota
	NtCreateSection
	InvalidAlloc
)
