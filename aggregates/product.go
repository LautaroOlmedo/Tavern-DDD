package aggregates

import (
	"TavernDDD/entity"
	"errors"
	"github.com/google/uuid"
)

var (
	ErrInvalidItem = errors.New("the name and description fields are required")
)

type Product struct {
	item     *entity.Item
	price    float64
	quantity int
}

func NewProduct(name, description string, price float64) (Product, error) {
	if name == "" || description == "" {
		return Product{}, ErrInvalidItem
	}
	return Product{
		item: &entity.Item{
			ID:          uuid.New(),
			Name:        name,
			Description: description,
		},
		price:    price,
		quantity: 0,
	}, nil
}
func (p Product) GetId() uuid.UUID {
	return p.item.ID
}

func (p Product) setId(id uuid.UUID) {
	p.item.ID = id
}

func (p Product) GetPrice() float64 {
	return p.price
}

func (p Product) GetItem() *entity.Item {
	return p.item
}
