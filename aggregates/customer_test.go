package aggregates_test

import (
	"TavernDDD/aggregates"
	"errors"
	"testing"
)

func Test_NewCustomer(t *testing.T) {
	type testCase struct {
		test          string
		name          string
		expectedError error
	}

	testCases := []testCase{
		{
			test:          "Empty name",
			name:          "",
			expectedError: aggregates.ErrInvalidPerson,
		}, {
			test:          "Valid name",
			name:          "Joe Doe",
			expectedError: nil,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.test, func(t *testing.T) {
			_, err := aggregates.NewCustomer(tc.name)
			if !errors.Is(err, tc.expectedError) {
				t.Errorf("Expected %v, got %v", tc.expectedError, err)
			}
		})
	}

}
