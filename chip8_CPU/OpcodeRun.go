package chip8_CPU

import (
	"fmt"
	"math/rand"
)

func (c *cpu) OpcodeRun(asm uint16) error {

	X := (0x0F00 & asm) >> 8
	Y := (0x00F0 & asm) >> 4
	NNN := (0x0FFF & asm)

	switch asm & 0xF000 {
	case 0x0000:
		return c.sys()

	case 0x1000:
		c.ALU.jump(NNN)
		return nil

	case 0x2000:
		c.ALU.call(NNN)
		return nil

	case 0x3000:
		c.ALU.se(NNN)
		return nil

	case 0x4000:
		c.ALU.sen(NNN)
		return nil

	case 0x5000:
		c.ALU.sevx(NNN)
		return nil

	case 0x6000:
		c.ALU.ld(NNN)
		return nil

	case 0x7000:
		c.ALU.add(NNN)
		return nil
	case 0x8000:
		c.ALU.math(asm)
		return nil

	case 0x9000:
		if c.ALU.vx[X] != c.ALU.vx[Y] {
			*pc++
		}
		*pc++
		comand = fmt.Sprintf("> %X -- %X-SNE", *pc, opcode)
		return comand

	case 0xA000:
		if c.ALU.vx[X] != c.ALU.vx[Y] {
			*pc++
		}
		*pc++
		comand = fmt.Sprintf("> %X -- %X-LD", *pc, opcode)
		return comand

	case 0xB000:
		*pc = (NNN + uint16(c.ALU.vx[0])) & 0x0FFF
		comand = fmt.Sprintf("> %X -- %X-JP", *pc, opcode)
		return comand

	case 0xC000:
		c.ALU.steck[X] = uint8(rand.Intn(256)) & uint8(0x00FF&opcode)
		comand = fmt.Sprintf("> %X -- %X-RND", *pc, opcode)
		return comand

	case 0xD000:
		return comand

	}
	return " "
}
