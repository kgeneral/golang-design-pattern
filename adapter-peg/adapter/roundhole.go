package adapter

type RoundHole interface {
	GetRadius() float64
	Fits(peg RoundPeg) bool
}

type RoundFloatHole struct {
	radius float64
}

func (r RoundFloatHole) GetRadius() float64 {
	return r.radius
}

func (r RoundFloatHole) Fits(peg RoundPeg) bool {
	return r.GetRadius() >= peg.GetRadius()
}

func NewRoundHole(radius float64) RoundHole {
	return RoundFloatHole{radius: radius}
}
