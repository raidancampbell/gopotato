package main

import (
	"fmt"
	"github.com/faiface/pixel/pixelgl"
)

var (
	keyPress   chan byte
	keyWaiting bool
)

func pollForKeys() {
	for key, _ := range keys {
		if disp.window.Pressed(key) {
			if !keys[key] && keyWaiting {
				keyPress <- keyToNibble(key)
				keyWaiting = false
				keyPress <- keyToNibble(key)
			}
			keys[key] = true
		} else {
			keys[key] = false
		}
	}
}

var keys = map[pixelgl.Button]bool{
	pixelgl.Key0: false,
	pixelgl.Key1: false,
	pixelgl.Key2: false,
	pixelgl.Key3: false,
	pixelgl.Key4: false,
	pixelgl.Key5: false,
	pixelgl.Key6: false,
	pixelgl.Key7: false,
	pixelgl.Key8: false,
	pixelgl.Key9: false,
	pixelgl.KeyA: false,
	pixelgl.KeyB: false,
	pixelgl.KeyC: false,
	pixelgl.KeyD: false,
	pixelgl.KeyE: false,
	pixelgl.KeyF: false,
}

func keyToNibble(key pixelgl.Button) byte {
	switch key {
	case pixelgl.Key0:
		return 0x00
	case pixelgl.Key1:
		return 0x01
	case pixelgl.Key2:
		return 0x02
	case pixelgl.Key3:
		return 0x03
	case pixelgl.Key4:
		return 0x04
	case pixelgl.Key5:
		return 0x05
	case pixelgl.Key6:
		return 0x06
	case pixelgl.Key7:
		return 0x07
	case pixelgl.Key8:
		return 0x08
	case pixelgl.Key9:
		return 0x09
	case pixelgl.KeyA:
		return 0x0A
	case pixelgl.KeyB:
		return 0x0B
	case pixelgl.KeyC:
		return 0x0C
	case pixelgl.KeyD:
		return 0x0D
	case pixelgl.KeyE:
		return 0x0E
	case pixelgl.KeyF:
		return 0x0F
	}
	panic(fmt.Sprintf("unexpected key passed to keyToNibble: %+v", key))
}

func isKeyPressed(nibble byte) bool {
	if nibble > 0x0F {
		panic(fmt.Sprintf("malformed nibble given to numToReg: %x", nibble))
	}

	switch nibble {
	case 0x00:
		return keys[pixelgl.Key0]
	case 0x01:
		return keys[pixelgl.Key1]
	case 0x02:
		return keys[pixelgl.Key2]
	case 0x03:
		return keys[pixelgl.Key3]
	case 0x04:
		return keys[pixelgl.Key4]
	case 0x05:
		return keys[pixelgl.Key5]
	case 0x06:
		return keys[pixelgl.Key6]
	case 0x07:
		return keys[pixelgl.Key7]
	case 0x08:
		return keys[pixelgl.Key8]
	case 0x09:
		return keys[pixelgl.Key9]
	case 0x0A:
		return keys[pixelgl.KeyA]
	case 0x0B:
		return keys[pixelgl.KeyB]
	case 0x0C:
		return keys[pixelgl.KeyC]
	case 0x0D:
		return keys[pixelgl.KeyD]
	case 0x0E:
		return keys[pixelgl.KeyE]
	case 0x0F:
		return keys[pixelgl.KeyF]
	default:
	}
	panic(fmt.Sprintf("impossible nibble given to numToReg: %x", nibble))
}
