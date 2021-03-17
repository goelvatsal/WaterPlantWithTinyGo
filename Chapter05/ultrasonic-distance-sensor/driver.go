package hcsr04

import (
	"machine"
	"time"
)

const speedOfSound = 0.0343 // cm / us

type HCSR04 interface {
	Configure()
	GetDistance() uint16
	GetDistanceFromPulseLength(pulseLength float32) uint16
}

type hcsr04 struct {
	trigger machine.Pin
	echo    machine.Pin
	timeout int64
}

func NewHCSR04(trigger, echo machine.Pin, maxDistance float32) HCSR04 {
	timeout := int64(maxDistance * 2 / speedOfSound)

	return &hcsr04{
		trigger: trigger,
		echo:    echo,
		timeout: timeout,
	}
}

func (sensor *hcsr04) Configure() {
	sensor.trigger.Configure(machine.PinConfig{Mode: machine.PinOutput})
	sensor.echo.Configure(machine.PinConfig{Mode: machine.PinInput})
}

func (sensor *hcsr04) sendPulse() {
	sensor.trigger.High()
	time.Sleep(10 * time.Microsecond)
	sensor.trigger.Low()
}

func (sensor *hcsr04) GetDistance() uint16 {
	i := 0
	timeoutTimer := time.Now()
	sensor.sendPulse()

	for {
		if sensor.echo.Get() {
			timeoutTimer = time.Now()
			break
		}
		i++
		if i > 15 {
			microseconds := time.Since(timeoutTimer).Microseconds()
			if microseconds > sensor.timeout {
				return 0
			}
		}
	}

	var pulseLength float32
	i = 0
	for {
		if !sensor.echo.Get() {
			microseconds := time.Since(timeoutTimer).Microseconds()
			pulseLength = float32(microseconds)
			break
		}

		i++
		if i > 15 {
			microseconds := time.Since(timeoutTimer).Microseconds()
			if microseconds > sensor.timeout {
				return 0
			}
		}
	}

	return sensor.GetDistanceFromPulseLength(pulseLength)
}

func (sensor *hcsr04) GetDistanceFromPulseLength(pulseLength float32) uint16 {
	pulseLength = pulseLength / 2
	result := pulseLength * speedOfSound

	return uint16(result)
}