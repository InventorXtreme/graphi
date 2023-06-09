package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"math"
	"unsafe"
)

// TO FINISH
// USE 0 as iso value to divide marching thing by
// need lang parser

type WindowSize struct {
	width  int32
	height int32
}

type QuadPix struct {
	ul float64
	ur float64
	ll float64
	lr float64
}

type vectorF64 struct {
	x float64
	y float64
}

type vectorI64 struct {
	x int64
	y int64
}

type vectorI struct {
	x int
	y int
}

type D2float64 struct {
	width  int
	height int
	data   [][]float64
}

func BuildD2float64(width, height int) D2float64 {
	temparray := make([][]float64, width)
	for i := range temparray {
		temparray[i] = make([]float64, height)
	}
	return D2float64{width, height, temparray}
}

func GetMinOfFloat64Array(a []float64) float64 {
	var m float64 = math.MaxFloat64
	for _, e := range a {
		if e < m {
			m = e
		}
	}
	return m
}

func (a D2float64) GetMinOfGraph() float64 {
	var m float64 = math.MaxFloat64
	for _, e := range a.data {
		if GetMinOfFloat64Array(e) < m {
			m = GetMinOfFloat64Array(e)
		}
	}

	return m
}

func GetMaxOfFloat64Array(a []float64) float64 {
	var m float64 = 0 - math.MaxFloat64
	for _, e := range a {
		if e > m {
			m = e
		}
	}
	return m
}

func (a D2float64) GetMaxOfGraph() float64 {
	var m float64 = 0 - math.MaxFloat64
	for _, e := range a.data {
		if GetMaxOfFloat64Array(e) > m {
			m = GetMaxOfFloat64Array(e)
		}
	}

	return m
}

type graphobj struct {
	scale_factor          float64
	center_point_offset_x float64
}

func convert_pix_point_to_graph_pix_point(window WindowSize, graph graphobj, vec vectorI) vectorI {
	var window_mid_w int
	window_mid_w = int(window.width) / 2

	var window_mid_h int
	window_mid_h = int(window.height) / 2

	tx := vec.x
	ty := vec.y

	tx = tx - window_mid_w
	ty = -1 * (ty - window_mid_h)

	return vectorI{tx, ty}

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
	graph := graphobj{1, 0}
	evalarea := BuildD2float64(int(MAIN_WINDOW_SIZE.width), int(MAIN_WINDOW_SIZE.height))
	for x := 0; x < evalarea.width; x++ {
		for y := 0; y < evalarea.height; y++ {
			inten := convert_pix_point_to_graph_pix_point(MAIN_WINDOW_SIZE, graph, vectorI{x, y})
			evalarea.data[x][y] = float64(inten.x-inten.y)

		}
	}

	evalareamax := evalarea.GetMaxOfGraph()
	evalareamin := evalarea.GetMinOfGraph()

	fmt.Println("evalarea max: ")
	fmt.Println(evalareamax)
	fmt.Println("evalarea min: ")
	fmt.Println(evalareamin)
	evalareamaxzeroadjusted := (-1 * evalareamin) + evalareamax

	fmt.Println("evalarea maxadj")
	fmt.Println(evalareamaxzeroadjusted)
	for true {
		for x := 0; x < int(MAIN_WINDOW_SIZE.width); x++ {
			for y := 0; y < int(MAIN_WINDOW_SIZE.height); y++ {
				setval := evalarea.data[x][y]
				t1 := setval + (-1 * (evalareamin))
				t2 := t1 / evalareamaxzeroadjusted
				if t2 > 1 {
					fmt.Println(setval, t1, evalareamaxzeroadjusted, evalareamax, t2)
				}
				trueset := int64(255 * t2)

				//fmt.Println(trueset)
				PixelSet(pixels, pitch, x, y, byte(trueset), 0, 0, 0)
				if setval > 0 {
					PixelSet(pixels, pitch, x, y, 0, 255, 0, 0)
				}
				if setval < 0 {
					PixelSet(pixels, pitch, x, y, 0, 0, 255, 0)
				}
			}
		}
		PixelSet(pixels, pitch, 0, 0, 255, 0, 255, 000)

		tex.Update(nil, unsafe.Pointer(&pixels[0]), pitch)

		tex.Unlock()
		renderer.Copy(tex, nil, nil)
		window.UpdateSurface()
	}

}
