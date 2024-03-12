package render

import "image/color"

func SetColor(c *color.RGBA, iteration uint32) {
	c.R = uint8(iteration % 255)
	c.G = uint8(iteration % 255)
	c.B = uint8(iteration % 255)
	c.A = 255
}
