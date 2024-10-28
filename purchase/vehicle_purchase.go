package main

import "fmt"

func main()  {
	fmt.Println(CalculateResellPrice(40000, 3))
}

func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck"{
        return true
        }
    return false
}

func ChooseVehicle(option1, option2 string) string {
	if option1 < option2 {
        return fmt.Sprintf("%s is clearly the better choice.", option1)
    }
    	return fmt.Sprintf("%s is clearly the better choice.", option2)
    
}

func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
        originalPrice = originalPrice * 80 /100
    } else if age >= 3 && age < 10 {
       originalPrice = originalPrice * 70 /100
    } else if age >= 10 {
        originalPrice = originalPrice * 50 /100
    }
    return originalPrice
}