package hs42561k

import (
	"errors"

	max7219spi "github.com/PacktPublishing/Programming-Microcontrollers-and-WebAssembly-with-TinyGo/ch5/max7219-spi"
)

type Driver interface {
	Configure()
	SetDigit(digit byte, character Character) error
	StartDisplayTest()
	StopDisplayTest()
	StartShutdownMode()
	StopShutdownMode()
}

type driver struct {
	digitNumber   uint8
	displayDriver max7219spi.Driver
}

func NewDriver(displayDriver max7219spi.Driver, digitNumber uint8) Driver {
	return &driver{
		displayDriver: displayDriver,
		digitNumber:   digitNumber,
	}
}

func (driver *driver) Configure() {
	driver.StopDisplayTest()

	switch driver.digitNumber {
	case 1: // only decode first digit
		driver.displayDriver.WriteCommand(byte(max7219spi.REG_DECODE_MODE), 0x01)
	case 2, 3, 4: //  decode digits 3-0
		driver.displayDriver.WriteCommand(byte(max7219spi.REG_DECODE_MODE), 0x0F)
	case 8: // decode 8 digits
		driver.displayDriver.WriteCommand(byte(max7219spi.REG_DECODE_MODE), 0xFF)
	}

	driver.StopShutdownMode()

	driver.displayDriver.WriteCommand(byte(max7219spi.REG_SCANLIMIT), byte(driver.digitNumber-1))
	for i := 1; i < int(driver.digitNumber); i++ {
		driver.displayDriver.WriteCommand(byte(i), byte(Blank))
	}
}

func (driver *driver) StartShutdownMode() {
	driver.displayDriver.WriteCommand(byte(max7219spi.REG_SHUTDOWN), 0x00)

}

func (driver *driver) StopShutdownMode() {
	driver.displayDriver.WriteCommand(byte(max7219spi.REG_SHUTDOWN), 0x01)
}

func (driver *driver) StartDisplayTest() {
	driver.displayDriver.WriteCommand(byte(max7219spi.REG_DISPLAY_TEST), 0x01)

}

func (driver *driver) StopDisplayTest() {
	driver.displayDriver.WriteCommand(byte(max7219spi.REG_DISPLAY_TEST), 0x00)
}

var ErrIllegalDigit = errors.New("Invalid digit selected")

func (driver *driver) SetDigit(digit byte, character Character) error {
	if uint8(digit) > driver.digitNumber {
		return ErrIllegalDigit
	}

	driver.displayDriver.WriteCommand(byte(digit), byte(character))

	return nil
}
