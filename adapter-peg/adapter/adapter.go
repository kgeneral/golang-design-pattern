package adapter

import "math"

type SquarePegAdapter struct {
	peg SquarePeg
}

func NewSquarePegAdapter(peg SquarePeg) RoundPeg {
	return SquarePegAdapter{peg: peg}
}

func (s SquarePegAdapter) GetRadius() float64 {
	return s.peg.GetWidth() * math.Sqrt(2.0) / 2
}
