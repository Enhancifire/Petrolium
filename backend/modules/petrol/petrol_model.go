package petrol

import (
	"strconv"
	"time"

	"github.com/Enhancifire/Petrolium/backend/modules/storage"
	"github.com/google/uuid"
	// "github.com/jinzhu/gorm"
)

type PetrolModel struct {
	// gorm.Model
	ID        string    `json:"id" gorm:"not null; unique; primary_key"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
	Quantity  float64   `json:"quantity" gorm:"not null"`
	Price     float64   `json:"price" gorm:"not null"`
	KmCount   int       `json:"kmCount" gorm:"not null"`
	Rate      float32   `json:"rate" gorm:"not null"`
	Date      time.Time `json:"date" gorm:"not null"`
}

func (p *PetrolModel) ToString() string {
	return "Quantity: " + strconv.FormatFloat(p.Quantity, 'f', -1, 64) + ", Price: " + strconv.FormatFloat(p.Price, 'f', -1, 64) + ", KmCount: " + strconv.Itoa(p.KmCount)
}

func (p *PetrolModel) copyWith(newItem *PetrolModel) *PetrolModel {
	if newItem.Price != 0 {
		p.Price = newItem.Price
	}

	if newItem.Quantity != 0 {
		p.Quantity = newItem.Quantity
	}

	if newItem.KmCount != 0 {
		p.KmCount = newItem.KmCount
	}

	if newItem.Rate != 0 {
		p.Rate = newItem.Rate
	}

	if !newItem.Date.IsZero() {
		p.Date = newItem.Date
	}

	return p

}

func (p *PetrolModel) SavePetrol() (*PetrolModel, error) {
	if p.Date.IsZero() {
		p.Date = time.Now()
	}

	p.ID = uuid.New().String()

	err := storage.DB.Create(&p).Error
	if err != nil {
		return &PetrolModel{}, err
	}
	return p, nil
}

func FetchList() ([]PetrolModel, error) {
	var models []PetrolModel
	err := storage.DB.Find(&models).Error
	if err != nil {
		return nil, err
	}
	return models, nil
}

func FetchIndividual(id string) (*PetrolModel, error) {
	instance := PetrolModel{
		ID: id,
	}

	err := storage.DB.First(&instance).Error
	if err != nil {
		return nil, err
	}
	return &instance, nil
}

func DeletePetrolItem(id string) error {
	instance := PetrolModel{
		ID: id,
	}

	err := storage.DB.Delete(&instance).Where("id = ?", id).Error

	if err != nil {
		return err
	}

	return nil
}

func UpdatePetrolItem(id string, item UpdatePetrolInput) (*PetrolModel, error) {
	var instance PetrolModel

	if err := storage.DB.Where("id = ?", id).First(&instance).Error; err != nil {
		return nil, err
	}

	storage.DB.Model(&instance).Updates(item)
	return &instance, nil
}

func ReturnModels() []any {
	return []any{
		PetrolModel{},
	}
}
