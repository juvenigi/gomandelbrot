package render

import (
	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"gomandelbrot/src/math"
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

	//rect := pixel.MakePictureData(pixel.R(0, 0, 800, 200))
	rect := pixel.MakePictureData(pixel.R(0, 0, 600, 720))

	for i, _ := range rect.Pix {
		rect.Pix[i] = colornames.Black
	}

	win.Clear(colornames.White)
	coords := pixel.R(-2, -1.5, .5, 1.5)
	//coords := pixel.R(-2, -0.1, -1.3, 0.1)
	math.UpdatePictureData(rect, &coords)
	pixel.NewSprite(rect, rect.Bounds()).Draw(win, pixel.IM.Moved(win.Bounds().Center()))

	for !win.Closed() {
		win.Update()
	}
}
