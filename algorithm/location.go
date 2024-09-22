package algorithm

import "math"

type Location struct {
	X float64
	Y float64
}

func (l Location) TimeToDepot() float64 {
	return math.Sqrt(l.X*l.X + l.Y*l.Y)
}

func (l Location) TimeToLocation(l2 Location) float64 {
	return math.Sqrt((l2.X-l.X)*(l2.X-l.X) + (l2.Y-l.Y)*(l2.Y-l.Y))
}