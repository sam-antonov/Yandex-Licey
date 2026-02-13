package tariff

type TaxiTariff string

const (
	Economy TaxiTariff = "Эконом"
	Comfort TaxiTariff = "Комфорт"
	Business TaxiTariff = "Бизнес"
)

func GetTariffMultiplier(clientTariff TaxiTariff) float64 {
	if clientTariff == Economy {
		return 1
	} else if clientTariff == Comfort {
		return 1.2
	} else {
		return 1.5
	}
}