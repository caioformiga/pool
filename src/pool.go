package src

import (
	"fmt"
	"log"
	"math"
	"strings"

	"github.com/caioformiga/pool/utils"
)

// Pool is an AMM that offers liquidity based on rules
// 		i)  k = base *	quote      (fixed value)
//		ii) price = quote / base   (variable)
type Pool struct {
	BaseAmount  float64
	QuoteAmount float64
	K           float64
}

func NewPool(baseAmount float64, quoteAmount float64) IPool {
	return &Pool{
		BaseAmount:  baseAmount,
		QuoteAmount: quoteAmount,
		K:           baseAmount * quoteAmount,
	}
}

// CommitByPrice commits new base and quote amounts to config pool on informed price
func (pool *Pool) CommitByPrice(priceToCommit float64) (err error) {
	log.Printf("before commit: %+v", pool.ToString())

	// calculate
	newBaseAmount, newQuoteAmount, err := pool.CalcNewPoolByPrice(priceToCommit)
	if err != nil {
		return err
	}

	// commit
	pool.BaseAmount = newBaseAmount
	pool.QuoteAmount = newQuoteAmount

	log.Printf("aftert commit: %+v", pool.ToString())
	return nil
}

// CalcNewPoolByPrice calculates new base and quote amounts to move pool to informed price
// 		new base amount = square matrix of k / price
//      new quote amount = k / new base amount
func (pool *Pool) CalcNewPoolByPrice(price float64) (newBaseAmount, newQuoteAmount float64, err error) {
	if price == float64(0) {
		err = fmt.Errorf(ErrMsgInvalidPrice)
		return
	}

	// calculate
	newBaseAmount = math.Sqrt(pool.K / price)
	newQuoteAmount = pool.K / newBaseAmount
	return
}

// CommitByAmount commits new base and quote amounts to config pool after move informed amount from base
func (pool *Pool) CommitByAmount(amountToCommit float64, side string) (err error) {
	log.Printf("before commit: %+v", pool.ToString())

	// calculate
	newBaseAmount, newQuoteAmount, err := pool.CalcNewPoolByAmount(amountToCommit, side)
	if err != nil {
		return err
	}

	// update
	pool.BaseAmount = newBaseAmount
	pool.QuoteAmount = newQuoteAmount

	log.Printf("aftert commit: %+v", pool.ToString())
	return
}

// CalcNewPoolByAmount calculates new base and quote amounts after move informed amout from base
//      new base amount = base amount - amount to commit (BUY) or base amount + amount to commit (SELL)
//      new quote amount = k / new  base amount
func (pool *Pool) CalcNewPoolByAmount(amount float64, side string) (newBaseAmount, newQuoteAmount float64, err error) {
	if amount == float64(0) {
		err = fmt.Errorf(ErrMsgInvalidAmount)
		return
	}

	// calculate
	newBaseAmount = pool.BaseAmount + amount
	if strings.EqualFold(side, "SELL") {
		newBaseAmount = pool.BaseAmount - amount
	}
	newQuoteAmount = pool.K / newBaseAmount

	return
}

// CaPoolPrice returns calculated pool price using base and quote amount
func (pool *Pool) CalcPoolPrice() float64 {
	if pool.BaseAmount == float64(0) {
		return float64(0)
	}

	if pool.QuoteAmount == float64(0) {
		return float64(0)
	}

	return pool.QuoteAmount / pool.BaseAmount
}

// CalcK returns calculated K using base and quote amount
func (pool *Pool) CalcK() float64 {
	if pool.BaseAmount == float64(0) {
		return float64(0)
	}

	if pool.QuoteAmount == float64(0) {
		return float64(0)
	}

	return pool.QuoteAmount * pool.BaseAmount
}

// GetBaseAmount returns base amount
func (pool *Pool) GetBaseAmount() float64 {
	return pool.BaseAmount
}

// GetQuoteAmount returns quote amount
func (pool *Pool) GetQuoteAmount() float64 {
	return pool.QuoteAmount
}

// GetK returns k
func (pool *Pool) GetK() float64 {
	return pool.K
}

// ToString returns pool base, quote and price as string formatted values with min digitis precision
func (pool *Pool) ToString() (s string) {
	s = fmt.Sprintf("{base: %s, quote: %s, price: %s",
		utils.FormatMinPrecisionToString(pool.GetBaseAmount()),
		utils.FormatMinPrecisionToString(pool.GetQuoteAmount()),
		utils.FormatMinPrecisionToString(pool.CalcPoolPrice()),
	)
	return s
}
