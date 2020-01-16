package rates

//LatestRate struct
type LatestRate struct {
	Rate map[string]string `json:"rate"`
}

//RateViaDate struct
type RateViaDate struct {
	Date string            `json:"date"`
	Rate map[string]string `json:"rate"`
}

//AverageRate struct
type AverageRate struct {
	Rate map[string]Average `json:"rate"`
}

//Average struct
type Average struct {
	Min string `json:"min"`
	Max string `json:"max"`
	Avg string `json:"avg"`
}
