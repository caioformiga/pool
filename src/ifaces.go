package src

// IPool is an AMM that offers liquidity based on algorithm that use K as base * quote
type IPool interface {
	// CalcNewPoolByPrice calculates new base and quote amounts to move pool to informed price
	// 		new base amount = square matrix of k / price
	//      new quote amount = k / new base amount
	CalcNewPoolByPrice(price float64) (newBaseAmount, newQuoteAmount float64, err error)

	// CommitByPrice commits new base and quote amounts to config pool on informed price
	CommitByPrice(priceToCommit float64) (err error)

	// CalcNewPoolByAmount calculates new base and quote amounts after move informed amout from base
	//      new base amount = base amount - amount to commit (BUY) or base amount + amount to commit (SELL)
	//      new quote amount = k / new  base amount
	CalcNewPoolByAmount(amount float64, side string) (newBaseAmount, newQuoteAmount float64, err error)

	// CommitByAmount commits new base and quote amounts to config pool after move informed amount from base
	CommitByAmount(amountToCommit float64, side string) (err error)

	// CalcK returns calculated K using base and quote amount
	CalcK() float64

	// CalcPoolPrice returns calculated pool price using base and quote amount
	CalcPoolPrice() float64

	// GetBaseAmount returns base amount
	GetBaseAmount() float64

	// GetQuoteAmount returns quote amount
	GetQuoteAmount() float64

	// GetK returns k
	GetK() float64

	// ToString returns pool base, quote and price as string formatted values with min digitis precision
	ToString() string
}

// IValidator define all rules for validate pool
type IValidator interface {
	// ValidatePool check if pool.K respect informed precision. In case, basePrecision 5 and quotePrecision 4
	// 9 decimals means 10 == 10.000000001
	ValidatePool(basePrecision, quotePrecision int32) bool
}
