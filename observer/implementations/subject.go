package implementations

import "github.com/k-cao/design-patterns-golang/observer/interfaces"

type WeatherStation struct {
	observers   []interfaces.Observer
	weatherData interfaces.WeatherData
}

func (w *WeatherStation) RegisterObserver(o interfaces.Observer) {
	w.observers = append(w.observers, o)
}

func (w *WeatherStation) RemoveObserver(o interfaces.Observer) {
	for i, obs := range w.observers {
		if obs == o {
			w.observers = append(w.observers[:i], w.observers[i+1:]...)
			return
		}
	}
}

func (w *WeatherStation) NotifyObservers() {
	for _, obs := range w.observers {
		obs.Update(w.weatherData)
	}
}

func (w *WeatherStation) measurementsChanged() {
	w.NotifyObservers()
}

func (w *WeatherStation) SetMeasurements(d interfaces.WeatherData) {
	w.weatherData = d
	w.measurementsChanged()
}
