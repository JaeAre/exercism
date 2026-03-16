// Package weather provides tools to manage weather conditions and forecasts.
package weather

var (
	// CurrentCondition holds the current weather condition as a string.
	CurrentCondition string
	// CurrentLocation holds the name of the current location as a string.
	CurrentLocation string
)

// Forecast returns a string describing the current weather condition for a given city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
