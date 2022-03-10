package gopotato

import (
	"fmt"
	"github.com/faiface/pixel/pixelgl"
)

var (
	k0, k1, k2, k3, k4, k5, k6, k7, k8, k9, ka, kb, kc, kd, ke, kf bool
	keyPress                                                       chan byte
	keyWaiting                                                     bool
)

func pollForKeys() {
	for key, _ := range keys {
		if disp.window.JustPressed(key) {
			keys[key] = true
			if keyWaiting {
				keyPress <- keyToNibble(key)
				keyWaiting = false
			}
		}
		if disp.window.JustReleased(key) {
			keys[key] = false
		}
	}
}

var keys = map[pixelgl.Button]bool{
	pixelgl.Key0: k0,
	pixelgl.Key1: k1,
	pixelgl.Key2: k2,
	pixelgl.Key3: k3,
	pixelgl.Key4: k4,
	pixelgl.Key5: k5,
	pixelgl.Key6: k6,
	pixelgl.Key7: k7,
	pixelgl.Key8: k8,
	pixelgl.Key9: k9,
	pixelgl.KeyA: ka,
	pixelgl.KeyB: kb,
	pixelgl.KeyC: kc,
	pixelgl.KeyD: kd,
	pixelgl.KeyE: ke,
	pixelgl.KeyF: kf,
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
		return k0
	case 0x01:
		return k1
	case 0x02:
		return k2
	case 0x03:
		return k3
	case 0x04:
		return k4
	case 0x05:
		return k5
	case 0x06:
		return k6
	case 0x07:
		return k7
	case 0x08:
		return k8
	case 0x09:
		return k9
	case 0x0A:
		return ka
	case 0x0B:
		return kb
	case 0x0C:
		return kc
	case 0x0D:
		return kd
	case 0x0E:
		return ke
	case 0x0F:
		return kf
	default:
	}
	panic(fmt.Sprintf("impossible nibble given to numToReg: %x", nibble))
}
