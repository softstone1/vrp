package algorithm

type Load struct {
	Number   int
	Pickup   Location
	DropOff  Location
	Assigned bool
}

func (l Load) Duration() float64 {
	return l.Pickup.TimeToLocation(l.DropOff)
}