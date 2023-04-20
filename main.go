package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

type WindowSize struct {
	width int
	height int
}

func main() {
	const MAIN_WINDOW_SIZE = WindowSize{800,600}

	fmt.Println("Started")
	sdl.Init(sdl.INIT_EVERYTHING)
}
