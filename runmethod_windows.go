package runsc

// Consts for supported run methods.
const (
	RtlCreateUserThread RunMethod = iota
	NtQueueApcThread
	NtQueueApcThreadEx
	InvalidRun
)
