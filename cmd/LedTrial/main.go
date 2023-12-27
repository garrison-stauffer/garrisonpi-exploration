package main

import (
	"fmt"
	"time"
)
import "github.com/stianeikeland/go-rpio/v4"

func main() {
	fmt.Println("Hello world")
	err := rpio.Open()
	defer func() {
		err = rpio.Close()
		if err != nil {
			fmt.Println("error closing rpio: ", err)
		}
	}()
	if err != nil {
		panic(fmt.Errorf("error initializing rpio: %w", err))
	}

	foo := rpio.Pin(16)
	foo.High()

	for i := 0; i < 50; i++ {
		time.Sleep(2000 * time.Millisecond)
		foo.Toggle()
	}
}
