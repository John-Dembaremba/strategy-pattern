package internal

type DiscountStrategy interface {
	Calculate(price float32) (float32, error)
}
