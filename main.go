package main

import (
	"fmt"
	"goproject/src"
)

func main() {
	chip := src.InitCPU("PONG")
	fmt.Println(chip.OpcodeRun())
	for i := 0; i < 140; i++ {
		s := chip.OpcodeText()
		if s != "1" {
			fmt.Println(s)
		}
	}
}
