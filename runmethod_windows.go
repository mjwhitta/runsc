//go:build windows

package runsc

// Consts for supported run methods.
const (
	RtlCreateUserThread RunMethod = iota
	NtQueueApcThread
	NtQueueApcThreadEx
	InvalidRun
)

func init() {
	runMethods[RtlCreateUserThread] = runUserThread
	runMethods[NtQueueApcThread] = runApcThread
	runMethods[NtQueueApcThreadEx] = runApcThreadEx
}
