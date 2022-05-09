package chip8_CPU

import "fmt"

func (cpu *cpu) sys() error {
	return fmt.Errorf("call ")
}

//Прыгнуть на адрес
func (alu *alu) jump(asm uint16) {
	alu.pc = asm & 0x0FFF
}

//Вызвать подпрограмму
func (alu *alu) call(asm uint16) {
	alu.steck[alu.sp] = byte(alu.pc)
	alu.sp++
	alu.pc = asm & 0x0FFF
}

//Сравнение vx[X] == 00NN
func (alu *alu) se(XNN uint16) {
	X := (0x0F00 & XNN) >> 8
	if alu.vx[X] == byte(XNN&0x00FF) {
		alu.pc += 4
	} else {
		alu.pc += 2
	}
}

//Сравнение vx[X] != 00NN
func (alu *alu) sen(XNN uint16) {
	X := (0x0F00 & XNN) >> 8

	if alu.vx[X] != byte(XNN&0x00FF) {
		alu.pc += 4
	} else {
		alu.pc += 2
	}
}

//Сравнение vx[X] == vx[Y]
func (alu *alu) sevx(XY0 uint16) {
	X := (0x0F00 & XY0) >> 8
	Y := (0x00F0 & XY0) >> 4

	if alu.vx[X] == alu.vx[Y] {
		alu.pc += 4
	} else {
		alu.pc += 2
	}
}

//Вставить значание в vx[X] из 00NN
func (alu *alu) ld(XNN uint16) {
	X := (0x0F00 & XNN) >> 8
	alu.vx[X] = byte(XNN & 0x00FF)
	alu.pc += 2
}

//Сложение X = X + NN
func (alu *alu) add(XNN uint16) {
	X := (0x0F00 & XNN) >> 8
	alu.vx[X] += byte(XNN & 0x00FF)
	alu.pc += 2
}

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