package main
func main()  {
	
}
func PreparationTime(layers []string, preparationTime int) int{
    if preparationTime == 0 {
        preparationTime = 2
    }
    return preparationTime * len(layers)
}
func Quantities(layers []string) (int, float64) {
    noodles := 0
    sauce := 0.0
    for _, layer := range layers {
        if layer == "noodles"{
            noodles +=50
        }
        if layer == "sauce" {
            sauce += 0.2
        }
    }
    return noodles, sauce
}

func AddSecretIngredient(friendsList []string, myList []string) {
    secretIngredient := friendsList[len(friendsList)-1]
    myList[len(myList)-1] = secretIngredient
}

func ScaleRecipe(amounts []float64, portions int) []float64 {
    scale := float64(portions) / 2.0
    scaledAmounts := make([]float64, len(amounts))
    
    for i, amount := range amounts {
        scaledAmounts[i] = amount * scale
    }
    
    return scaledAmounts
}