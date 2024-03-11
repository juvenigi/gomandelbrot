package math

import (
	"github.com/gopxl/pixel"
	"image/color"
	"math"
	"math/cmplx"
	_ "math/cmplx"
	"sync"
)

func (cell *IteratorCell) IterateUntilDone(c complex128) {
	//for cmplx.Abs(cell.Value) < 2 && cell.Iteration < math.MaxUint32 {
	//	cell.Value = cell.Value*cell.Value + c
	//	cell.Iteration++
	//}

	//for getAbsSquared(cell.Value) < 4 && cell.Iteration < math.MaxUint32 {
	//	cell.Value = cell.Value*cell.Value + c
	//	cell.Iteration++
	//}

	for cell.Value.GetAbsSquared() < 4 && cell.Iteration < math.MaxUint32 {
		cell.Value = cell.Value*cell.Value + MyComplex(c)
		cell.Iteration++
	}
}

// UpdatePictureData calculates Mandelbrot set for each pixel in the picture data
// this is achieved by converting the address of rect.Pix to a complex number (treated as C constant)
func UpdatePictureData(rect *pixel.PictureData, coords *pixel.Rect) {

	resacleRe := coords.W() / rect.Rect.W()
	resacleIm := coords.H() / rect.Rect.H()

	pixW := rect.Stride
	pixSize := len(rect.Pix)

	shiftW := coords.Bounds().Min.X
	shiftH := coords.Bounds().Min.Y

	// left here so that I can use the function while debugging
	cmplx.Abs(complex(0, 0))

	var cells = make([]IteratorCell, pixSize)
	var wg sync.WaitGroup
	for i, _ := range rect.Pix {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			iterationC := complex(float64(i%pixW)*resacleRe+shiftW, float64(i/pixW)*(-resacleIm)+shiftH)
			cells[i].IterateUntilDone(iterationC)
			rect.Pix[i] = obtainColor(cells[i].Iteration)
		}(i)
	}
	wg.Wait()
}

func obtainColor(iteration uint32) color.RGBA {
	return color.RGBA{
		R: uint8(iteration % 255),
		G: uint8(iteration % 255),
		B: uint8(iteration % 255),
		A: 255,
	}
}
