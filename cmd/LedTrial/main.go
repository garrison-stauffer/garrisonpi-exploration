package main

import (
	"fmt"
	"os"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/conn/v3/spi"
	"periph.io/x/conn/v3/spi/spireg"
	"periph.io/x/devices/v3/nrzled"
	"periph.io/x/host/v3"
)

func main() {
	//spi.
	foo, err := host.Init()
	fmt.Printf("registered host: %+v\n", foo)
	if err != nil {
		fmt.Printf("error initializing host: %v\n", err)
		os.Exit(1)
		return
	}
	pin, err := spireg.Open("")
	fmt.Printf("pin loaded: %+v\n", pin)
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
		Freq:      2500 * physic.KiloHertz,
	})
	if err != nil {
		fmt.Printf("error opening nrzled device: %v\n", err)
		os.Exit(1)
	}

	defer device.Halt()

	for {
		_, err := device.Write([]byte{255, 255, 255, 255})
		if err != nil {
			return
		}
	}
}
