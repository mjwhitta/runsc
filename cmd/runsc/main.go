package main

import (
	"encoding/hex"
	"strings"
	"time"

	"github.com/mjwhitta/log"
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
