package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"

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
	nocolor bool
	pid     uint64
	run     string
	suspend bool
	times   uint64
	verbose bool
	version bool
	write   string
}

func init() {
	// Configure cli package
	cli.Align = true
	cli.Authors = []string{"Miles Whittaker <mj@whitta.dev>"}
	cli.Banner = filepath.Base(os.Args[0]) + " [OPTIONS]"
	cli.BugEmail = "runsc.bugs@whitta.dev"

	cli.ExitStatus(
		"Normally the exit status is 0. In the event of an error the",
		"exit status will be one of the below:\n\n",
		fmt.Sprintf("%d: Invalid option\n", InvalidOption),
		fmt.Sprintf("%d: Missing option\n", MissingOption),
		fmt.Sprintf("%d: Invalid argument\n", InvalidArgument),
		fmt.Sprintf("%d: Missing arguments\n", MissingArgument),
		fmt.Sprintf("%d: Extra arguments\n", ExtraArgument),
		fmt.Sprintf("%d: Exception", Exception),
	)
	cli.Info(
		"This tool will launch calc shellcode using the allocation,",
		"write, and run methods specified.",
	)
	cli.SectionAligned("ALLOCATION METHODS", "|", stringify(aMethods))
	cli.SectionAligned("RUN METHODS", "|", stringify(rMethods))
	cli.SectionAligned("WRITE METHODS", "|", stringify(wMethods))

	cli.Title = "runsc"

	// Parse cli flags
	cli.Flag(
		&flags.alloc,
		"a",
		"alloc",
		defAlloc,
		fmt.Sprintf(
			"Specify memory allocation method (default: %s).",
			defAlloc,
		),
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
		"Specify uint32 PID to inject calc into (default: 0 = self).",
	)
	cli.Flag(
		&flags.run,
		"r",
		"run",
		defRun,
		fmt.Sprintf(
			"Specify execution method (default: %s).",
			defRun,
		),
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
		&flags.write,
		"w",
		"write",
		defWrite,
		fmt.Sprintf(
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
		fmt.Println(
			filepath.Base(os.Args[0]) + " version " + runsc.Version,
		)
		os.Exit(Good)
	}

	if flags.pid > math.MaxUint32 {
		cli.Usage(InvalidOption)
	}

	// Validate cli flags
	if cli.NArg() > 0 {
		cli.Usage(ExtraArgument)
	}
}
