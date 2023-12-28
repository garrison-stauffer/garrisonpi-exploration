package main

import (
	"fmt"
	"os"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/conn/v3/spi"
	"periph.io/x/conn/v3/spi/spireg"
	"periph.io/x/devices/v3/nrzled"
)

func main() {
	//spi.
	var pin spi.PortCloser
	pin, err := spireg.Open("12")
	if err != nil {
		fmt.Printf("error opening spi port 12: %v\n", err)
		os.Exit(1)
		return
	}
	defer func(pin spi.PortCloser) {
		err := pin.Close()
		if err != nil {
			fmt.Printf("error closing spi port: %v\n", err)
			os.Exit(1)
		}
	}(pin)

	device, err := nrzled.NewSPI(pin, &nrzled.Opts{
		NumPixels: 15,
		Channels:  4,
		Freq:      1100 * physic.Hertz,
	})
	if err != nil {
		fmt.Printf("error opening nrzled device: %v\n", err)
		os.Exit(1)
	}

	defer device.Halt()

	for {
		device.Write()
		write, err := device.Write([]byte{255, 255, 255, 255})
		if err != nil {
			return
		}
	}
}
