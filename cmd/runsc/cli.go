package main

import (
	"os"
	"strings"

	"gitlab.com/mjwhitta/cli"
	hl "gitlab.com/mjwhitta/hilighter"
	"gitlab.com/mjwhitta/runsc"
)

// Flags
type cliFlags struct {
	min     bool
	nocolor bool
	pid     uint64
	suspend bool
	times   uint64
	verbose bool
	version bool
	wait    bool
}

var flags cliFlags

func init() {
	// Configure cli package
	cli.Align = true
	cli.Authors = []string{"Miles Whittaker <mj@whitta.dev>"}
	cli.Banner = hl.Sprintf(
		"%s [OPTIONS] <method>",
		os.Args[0],
	)
	cli.BugEmail = "runsc.bugs@whitta.dev"
	cli.ExitStatus = strings.Join(
		[]string{
			"Normally the exit status is 0. In the event of an error",
			"the exit status will be one of the below:\n\n",
			hl.Sprintf("%d: Invalid option\n", InvalidOption),
			hl.Sprintf("%d: Invalid argument\n", InvalidArgument),
			hl.Sprintf("%d: Missing arguments\n", MissingArguments),
			hl.Sprintf("%d: Extra arguments\n", ExtraArguments),
			hl.Sprintf("%d: Exception", Exception),
		},
		" ",
	)
	cli.Info = strings.Join([]string{
		"This tool will launch calc shellcode using the method",
		"specified on the command-line.",
	},
		" ",
	)
	cli.Section(
		"METHODS",
		strings.Join(
			[]string{
				"navm: NtAllocateVirtualMemory",
				"ncs: NtCreateSection",
				"nqat: NtQueueApcThread",
				"nqate: NtQueueApcThreadEx",
			},
			"\n",
		),
	)
	cli.Title = "runsc"

	// Parse cli flags
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
		"Specify PID to inject calc into (default: 0 - self).",
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
		"Show show stacktrace if error.",
	)
	cli.Flag(&flags.version, "V", "version", false, "Show version.")
	cli.Flag(
		&flags.wait,
		"w",
		"wait",
		false,
		"Wait 20 secs before and 10 mins after (for debugging).",
	)
	cli.Parse()
}

// Process cli flags and ensure no issues
func validate() {
	hl.Disable(flags.nocolor)
	runsc.Suspended = flags.suspend

	// Short circuit if version was requested
	if flags.version {
		hl.Printf("runsc version %s\n", runsc.Version)
		os.Exit(Good)
	}

	// Validate cli flags
	if cli.NArg() == 0 {
		cli.Usage(MissingArguments)
	} else if cli.NArg() > 1 {
		cli.Usage(ExtraArguments)
	}
}
