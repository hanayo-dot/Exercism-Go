// Package weather provides functions for forecasting the current weather conditions of various cities in Goblinocus.
package weather

var (
    // CurrentCondition stores the current weather condition of the selected city.
	CurrentCondition string

    // CurrentLocation stores the name of the city for which the current weather is provided.
	CurrentLocation  string
)
// Forecast returns the current weather condition for the specified city.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
