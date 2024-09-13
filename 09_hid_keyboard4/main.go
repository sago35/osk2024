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

	layer := 0
	keymap := []map[machine.Pin]keyboard.Keycode{
		{
			machine.BUTTON_1: keyboard.Key1,
			machine.BUTTON_2: keyboard.Key2,
		},
		{
			machine.BUTTON_1: keyboard.KeyA,
			machine.BUTTON_2: keyboard.KeyBackspace,
		},
	}

	for {
		if layer < len(keymap) {
			for k, v := range keymap[layer] {
				if !k.Get() {
					kb.Down(v)
				} else {
					kb.Up(v)
				}
			}
		}

		if !machine.BUTTON_3.Get() {
			layer = 1
		} else {
			layer = 0
		}
	}
}
