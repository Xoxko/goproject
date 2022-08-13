package chip8_CPU

import "fmt"

func (r *ram) Write(addres uint16, data uint8) error {
	buffer := addres & 0xf000
	if buffer != 0 {
		return fmt.Errorf("Адрес %v должен находиться в дипазоне 0x0000 - 0x0fff", addres)
	}
	r.memory[addres] = data
	return nil
}

func (r *ram) Read(addres uint16) (uint8, error) {
	buffer := addres & 0xf000
	if buffer != 0 {
		return 0, fmt.Errorf("Адрес %v должен находиться в дипазоне 0x0000 - 0x0fff", addres)
	}

	return r.memory[addres], nil
}
