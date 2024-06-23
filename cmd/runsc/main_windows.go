//go:build windows

package main

import "github.com/mjwhitta/runsc"

var (
	aMethods map[string]method = map[string]method{
		"heap": {
			desc: "HeapCreate, HeapAlloc (PID 0 only)",
			val:  int(runsc.HeapAlloc),
		},
		"navm": {
			desc: "NtAllocateVirtualMemory",
			val:  int(runsc.NtAllocateVirtualMemory),
		},
		"ncs": {
			desc: "NtCreateSeciton",
			val:  int(runsc.NtCreateSection),
		},
	}
	defAlloc string            = "navm"
	defRun   string            = "ncte"
	defWrite string            = "nwvm"
	rMethods map[string]method = map[string]method{
		"ceps": {
			desc: "CertEnumPhysicalStore (no suspend)",
			val:  int(runsc.CertEnumPhysicalStore),
		},
		"cess": {
			desc: "CertEnumSystemStore (dangerous, no suspend)",
			val:  int(runsc.CertEnumSystemStore),
		},
		"cf2": {
			desc: "CopyFile2 (PID 0 only, no suspend)",
			val:  int(runsc.CopyFile2),
		},
		"ncte": {
			desc: "NtCreateThreadEx",
			val:  int(runsc.NtCreateThreadEx),
		},
		"nqat": {
			desc: "NtQueueApcThread",
			val:  int(runsc.NtQueueApcThread),
		},
		"nqate": {
			desc: "NtQueueApcThreadEx (unstable)",
			val:  int(runsc.NtQueueApcThreadEx),
		},
		"rcut": {
			desc: "RtlCreateUserThread",
			val:  int(runsc.RtlCreateUserThread),
		},
	}
	wMethods map[string]method = map[string]method{
		"nwvm": {
			desc: "NtWriteVirtualMemory",
			val:  int(runsc.NtWriteVirtualMemory),
		},
	}
)
