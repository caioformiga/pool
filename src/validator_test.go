package src

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnit_ValidatePool(t *testing.T) {
	assert := assert.New(t)

	t.Run("unhappy path - precision zero", func(t *testing.T) {
		baseAmount := float64(10000.54321)
		quoteAmount := float64(10.4321)

		pool := NewPool(baseAmount, quoteAmount)

		validator := NewValidator(pool)

		// KLV/USDT
		isValid := validator.ValidatePool(int32(0), int32(0))
		assert.False(isValid)
	})

	t.Run("success - valid precision", func(t *testing.T) {

		baseAmount := float64(10000.54321)
		quoteAmount := float64(10.4321)

		pool := NewPool(baseAmount, quoteAmount)

		validator := NewValidator(pool)

		// KLV/USDT
		isValid := validator.ValidatePool(int32(5), int32(4))
		assert.True(isValid)
	})
}
