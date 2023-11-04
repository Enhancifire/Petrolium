package petrol

import "time"

type CreatePetrolInput struct {
	Quantity float64   `json:"quantity"`
	Price    float64   `json:"price"`
	KmCount  int       `json:"kmCount"`
	Date     time.Time `json:"date"`
	Rate     float32   `json:"rate"`
}

type UpdatePetrolInput struct {
	Quantity float64   `json:"quantity"`
	Price    float64   `json:"price"`
	KmCount  int       `json:"kmCount"`
	Date     time.Time `json:"date"`
	Rate     float32   `json:"rate"`
}
