package chip8_CPU

//Вставить значание в vx[X] из 00NN
func (a *alu) ld() {
	a.vx[a.opcode[0]&0x0F] = a.opcode[1]
}

//Сложение X = X + NN
func (a *alu) add() {
	X := a.opcode[0] & 0x0F
	a.vx[X] += a.opcode[1]
}

/*
//Математические и бинарные выполнение команд
func (alu *alu) math(asm uint16) {
	X := (0x0F00 & asm) >> 8
	Y := (0x00F0 & asm) >> 4

	switch asm & 0x000F {
	case 0x0000:
		alu.vx[X] = alu.vx[Y]
		alu.pc += 2

	case 0x0001:
		alu.vx[X] |= alu.vx[Y]
		alu.pc += 2

	case 0x0002:
		alu.vx[X] &= alu.vx[Y]
		alu.pc += 2

	case 0x0003:
		alu.vx[X] ^= alu.vx[Y]
		alu.pc += 2

	case 0x0004:
		buf := uint16(alu.vx[X]) + uint16(alu.vx[Y])
		alu.pc += 2
		alu.vx[X] = uint8(0x00FF & buf)
		alu.vf = bool((0xFF00 & buf) > 255)

	case 0x0005:
		buf := uint16(alu.vx[X]) - uint16(alu.vx[Y])
		alu.vx[X] = uint8(0x00FF & buf)
		alu.pc += 2
		alu.vf = alu.vx[X] >= alu.vx[Y]

	case 0x0006:
		alu.vf = (alu.vx[X] & 0x1) == 1
		alu.vx[X] >>= 1
		alu.pc += 2

	case 0x0007:
		alu.vf = alu.vx[Y] > alu.vx[X]
		alu.vx[X] = alu.vx[Y] - alu.vx[X]

	case 0x000E:
		alu.vf = ((alu.vx[X] >> 7) & 0x1) == 1
		alu.vx[X] <<= 1
		alu.pc += 2
	}
}
*/
