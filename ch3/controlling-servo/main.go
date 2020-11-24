package main

import (
	"machine"
	"math"
	"time"
)

const twoPi = float32(2 * math.Pi)

// Left
const pos0DutyCycle = 1500 * time.Microsecond
const pos0RemainingPeriod = 18500 * time.Microsecond

// Middle
const pos1DutyCycle = 1000 * time.Microsecond
const pos1RemainingPeriod = 19000 * time.Microsecond

// Right
const pos2DutyCycle = 2000 * time.Microsecond
const pos2RemainingPeriod = 18000 * time.Microsecond

func main() {
	machine.InitPWM()

	servoPin := machine.PWM{Pin: machine.D3}
	servoPin.Configure()

	position := 0

	time.Sleep(time.Second)

	for {
		for position = 0; position < 180; position++ {
			servoPin.Pin.High()
			time.Sleep(pos0DutyCycle)
			servoPin.Pin.Low()
			time.Sleep(pos0RemainingPeriod)

			time.Sleep(100 * time.Millisecond)
		}

		time.Sleep(time.Second)

		for position = 0; position >= 1; position-- {
			servoPin.Pin.High()
			time.Sleep(pos2DutyCycle)
			servoPin.Pin.Low()
			time.Sleep(pos2RemainingPeriod)

			time.Sleep(100 * time.Millisecond)
		}
		return
	}
}