package main

import (
	"fmt"
	"goproject/src"
)

func main() {
	chip := src.InitCPU("PONG")

	fmt.Println(chip.OpcodeRun())
	fmt.Println(chip.OpcodeRun())
	fmt.Println(chip.OpcodeRun())
	fmt.Println(chip.OpcodeRun())
	fmt.Println(chip.OpcodeRun())
	fmt.Println(chip.OpcodeRun())
	fmt.Println(chip.OpcodeRun())
	fmt.Println(chip.OpcodeRun())
	fmt.Println(chip.OpcodeRun())
	fmt.Println(chip.OpcodeRun())
}
