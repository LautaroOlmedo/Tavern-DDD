package memory

import (
	"TavernDDD/aggregates"
	"TavernDDD/domain/customer"
	"fmt"
	"github.com/google/uuid"
	"sync"
)

type MemoryRepository struct {
	customers map[uuid.UUID]aggregates.Customer
	sync.Mutex
}

func New() *MemoryRepository {
	return &MemoryRepository{
		customers: make(map[uuid.UUID]aggregates.Customer),
	}
}

func (mr *MemoryRepository) Get(id uuid.UUID) (aggregates.Customer, error) {

	if myCustomer, ok := mr.customers[id]; ok {
		return myCustomer, nil
	}
	return aggregates.Customer{}, customer.ErrCustomerNotFound
}

func (mr *MemoryRepository) Add(c aggregates.Customer) error {
	// ---> Check if the map is initialized, if not initialize the map
	if mr.customers == nil {
		mr.Lock()
		mr.customers = make(map[uuid.UUID]aggregates.Customer)
		mr.Unlock()
	}
	// ---> Make sure customer is already in repo
	if _, ok := mr.customers[c.GetId()]; ok {
		return fmt.Errorf("customer already exists %w", customer.ErrFailedAddCustomer)
	}
	// ---> Add customer
	mr.Lock()
	mr.customers[c.GetId()] = c
	mr.Unlock()
	return nil
}

func (mr *MemoryRepository) Update(c aggregates.Customer) error {
	if _, ok := mr.customers[c.GetId()]; !ok {
		return fmt.Errorf("customer does not exists %w", customer.ErrFailedUpdateCustomer)
	}
	mr.Lock()
	mr.customers[c.GetId()] = c
	mr.Unlock()
	return nil
}
