package internal

type StrategyContext struct {
	strategy DiscountStrategy
}

func StrategyContextHandler(strategy DiscountStrategy) StrategyContext {
	return StrategyContext{strategy: strategy}
}

func (c *StrategyContext) CalculateDiscountPrice(price float32) (float32, error) {
	dp, err := c.strategy.Calculate(price)
	return dp, err
}
