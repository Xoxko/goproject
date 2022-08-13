package chip8_CPU

func (s *steck) push(pc uint16) {
	s.steck[s.sp] = pc & 0x0FFF
	s.sp++
}

func (s *steck) pop() {

}
