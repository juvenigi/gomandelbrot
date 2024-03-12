package render

import (
	"github.com/gopxl/pixel"
	"golang.org/x/image/colornames"
	"gomandelbrot/src/math"
	"image/color"
)

// todo: add more visualisation configuration and functionality.

// getMandelbrotSprite returns a pixel.Sprite of the Mandelbrot set
func getMandelbrotSprite() *pixel.Sprite {
	//rect := pixel.MakePictureData(pixel.R(0, 0, 800, 200))
	rect := pixel.MakePictureData(pixel.R(0, 0, 600, 720))

	for i, _ := range rect.Pix {
		rect.Pix[i] = colornames.Black
	}
	coords := pixel.R(-2, -1.5, .5, 1.5)
	//coords := pixel.R(-2, -0.1, -1.3, 0.1)
	iterationValues := math.CalculateIterations(rect, &coords)
	for i, _ := range rect.Pix {
		setColor(&rect.Pix[i], (*iterationValues)[i].Iteration)
	}
	return pixel.NewSprite(rect, rect.Bounds())
}

// fixme: this visualisation is only temporary, as this cycles through the colors too fast.
func setColor(c *color.RGBA, iteration uint32) {
	c.R = uint8(iteration % 255)
	c.G = uint8(iteration % 255)
	c.B = uint8(iteration % 255)
	c.A = 255
}
