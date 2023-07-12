package customer

import (
	"TavernDDD/aggregates"
	"errors"
	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound     = errors.New("the customer was not found in the repository")
	ErrFailedAddCustomer    = errors.New("failed to add the customer")
	ErrFailedUpdateCustomer = errors.New("failed to update the customer")
)

type Repository interface {
	Get(uuid uuid.UUID) (aggregates.Customer, error)
	Add(aggregates.Customer) error
	Update(aggregates.Customer) error
}
