package solution

import (
	"cmp"
	"errors"
	"slices"
)

type Load struct {
	Number   int
	Pickup   Location
	DropOff  Location
	Complete bool
}

func (l Load) Duration() float64 {
	return l.Pickup.TimeToLocation(l.DropOff)
}

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

type Shift struct {
	Loads   []Load
	Drivers []Driver
}

func NewShift(loads []Load) (*Shift, error) {
	if len(loads) == 0 {
		return nil, errors.New("no load to schedule")
	}
	if len(loads) > 200 {
		return nil, errors.New("number of loads is over the limit")
	}
	SortLoadsByPickup(loads)
	loads[0].Complete = true
	firstDriver := NewDriver(loads[0])
	return &Shift{
		Loads:   loads,
		Drivers: []Driver{firstDriver},
	}, nil
}

func (s *Shift) AddDriver(driver Driver) {
	s.Drivers = append(s.Drivers, driver)
}

func (s *Shift) NextLoad() bool {
	var maxTime float64 = 720 // 12 hours time limit
	var newDriverTimeFactor float64 = 500
	var nextLoad *Load
	var nextDriver *Driver
	var minTime float64 = -1

	for i, load := range s.Loads {
		if load.Complete {
			continue
		}
		// add a new driver
		minTime = newDriverTimeFactor + load.Pickup.TimeToDepot()
		nextLoad = &s.Loads[i]
	}

	for i, driver := range s.Drivers {
		if driver.TotalTime >= maxTime {
			continue
		}
		for j, load := range s.Loads {
			if load.Complete {
				continue
			}
			time := driver.LastLoad().DropOff.TimeToLocation(load.Pickup) 

			if driver.TotalTime + time + load.Duration() + load.DropOff.TimeToDepot() - driver.LastLoad().DropOff.TimeToDepot() > maxTime {
				continue
			}
			if minTime == -1 || minTime > time {
				minTime = time
				nextLoad = &s.Loads[j]
				nextDriver = &s.Drivers[i]
			}
		}
	}

	if nextLoad != nil {
		nextLoad.Complete = true
		if nextDriver != nil {
			nextDriver.AddLoad(*nextLoad)
		} else {
			s.AddDriver(NewDriver(*nextLoad))
		}
		return true
	}
	return false
}

func SortLoadsByPickup(loads []Load) {
	slices.SortFunc(loads, func(l1, l2 Load) int {
		return cmp.Compare(l1.Pickup.TimeToDepot() + l1.Duration() +l1.DropOff.TimeToDepot(), 
		l2.Pickup.TimeToDepot()+ l2.Duration() + l2.DropOff.TimeToDepot())
	})
}
