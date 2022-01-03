package src

import "github.com/veandco/go-sdl2/sdl"

const (
	X = 64 * 6
	Y = 32 * 6
)

type Display struct {
	win *sdl.Window
	sur *sdl.Surface
}

func (D *Display) InitDis() {

	err := sdl.Init(sdl.INIT_VIDEO)
	if err != nil {
		panic(err)
	}

	D.win, err = sdl.CreateWindow("test",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		X, Y, sdl.WINDOW_SHOWN,
	)
	if err != nil {
		panic(err)
	}

	D.sur, err = D.win.GetSurface()
	if err != nil {
		panic(err)
	}
}

func (D *Display) Spr(x, y int32) {
	rect := sdl.Rect{x * 6, y * 6, 6, 6}
	D.sur.FillRect(&rect, 0xffff0000)
	D.win.UpdateSurface()
}

func (D *Display) Free() {
	rect := sdl.Rect{0, 0, X, Y}
	D.sur.FillRect(&rect, 0x0)
	D.win.UpdateSurface()
}

func (D *Display) Close() {
	D.win.Destroy()
	sdl.Quit()
}
