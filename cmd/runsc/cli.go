package main

import (
	"os"
	"runtime"

	"github.com/mjwhitta/cli"
	hl "github.com/mjwhitta/hilighter"
	"github.com/mjwhitta/runsc"
)

// Exit status
const (
	Good = iota
	InvalidOption
	MissingOption
	InvalidArgument
	MissingArgument
	ExtraArgument
	Exception
)

// Flags
var flags struct {
	alloc   string
	min     bool
	nocolor bool
	pid     uint64
	run     string
	suspend bool
	times   uint64
	verbose bool
	version bool
	wait    bool
	write   string
}

func init() {
	var defAlloc string = "n/a"
	var defRun string = "n/a"
	var defWrite string = "n/a"

	// Configure cli package
	cli.Align = true
	cli.Authors = []string{"Miles Whittaker <mj@whitta.dev>"}
	cli.Banner = hl.Sprintf("%s [OPTIONS]", os.Args[0])
	cli.BugEmail = "runsc.bugs@whitta.dev"
	cli.ExitStatus(
		"Normally the exit status is 0. In the event of an error the",
		"exit status will be one of the below:\n\n",
		hl.Sprintf("%d: Invalid option\n", InvalidOption),
		hl.Sprintf("%d: Missing option\n", MissingOption),
		hl.Sprintf("%d: Invalid argument\n", InvalidArgument),
		hl.Sprintf("%d: Missing arguments\n", MissingArgument),
		hl.Sprintf("%d: Extra arguments\n", ExtraArgument),
		hl.Sprintf("%d: Exception", Exception),
	)
	cli.Info(
		"This tool will launch calc shellcode using the allocation,",
		"write, and run methods specified.",
	)

	switch runtime.GOOS {
	case "windows":
		defAlloc = "navm"
		defRun = "rcut"
		defWrite = "nwvm"

		cli.Section(
			"ALLOCATION METHODS",
			"heap: HeapCreate, HeapAlloc (PID 0 only)\n",
			"navm: NtAllocateVirtualMemory\n",
			"ncs: NtCreateSection",
		)
		cli.Section(
			"RUN METHODS",
			"nqat: NtQueueApcThread\n",
			"nqate: NtQueueApcThreadEx (unstable)\n",
			"rcut: RtlCreateUserThread",
		)
		cli.Section(
			"WRITE METHODS",
			"nwvm: NtWriteVirtualMemory",
		)
	}

	cli.Title = "runsc"

	// Parse cli flags
	cli.Flag(
		&flags.alloc,
		"a",
		"alloc",
		defAlloc,
		hl.Sprintf(
			"Specify memory allocation method (default: %s).",
			defAlloc,
		),
	)
	cli.Flag(
		&flags.min,
		"m",
		"min",
		false,
		"Use minimal calc shellcode.",
	)
	cli.Flag(
		&flags.nocolor,
		"no-color",
		false,
		"Disable colorized output.",
	)
	cli.Flag(
		&flags.pid,
		"p",
		"pid",
		0,
		"Specify PID to inject calc into (default: 0 = self).",
	)
	cli.Flag(
		&flags.run,
		"r",
		"run",
		defRun,
		hl.Sprintf("Specify execution method (default: %s).", defRun),
	)
	cli.Flag(
		&flags.suspend,
		"s",
		"suspend",
		false,
		"Suspend created threads.",
	)
	cli.Flag(
		&flags.times,
		"t",
		"times",
		1,
		"Specify how many times to launch calc.",
	)
	cli.Flag(
		&flags.verbose,
		"v",
		"verbose",
		false,
		"Show stacktrace, if error.",
	)
	cli.Flag(&flags.version, "V", "version", false, "Show version.")
	cli.Flag(
		&flags.wait,
		"wait",
		false,
		"Wait 10 secs before and 10 mins after (for debugging).",
	)
	cli.Flag(
		&flags.write,
		"w",
		"write",
		defWrite,
		hl.Sprintf(
			"Specify memory write method (default: %s).",
			defWrite,
		),
	)
	cli.Parse()
}

// Process cli flags and ensure no issues
func validate() {
	hl.Disable(flags.nocolor)

	// Short circuit if version was requested
	if flags.version {
		hl.Printf("runsc version %s\n", runsc.Version)
		os.Exit(Good)
	}

	// Validate cli flags
	if cli.NArg() > 0 {
		cli.Usage(ExtraArgument)
	}
}
