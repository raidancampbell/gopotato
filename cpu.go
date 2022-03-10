package main

import (
	"fmt"
	"time"
)

type reg *byte

var (
	v0, v1, v2, v3, v4, v5, v6, v7, v8, v9, va, vb, vc, vd, ve reg

	vf     reg // flag register
	i      uint16
	dt, st reg // delay and sound timers
	pc     uint16
	sp     byte
	stack  [16]uint16
)

func Init() {
	initDisp()
	initRAM()
	go timerTick()
	go tick()
}

// emulate the CPU at 512hz
func tick() {
	tim := time.NewTimer(1953 * time.Microsecond)
	for {
		select {
		case <-tim.C:
		}
	}
}

// timerTick controls
func timerTick() {
	tim := time.NewTimer(16667 * time.Microsecond)

	for {
		select {
		case <-tim.C:
			pollForKeys() // abusively putting this in the timer code.  we don't need to poll that often
			if *dt != 0x00 {
				*dt--
			}
			if *st != 0x00 {
				*st--
			}
		}
	}
}

func numToReg(nibble byte) reg {
	if nibble > 0x0F {
		panic(fmt.Sprintf("malformed nibble given to numToReg: %x", nibble))
	}

	switch nibble {
	case 0x00:
		return v0
	case 0x01:
		return v1
	case 0x02:
		return v2
	case 0x03:
		return v3
	case 0x04:
		return v4
	case 0x05:
		return v5
	case 0x06:
		return v6
	case 0x07:
		return v7
	case 0x08:
		return v8
	case 0x09:
		return v9
	case 0x0A:
		return va
	case 0x0B:
		return vb
	case 0x0C:
		return vc
	case 0x0D:
		return vd
	case 0x0E:
		return ve
	case 0x0F:
		return vf
	default:
	}
	panic(fmt.Sprintf("impossible nibble given to numToReg: %x", nibble))
}

func intToHex(i int) byte {
	switch i {
	case 0:
		return 0x00
	case 1:
		return 0x01
	case 2:
		return 0x02
	case 3:
		return 0x03
	case 4:
		return 0x04
	case 5:
		return 0x05
	case 6:
		return 0x06
	case 7:
		return 0x07
	case 8:
		return 0x08
	case 9:
		return 0x09
	}
	panic(fmt.Sprintf("unexpected int %d given to intToHex", i))
}
