//go:build windows

package main

import "github.com/mjwhitta/runsc"

var (
	aMethods map[string]method = map[string]method{
		"heap": {
			desc: "HeapCreate, HeapAlloc (PID 0 only)",
			val:  uint(runsc.HeapAlloc),
		},
		"navm": {
			desc: "NtAllocateVirtualMemory",
			val:  uint(runsc.NtAllocateVirtualMemory),
		},
		"ncs": {
			desc: "NtCreateSeciton",
			val:  uint(runsc.NtCreateSection),
		},
	}
	defAlloc string            = "navm"
	defRun   string            = "ncte"
	defWrite string            = "nwvm"
	rMethods map[string]method = map[string]method{
		"ceps": {
			desc: "CertEnumPhysicalStore (no suspend)",
			val:  uint(runsc.CertEnumPhysicalStore),
		},
		"cess": {
			desc: "CertEnumSystemStore (dangerous, no suspend)",
			val:  uint(runsc.CertEnumSystemStore),
		},
		"cf2": {
			desc: "CopyFile2 (PID 0 only, no suspend)",
			val:  uint(runsc.CopyFile2),
		},
		"ewsw": {
			desc: "EnumWindowStationsW (PID 0 only, no suspend)",
			val:  uint(runsc.EnumWindowStationsW),
		},
		"ncte": {
			desc: "NtCreateThreadEx",
			val:  uint(runsc.NtCreateThreadEx),
		},
		"nqat": {
			desc: "NtQueueApcThread",
			val:  uint(runsc.NtQueueApcThread),
		},
		"nqate": {
			desc: "NtQueueApcThreadEx (unstable)",
			val:  uint(runsc.NtQueueApcThreadEx),
		},
		"rcut": {
			desc: "RtlCreateUserThread",
			val:  uint(runsc.RtlCreateUserThread),
		},
	}
	wMethods map[string]method = map[string]method{
		"nwvm": {
			desc: "NtWriteVirtualMemory",
			val:  uint(runsc.NtWriteVirtualMemory),
		},
	}
)
