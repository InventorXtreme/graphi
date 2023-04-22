package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type WindowSize struct {
	width int32
	height int32
}

func main() {
	MAIN_WINDOW_SIZE := WindowSize{800,600}

	fmt.Println("Started")
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()
	fmt.Println("SDL INIT DONE")

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,MAIN_WINDOW_SIZE.width, MAIN_WINDOW_SIZE.height, sdl.WINDOW_SHOWN)

	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	renderer,err := sdl.CreateSoftwareRenderer(surface)
	if err != nil {
		panic(err)
	}
	rendertex := renderer.CreateTexture(sdl.PIXELFORMAT_ARGB8888,sdl.TEXTUREACCESS_STREAMING,MAIN_WINDOW_SIZE.width,MAIN_WINDOW_SIZE.height)

	}
