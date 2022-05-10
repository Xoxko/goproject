package main

import (
	cpu "CHIP8/chip8_CPU"
)

func main() {
	chip := cpu.InitCPU()
	chip.OpcodeRun(0x3333)
	chip.OpcodeRun(0x4333)
	chip.OpcodeRun(0x5333)
	chip.OpcodeRun(0x6333)
	chip.OpcodeRun(0x7333)
	chip.OpcodeRun(0x8333)
	chip.OpcodeRun(0x9333)
	chip.OpcodeRun(0xA333)
	chip.OpcodeRun(0xB333)
}
