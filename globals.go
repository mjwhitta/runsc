package runsc

type clientID struct {
	UniqueProcess uintptr
	UniqueThread  uintptr
}

type objectAttrs struct {
	Length                   uintptr
	RootDirectory            uintptr
	ObjectName               uintptr
	Attributes               uintptr
	SecurityDescriptor       uintptr
	SecurityQualityOfService uintptr
}

// Why are these not defined in the windows pkg?!

// ProcessAllAccess is the PROCESS_ALL_ACCESS const from winnt.h
const ProcessAllAccess = 0x1fffff

// SecCommit is the SEC_COMMIT const from winnt.h
const SecCommit = 0x08000000

// SectionWrite is the SECTION_MAP_WRITE const from winnt.h
const SectionWrite = 0x2

// SectionRead is the SECTION_MAP_READ const from winnt.h
const SectionRead = 0x4

// SectionExecute is the SECTION_MAP_EXECUTE const from winnt.h
const SectionExecute = 0x8

// SectionRWX is the combination of READ, WRITE, and EXECUTE
const SectionRWX = SectionWrite | SectionRead | SectionExecute

// Suspended can be set to true to suspend created threads
var Suspended = false

// Version is the package version
const Version = "1.0.10"
