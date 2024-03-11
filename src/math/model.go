package math

type IteratorCell struct {
	Value     MyComplex
	Iteration uint32
}

type IteratorRect struct {
	Cells []IteratorCell
	Width uint
}

// CoordinateBounds is a struct that holds the bounds of the Mandelbrot set
type CoordinateBounds struct {
	/*
	   +-----------+-----------+
	   |    Second |   First   |
	   |  Quadrant |  Quadrant |
	   |           |           |
	   +-----------0---(real)-->
	   |    Third  |  Fourth   |
	   |  Quadrant |  Quadrant |
	   |           |           |
	   +-----------------------+
	*/
	ThirdQuadrant complex128
	FirstQuadrant complex128
}
