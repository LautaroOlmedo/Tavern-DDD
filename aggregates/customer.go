package aggregates

import (
	"TavernDDD/entity"
	"TavernDDD/valueObject"
	"errors"
	"github.com/google/uuid"
)

var (
	ErrInvalidPerson = errors.New("a customer has to have a valid name")
)

type Customer struct {
	person       *entity.Person
	products     []*entity.Item
	transactions []valueObject.Transaction
}

// NewCustomer is a factory to create a new Customer aggregate
func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}
	person := &entity.Person{
		ID:   uuid.New(),
		Name: name,
	}

	return Customer{
		person:       person,
		products:     make([]*entity.Item, 0),
		transactions: make([]valueObject.Transaction, 0),
	}, nil

}

func (c *Customer) GetId() uuid.UUID {
	return c.person.ID
}

func (c *Customer) setId(id uuid.UUID) {
	c.person.ID = id
}
