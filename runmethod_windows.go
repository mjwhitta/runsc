//go:build windows

package runsc

// Consts for supported run methods.
const (
	CertEnumPhysicalStore RunMethod = iota + 1
	CertEnumSystemStore
	NtCreateThreadEx
	NtQueueApcThread
	NtQueueApcThreadEx
	RtlCreateUserThread
)

var runMethods map[RunMethod]rFunc = map[RunMethod]rFunc{
	CertEnumPhysicalStore: runCertPhysical,
	CertEnumSystemStore:   runCertSystem,
	NtCreateThreadEx:      runCreateThreadEx,
	NtQueueApcThread:      runApcThread,
	NtQueueApcThreadEx:    runApcThreadEx,
	RtlCreateUserThread:   runUserThread,
}
