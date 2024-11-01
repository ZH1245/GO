package main

import (
	"cars-assemble/assemble"
	"fmt"
)

func main() {
	var successRate, carsProducedPerHour, numberOfCarsToProduce int
	// Read input from the user until newline
	fmt.Print("Enter success rate and card produced per hour: ")
	_, err := fmt.Scanln(&successRate, &carsProducedPerHour)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	perHour, err := assemble.CalculateWorkingCarsPerHour(carsProducedPerHour, successRate)
	perMin, err3 := assemble.CalculateWorkingCarsPerMinute(carsProducedPerHour, successRate)
	fmt.Println();
	if err == nil && err3 == nil {
		fmt.Printf("Number of successful cars per hour: %.1f\nNumber of successful cars per minute: %.1f\n", perHour, perMin)
	}
	fmt.Println();
	fmt.Print("Enter Number of cars to Produce: ")
	_, err2 := fmt.Scanln(&numberOfCarsToProduce)
	if err2 != nil {
		fmt.Println("Error:", err2)
		return
	}
	fmt.Println();
	res2 := assemble.CalculateCost(numberOfCarsToProduce)
	if err == nil {
		fmt.Printf("Cost for %d cars : %d\n", numberOfCarsToProduce, res2)
	}
}
