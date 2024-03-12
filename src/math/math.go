package math

import (
	"fmt"
	"github.com/gopxl/pixel"
	"image/color"
	"math/cmplx"
	_ "math/cmplx"
	"sync"
)

type IteratorCell struct {
	Value     MyComplex
	Iteration uint32
}

// todo: make a smarter breakoff condition (e.g. add deadline to the function)
var breakoff uint32 = 40000

// IterateUntilDone iterates over the complex number until it's abs is greater than 2.
// This is the core of the Mandelbrot fractal calculation.
func (cell *IteratorCell) IterateUntilDone(c complex128) {
	for cell.Value.GetAbsSquared() < 4 && cell.Iteration < breakoff {
		//for cmplx.Abs(complex128(cell.Value)) < 2 && cell.Iteration < math.MaxUint32 {
		// z = z^2 + c
		cell.Value = cell.Value*cell.Value + MyComplex(c)
		cell.Iteration++
	}
}

// CalculateIterations calculates Mandelbrot set for each pixel in the picture data
// this is achieved by converting the address of rect.Pix to a complex number (treated as C constant)
func CalculateIterations(rect *pixel.PictureData, coords *pixel.Rect) *[]IteratorCell {

	resacleRe := coords.W() / rect.Rect.W()
	resacleIm := coords.H() / rect.Rect.H()

	pixW := rect.Stride
	pixSize := len(rect.Pix)

	shiftW := coords.Bounds().Min.X
	shiftH := coords.Bounds().Min.Y

	fmt.Println("computing the mandelbrot set with breakoff", breakoff, "(might take a while)")
	fmt.Println("coords.Re:", coords.Min.X, coords.Max.X, "coords.Im:", coords.Min.Y, coords.Max.Y)
	fmt.Print(pixW, " pixH: ", pixSize/pixW, "resacleReFactor: ", resacleRe, " resacleImFactor: ", resacleIm, " shiftW: ", "\n")
	// left here so that I can use the function while debugging
	cmplx.Abs(complex(0, 0))

	var cells = make([]IteratorCell, pixSize)
	var wg sync.WaitGroup
	for i, _ := range rect.Pix {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			iterationC := complex(float64(i%pixW)*resacleRe+shiftW, float64(i/pixW)*resacleIm+shiftH)
			cells[i].IterateUntilDone(iterationC)
			SetColor(&rect.Pix[i], cells[i].Iteration)
		}(i)
	}
	wg.Wait()
	return &cells
}

func SetColor(c *color.RGBA, iteration uint32) {
	c.R = uint8(iteration % 255)
	c.G = uint8(iteration % 255)
	c.B = uint8(iteration % 255)
	c.A = 255
}
