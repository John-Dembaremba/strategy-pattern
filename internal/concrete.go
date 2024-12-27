package internal

import "errors"

func isPriceNegative(price float32) error {
	if price < 0 {
		return errors.New("negative price not allowed")
	}
	return nil
}

type NoDiscount struct{}

func (d NoDiscount) Calculate(price float32) (float32, error) {
	if err := isPriceNegative(price); err != nil {
		return 0, err
	}
	return price, nil
}

type PercentageDiscount struct {
	Percentage float32
}

func (p PercentageDiscount) Calculate(price float32) (float32, error) {
	if err := isPriceNegative(price); err != nil {
		return 0, err
	}
	dp := price - (p.Percentage / 100 * price)
	return dp, nil
}

type FixedDiscount struct {
	FixedPrice float32
}

func (p FixedDiscount) Calculate(price float32) (float32, error) {
	if err := isPriceNegative(price); err != nil {
		return 0, err
	}

	if price < p.FixedPrice {
		return 0, errors.New("fixed discount is greater than price")
	}
	dp := price - p.FixedPrice
	return dp, nil
}
