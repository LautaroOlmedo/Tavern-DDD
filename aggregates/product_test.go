package aggregates_test

import (
	"TavernDDD/aggregates"
	"errors"
	"testing"
)

func Test_NewProduct(t *testing.T) {
	type testCase struct {
		test          string
		name          string
		description   string
		price         float64
		quantity      int
		expectedError error
	}

	testCases := []testCase{
		{
			test:          "empty name",
			name:          "",
			description:   "valid description",
			price:         2.1,
			quantity:      0,
			expectedError: aggregates.ErrInvalidItem,
		},
		{
			test:          "empty description",
			name:          "beer",
			description:   "",
			price:         2.1,
			quantity:      0,
			expectedError: aggregates.ErrInvalidItem},
		{
			test:          "valid test",
			name:          "beer",
			description:   "schneideeeeer",
			price:         2.1,
			quantity:      0,
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := aggregates.NewProduct(tc.name, tc.description, tc.price)
			if !errors.Is(err, tc.expectedError) {
				t.Errorf("Expected %v, got %v", tc.expectedError, err)
			}
		})
	}
}
