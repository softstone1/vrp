package data

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/softstone1/vrp/algorithm"
)

func ExtractLoads(file io.Reader ) ([]algorithm.Load, error) {
	var loads []algorithm.Load
	scanner := bufio.NewScanner(file)
	scanner.Scan() // skipp the header

	for scanner.Scan() {
		line := scanner.Text()
		load, err := extractLoad(line)
		if err != nil {
			return nil, fmt.Errorf("error extracting load data: %v", err)
		}
		loads = append(loads, *load)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning file: %v", err)
	}
	return loads, nil
}

func extractLoad(line string) (*algorithm.Load, error) {
	parts := strings.Fields(line)
	if len(parts) < 3 {
		return nil, errors.New("invalid data format")
	}
	number, err := strconv.Atoi(parts[0])
	if err != nil {
		return nil, fmt.Errorf("failed to parse load id: %v", err)
	}
	pickupLocation, err := parseLocation(parts[1])
	if err != nil {
		return nil, err
	}
	dropoffLocation, err := parseLocation(parts[2])
	if err != nil {
		return nil, err
	}
	return &algorithm.Load{
		Number:  number,
		Pickup:  *pickupLocation,
		DropOff: *dropoffLocation,
	}, nil

}

func parseLocation(location string) (*algorithm.Location, error) {
	coordinates := strings.Split(strings.Trim(location, "()"), ",")
	x, err := strconv.ParseFloat(coordinates[0], 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse location coordinate x: %v", err)
	}
	y, err := strconv.ParseFloat(coordinates[1], 64)
	if err != nil {
		return nil, fmt.Errorf("failed to parse location coordinate y: %v", err)
	}
	return &algorithm.Location{X: x, Y: y}, nil
}