package main

import (
	cpu "CHIP8/chip8_CPU"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

const CHIP_8_WIDTH int32 = 64
const CHIP_8_HEIGHT int32 = 32

func main() {
	chip := cpu.InitCPU()
	var midifick int32 = 10

	var name string = ""
	if sdlErr := sdl.Init(sdl.INIT_EVERYTHING); sdlErr != nil {
		panic("sdl not Init sisteam")
	}
	defer sdl.Quit()

	window, windowerr := sdl.CreateWindow("chip8 + "+name, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, CHIP_8_WIDTH*midifick, CHIP_8_HEIGHT*midifick, sdl.WINDOW_SHOWN)
	if windowerr != nil {
		panic("load window error sisteam")
	}
	defer window.Destroy()

	canvas, canvasErr := sdl.CreateRenderer(window, -1, 0)
	if canvasErr != nil {
		panic("Create Render error")
	}
	defer canvas.Destroy()
	canvas.SetDrawColor(255, 0, 0, 100)
	canvas.Clear()

	for eventT := sdl.PollEvent(); eventT != nil; eventT = sdl.PollEvent() {
		switch et := eventT.(type) {
		case *sdl.QuitEvent:
			os.Exit(0)
		case *sdl.KeyboardEvent:
			if et.Type == sdl.KEYUP {
				switch et.Keysym.Sym {
				case sdl.K_1:
					chip.Key(0x1, false)
				case sdl.K_2:
					chip.Key(0x2, false)
				case sdl.K_3:
					chip.Key(0x3, false)
				case sdl.K_4:
					chip.Key(0xC, false)
				case sdl.K_q:
					chip.Key(0x4, false)
				case sdl.K_w:
					chip.Key(0x5, false)
				case sdl.K_e:
					chip.Key(0x6, false)
				case sdl.K_r:
					chip.Key(0xD, false)
				case sdl.K_a:
					chip.Key(0x7, false)
				case sdl.K_s:
					chip.Key(0x8, false)
				case sdl.K_d:
					chip.Key(0x9, false)
				case sdl.K_f:
					chip.Key(0xE, false)
				case sdl.K_z:
					chip.Key(0xA, false)
				case sdl.K_x:
					chip.Key(0x0, false)
				case sdl.K_c:
					chip.Key(0xB, false)
				case sdl.K_v:
					chip.Key(0xF, false)
				}
			} else if et.Type == sdl.KEYDOWN {
				switch et.Keysym.Sym {
				case sdl.K_1:
					chip.Key(0x1, true)
				case sdl.K_2:
					chip.Key(0x2, true)
				case sdl.K_3:
					chip.Key(0x3, true)
				case sdl.K_4:
					chip.Key(0xC, true)
				case sdl.K_q:
					chip.Key(0x4, true)
				case sdl.K_w:
					chip.Key(0x5, true)
				case sdl.K_e:
					chip.Key(0x6, true)
				case sdl.K_r:
					chip.Key(0xD, true)
				case sdl.K_a:
					chip.Key(0x7, true)
				case sdl.K_s:
					chip.Key(0x8, true)
				case sdl.K_d:
					chip.Key(0x9, true)
				case sdl.K_f:
					chip.Key(0xE, true)
				case sdl.K_z:
					chip.Key(0xA, true)
				case sdl.K_x:
					chip.Key(0x0, true)
				case sdl.K_c:
					chip.Key(0xB, true)
				case sdl.K_v:
					chip.Key(0xF, true)
				}
			}
		}

	}
}
