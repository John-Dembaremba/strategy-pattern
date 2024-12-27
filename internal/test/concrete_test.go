package test

import (
	"errors"
	"testing"

	"github.com/John-Dembaremba/strategy-pattern/internal"
)

type TestCases struct {
	name                  string
	price                 float32
	expectedDiscountPrice float32
	expectedErr           error
}

func TestConcreteImplementation(t *testing.T) {

	assertValues := func(t testing.TB, got, want float32) {
		t.Helper()
		if got != want {
			t.Errorf("expected discount %v, got %v", want, got)
		}
	}

	assertErrors := func(t testing.TB, got, want error) {
		t.Helper()
		if want != nil {
			switch got {
			case nil:
				t.Error("expected error, got nil")
			default:
				if want.Error() != got.Error() {
					t.Errorf("expected error %v, got %v", want, got)
				}
			}
		}
	}

	t.Run("no discount", func(t *testing.T) {
		testCases := []TestCases{
			{
				name:                  "success",
				price:                 120,
				expectedDiscountPrice: 120,
				expectedErr:           nil,
			},
			{
				name:                  "error - negative price",
				price:                 -10,
				expectedDiscountPrice: 0,
				expectedErr:           errors.New("negative price not allowed"),
			},
		}

		noDiscountHandler := internal.NoDiscount{}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				got, err := noDiscountHandler.Calculate(tc.price)
				assertValues(t, got, tc.expectedDiscountPrice)
				assertErrors(t, err, tc.expectedErr)
			})
		}
	})

	t.Run("percentage discount", func(t *testing.T) {
		testCases := []TestCases{
			{
				name:                  "success",
				price:                 100,
				expectedDiscountPrice: 70,
				expectedErr:           nil,
			},
			{
				name:                  "error - negative price",
				price:                 -100,
				expectedDiscountPrice: 0.0,
				expectedErr:           errors.New("negative price not allowed"),
			},
		}

		percentageP := float32(30)
		percDiscountHandler := internal.PercentageDiscount{
			Percentage: percentageP,
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				got, err := percDiscountHandler.Calculate(tc.price)
				assertValues(t, got, tc.expectedDiscountPrice)
				assertErrors(t, err, tc.expectedErr)
			})
		}

	})

	t.Run("fixed discount", func(t *testing.T) {
		testCases := []TestCases{
			{
				name:                  "success",
				price:                 100,
				expectedDiscountPrice: 80,
				expectedErr:           nil,
			},
			{
				name:                  "error - negative price",
				price:                 -100,
				expectedDiscountPrice: 0,
				expectedErr:           errors.New("negative price not allowed"),
			},
			{
				name:                  "error - excess fixed dicsount",
				price:                 10,
				expectedDiscountPrice: 0,
				expectedErr:           errors.New("fixed discount is greater than price"),
			},
		}

		fixedPrice := float32(20)
		handler := internal.FixedDiscount{
			FixedPrice: fixedPrice,
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				got, err := handler.Calculate(tc.price)
				assertValues(t, got, tc.expectedDiscountPrice)
				assertErrors(t, err, tc.expectedErr)
			})
		}
	})
}
