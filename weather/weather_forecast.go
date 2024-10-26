// Package weather provides current weather information.
package main

func main()  {
	
}

// CurrentCondition describes the current weather condition.
var CurrentCondition string

// CurrentLocation holds the current location for the weather condition.
var CurrentLocation string

// Forecast sets the location and weather condition, returning a formatted forecast.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}