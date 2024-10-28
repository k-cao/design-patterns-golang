package implementations

import (
	"fmt"

	"github.com/k-cao/design-patterns-golang/observer/interfaces"
)

type ConditionsDisplay struct {
	weatherData interfaces.WeatherData
}

func (d *ConditionsDisplay) Update(w interfaces.WeatherData) {
	d.weatherData = w
	d.display()
}

func (d *ConditionsDisplay) display() {
	fmt.Printf(
		"Current conditions: %v F degrees, %v humidity, and %v pressure\n",
		d.weatherData.Temperature, d.weatherData.Humidity, d.weatherData.Pressure)
}
