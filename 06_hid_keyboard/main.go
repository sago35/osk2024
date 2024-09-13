package main

import (
	"machine"
	"machine/usb/hid/keyboard"
)

func main() {
	machine.BUTTON_1.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	kb := keyboard.Port()

	for {
		if !machine.BUTTON_1.Get() {
			kb.Down(keyboard.KeyA)
		} else {
			kb.Up(keyboard.KeyA)
		}
	}
}
