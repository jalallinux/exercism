package lasagna

// TODO: define the 'OvenTime' constant
const OvenTime = 40;

// RemainingOvenTime returns the remaining minutes based on the `actual` minutes already in the oven.
func RemainingOvenTime(actualMinutesInOven int) int {
    return OvenTime - actualMinutesInOven;
}

// PreparationTime calculates the time needed to prepare the lasagna based on the amount of layers.
func PreparationTime(numberOfLayers int) int {
	return numberOfLayers * 2
    
	// result := numberOfLayers * 2
 //    return result
    
	// var result = numberOfLayers * 2
 //    return result
    
	// var result int = numberOfLayers * 2
 //    return result
    
 //    var result int
 //    result = numberOfLayers * 2
	// return result
}

// ElapsedTime calculates the time elapsed cooking the lasagna. This time includes the preparation time and the time the lasagna is baking in the oven.
func ElapsedTime(numberOfLayers int, actualMinutesInOven int) int {
    return PreparationTime(numberOfLayers) + actualMinutesInOven
}
