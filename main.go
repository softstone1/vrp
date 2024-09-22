package main

import (
	"fmt"
	"os"

	"github.com/softstone1/vrp/algorithm"
	"github.com/softstone1/vrp/data"
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
	
	loads, err := data.ExtractLoads(file)
	if err != nil {
		fmt.Printf("error extracting loads from problem file: %v\n", err)
		return
	}

	shift, err := algorithm.NewShift(loads)
	if err != nil {
		fmt.Printf("error creating a new shift: %v\n", err)
		return
	}

	for shift.NextLoad() {} // processing all available loads
	
	fmt.Print(shift.Output())
}


