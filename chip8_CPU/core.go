package chip8_CPU

import "log"

//память хранения программы 0x000-0x200 зарезервированно для fontSet
type ram struct {
	memory [0x1000]uint8
}

type chip8 struct {
	MEM     ram           // Память
	CPU     cpu_chip8     // ЦПУ
	display [32][64]uint8 // Дисплей
}

type cpu_chip8 struct {
	opcode [2]uint8 // Регистр инструкции
	al     alu      // Арифметико-логическое устройство

	vx register // Регистр хранение данных
	st steck    // Регистр вызовов подпрограмм
	pc uint16   // Регистр счетчик команды
	ic uint16   // Регистр адресов 12-bit
	vf bool     // Регистр переполнение
}

type alu struct {
	vx     *register
	opcode *[2]uint8
}

//стек вызовов подпрограмм
type steck struct {
	steck [16]uint16
	sp    uint16
}

//регистры Vx and Vy
type register [16]uint8

func Init_chip8() *chip8 {
	chip := chip8{}
	chip.CPU.al.vx = &chip.CPU.vx
	chip.CPU.al.opcode = &chip.CPU.opcode
	return &chip
}

func (c *chip8) Step() {
	cp := &c.CPU
	var err error
	for i := range cp.opcode {
		cp.opcode[i], err = c.MEM.Read(cp.pc)
		if err != nil {
			log.Println(err)
		}
		cp._pc(1)
	}
}

var fontSet = [80]uint8{
	0xF0, 0x90, 0x90, 0x90, 0xF0, //0
	0x20, 0x60, 0x20, 0x20, 0x70, //1
	0xF0, 0x10, 0xF0, 0x80, 0xF0, //2
	0xF0, 0x10, 0xF0, 0x10, 0xF0, //3
	0x90, 0x90, 0xF0, 0x10, 0x10, //4
	0xF0, 0x80, 0xF0, 0x10, 0xF0, //5
	0xF0, 0x80, 0xF0, 0x90, 0xF0, //6
	0xF0, 0x10, 0x20, 0x40, 0x40, //7
	0xF0, 0x90, 0xF0, 0x90, 0xF0, //8
	0xF0, 0x90, 0xF0, 0x10, 0xF0, //9
	0xF0, 0x90, 0xF0, 0x90, 0x90, //A
	0xE0, 0x90, 0xE0, 0x90, 0xE0, //B
	0xF0, 0x80, 0x80, 0x80, 0xF0, //C
	0xE0, 0x90, 0x90, 0x90, 0xE0, //D
	0xF0, 0x80, 0xF0, 0x80, 0xF0, //E
	0xF0, 0x80, 0xF0, 0x80, 0x80, //F
}
