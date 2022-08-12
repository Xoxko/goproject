package main

import (
	cpu "CHIP8/chip8_CPU"
	"fmt"
)

func main() {
	chip := cpu.Init_chip8()

	fmt.Println(chip.CPU.GetRegisterCPU())
}
