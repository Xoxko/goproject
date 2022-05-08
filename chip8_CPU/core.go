package chip8

import (
	"fmt"
	"math/rand"
	"os"
)

type cpu struct {
	dis Display
	ALU alu
	mem []uint16
}

type alu struct {
	steck  [16]byte
	vx     [16]byte
	st, vf byte
	pc, I  uint16
}

func errorFunc(s string, err error) {
	if err != nil {
		fmt.Println(s)
	}
}
func openFile(s string, C *cpu) {
	file, err := os.Open(s)
	errorFunc("Open File", err)
	defer file.Close()
	stat, err := file.Stat()
	errorFunc("Stat File", err)
	str := make([]byte, stat.Size())
	file.Read(str)

	var pos uint16
	for i := 1; i < int(stat.Size()); i += 2 {
		pos = uint16(str[i-1]) << 8
		pos += uint16(str[i])
		C.mem[0x1ff+(i+1)/2] = pos
		pos = 0
	}
}

func InitCPU(s string) *cpu {
	// Open file read
	c := cpu{mem: make([]uint16, 0xfff)}
	c.ALU.pc = 0x200
	openFile(s, &c)
	c.dis.InitDis()
	return &c
}

func (c *cpu) OpcodeRun() string {
	var comand string

	mem := c.mem
	pc := &c.ALU.pc
	opcode := uint16(mem[*pc])

	X := (0x0F00 & opcode) >> 8
	Y := (0x00F0 & opcode) >> 4
	NNN := (0x0FFF & opcode)

	switch opcode & 0xF000 {
	case 0x0000:
		c.dis.Free()
		*pc++
		comand = fmt.Sprintf("> %X -- 0x00E0-CLS", *pc)
		return comand

	case 0x1000:
		*pc = opcode & 0x0FFF
		comand = fmt.Sprintf("> %X -- %X-JP", *pc, opcode)
		return comand

	case 0x2000:
		c.ALU.steck[c.ALU.st] = byte(*pc)
		c.ALU.st++
		*pc = opcode & 0x0FFF
		comand = fmt.Sprintf("> %X -- %X-CALL", *pc, opcode)
		return comand

	case 0x3000:
		if c.ALU.vx[X] == byte(opcode&0x00FF) {
			*pc += 2
			comand = fmt.Sprintf("> %X -- %X-SE+2", *pc, opcode)
			return comand
		} else {
			*pc++
			comand = fmt.Sprintf("> %X -- %X-SE+1", *pc, opcode)
			return comand
		}

	case 0x4000:
		if c.ALU.vx[X] != byte(opcode&0x00FF) {
			*pc += 2
			comand = fmt.Sprintf("> %X -- %X-SE+2", *pc, opcode)
			return comand
		} else {
			*pc++
			comand = fmt.Sprintf("> %X -- %X-SE+1", *pc, opcode)
			return comand
		}

	case 0x5000:
		if c.ALU.vx[X] == c.ALU.vx[Y] {
			*pc += 2
			comand = fmt.Sprintf("> %X -- %X-SE+2", *pc, opcode)
			return comand
		} else {
			*pc++
			comand = fmt.Sprintf("> %X -- %X-SE+1", *pc, opcode)
			return comand
		}

	case 0x6000:
		c.ALU.vx[X] = byte(mem[*pc])
		*pc++
		comand = fmt.Sprintf("> %X -- %X-SET", *pc, opcode)
		return comand

	case 0x7000:
		c.ALU.vx[X] += byte(mem[*pc])
		*pc++
		comand = fmt.Sprintf("> %X -- %X-SET", *pc, opcode)
		return comand

	case 0x8000:
		//==================================================================================
		switch opcode & 0x000F {
		case 0x0000:
			c.ALU.vx[X] = c.ALU.vx[Y]
			*pc++
			comand = fmt.Sprintf("> %X -- %X-OR", *pc, opcode)
			return comand

		case 0x0001:
			c.ALU.vx[X] |= c.ALU.vx[Y]
			*pc++
			comand = fmt.Sprintf("> %X -- %X-OR", *pc, opcode)
			return comand

		case 0x0002:
			c.ALU.vx[X] &= c.ALU.vx[Y]
			*pc++
			comand = fmt.Sprintf("> %X -- %X-AND", *pc, opcode)
			return comand

		case 0x0003:
			c.ALU.vx[X] ^= c.ALU.vx[Y]
			*pc++
			comand = fmt.Sprintf("> %X -- %X-XOR", *pc, opcode)
			return comand

		case 0x0004:
			buf := uint16(c.ALU.vx[X]) + uint16(c.ALU.vx[Y])
			*pc++
			c.ALU.vx[X] = uint8(0x00FF & buf)
			c.ALU.vf = uint8((0xFF00 & buf) >> 8)
			comand = fmt.Sprintf("> %X -- %X-ADD(vx,vf)", *pc, opcode)
			return comand

		case 0x0005:
			buf := uint16(c.ALU.vx[X]) - uint16(c.ALU.vx[Y])
			c.ALU.vx[X] = uint8(0x00FF & buf)
			*pc++
			if c.ALU.vx[X] >= c.ALU.vx[Y] {
				c.ALU.vf = 1
			} else {
				c.ALU.vf = 0
			}
			comand = fmt.Sprintf("> %X -- %X-SUB(vx,vf)", *pc, opcode)
			return comand

		case 0x0006:
			c.ALU.vf = c.ALU.vx[X] & 0x1
			*pc++
			c.ALU.vx[X] >>= 1
			comand = fmt.Sprintf("> %X -- %X-SHR(vx,vf)", *pc, opcode)
			return comand

		case 0x0007:
			if c.ALU.vx[Y] > c.ALU.vx[X] {
				c.ALU.vf = 1
			} else {
				c.ALU.vf = 0
			}
			c.ALU.vx[X] -= c.ALU.vx[Y]
			comand = fmt.Sprintf("> %X -- %X-SUBN(vx,vf)", *pc, opcode)
			return comand

		case 0x000E:
			c.ALU.vf = c.ALU.vx[X] & 0x1
			*pc++
			c.ALU.vx[X] <<= 1
			comand = fmt.Sprintf("> %X -- %X-SHL(vx,vf)", *pc, opcode)
			return comand
		}
		//==================================================================================
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

		// case 0xD000:
		// 	hight := opcode & 0x000F
		// 	var pixel uint16
		// 	for yline := 0; yline < int(hight); yline++ {
		// 		pixel = mem[int(c.ALU.I)+yline]
		// 		for xline := 0; xline < 8; xline++ {
		// 		}
		// 	}

		// 	return comand
	}
	return " "
}

func (c *cpu) OpcodeText() string {
	var comand string

	mem := c.mem
	pc := &c.ALU.pc
	opcode := uint16(mem[*pc])

	X := (0x0F00 & opcode) >> 8
	Y := (0x00F0 & opcode) >> 4

	switch opcode & 0xF000 {
	case 0x0000:
		*pc++
		comand = fmt.Sprintf("> %X -- 0x00E0-CLS", *pc)
		return comand

	case 0x1000:
		*pc++
		comand = fmt.Sprintf("> %X -- %X-JP", *pc, opcode)
		return comand

	case 0x2000:
		*pc++
		comand = fmt.Sprintf("> %X -- %X-CALL", *pc, opcode)
		return comand

	case 0x3000:
		if c.ALU.vx[X] == byte(opcode&0x00FF) {
			*pc++
			comand = fmt.Sprintf("> %X -- %X-SE+2==", *pc, opcode)
			return comand
		} else {
			*pc++
			comand = fmt.Sprintf("> %X -- %X-SE+1==", *pc, opcode)
			return comand
		}

	case 0x4000:
		if c.ALU.vx[X] != byte(opcode&0x00FF) {
			*pc++
			comand = fmt.Sprintf("> %X -- %X-SE+2!=", *pc, opcode)
			return comand
		} else {
			*pc++
			comand = fmt.Sprintf("> %X -- %X-SE+1!=", *pc, opcode)
			return comand
		}

	case 0x5000:
		if c.ALU.vx[X] == c.ALU.vx[Y] {
			*pc++
			comand = fmt.Sprintf("> %X -- %X-SE+2", *pc, opcode)
			return comand
		} else {
			*pc++
			comand = fmt.Sprintf("> %X -- %X-SE+1", *pc, opcode)
			return comand
		}

	case 0x6000:
		*pc++
		comand = fmt.Sprintf("> %X -- %X-SET", *pc, opcode)
		return comand

	case 0x7000:
		*pc++
		comand = fmt.Sprintf("> %X -- %X-SET", *pc, opcode)
		return comand

	case 0x8000:
		//==================================================================================
		switch opcode & 0x000F {
		case 0x0000:
			*pc++
			comand = fmt.Sprintf("> %X -- %X-OR", *pc, opcode)
			return comand

		case 0x0001:
			*pc++
			comand = fmt.Sprintf("> %X -- %X-OR", *pc, opcode)
			return comand

		case 0x0002:
			*pc++
			comand = fmt.Sprintf("> %X -- %X-AND", *pc, opcode)
			return comand

		case 0x0003:
			*pc++
			comand = fmt.Sprintf("> %X -- %X-XOR", *pc, opcode)
			return comand

		case 0x0004:
			*pc++
			comand = fmt.Sprintf("> %X -- %X-ADD(vx,vf)", *pc, opcode)
			return comand

		case 0x0005:
			*pc++
			comand = fmt.Sprintf("> %X -- %X-SUB(vx,vf)", *pc, opcode)
			return comand

		case 0x0006:
			*pc++
			comand = fmt.Sprintf("> %X -- %X-SHR(vx,vf)", *pc, opcode)
			return comand

		case 0x0007:
			*pc++
			comand = fmt.Sprintf("> %X -- %X-SUBN(vx,vf)", *pc, opcode)
			return comand

		case 0x000E:
			*pc++
			comand = fmt.Sprintf("> %X -- %X-SHL(vx,vf)", *pc, opcode)
			return comand
		}
		//==================================================================================
	case 0x9000:
		*pc++
		comand = fmt.Sprintf("> %X -- %X-SNE", *pc, opcode)
		return comand

	case 0xA000:
		*pc++
		comand = fmt.Sprintf("> %X -- %X-LD", *pc, opcode)
		return comand

	case 0xB000:
		*pc++
		comand = fmt.Sprintf("> %X -- %X-JP", *pc, opcode)
		return comand

	case 0xC000:
		*pc++
		comand = fmt.Sprintf("> %X -- %X-RND", *pc, opcode)
		return comand

	case 0xD000:
		*pc++
		comand = fmt.Sprintf("> %X -- %X-WIND", *pc, opcode)
		return comand

	}
	*pc++
	return "1"
}
