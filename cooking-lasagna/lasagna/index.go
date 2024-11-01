/*
In this exercise you're going to write some code to help you cook a brilliant lasagna from your favorite cooking book.

 1. Define the expected oven time in minutes
 2. Calculate the remaining oven time in minutes
 3. Calculate the preparation time in minutes
 4. Calculate the elapsed working time in minutes
*/
package lasagna

/*
Oven time in minutes (m)
*/
const ovenTime = 50

// The number of minutes the lasagna has been in the oven
func RemainingOvenTime(actualMins int) int {
	return ovenTime - actualMins
}

// Time is calculated as numberOfLayers * 2
func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * 2
}

func ElapsedTime(numberOfLayers, actualMinsInOven int) int {
	remainingTime := RemainingOvenTime(actualMinsInOven)
	timeToProcessLayers := PreparationTime(3)
	return remainingTime + timeToProcessLayers
}
