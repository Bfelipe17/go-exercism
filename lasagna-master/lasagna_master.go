package lasagna

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avgPrepTime int) int {
	if avgPrepTime == 0 {
		avgPrepTime = 2
	}

	return len(layers) * avgPrepTime
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (int, float64) {
	var noodles int
	var sauceLayer int
	const noodlesPerLayer = 50

	for _, layer := range layers {
		if layer == "noodles" {
			noodles++
		} else if layer == "sauce" {
			sauceLayer++
		}
	}

	return noodles * noodlesPerLayer, float64(sauceLayer) * 0.2
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friendList []string, myList []string) {
	(myList)[len(myList)-1] = friendList[len(friendList)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portion int) []float64 {
	newList := make([]float64, len(quantities))

	portionVal := float64(portion) / 2.0
	for index, val := range quantities {
		newList[index] = val * portionVal
	}

	return newList
}

// Your first steps could be to read through the tasks, and create
// these functions with their correct parameter lists and return types.
// The function body only needs to contain `panic("")`.
//
// This will make the tests compile, but they will fail.
// You can then implement the function logic one by one and see
// an increasing number of tests passing as you implement more
// functionality.
