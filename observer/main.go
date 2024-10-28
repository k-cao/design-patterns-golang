package main

/*
Pattern: Observer
* Objects can even decide at runtime whether they want to be kept informed
*/

import (
	"github.com/k-cao/design-patterns-golang/observer/implementations"
	"github.com/k-cao/design-patterns-golang/observer/interfaces"
)

func main() {
	weatherStation := &implementations.WeatherStation{}
	display := &implementations.ConditionsDisplay{}

	weatherStation.RegisterObserver(display)

	weatherData := interfaces.WeatherData{
		Temperature: 25.3,
		Humidity:    40.1,
		Pressure:    20.2,
	}
	weatherStation.SetMeasurements(weatherData)
	weatherData = interfaces.WeatherData{
		Temperature: 35.3,
		Humidity:    50.1,
		Pressure:    30.2,
	}
	weatherStation.SetMeasurements(weatherData)

	weatherStation.RemoveObserver(display)

	weatherData = interfaces.WeatherData{
		Temperature: 0,
		Humidity:    0,
		Pressure:    0,
	}
	weatherStation.SetMeasurements(weatherData)
}
