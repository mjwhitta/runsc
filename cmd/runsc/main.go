package main

import (
	"encoding/hex"
	"strings"
	"time"

	"github.com/mjwhitta/log"
	"github.com/mjwhitta/runsc"
)

func launch(sc []byte) error {
	var l *runsc.Launcher = runsc.New()

	if _, ok := aMethods[flags.alloc]; ok {
		l.AllocVia(runsc.AllocMethod(aMethods[flags.alloc].val))
	}

	l.InPID(uint32(flags.pid))

	if _, ok := rMethods[flags.run]; ok {
		l.RunVia(runsc.RunMethod(rMethods[flags.run].val))
	}

	l.Suspend(flags.suspend)

	if _, ok := wMethods[flags.write]; ok {
		l.WriteVia(runsc.WriteMethod(wMethods[flags.write].val))
	}

	log.Infof(
		"Launching calc in %d: %s->%s->%s",
		flags.pid,
		flags.alloc,
		flags.write,
		flags.run,
	)
	if e := l.Exe(sc); e != nil {
		return e
	}

	return nil
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
	var sc []byte

	validate()

	// Convert hex to shellcode
	if sc, e = hex.DecodeString(strings.Join(calc, "")); e != nil {
		panic(e)
	}

	for i := 0; i < int(flags.times); i++ {
		if e = launch(sc); e != nil {
			panic(e)
		}

		time.Sleep(time.Second)
	}
}
