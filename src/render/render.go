package render

import (
	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/pixelgl"
	"golang.org/x/image/colornames"
)

func Run() {
	cfg := pixelgl.WindowConfig{
		Title: "Mandelbrot Fractal",
		//Icon:                   nil,
		Bounds: pixel.R(0, 0, 1920, 1080),
		//Position:               pixel.Vec{},
		//Monitor:                nil,
		//Resizable:              false,
		//Undecorated:            false,
		//NoIconify:              false,
		//AlwaysOnTop:            false,
		//TransparentFramebuffer: false,
		VSync:       true,
		Maximized:   true,
		Invisible:   false,
		SamplesMSAA: 8,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}
	win.Clear(colornames.Black)
	getMandelbrotSprite().Draw(win, pixel.IM.Moved(win.Bounds().Center()))

	for !win.Closed() {
		win.Update()
	}
}
