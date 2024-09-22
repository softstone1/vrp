package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/softstone1/vrp/solution"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("problem file is required")
		return
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Printf("error opening problem file: %v\n", err)
		return
	}
	defer file.Close()
	var loads []solution.Load
	scanner := bufio.NewScanner(file)
	scanner.Scan() // skipp the header

	for scanner.Scan() {
		line := scanner.Text()
		load, err := extractLoad(line)
		if err != nil {
			fmt.Printf("error extracting load data: %v\n", err)
			return
		}
		loads = append(loads, *load)
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("error scanning file: %v\n", err)
	}
	shift, err := solution.NewShift(loads)
	if err != nil {
		fmt.Printf("error creating a new shift: %v\n", err)
		return
	}
	for shift.NextLoad(){}
	for _, driver := range shift.Drivers{
		var sb strings.Builder
		sb.Grow(len(driver.Loads))
		sb.WriteString("[")
		for i, load := range driver.Loads {
			if i == 0 {
				sb.WriteString(fmt.Sprintf("%d", load.Number))
				continue
			} 
			sb.WriteString(fmt.Sprintf(",%d", load.Number))
		}
		sb.WriteString("]")
		fmt.Println(sb.String())
	}
}

func extractLoad(line string) (*solution.Load, error) {
	parts := strings.Fields(line)
	if len(parts) < 3 {
		return nil, errors.New("invalid data format")
	}
	number, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, fmt.Errorf("failed to parse load id: %v", err)
	}
	pickupLocation := strings.Split(strings.Trim(parts[1], "()"), ",")
	pickupX, err := strconv.ParseFloat(pickupLocation[0], 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse location: %v", err)
	}
	pickupY, err := strconv.ParseFloat(pickupLocation[1], 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse location: %v", err)
	}
	dropoffLocation := strings.Split(strings.Trim(parts[2], "()"), ",")
	dropoffX, err := strconv.ParseFloat(dropoffLocation[0], 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse location: %v", err)
	}
	dropoffY, err := strconv.ParseFloat(dropoffLocation[1], 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse location: %v", err)
	}
	return &solution.Load{
		Number:  number,
		Pickup:  solution.Location{X: pickupX, Y: pickupY},
		DropOff: solution.Location{X: dropoffX, Y: dropoffY},
	}, nil

}
