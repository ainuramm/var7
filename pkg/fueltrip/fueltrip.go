package fueltrip

import "fmt"

func FuelNeeded(distanceKm, litersPer100 float64) (float64, error) {
	if distanceKm <= 0 {
		return 0, fmt.Errorf("distance must be positive")
	}
	if litersPer100 <= 0 {
		return 0, fmt.Errorf("fuel consumption must be positive")
	}

	return distanceKm * litersPer100 / 100, nil
}

func TripCost(fuelLiters, pricePerLiter float64) (float64, error) {
	if fuelLiters <= 0 {
		return 0, fmt.Errorf("fuel must be positive")
	}
	if pricePerLiter <= 0 {
		return 0, fmt.Errorf("price must be positive")
	}

	return fuelLiters * pricePerLiter, nil
}

func AddTrafficReserve(fuel *float64, percent float64) error {
	if fuel == nil {
		return fmt.Errorf("fuel pointer is nil")
	}
	if percent < 0 || percent > 100 {
		return fmt.Errorf("percent must be between 0 and 100")
	}

	*fuel += *fuel * percent / 100
	return nil
}

func FormatTripReport(route string, fuel, cost float64) (string, error) {
	if route == "" {
		return "", fmt.Errorf("route cannot be empty")
	}
	if fuel <= 0 || cost <= 0 {
		return "", fmt.Errorf("fuel and cost must be positive")
	}

	report := fmt.Sprintf(
		"Route: %s\nFuel: %.2f liters\nCost: %.2f\n",
		route, fuel, cost,
	)

	return report, nil
}
