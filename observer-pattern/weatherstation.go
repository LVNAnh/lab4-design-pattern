package main

import (
	"fmt"
)

type WeatherStation struct {
	observers   []Observer
	temperature float64
}

func NewWeatherStation() *WeatherStation {
	return &WeatherStation{
		observers: make([]Observer, 0),
	}
}

func (w *WeatherStation) RegisterObserver(observer Observer) {
	w.observers = append(w.observers, observer)
}

func (w *WeatherStation) RemoveObserver(observer Observer) {
	for i, obs := range w.observers {
		if obs == observer {
			w.observers = append(w.observers[:i], w.observers[i+1:]...)
			break
		}
	}
}

func (w *WeatherStation) NotifyObservers() {
	for _, observer := range w.observers {
		observer.Update(w.temperature)
	}
}

func (w *WeatherStation) SetTemperature(temp float64) {
	fmt.Printf("Trạm thời tiết: Nhiệt độ đã thay đổi thành %.1f°C\n", temp)
	w.temperature = temp
	w.NotifyObservers()
}
