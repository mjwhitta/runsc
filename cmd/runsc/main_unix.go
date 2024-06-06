//go:build !windows

package main

var (
	aMethods map[string]method = map[string]method{
		"none": {"Unsupported OS", 0},
	}
	defAlloc string            = "n/a"
	defRun   string            = "n/a"
	defWrite string            = "n/a"
	rMethods map[string]method = map[string]method{
		"none": {"Unsupported OS", 0},
	}
	wMethods map[string]method = map[string]method{
		"none": {"Unsupported OS", 0},
	}
)
