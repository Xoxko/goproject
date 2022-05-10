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
		c.ALU.math(NNN)
		return nil

	case 0x9000:
		c.ALU.sne(NNN)
		return nil

	case 0xA000:
		c.ALU.ldd(NNN)
		return nil

	case 0xB000:
		c.ALU.jm(NNN)
		return nil

	case 0xC000:
		c.ALU.steck[X] = uint8(rand.Intn(256)) & uint8(0x00FF&opcode)
		comand = fmt.Sprintf("> %X -- %X-RND", *pc, opcode)
		return comand

	case 0xD000:
		return comand

	}
	return " "
}
