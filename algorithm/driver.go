package algorithm

type Driver struct {
	Loads     []Load
	TotalTime float64
	MaxTime   float64
	Done      bool
}

func NewDriver(load Load) Driver {
	totalTime := load.Pickup.TimeToDepot() + load.Duration() + load.DropOff.TimeToDepot()
	return Driver{
		Loads:     []Load{load},
		TotalTime: totalTime,
		MaxTime:   12*60,
	}
}
func (d *Driver) CanAddLoad(load Load) bool {
	if d.Done {
		return false
	}
	if d.TotalTimeWith(load) >= d.MaxTime {
		return false
	}
	return true
}

func (d *Driver) AddLoad(load Load)  {
	d.TotalTime = d.TotalTimeWith(load)
	d.Done = d.TotalTime >= d.MaxTime
	d.Loads = append(d.Loads, load)
}

func (d *Driver) TotalTimeWith(load Load) float64 {
	return d.TotalTime - d.LastLocation().TimeToDepot() +
		d.LastLocation().TimeToLocation(load.Pickup) + load.Duration() + load.DropOff.TimeToDepot()
}

func (d *Driver) LastLocation() Location {
	return d.Loads[len(d.Loads)-1].DropOff
}
