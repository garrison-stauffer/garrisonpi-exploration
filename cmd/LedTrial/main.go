package main

import (
	"fmt"
	ws281x "github.com/mcuadros/go-rpi-ws281x"
	"image/color"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	foo, _ := ws281x.NewCanvas(5, 1, &ws281x.HardwareConfig{
		Pin:       18,
		StripType: ws281x.StripRGBW,
	})
	_ = foo.Initialize()
	foo.Set(0, 0, color.RGBA{
		R: 255,
		A: 128,
	})

	time.Sleep(5 * time.Second)
	foo.Set(0, 0, color.RGBA{
		B: 255,
		A: 128,
	})

	exit := make(chan os.Signal, 1) // we need to reserve to buffer size 1, so the notifier are not blocked
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM)

	//selectOptions :=
	fmt.Println("awaiting exit?")

	<-exit
	fmt.Println("Closing...?")
	_ = foo.Close()

	return
}
