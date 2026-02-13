package utils

const (
	minPrice = 99.0
	maxPrice = 20000
)

func ApplyPriceLimits(price float64) float64 {
	switch {
	case price < minPrice:
		return minPrice
	case price > maxPrice:
		return maxPrice
	default:
		return price
	}
}