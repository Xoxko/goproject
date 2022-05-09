package main

import (
	cpu "CHIP8/chip8_CPU"
)

func main() {
	chip := cpu.InitCPU("program/PONG")
	chip.OpcodeRun()
}
