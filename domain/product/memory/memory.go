package memory

import (
	"TavernDDD/aggregates"
	"TavernDDD/domain/product"
	"fmt"
	"github.com/google/uuid"
	"sync"
)

type MemRepository struct {
	products map[uuid.UUID]aggregates.Product
	sync.Mutex
}

func New() *MemRepository {
	return &MemRepository{
		products: make(map[uuid.UUID]aggregates.Product),
	}
}

func (mr *MemRepository) GetAll() []aggregates.Product {
	var products []aggregates.Product

	for _, p := range mr.products {
		products = append(products, p)
	}
	return products
}

func (mr *MemRepository) GetByID(id uuid.UUID) (aggregates.Product, error) {
	if myProduct, ok := mr.products[id]; ok {
		return myProduct, nil
	}
	return aggregates.Product{}, product.ErrProductNotFound
}

func (mr *MemRepository) Add(p aggregates.Product) error {
	// ---> Check if the map is initialized, if not initialize the map
	if mr.products == nil {
		mr.Lock()
		mr.products = make(map[uuid.UUID]aggregates.Product)
		mr.Unlock()
	}

	// ---> Make sure product is already in repo
	if _, ok := mr.products[p.GetId()]; ok {
		return fmt.Errorf("product already exists %w", product.ErrFailedAddProduct)
	}

	// ---> Add product
	mr.Lock()
	mr.products[p.GetId()] = p
	mr.Unlock()
	return nil
}

func (mr *MemRepository) Update(p aggregates.Product) error {
	if _, ok := mr.products[p.GetId()]; !ok {
		return fmt.Errorf("product does not exists %w", product.ErrFailedUpdateProduct)
	}
	mr.Lock()
	mr.products[p.GetId()] = p
	mr.Unlock()
	return nil
}

func (mr *MemRepository) Delete(id uuid.UUID) error {
	mr.Lock()
	defer mr.Unlock()

	if _, ok := mr.products[id]; !ok {
		return product.ErrFailedDeleteProduct
	}
	delete(mr.products, id)
	return nil
}
