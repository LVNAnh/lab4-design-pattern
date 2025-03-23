package main

import (
	"fmt"
)

func main() {
	weatherStation := NewWeatherStation()

	display1 := NewTemperatureDisplay("Phòng khách")
	display2 := NewTemperatureDisplay("Văn phòng")
	phoneApp := NewPhoneApp("NguyenVanA")

	weatherStation.RegisterObserver(display1)
	weatherStation.RegisterObserver(display2)
	weatherStation.RegisterObserver(phoneApp)

	weatherStation.SetTemperature(25.5)

	fmt.Println("\n--- Sau khi gỡ bỏ màn hình văn phòng ---")
	weatherStation.RemoveObserver(display2)

	weatherStation.SetTemperature(26.8)
}
