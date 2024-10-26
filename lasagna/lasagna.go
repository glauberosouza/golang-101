package main

import "fmt"

func main(){
	fmt.Println("Remaining Oven Time:", RemainingOvenTime(10), "minutes")
	fmt.Println("Preparation Time:", PreparationTime(2), "minutes")
	fmt.Println("Elapsed Time:", ElapsedTime(2, 10), "minutes")
}

const OvenTime = 40

func RemainingOvenTime(actualMinutesInOven int) int {
    return OvenTime - actualMinutesInOven
}

func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * 2
}

func ElapsedTime(numberOfLayers, actualMinutesInOven int) int {
	return  PreparationTime(numberOfLayers) + actualMinutesInOven
}
