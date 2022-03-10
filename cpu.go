package main

import (
	"encoding/binary"
	"fmt"
	"sync"
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

func init() {
	pc = 0x200
	initRAM()
	v0 = new(byte)
	v1 = new(byte)
	v2 = new(byte)
	v3 = new(byte)
	v4 = new(byte)
	v5 = new(byte)
	v6 = new(byte)
	v7 = new(byte)
	v8 = new(byte)
	v9 = new(byte)
	va = new(byte)
	vb = new(byte)
	vc = new(byte)
	vd = new(byte)
	ve = new(byte)
	vf = new(byte)
	dt = new(byte)
	st = new(byte)
	disp = display{
		Mutex: &sync.Mutex{},
	}
}

// emulate the CPU at 512hz
func tick() {
	tim := time.NewTicker(19530 * time.Microsecond)
	for {
		select {
		case <-tim.C:
			for itr := 0; itr < 16; itr++ {
				opWord := binary.BigEndian.Uint16([]byte{mem[pc], mem[pc+1]})
				var op opcode
				found := false
				for opItr := range opcodes {
					if opcodes[opItr].matches(opWord) {
						op = opcodes[opItr]
						found = true
						break
					}
				}
				if !found {
					panic(fmt.Sprintf("failed to find opcode %x", opWord))
				}
				fmt.Printf("executing opcode %x as: %s\n", opWord, op.name)
				op.exec(opWord)
			}
		}
	}
}

// timerTick controls
func timerTick() {
	tim := time.NewTicker(16667 * time.Microsecond)

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
