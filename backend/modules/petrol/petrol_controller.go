package petrol

import (
// "github.com/Enhancifire/Petrolium/backend/modules/storage"
)

// var petrolList = []PetrolModel{}

func CreatePetrolModel(p *CreatePetrolInput) (*PetrolModel, error) {
	petrol := PetrolModel{
		Quantity: p.Quantity,
		Price:    p.Price,
		KmCount:  p.KmCount,
		Rate:     p.Rate,
		Date:     p.Date,
	}
	data, err := petrol.SavePetrol()
	return data, err
}

func GetPetrolList() ([]PetrolModel, error) {
	list, err := FetchList()
	if err != nil {
		return nil, err
	}

	return list, nil
}

func GetPetrolIndividual(id string) (*PetrolModel, error) {
	petrol, err := FetchIndividual(id)
	if err != nil {
		return nil, err
	}
	return petrol, nil
}

func UpdatePetrolModel(id string, p UpdatePetrolInput) (*PetrolModel, error) {
	return UpdatePetrolItem(id, p)
}

func DeletePetrolModel(id string) error {
	err := DeletePetrolItem(id)
	return err
}

func NewPetrolModel(id string, quantity float64, price float64, kmCount int) *PetrolModel {
	return &PetrolModel{
		Quantity: quantity,
		Price:    price,
		KmCount:  kmCount,
	}
}
