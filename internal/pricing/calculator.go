package pricing

import (
	"math"
	"time"
	"taxi/internal/taxi_time"
	"taxi/internal/traffic"
	"taxi/internal/trip"
	"taxi/internal/utils"
	"taxi/internal/weather"
	"taxi/internal/tariff"
)

type PriceCalculator struct {
	TrafficClient traffic.TrafficClient
}

func (c *PriceCalculator) CalculatePrice(
	clientTrip trip.TripParameters,
	now time.Time,
	currentWeather weather.WeatherData,
	lat, lng float64,
	tar tariff.TaxiTariff,
) float64 {
	base := trip.CalculateBasePrice(clientTrip)
	timeMult := taxi_time.GetTimeMultiplier(now)
	weatherMult := weather.GetWeatherMultiplier(currentWeather)
	trafficMult := traffic.GetTrafficMultiplier(c.TrafficClient.GetTrafficLevel(lat, lng))
	tariffMult := tariff.GetTariffMultiplier(tar)

	finalPrice := base * timeMult * weatherMult * trafficMult * tariffMult
	return utils.ApplyPriceLimits(math.Round(finalPrice))
}