//go:build !windows

package main

var (
	aMethods map[string]method = map[string]method{
		"none": {"Unsupported OS", 0},
	}
	defAlloc string            = "none"
	defRun   string            = "none"
	defWrite string            = "none"
	rMethods map[string]method = map[string]method{
		"none": {"Unsupported OS", 0},
	}
	wMethods map[string]method = map[string]method{
		"none": {"Unsupported OS", 0},
	}
)
