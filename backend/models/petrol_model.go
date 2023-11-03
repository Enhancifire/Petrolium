package models

import "strconv"

type PetrolModel struct {
	ID       string  `json:"id"`
	Quantity float64 `json:"quantity"`
	Price    float64 `json:"price"`
	KmCount  int     `json:"kmCount"`
}

func NewPetrolModel(id string, quantity float64, price float64, kmCount int) *PetrolModel {
	return &PetrolModel{
		ID:       id,
		Quantity: quantity,
		Price:    price,
		KmCount:  kmCount,
	}
}

func (p *PetrolModel) toString() string {
	return "ID: " + p.ID + ", Quantity: " + strconv.FormatFloat(p.Quantity, 'f', -1, 64) + ", Price: " + strconv.FormatFloat(p.Price, 'f', -1, 64) + ", KmCount: " + strconv.Itoa(p.KmCount)
}
