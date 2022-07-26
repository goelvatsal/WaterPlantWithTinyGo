package main

import (
	"github.com/PacktPublishing/Creative-DIY-Microcontroller-Projects-with-TinyGo-and-WebAssembly/Chapter04/buzzer"
	"github.com/tinygo-org/tinygo/tree/release/src/runtime/volatile"
	"machine"
	"time"
)

func main() {
	newBuzzer := buzzer.NewBuzzer(machine.D4)
	newBuzzer.Configure()

	for {
		newBuzzer.Beep(time.Millisecond*100, 3)
		time.Sleep(3 * time.Second)
	}
}
