package entities

import "math"

type Point struct {
	X, Y float64
}

func (p *Point) DistanceTo(other Point) float64 {
	return math.Sqrt(math.Pow(p.X-other.X, 2) + math.Pow(p.Y-other.Y, 2))
}
