package algorithm

type Driver struct {
	Loads     []Load
	TotalTime float64
}

func NewDriver(load Load) Driver {
	totalTime := load.Pickup.TimeToDepot() + load.Duration() + load.DropOff.TimeToDepot()
	return Driver{
		Loads:     []Load{load},
		TotalTime:  totalTime,
	}
}

func (d *Driver) AddLoad(load Load)  {
	d.TotalTime -= d.LastLoad().DropOff.TimeToDepot()
	d.TotalTime += d.LastLoad().DropOff.TimeToLocation(load.Pickup) + load.Duration() + load.DropOff.TimeToDepot() 
	d.Loads = append(d.Loads, load)
}

func (d *Driver) LastLoad() Load {
	return d.Loads[len(d.Loads)-1]
}