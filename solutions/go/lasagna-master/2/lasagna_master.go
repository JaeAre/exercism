package lasagnamaster

// TODO: define the 'PreparationTime()' function
func PreparationTime(layers []string, avg_time int) int {
	if avg_time == 0 {
		avg_time = 2
	}
	return len(layers) * avg_time
}

// TODO: define the 'Quantities()' function
func Quantities(layers []string) (noodles int, sauce float64) {
	for _, layer := range layers {
		if layer == "noodles" {
			noodles += 50
		} else if layer == "sauce" {
			sauce += 0.2
		}
	}
	return
}

// TODO: define the 'AddSecretIngredient()' function
func AddSecretIngredient(friend_list, my_list []string) {
	my_list[len(my_list)-1] = friend_list[len(friend_list)-1]
}

// TODO: define the 'ScaleRecipe()' function
func ScaleRecipe(quantities []float64, portions int) []float64 {
	scaled := make([]float64, len(quantities))
	for idx, quantity := range quantities {
		scaled[idx] = quantity * float64(portions) / 2
	}
	return scaled
}
