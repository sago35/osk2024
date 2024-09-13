package main

import (
	"machine"
	"machine/usb/hid/keyboard"
)

func main() {
	machine.BUTTON_1.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	machine.BUTTON_2.Configure(machine.PinConfig{Mode: machine.PinInputPullup})
	machine.BUTTON_3.Configure(machine.PinConfig{Mode: machine.PinInputPullup})

	kb := keyboard.Port()

	for {
		if !machine.BUTTON_1.Get() {
			kb.Down(keyboard.KeyA)
		} else {
			kb.Up(keyboard.KeyA)
		}

		if !machine.BUTTON_2.Get() {
			kb.Down(keyboard.KeyB)
		} else {
			kb.Up(keyboard.KeyB)
		}

		if !machine.BUTTON_3.Get() {
			kb.Down(keyboard.KeyC)
		} else {
			kb.Up(keyboard.KeyC)
		}
	}
}
