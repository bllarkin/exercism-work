package lasagna

func PreparationTime(layers []string, avgTime int) int {
	if avgTime == 0 {
		avgTime = 2
	}

	return len(layers) * avgTime
}

func Quantities(layers []string) (int, float64) {
	totalNoodles := 0
	totalSauce := 0

	for i := 0; i < len(layers); i++ {
		if layers[i] == "noodles" {
			totalNoodles++
		} else if layers[i] == "sauce" {
			totalSauce++
		}
	}

	return int(totalNoodles * 50), float64(float64(totalSauce) * 0.2)
}

func AddSecretIngredient(friendsList []string, myList []string) {
	myList[len(myList)-1] = friendsList[len(friendsList)-1]
}

func ScaleRecipe(origAmounts []float64, portions int) []float64 {
	var scaledAmounts []float64

	for i := 0; i < len(origAmounts); i++ {
		scaledAmounts = append(scaledAmounts, origAmounts[i]*(float64(portions)/2.0))
	}

	return scaledAmounts
}
