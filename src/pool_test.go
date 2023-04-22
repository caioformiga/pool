package src

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_CommitByPrice(t *testing.T) {
	assert := assert.New(t)

	baseAmount := float64(10000)
	quoteAmount := float64(10)

	t.Run("success", func(t *testing.T) {
		t.Run("BUY", func(t *testing.T) {
			pool := NewPool(baseAmount, quoteAmount)

			prevBaseAmount := pool.GetBaseAmount()
			prevQuoteAmount := pool.GetQuoteAmount()
			prevCurrentPrice := pool.CalcPoolPrice()

			priceToCommit := float64(0.00095)

			err := pool.CommitByPrice(priceToCommit)
			assert.Nil(err)

			// each time a buy order is executed price decrease, prev is more than current
			assert.GreaterOrEqual(prevCurrentPrice, pool.CalcPoolPrice())

			// each time a buy order is executed get more base amount, prev is less than current
			assert.LessOrEqual(prevBaseAmount, pool.GetBaseAmount())

			// each time a buy order is executed less quote amount, prev is more than current
			assert.GreaterOrEqual(prevQuoteAmount, pool.GetQuoteAmount())
		})

		t.Run("SELL", func(t *testing.T) {
			pool := NewPool(baseAmount, quoteAmount)

			prevBaseAmount := pool.GetBaseAmount()
			prevQuoteAmount := pool.GetQuoteAmount()
			prevCurrentPrice := pool.CalcPoolPrice()

			priceToCommit := float64(0.00105)

			err := pool.CommitByPrice(priceToCommit)
			assert.Nil(err)

			// each time a sell order is executed price increase, prev is less than current
			assert.LessOrEqual(prevCurrentPrice, pool.CalcPoolPrice())

			// each time a sell order is executed get less base amount, prev is more than current
			assert.GreaterOrEqual(prevBaseAmount, pool.GetBaseAmount())

			// each time a sell order is executed more quote amount, prev is less than current
			assert.LessOrEqual(prevQuoteAmount, pool.GetQuoteAmount())
		})
	})

	t.Run("unhappy path - zero price to commit", func(t *testing.T) {
		pool := NewPool(baseAmount, quoteAmount)

		err := pool.CommitByPrice(float64(0))
		assert.NotNil(err)
		assert.Contains(err.Error(), ErrMsgInvalidPrice)
	})
}

func TestUnit_CommitByAmount(t *testing.T) {
	assert := assert.New(t)

	baseAmount := float64(10000)
	quoteAmount := float64(10)

	t.Run("success", func(t *testing.T) {
		t.Run("BUY", func(t *testing.T) {
			pool := NewPool(baseAmount, quoteAmount)

			prevBaseAmount := pool.GetBaseAmount()
			prevQuoteAmount := pool.GetQuoteAmount()
			prevCurrentPrice := pool.CalcPoolPrice()

			amountToCommit := float64(0.2532)

			err := pool.CommitByAmount(amountToCommit, "BUY")
			assert.Nil(err)

			// each time a buy order is executed price decrease, prev is more than current
			assert.GreaterOrEqual(prevCurrentPrice, pool.CalcPoolPrice())

			// each time a buy order is executed get more base amount, prev is less than current
			assert.LessOrEqual(prevBaseAmount, pool.GetBaseAmount())

			// each time a buy order is executed less quote amount, prev is more than current
			assert.GreaterOrEqual(prevQuoteAmount, pool.GetQuoteAmount())
		})

		//       base 			quote		   k		 	price
		// 9753.05270549682		10.2532		100000		0.00105
		t.Run("SELL", func(t *testing.T) {
			pool := NewPool(baseAmount, quoteAmount)

			prevBaseAmount := pool.GetBaseAmount()
			prevQuoteAmount := pool.GetQuoteAmount()
			prevCurrentPrice := pool.CalcPoolPrice()

			amountToCommit := float64(0.2532)

			err := pool.CommitByAmount(amountToCommit, "SELL")
			assert.Nil(err)

			// each time a sell order is executed price increase, prev is less than current
			assert.LessOrEqual(prevCurrentPrice, pool.CalcPoolPrice())

			// each time a sell order is executed get less base amount, prev is more than current
			assert.GreaterOrEqual(prevBaseAmount, pool.GetBaseAmount())

			// each time a sell order is executed more quote amount, prev is less than current
			assert.LessOrEqual(prevQuoteAmount, pool.GetQuoteAmount())
		})
	})

	t.Run("unhappy path - zero price to commit", func(t *testing.T) {
		pool := NewPool(baseAmount, quoteAmount)

		err := pool.CommitByAmount(float64(0), "")
		assert.NotNil(err)
		assert.Contains(err.Error(), ErrMsgInvalidAmount)
	})
}

func TestUnit_CalcPoolPrice(t *testing.T) {
	assert := assert.New(t)

	t.Run("success", func(t *testing.T) {
		baseAmount := float64(10259.78352085154)
		quoteAmount := float64(9.746794344808965)

		pool := NewPool(baseAmount, quoteAmount)

		assert.Equal(float64(0.0009500000000000001), pool.CalcPoolPrice())
	})

	t.Run("success", func(t *testing.T) {
		baseAmount := float64(9759.000729485331)
		quoteAmount := float64(10.2469507659596)

		pool := NewPool(baseAmount, quoteAmount)

		assert.Equal(float64(0.0010500000000000002), pool.CalcPoolPrice())
	})
}

func TestUnit_ToString(t *testing.T) {
	assert := assert.New(t)

	baseAmount := float64(10000)
	quoteAmount := float64(10)

	pool := NewPool(baseAmount, quoteAmount)

	s := pool.ToString()
	assert.NotNil(s)

	println(s)
}
