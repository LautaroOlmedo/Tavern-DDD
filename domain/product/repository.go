package product

import (
	"TavernDDD/aggregates"
	"errors"
	"github.com/google/uuid"
)

var (
	ErrProductNotFound     = errors.New("the product was not found in the repository")
	ErrFailedAddProduct    = errors.New("failed to add the product")
	ErrFailedUpdateProduct = errors.New("failed to update the product")
	ErrFailedDeleteProduct = errors.New("cannot delete de product")
)

type Repository interface {
	GetAll() []aggregates.Product
	GetByID(id uuid.UUID) (aggregates.Product, error)
	Add(product aggregates.Product) error
	Update(product aggregates.Product) error
	Delete(id uuid.UUID)
}
