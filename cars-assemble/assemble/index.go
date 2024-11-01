/*
In this exercise you'll be writing code to analyze the production in a car factory.

1. Calculate the number of working cars produced per hour
2. Calculate the number of working cars produced per minute
3. Calculate the cost of production
*/
package assemble

import (
	"errors"
	"math"
)

/*
Function that takes in the number of cars produced per hour and the success rate and calculates the number of successful cars made per hour. The success rate is given as a percentage, from 0 to 100
*/
func CalculateWorkingCarsPerHour(carsProducedPerHour, successRate int) (float64, error) {
	if successRate < 0 && successRate > 100 {
		return -1, errors.New("Success Rate must be between 0 and 100")
	}
	return math.Floor(float64((successRate * carsProducedPerHour) / 100)), nil
}

/*
Implement a function that takes in the number of cars produced per hour and the success rate and calculates how many cars are successfully produced each minute:
*/
func CalculateWorkingCarsPerMinute(carsProducedPerHour, successRate int) (float64, error) {
	successfulCarsPerHour, err := CalculateWorkingCarsPerHour(carsProducedPerHour, successRate)
	if err != nil {
		return -1, err
	}
	return math.Floor(float64(successfulCarsPerHour / 60)), nil
}

/*
Each car normally costs $10,000 to produce individually, regardless of whether it is successful or not. But with a bit of planning, 10 cars can be produced together for $95,000.

For example, 37 cars can be produced in the following way: 37 = 3 x groups of ten + 7 individual cars

So the cost for 37 cars is: 3*95,000+7*10,000=355,000

Implement the function CalculateCost that calculates the cost of producing a number of cars, regardless of whether they are successful:
*/
func CalculateCost(numberOfCars int) int {
	const costOfTenCars = 95000
	const costOfOneCar = 10000
	batchOfTen := int(math.Floor(float64(numberOfCars) / 10))
	return (batchOfTen * costOfTenCars) + ((numberOfCars % 10) * costOfOneCar)
}
