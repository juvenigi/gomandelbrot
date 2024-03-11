package render

import (
	"github.com/gopxl/pixel"
	"github.com/gopxl/pixel/pixelgl"
	"golang.org/x/image/colornames"
	"gomandelbrot/src/math"
)

func Run() {
	cfg := pixelgl.WindowConfig{
		Title:  "Pixel Rocks!",
		Bounds: pixel.R(0, 0, 1024, 768),
		VSync:  true,
	}
	win, err := pixelgl.NewWindow(cfg)
	if err != nil {
		panic(err)
	}

	//abspath, _ := filepath.Abs("res/image/achievement.png")
	//abspath, _ := filepath.Abs("res/image/serrogo.jpg")
	//pic, err := loadPicture(abspath)
	//if err != nil {
	//	panic(err)
	//}

	rect := pixel.MakePictureData(pixel.R(0, 0, 300, 300))

	for i, _ := range rect.Pix {
		rect.Pix[i] = colornames.Black
	}

	win.Clear(colornames.White)
	coords := pixel.R(-2, -1.5, .5, 1.5)
	math.UpdatePictureData(rect, &coords)
	pixel.NewSprite(rect, rect.Bounds()).Draw(win, pixel.IM.Moved(win.Bounds().Center()))

	for !win.Closed() {
		win.Update()
	}
}
