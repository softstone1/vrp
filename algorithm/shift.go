package algorithm

import (
	"cmp"
	"errors"
	"fmt"
	"slices"
	"strings"
)


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
	DescSortLoadsByPickup(loads)
	loads[0].Assigned = true
	driver1 := NewDriver(loads[0])
	loads[len(loads)/2].Assigned = true
	driver2 := NewDriver(loads[len(loads)/2])
	loads[len(loads)-1].Assigned = true
	driver3 := NewDriver(loads[len(loads)-1])
	return &Shift{
		Loads:   loads,
		Drivers: []Driver{driver1, driver2, driver3},
	}, nil
}

func (s *Shift) AddDriver(driver Driver) {
	s.Drivers = append(s.Drivers, driver)
}

func (s *Shift) NextLoad() bool {
	var maxTime float64 = 12*60 // 12 hours time limit
	var newDriverTimeFactor float64 = 500
	var nextLoad *Load
	var nextDriver *Driver
	var minPickupTime float64 = -1

	for i, load := range s.Loads {
		if load.Assigned {
			continue
		}
		// add a new driver
		minPickupTime = newDriverTimeFactor + load.Pickup.TimeToDepot()
		nextLoad = &s.Loads[i]
	}

	for i, driver := range s.Drivers {
		if driver.TotalTime >= maxTime {
			continue
		}
		for j, load := range s.Loads {
			if load.Assigned {
				continue
			}
			pickupTime := driver.LastLoad().DropOff.TimeToLocation(load.Pickup)

			if driver.TotalTime - driver.LastLoad().DropOff.TimeToDepot() + pickupTime  + load.Duration() + load.DropOff.TimeToDepot() > maxTime {
				continue
			}
			if minPickupTime == -1 || minPickupTime > pickupTime {
				minPickupTime = pickupTime
				nextLoad = &s.Loads[j]
				nextDriver = &s.Drivers[i]
			}
		}
	}

	if nextLoad != nil {
		nextLoad.Assigned = true
		if nextDriver != nil {
			nextDriver.AddLoad(*nextLoad)
		} else {
			s.AddDriver(NewDriver(*nextLoad))
		}
		return true
	}
	return false // no more available load
}

func (s *Shift) Output() string {
	var sb strings.Builder
	for _, driver := range s.Drivers {
		
		sb.WriteString("[")
		for i, load := range driver.Loads {
			if i == 0 {
				sb.WriteString(fmt.Sprintf("%d", load.Number))
				continue
			}
			sb.WriteString(fmt.Sprintf(",%d", load.Number))
		}
		sb.WriteString("]\n")
	}
	return sb.String()
}

func DescSortLoadsByPickup(loads []Load) {
	slices.SortStableFunc(loads, func(l1, l2 Load) int {
		return cmp.Compare(l2.Pickup.TimeToDepot(), l1.Pickup.TimeToDepot())
	})
}
