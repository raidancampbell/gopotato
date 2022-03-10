package main

import (
	"crypto/rand"
)

type opcode struct {
	matches             func(op uint16) bool
	exec                func(op uint16)
	elapsedMicroseconds int
	name                string
	description         string
}

/*
nnn or addr - A 12-bit value, the lowest 12 bits of the instruction
n or nibble - A 4-bit value, the lowest 4 bits of the instruction
x - A 4-bit value, the lower 4 bits of the high byte of the instruction
y - A 4-bit value, the upper 4 bits of the low byte of the instruction
kk or byte - An 8-bit value, the lowest 8 bits of the instruction
*/

var opcodes = []opcode{
	//{
	//	matches: func(op uint16) bool {
	//		// This instruction is only used on the old computers on which Chip-8 was originally implemented. It is ignored by modern interpreters.
	//		return false
	//	},
	//	exec: func(op uint16) {
	//
	//	},
	//	elapsedMicroseconds: 0,
	//	name:                "0nnn: SYS addr",
	//	description:         "Jump to a machine code routine at nnn",
	//},
	{
		matches: func(op uint16) bool {
			return op == 0x00E0
		},
		exec: func(op uint16) {
			pc++
			disp.fb = framebuffer{}
		},
		elapsedMicroseconds: 109,
		name:                "00E0: CLS",
		description:         "Clear the display",
	},
	{
		matches: func(op uint16) bool {
			return op == 0x00EE
		},
		exec: func(op uint16) {
			pc = stack[sp]
			sp--
		},
		elapsedMicroseconds: 105,
		name:                "00EE: RET",
		description:         "Return from a subroutine",
	},
	{
		matches: func(op uint16) bool {
			return op >= 0x1000 && op < 0x2000
		},
		exec: func(op uint16) {
			pc = op - 0x1000
		},
		elapsedMicroseconds: 105,
		name:                "1nnn: JP addr",
		description:         "Jump to location nnn.",
	},
	{
		matches: func(op uint16) bool {
			return op >= 0x2000 && op < 0x3000
		},
		exec: func(op uint16) {
			sp++
			stack[sp] = pc
			pc = op - 0x2000
		},
		elapsedMicroseconds: 105,
		name:                "2nnn: CALL addr",
		description:         "Call subroutine at nnn.",
	},
	{
		matches: func(op uint16) bool {
			return op >= 0x3000 && op < 0x4000
		},
		exec: func(op uint16) {
			v := numToReg(byte((op & 0x0F00) >> 2))
			val := byte(op & 0x00FF)
			if *v == val {
				pc++
			}
			pc++
		},
		elapsedMicroseconds: 55,
		name:                "3xkk: SE Vx, byte",
		description:         "Skip next instruction if Vx = kk.",
	},
	{
		matches: func(op uint16) bool {
			return op >= 0x4000 && op < 0x5000
		},
		exec: func(op uint16) {
			v := numToReg(byte((op & 0x0F00) >> 2))
			val := byte(op & 0x00FF)
			if *v != val {
				pc++
			}
			pc++
		},
		elapsedMicroseconds: 55,
		name:                "4xkk: SNE Vx, byte",
		description:         "Skip next instruction if Vx != kk.",
	},
	{
		matches: func(op uint16) bool {
			return op >= 0x5000 && op < 0x6000
		},
		exec: func(op uint16) {
			rx := numToReg(byte((op & 0x0F00) >> 2))
			ry := numToReg(byte((op & 0x00F0) >> 1))
			if *rx == *ry {
				pc++
			}
			pc++
		},
		elapsedMicroseconds: 73,
		name:                "5xy0: SE Vx, Vy",
		description:         "Skip next instruction if Vx = Vy.",
	},
	{
		matches: func(op uint16) bool {
			return op >= 0x6000 && op < 0x7000
		},
		exec: func(op uint16) {
			r := numToReg(byte((op & 0x0F00) >> 2))
			val := byte(op & 0x00FF)
			*r = val
			pc++
		},
		elapsedMicroseconds: 27,
		name:                "6xkk: LD Vx, byte",
		description:         "Set Vx = kk.",
	},
	{
		matches: func(op uint16) bool {
			return op >= 0x7000 && op < 0x8000
		},
		exec: func(op uint16) {
			r := numToReg(byte((op & 0x0F00) >> 2))
			val := byte(op & 0x00FF)
			*r = *r + val
			pc++
		},
		elapsedMicroseconds: 45,
		name:                "7xkk: ADD Vx, byte",
		description:         "Set Vx = Vx + kk.",
	},
	{
		matches: func(op uint16) bool {
			return op >= 0x8000 && op < 0x9000 && op&0x000F == 0x0000
		},
		exec: func(op uint16) {
			rx := numToReg(byte((op & 0x0F00) >> 2))
			ry := numToReg(byte((op & 0x00F0) >> 1))
			*rx = *ry
			pc++
		},
		elapsedMicroseconds: 200,
		name:                "8xy0: LD Vx, Vy",
		description:         "Set Vx = Vy.",
	},
	{
		matches: func(op uint16) bool {
			return op >= 0x8000 && op < 0x9000 && op&0x000F == 0x0001
		},
		exec: func(op uint16) {
			rx := numToReg(byte((op & 0x0F00) >> 2))
			ry := numToReg(byte((op & 0x00F0) >> 1))
			*rx |= *ry
			pc++
		},
		elapsedMicroseconds: 200,
		name:                "8xy1: OR Vx, Vy",
		description:         "Set Vx = Vx OR Vy.",
	},
	{
		matches: func(op uint16) bool {
			return op >= 0x8000 && op < 0x9000 && op&0x000F == 0x0002
		},
		exec: func(op uint16) {
			rx := numToReg(byte((op & 0x0F00) >> 2))
			ry := numToReg(byte((op & 0x00F0) >> 1))
			*rx &= *ry
			pc++
		},
		elapsedMicroseconds: 200,
		name:                "8xy2: AND Vx, Vy",
		description:         "Set Vx = Vx AND Vy.",
	},
	{
		matches: func(op uint16) bool {
			return op >= 0x8000 && op < 0x9000 && op&0x000F == 0x0003
		},
		exec: func(op uint16) {
			rx := numToReg(byte((op & 0x0F00) >> 2))
			ry := numToReg(byte((op & 0x00F0) >> 1))
			*rx ^= *ry
			pc++
		},
		elapsedMicroseconds: 200,
		name:                "8xy3: XOR Vx, Vy",
		description:         "Set Vx = Vx XOR Vy.",
	},
	{
		matches: func(op uint16) bool {
			return op >= 0x8000 && op < 0x9000 && op&0x000F == 0x0004
		},
		exec: func(op uint16) {
			rx := numToReg(byte((op & 0x0F00) >> 2))
			ry := numToReg(byte((op & 0x00F0) >> 1))
			b := byte(0x00)
			if *rx+*ry > 255 {
				b = byte(0x01)
			}
			vf = &b
			*rx += *ry
			pc++
		},
		elapsedMicroseconds: 200,
		name:                "8xy4: ADD Vx, Vy",
		description:         "Set Vx = Vx + Vy, set VF = carry.",
	},
	{
		matches: func(op uint16) bool {
			return op >= 0x8000 && op < 0x9000 && op&0x000F == 0x0005
		},
		exec: func(op uint16) {
			rx := numToReg(byte((op & 0x0F00) >> 2))
			ry := numToReg(byte((op & 0x00F0) >> 1))
			b := byte(0x00)
			if *rx > *ry {
				b = byte(0x01)
			}
			vf = &b
			*rx -= *ry
			pc++
		},
		elapsedMicroseconds: 200,
		name:                "8xy5: SUB Vx, Vy",
		description:         "Set Vx = Vx - Vy, set VF = NOT borrow.",
	},
	{
		matches: func(op uint16) bool {
			return op >= 0x8000 && op < 0x9000 && op&0x000F == 0x0006
		},
		exec: func(op uint16) {
			rx := numToReg(byte((op & 0x0F00) >> 2))
			b := byte(op & 0x01)
			vf = &b
			*rx >>= 1
			pc++
		},
		elapsedMicroseconds: 200,
		name:                "8xy6: SHR Vx {, Vy}",
		description:         "Set Vx = Vx SHR 1.",
	},
	{
		matches: func(op uint16) bool {
			return op >= 0x8000 && op < 0x9000 && op&0x000F == 0x0007
		},
		exec: func(op uint16) {
			rx := numToReg(byte((op & 0x0F00) >> 2))
			ry := numToReg(byte((op & 0x00F0) >> 1))
			b := byte(0x00)
			if *ry > *rx {
				b = byte(0x01)
			}
			vf = &b
			*rx = *ry - *rx
			pc++
		},
		elapsedMicroseconds: 200,
		name:                "8xy7: SUBN Vx, Vy",
		description:         "Set Vx = Vy - Vx, set VF = NOT borrow.",
	},
	{
		matches: func(op uint16) bool {
			return op >= 0x8000 && op < 0x9000 && op&0x000F == 0x000E
		},
		exec: func(op uint16) {
			rx := numToReg(byte((op & 0x0F00) >> 2))
			b := byte((op & 0x8000) >> 2)
			if b > 0x00 {
				b = 0x01
			}
			vf = &b
			*rx <<= 1
			pc++
		},
		elapsedMicroseconds: 200,
		name:                "8xyE: SHL Vx {, Vy}",
		description:         "Set Vx = Vx SHL 1.",
	},
	{
		matches: func(op uint16) bool {
			return op >= 0x9000 && op < 0xA000 && op&0x000F == 0x0000
		},
		exec: func(op uint16) {
			rx := numToReg(byte((op & 0x0F00) >> 2))
			ry := numToReg(byte((op & 0x00F0) >> 1))
			if *rx != *ry {
				pc++
			}
			pc++
		},
		elapsedMicroseconds: 73,
		name:                "9xy0: SNE Vx, Vy",
		description:         "Skip next instruction if Vx != Vy.",
	},
	{
		matches: func(op uint16) bool {
			return op >= 0xA000 && op < 0xB000
		},
		exec: func(op uint16) {
			i = op & 0x0FFF
			pc++
		},
		elapsedMicroseconds: 55,
		name:                "Annn: LD I, addr",
		description:         "Set I = nnn.",
	},
	{
		matches: func(op uint16) bool {
			return op >= 0xB000 && op < 0xC000
		},
		exec: func(op uint16) {
			pc = 0x0FFF + uint16(*v0)
			//TODO: increment PC?
		},
		elapsedMicroseconds: 105,
		name:                "Bnnn: JP V0, addr",
		description:         "Jump to location nnn + V0.",
	},
	{
		matches: func(op uint16) bool {
			return op >= 0xB000 && op < 0xC000
		},
		exec: func(op uint16) {
			rx := numToReg(byte((op & 0x0F00) >> 2))
			randByte := make([]byte, 1)
			rand.Read(randByte)
			*rx = randByte[0] & byte(op&0x00FF)
			pc++
		},
		elapsedMicroseconds: 164,
		name:                "Cxkk: RND Vx, byte",
		description:         "Set Vx = random byte AND kk.",
	},
	{
		matches: func(op uint16) bool {
			return op >= 0xD000 && op < 0xE000
		},
		exec: func(op uint16) {
			n := byte(op & 0x000F)
			sprite := mem[i:n]
			rx := numToReg(byte((op & 0x0F00) >> 2))
			ry := numToReg(byte((op & 0x00F0) >> 1))

			drawSprite(sprite, *rx, *ry)
			pc++
		},
		elapsedMicroseconds: 22734,
		name:                "Dxyn: DRW Vx, Vy, nibble",
		description:         "Display n-byte sprite starting at memory location I at (Vx, Vy), set VF = collision.",
	},
	{
		matches: func(op uint16) bool {
			return op&0xF0FF == 0xE09E
		},
		exec: func(op uint16) {
			k := isKeyPressed(byte((op & 0x0F00) >> 2))
			if k {
				pc++
			}
			pc++
		},
		elapsedMicroseconds: 73,
		name:                "Ex9E: SKP Vx",
		description:         "Skip next instruction if key with the value of Vx is pressed.",
	},
	{
		matches: func(op uint16) bool {
			return op&0xF0FF == 0xE0A1
		},
		exec: func(op uint16) {
			k := isKeyPressed(byte((op & 0x0F00) >> 2))
			if !k {
				pc++
			}
			pc++
		},
		elapsedMicroseconds: 73,
		name:                "ExA1: SKNP Vx",
		description:         "Skip next instruction if key with the value of Vx is not pressed.",
	},
	{
		matches: func(op uint16) bool {
			return op&0xF0FF == 0xF007
		},
		exec: func(op uint16) {
			rx := numToReg(byte((op & 0x0F00) >> 2))
			*rx = *dt
			pc++
		},
		elapsedMicroseconds: 45,
		name:                "Fx07: LD Vx, DT",
		description:         "Set Vx = delay timer value.",
	},
	{
		matches: func(op uint16) bool {
			return op&0xF0FF == 0xF00A
		},
		exec: func(op uint16) {
			rx := numToReg(byte((op & 0x0F00) >> 2))
			keyWaiting = true
			*rx = <-keyPress
			pc++
		},
		elapsedMicroseconds: 0,
		name:                "Fx0A: LD Vx, K",
		description:         "Wait for a key press, store the value of the key in Vx.",
	},
	{
		matches: func(op uint16) bool {
			return op&0xF0FF == 0xF015
		},
		exec: func(op uint16) {
			rx := numToReg(byte((op & 0x0F00) >> 2))
			*dt = *rx
			pc++
		},
		elapsedMicroseconds: 45,
		name:                "Fx15: LD DT, Vx",
		description:         "Set delay timer = Vx.",
	},
	{
		matches: func(op uint16) bool {
			return op&0xF0FF == 0xF018
		},
		exec: func(op uint16) {
			rx := numToReg(byte((op & 0x0F00) >> 2))
			*st = *rx
			pc++
		},
		elapsedMicroseconds: 45,
		name:                "Fx18: LD ST, Vx",
		description:         "Set sound timer = Vx.",
	},
	{
		matches: func(op uint16) bool {
			return op&0xF0FF == 0xF01E
		},
		exec: func(op uint16) {
			rx := numToReg(byte((op & 0x0F00) >> 2))
			i += uint16(*rx)
			pc++
		},
		elapsedMicroseconds: 86,
		name:                "Fx1E: ADD I, Vx",
		description:         "Set I = I + Vx.",
	},
	{
		matches: func(op uint16) bool {
			return op&0xF0FF == 0xF029
		},
		exec: func(op uint16) {
			rx := numToReg(byte((op & 0x0F00) >> 2))
			i = byteToFontLoc(*rx)
			pc++
		},
		elapsedMicroseconds: 91,
		name:                "Fx29: LD F, Vx",
		description:         "Set I = location of sprite for digit Vx.",
	},
	{
		matches: func(op uint16) bool {
			return op&0xF0FF == 0xF033
		},
		exec: func(op uint16) {
			rx := numToReg(byte((op & 0x0F00) >> 2))
			if *rx >= 100 {
				hundreds := *rx / 100
				mem[i] = intToHex(int(hundreds))
			}
			if *rx >= 10 {
				hundreds := *rx / 100
				tens := (*rx - hundreds) / 10
				mem[i+1] = intToHex(int(tens))
			}
			if *rx >= 1 {
				hundreds := *rx / 100
				tens := (*rx - hundreds) / 10
				ones := *rx - hundreds - tens
				mem[i+2] = intToHex(int(ones))
			}
			pc++
		},
		elapsedMicroseconds: 927,
		name:                "Fx33: LD B, Vx",
		description:         "Store BCD representation of Vx in memory locations I, I+1, and I+2.",
	},
	{
		matches: func(op uint16) bool {
			return op&0xF0FF == 0xF055
		},
		exec: func(op uint16) {
			maxReg := byte((op & 0x0F00) >> 2)
			for itr := byte(0); itr <= maxReg; itr++ {
				rx := numToReg(itr)
				mem[i+uint16(itr)] = *rx
			}
		},
		elapsedMicroseconds: 605,
		name:                "Fx55: LD [I], Vx",
		description:         "Store registers V0 through Vx in memory starting at location I.",
	},
	{
		matches: func(op uint16) bool {
			return op&0xF0FF == 0xF065
		},
		exec: func(op uint16) {
			maxReg := byte((op & 0x0F00) >> 2)
			for itr := byte(0); itr <= maxReg; itr++ {
				rx := numToReg(itr)
				*rx = mem[i+uint16(itr)]
			}
		},
		elapsedMicroseconds: 605,
		name:                "Fx65: LD Vx, [I]",
		description:         "Read registers V0 through Vx from memory starting at location I.",
	},
}
