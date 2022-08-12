 
	/*


		package chip8_CPU

		func (c *cpu) OpcodeRun(asm uint16) error {

			NNN := (0x0FFF & asm)

			switch asm & 0xF000 {
			case 0x0000:
				return c.sys()

			case 0x1000:
				c.ALU.jump(NNN)
				return nil

			case 0x2000:
				c.ALU.call(NNN)
				return nil

			case 0x3000:
				c.ALU.se(NNN)
				return nil

			case 0x4000:
				c.ALU.sen(NNN)
				return nil

			case 0x5000:
				c.ALU.sevx(NNN)
				return nil

			case 0x6000:
				c.ALU.ld(NNN)
				return nil

			case 0x7000:
				c.ALU.add(NNN)
				return nil
			case 0x8000:
				c.ALU.math(NNN)
				return nil

			case 0x9000:
				c.ALU.sne(NNN)
				return nil

			case 0xA000:
				c.ALU.ldd(NNN)
				return nil

			case 0xB000:
				c.ALU.jm(NNN)
				return nil

			case 0xC000:
				c.ALU.rand(NNN)
				return nil

			case 0xD000:
				c.draws(NNN)
				return nil

			case 0xE000:

				return nil

			}
			return nil
		}

	*/