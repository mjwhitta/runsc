package main

import (
	"encoding/hex"
	"strings"
	"time"

	hl "github.com/mjwhitta/hilighter"
	"github.com/mjwhitta/log"
	"github.com/mjwhitta/runsc"
)

func getPayload() string {
	var payload []string = calc

	if flags.min {
		payload = calcMin
	}

	return strings.Join(payload, "")
}

func launch(sc []byte) error {
	var l *runsc.Launcher = runsc.New()

	if _, ok := aMethods[flags.alloc]; !ok {
		return hl.Errorf("unsupported alloc method: %s", flags.alloc)
	}

	if _, ok := rMethods[flags.run]; !ok {
		return hl.Errorf("unsupported run method: %s", flags.run)
	}

	if _, ok := wMethods[flags.write]; !ok {
		return hl.Errorf("unsupported write method: %s", flags.write)
	}

	l.AllocVia(runsc.AllocMethod(aMethods[flags.alloc].val))
	l.InPID(uint32(flags.pid))
	l.RunVia(runsc.RunMethod(rMethods[flags.run].val))
	l.Suspend(flags.suspend)
	l.WriteVia(runsc.WriteMethod(wMethods[flags.write].val))

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
	if sc, e = hex.DecodeString(getPayload()); e != nil {
		panic(e)
	}

	if flags.wait {
		time.Sleep(10 * time.Second)
	}

	for i := 0; i < int(flags.times); i++ {
		if e = launch(sc); e != nil {
			panic(e)
		}

		time.Sleep(time.Second)
	}

	if flags.wait {
		time.Sleep(10 * time.Minute)
	}
}
