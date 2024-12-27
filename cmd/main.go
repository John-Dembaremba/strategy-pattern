package main

import (
	"fmt"

	"github.com/John-Dembaremba/strategy-pattern/internal"
)

func main() {

	noDiscountHandler := internal.NoDiscount{}
	percentageDisc := 20
	percentageDiscountHandler := internal.PercentageDiscount{
		Percentage: float32(percentageDisc),
	}
	fixedDisc := 30
	fixedDiscountHandler := internal.FixedDiscount{
		FixedPrice: float32(fixedDisc),
	}

	var price float32
	fmt.Print("Enter price:$ ")
	_, err := fmt.Scan(&price)
	if err != nil {
		fmt.Println("Invalid input. Please enter a numeric value.")
		return
	}
	fmt.Printf("Discount calculations for price: $%.2f\n", price)

	noDiscountContext := internal.StrategyContextHandler(noDiscountHandler)
	dp, err := noDiscountContext.CalculateDiscountPrice(float32(price))
	if err != nil {
		fmt.Printf("NoDiscount calculation encounted error: %v\n", err)
	} else {
		fmt.Printf("NoDiscount price: $%.2f\n", dp)
	}
	fmt.Println("=============================================")

	percentageDiscContext := internal.StrategyContextHandler(percentageDiscountHandler)
	fmt.Printf("Percentage Discount value: %v\n", percentageDisc)
	dp, err = percentageDiscContext.CalculateDiscountPrice(float32(price))
	if err != nil {
		fmt.Printf("Percentage discount calculation encounted error: %v\n", err)
	} else {
		fmt.Printf("Percentage discount price: $%.2f\n", dp)
	}
	fmt.Println("=============================================")

	fixedDiscContext := internal.StrategyContextHandler(fixedDiscountHandler)
	fmt.Printf("Fixed Discount value: %v\n", fixedDisc)
	dp, err = fixedDiscContext.CalculateDiscountPrice(float32(price))
	if err != nil {
		fmt.Printf("Fixed discount calculation encounted error: %v\n", err)
	} else {
		fmt.Printf("Fixed discount price: $%.2f\n", dp)
	}

}
