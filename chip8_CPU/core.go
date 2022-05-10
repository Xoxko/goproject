package chip8_CPU

type cpu struct {
	ALU       alu
	mem       [4096]byte
	display   [32][64]uint8
	sholdDraw bool
}

type alu struct {
	steck [16]byte
	vx    [16]byte
	sp    byte //steck pointer
	vf    bool
	pc    uint16 //program counter
	i     uint16 //index register
}

func InitCPU() cpu {
	return cpu{}
}
