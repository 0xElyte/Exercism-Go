// Package weather provides the current weather information for the city.
package weather

var (
	// CurrentCondition is a string variable that stores the weather condition.
	CurrentCondition string
	// CurrentLocation is a string  variable that stores the city's name.
	CurrentLocation  string
)

// Forecast returns the weather information of the provided city and condition.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
