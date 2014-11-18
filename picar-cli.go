package main

import (
	"github.com/mrmorphic/hwio"
	"time"
)

func main() {
	// For some reason my pi isn't autodetected, so be explicit here
	hwio.SetDriver(new(hwio.RaspberryPiDTDriver))

	steeringOne, e := hwio.GetPinWithMode("gpio18", hwio.OUTPUT)
	if e != nil {
	}
	steeringTwo, e := hwio.GetPinWithMode("gpio23", hwio.OUTPUT)
	if e != nil {
	}

	go func() {
		drivingOne, e := hwio.GetPinWithMode("gpio4", hwio.OUTPUT)
		if e != nil {
		}

		drivingTwo, e := hwio.GetPinWithMode("gpio17", hwio.OUTPUT)
		if e != nil {
		}

		// Initial state is both off
		hwio.DigitalWrite(drivingOne, hwio.LOW)
		hwio.DigitalWrite(drivingTwo, hwio.LOW)

		// Go forwards?
		hwio.DigitalWrite(drivingOne, hwio.HIGH)
		hwio.DigitalWrite(drivingTwo, hwio.LOW)

		time.Sleep(2 * time.Second)

		// Go backwards?
		hwio.DigitalWrite(drivingOne, hwio.LOW)
		hwio.DigitalWrite(drivingTwo, hwio.HIGH)

		time.Sleep(2 * time.Second)

		// Stop
		hwio.DigitalWrite(drivingOne, hwio.LOW)
		hwio.DigitalWrite(drivingTwo, hwio.LOW)


	}()

	// Initial state is both off
	hwio.DigitalWrite(steeringOne, hwio.LOW)
	hwio.DigitalWrite(steeringTwo, hwio.LOW)

	// Go left?
	hwio.DigitalWrite(steeringOne, hwio.HIGH)
	hwio.DigitalWrite(steeringTwo, hwio.LOW)

	time.Sleep(3 * time.Second)

	// Go right?
	hwio.DigitalWrite(steeringOne, hwio.LOW)
	hwio.DigitalWrite(steeringTwo, hwio.HIGH)

	time.Sleep(3 * time.Second)

	// Center
	hwio.DigitalWrite(steeringOne, hwio.LOW)
	hwio.DigitalWrite(steeringTwo, hwio.LOW)

	defer hwio.CloseAll()
}
