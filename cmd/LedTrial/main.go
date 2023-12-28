package main

import (
	"fmt"
	"os"
	"os/signal"
	"periph.io/x/conn/v3/physic"
	"periph.io/x/conn/v3/spi"
	"periph.io/x/conn/v3/spi/spireg"
	"periph.io/x/devices/v3/nrzled"
	"periph.io/x/host/v3"
	"syscall"
	"time"
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

	exit := make(chan os.Signal, 1) // we need to reserve to buffer size 1, so the notifier are not blocked
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)
	go func() {

	}()

	selectOptions := [][]byte{
		[]byte{255, 255, 255, 128},
		[]byte{0, 0, 255, 128},
		[]byte{255, 0, 0, 128},
		[]byte{0, 255, 0, 128},
	}

	var option = 0

	_, err = device.Write([]byte{255, 0, 0, 128})
	time.Sleep(3 * time.Second)
	device.Halt()
	time.Sleep(3 * time.Second)
	_, err = device.Write([]byte{0, 0, 0, 0, 255, 0, 0, 128})
	time.Sleep(3 * time.Second)
	device.Halt()
	time.Sleep(3 * time.Second)
	_, err = device.Write([]byte{0, 0, 0, 0, 0, 0, 0, 0, 255, 0, 0, 128})
	time.Sleep(3 * time.Second)
	device.Halt()
	time.Sleep(3 * time.Second)
	_, err = device.Write([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 0, 0, 128})
	time.Sleep(3 * time.Second)
	device.Halt()
	time.Sleep(3 * time.Second)
	_, err = device.Write([]byte{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 255, 0, 0, 128})
	device.Halt()
	time.Sleep(3 * time.Second)

	for {
		select {
		case <-exit:
			return
		default:
			buf := make([]byte, 0, 4*15)
			for i := 0; i < 15; i++ {
				buf = append(buf, selectOptions[option]...)
			}
			fmt.Printf("buf is %+v\n", buf)
			option = (option + 1) % 4
			_, err := device.Write(buf)
			if err != nil {
				fmt.Printf("hmmm, error writing buffer %v\n", err)
				return
			}
			time.Sleep(time.Second)
		}
	}
}
