package main

import (
	"machine"
	"time"
)

func main() {
	led := machine.LCD_BACKLIGHT
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})
	for {
		led.Toggle()
		time.Sleep(time.Millisecond * 500)
	}
}
