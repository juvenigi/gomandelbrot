package math

type MyComplex complex128

func (c MyComplex) GetAbsSquared() float64 {
	return real(c)*real(c) + imag(c)*imag(c)
}
