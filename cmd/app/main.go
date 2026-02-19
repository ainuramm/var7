package main

import (
	"fmt"

	"var7/pkg/fueltrip"

	"github.com/fatih/color"
	"github.com/google/uuid"
)

func main() {

	tripID := uuid.New()
	color.Cyan("Trip ID: %s\n", tripID.String())

	distance := 300.0
	consumption := 8.0
	price := 1.5

	fuel, err := fueltrip.FuelNeeded(distance, consumption)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = fueltrip.AddTrafficReserve(&fuel, 10)
	if err != nil {
		fmt.Println(err)
		return
	}

	cost, err := fueltrip.TripCost(fuel, price)
	if err != nil {
		fmt.Println(err)
		return
	}

	report, err := fueltrip.FormatTripReport("City A → City B", fuel, cost)
	if err != nil {
		fmt.Println(err)
		return
	}

	color.Green("\n=== Trip Report ===")
	fmt.Printf("%s", report)
}