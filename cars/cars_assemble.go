package main

func main()  {
	
}

func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
    carPerHour := float64(productionRate) * successRate /100
	return  carPerHour
}

func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	carsByHour := CalculateWorkingCarsPerHour(productionRate, successRate)
    carByMinutes := int(carsByHour) / 60
    return carByMinutes
}

func CalculateCost(carsCount int) uint {
    groupsOfTen := carsCount / 10
    individualCars := carsCount % 10

    cost := (groupsOfTen * 95000) + (individualCars * 10000)
    return uint(cost)
}