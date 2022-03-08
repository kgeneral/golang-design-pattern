package adapter

type SquarePeg interface {
	GetWidth() float64
}

type SquareFloatPeg struct {
	width float64
}

func NewSquarePeg(width float64) SquarePeg {
	return SquareFloatPeg{width: width}
}

func (r SquareFloatPeg) GetWidth() float64 {
	return r.width
}
