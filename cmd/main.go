package main

import (
	"fmt"
	"taxi/internal/pricing"
	"taxi/internal/tariff"
	"taxi/internal/traffic"
	"taxi/internal/trip"
	"taxi/internal/weather"
	"time"
)

func main() {
	calculator := pricing.PriceCalculator{
		TrafficClient: &traffic.RealTrafficClient{}, // В продакшене используется настоящий клиент, а мы подключим структуру-заглушку для имитации его работы
	}

	price := calculator.CalculatePrice(
		trip.TripParameters{Distance: 8.5, Duration: 20},
		time.Now(),
		weather.WeatherData{Condition: weather.HeavyRain, WindSpeed: 10},
		55.751244, 37.618423,
		tariff.Comfort,
	)

	fmt.Printf("Ваша цена: %.0f руб.\n", price)
}
