package main

import (
	"fmt"
)

type TemperatureDisplay struct {
	id string
}

func NewTemperatureDisplay(id string) *TemperatureDisplay {
	return &TemperatureDisplay{id: id}
}

func (d *TemperatureDisplay) Update(temperature float64) {
	fmt.Printf("Màn hình %s: Hiển thị nhiệt độ mới: %.1f°C\n", d.id, temperature)
}

type PhoneApp struct {
	username string
}

func NewPhoneApp(username string) *PhoneApp {
	return &PhoneApp{username: username}
}

func (p *PhoneApp) Update(temperature float64) {
	fmt.Printf("Ứng dụng (%s): Thông báo - Nhiệt độ hiện tại là %.1f°C\n", p.username, temperature)
}
