package memory

import (
	"TavernDDD/aggregates"
	"TavernDDD/domain/customer"
	"errors"
	"github.com/google/uuid"
	"testing"
)

func TestMemory_GetCustom(t *testing.T) {
	type testCase struct {
		name          string
		id            uuid.UUID
		expectedError error
	}
	cust, err := aggregates.NewCustomer("Goku")
	if err != nil {
		t.Fatal(err)
	}

	id := cust.GetId()

	repo := MemoryRepository{
		customers: map[uuid.UUID]aggregates.Customer{
			id: cust,
		},
	}

	testCases := []testCase{
		{
			name:          "Not customer by id",
			id:            uuid.MustParse("519e68a5-6c1d-4668-8232-8d0dedad9771"),
			expectedError: customer.ErrCustomerNotFound,
		},
		{
			name:          "Customer by id",
			id:            id,
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := repo.Get(tc.id)
			if !errors.Is(err, tc.expectedError) {
				t.Errorf("Expected error %v, got %v", tc.expectedError, err)
			}
		})
	}
}
