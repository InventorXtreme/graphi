package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"unsafe"
)

type WindowSize struct {
	width  int32
	height int32
}

func PixelSet(pixarray []byte, pitch int, x int, y int, r, g, b, a byte) {
	index := y*(pitch) + x*4
	pixarray[index] = r
	pixarray[index+1] = g
	pixarray[index+2] = b
	pixarray[index+3] = a

}

func main() {
	MAIN_WINDOW_SIZE := WindowSize{800, 600}

	fmt.Println("Started")
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
	defer sdl.Quit()
	fmt.Println("SDL INIT DONE")

	window, err := sdl.CreateWindow("test", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED, MAIN_WINDOW_SIZE.width, MAIN_WINDOW_SIZE.height, sdl.WINDOW_SHOWN)

	if err != nil {
		panic(err)
	}
	defer window.Destroy()

	surface, err := window.GetSurface()
	if err != nil {
		panic(err)
	}

	renderer, err := sdl.CreateSoftwareRenderer(surface)
	if err != nil {
		panic(err)
	}
	tex, err := renderer.CreateTexture(sdl.PIXELFORMAT_ARGB8888, sdl.TEXTUREACCESS_STREAMING, MAIN_WINDOW_SIZE.width, MAIN_WINDOW_SIZE.height)
	if err != nil {
		panic(err)
	}
	pixels, pitch, err := tex.Lock(nil)
	if err != nil {
		panic(err)
	}

	for true {
		//for u := 0; u < int(MAIN_WINDOW_SIZE.width); u++ {
		//	for v := 0; v < int(MAIN_WINDOW_SIZE.height); v++ {
		//		PixelSet(pixels, pitch, u, v, 255, 0, 255, 0)
		//	}
		//}
		PixelSet(pixels, pitch, 0, 0, 255, 0, 255, 000)

		tex.Update(nil, unsafe.Pointer(&pixels[0]), pitch)

		tex.Unlock()
		renderer.Copy(tex, nil, nil)
		window.UpdateSurface()
	}

}
