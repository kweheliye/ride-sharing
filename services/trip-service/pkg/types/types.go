package types

type OsrmApiResponse struct {
	Routes []struct {
		Distance float64 `json:"distance"`
		Duration float64 `json:"duration"`
		Geometry struct {
			Coordinates [][]float64 `json:"coordinates"`
		} `json:"geometry"`
	} `json:"routes"`
}

type PricingConfig struct {
	PricePerUnitOfDistance float64
	PricingPerMinute       float64
}

func DefaultPricingConfig() *PricingConfig {
	return &PricingConfig{
		PricePerUnitOfDistance: 1.5,
		PricingPerMinute:       0.25,
	}
}
