// Package weather provides information about the weather.
package weather

// CurrentCondition is the current weather condition e..g "sunny".
var CurrentCondition string

// CurrentLocation is the current location e.g. "Sydney".
var CurrentLocation string

// Forecast gives a forecast for a city given the current conditions.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}
