package main

import (
	"fmt"
	"goproject/src"
)

func main() {
	chip := src.InitCPU("PONG")
	for i := 25; i == 0; i-- {
		fmt.Println(chip.OpcodeText())
	}
}
