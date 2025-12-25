// Package weather is a package that provides a function to forecast the weather for a given city.
package weather

var (
	// CurrentCondition is the current weather condition for the given city.
	CurrentCondition string
	// CurrentLocation is the current location for the given city.
	CurrentLocation string
)

// Forecast function returns the current weather condition for the given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
