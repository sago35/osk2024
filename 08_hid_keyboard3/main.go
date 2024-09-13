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

	keymap := map[machine.Pin]keyboard.Keycode{
		machine.BUTTON_1: keyboard.Key1,
		machine.BUTTON_2: keyboard.Key2,
		machine.BUTTON_3: keyboard.Key3,
	}

	for {
		for k, v := range keymap {
			if !k.Get() {
				kb.Down(v)
			} else {
				kb.Up(v)
			}
		}
	}
}
