package adapter

type RoundPeg interface {
	GetRadius() float64
}

type RoundFloatPeg struct {
	radius float64
}

func NewRoundPeg(radius float64) RoundPeg {
	return RoundFloatPeg{radius: radius}
}

func (r RoundFloatPeg) GetRadius() float64 {
	return r.radius
}
