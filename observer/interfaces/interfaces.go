package interfaces

type WeatherData struct {
	Temperature float32
	Humidity    float32
	Pressure    float32
}

type Observer interface {
	Update(w WeatherData)
}

type Subject interface {
	RegisterObserver(o Observer)
	RemoveObserver(o Observer)
	NotifyObservers()
}
