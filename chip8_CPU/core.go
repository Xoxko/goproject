package chip8_CPU

import (
	"fmt"
	"os"
)

type cpu struct {
	ALU alu
	mem []uint16
}

type alu struct {
	steck [16]byte
	vx    [16]byte
	sp    byte //steck pointer
	vf    bool
	pc    uint16 //program counter
	i     uint16
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
	//c.dis.InitDis()
	return &c
}
