package chip8_CPU

//Возвращает все регистры
func (c *cpu_chip8) GetRegisterCPU() (vx [16]uint8, st [16]uint8, pc uint16, ic uint16, vf bool) {
	return c.vx, c.st, c.pc, c.ic, c.vf
}

// Регистр хранение данных
func (c *cpu_chip8) GetRegisterVx() (vx [16]uint8) {
	return c.vx
}

// Регистр вызовов подпрограмм
func (c *cpu_chip8) GetRegisterSteck() (st [16]uint8) {
	return c.st
}

// Регистр счетчик команды
func (c *cpu_chip8) GetRegisterPc() (pc uint16) {
	return c.pc
}

// Регистр адресов 12-bit
func (c *cpu_chip8) GetRegisterIc() (ic uint16) {
	return c.ic
}

// Регистр переполнение
func (c *cpu_chip8) GetRegisterVf() (vf bool) {
	return c.vf
}
