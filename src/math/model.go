package math

import "github.com/gopxl/pixel"

type Complex struct {
	Im float64
	Re float64
}

type IteratorCell struct {
	Value     Complex
	Iteration uint32
}

type IteratorRect struct {
	Cells []IteratorCell
	Width uint
}

// todo implement
func CreateIteratorRect(rect *pixel.PictureData) {
	//height := rect.Bounds().H()
	//width := rect.Bounds().W()
	//rect.Pix
}

func (cell *IteratorCell) IterateUntilDone() {

}
