package chip8_CPU

import (
	"fmt"
)

//Возвращает все регистры
func (c cpu_chip8) GetRegisterCPU() (vx [16]uint8, st [16]uint16, pc uint16, ic uint16, vf bool) {
	return c.vx, c.st.steck, c.pc, c.ic, c.vf
}

// Регистр хранение данных
func (c cpu_chip8) GetRegisterVx() (vx [16]uint8) {
	return c.vx
}

// Регистр вызовов подпрограмм
func (c cpu_chip8) GetRegisterSteck() (st [16]uint16) {
	return c.st.steck
}

// Регистр счетчик команды
func (c cpu_chip8) GetRegisterPc() (pc uint16) {
	return c.pc
}

// Регистр адресов 12-bit
func (c cpu_chip8) GetRegisterIc() (ic uint16) {
	return c.ic
}

// Регистр переполнение
func (c cpu_chip8) GetRegisterVf() (vf bool) {
	return c.vf
}

//Установка регистра
func (c *cpu_chip8) SetPc(pc uint16) error {
	buf := pc & 0xf000
	if buf != 0 {
		return fmt.Errorf("Переполнение регистра счетчик команд %v", pc)
	}
	c.pc = pc
	return nil
}

func (c *cpu_chip8) _pc(pc uint16) {
	c.pc += (pc & 0x000F)
}

//Прыгнуть на адрес
func (c *cpu_chip8) jump() {
	buf := uint16(c.opcode[0]&0x0F) << 8
	buf |= uint16(c.opcode[1])
	c.SetPc(buf)
}

//Вызвать подпрограмму
func (c *cpu_chip8) call() {
	c.st.push(c.pc)
}

//Пропуск следующей команды если VX равен NN.
func (c *cpu_chip8) se() {
	X := c.opcode[0] & 0x0F
	NN := c.opcode[1]

	if c.vx[X] == NN {
		c._pc(2)
	}
}

//Сравнение vx[X] != 00NN
func (c *cpu_chip8) sen() {
	X := c.opcode[0] & 0x0F
	NN := c.opcode[1]

	if c.vx[X] != NN {
		c._pc(2)
	}
}

//Сравнение vx[X] == vx[Y]
func (c *cpu_chip8) sevx() {
	X := c.opcode[0] & 0x0F
	Y := (c.opcode[1] & 0xF0) >> 4

	if c.vx[X] == c.vx[Y] {
		c._pc(2)
	}
}
