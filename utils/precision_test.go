package utils

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFormatManyPrecisionToString(t *testing.T) {
	assert := assert.New(t)

	t.Run("With empty slice", func(t *testing.T) {
		floatPrecision := 13

		formatedValues := FormatManyPrecisionToString([]float64{}, floatPrecision)
		assert.Len(formatedValues, 0)
	})

	t.Run("With 0 values", func(t *testing.T) {
		floatPrecision := 13

		apr := 0.0
		avgDailyRewards := 0.0
		cumulativeFee := 0.0
		initialInvestment := 0.0

		values := []float64{apr, avgDailyRewards, cumulativeFee, initialInvestment}
		formatedValues := FormatManyPrecisionToString(values, floatPrecision)
		assert.Len(formatedValues, 4)
	})

	t.Run("With many values", func(t *testing.T) {
		floatPrecision := 13

		apr := 100.0
		avgDailyRewards := apr / 365
		cumulativeFee := 0.002
		initialInvestment := 20000.00

		values := []float64{apr, avgDailyRewards, cumulativeFee, initialInvestment}
		formatedValues := FormatManyPrecisionToString(values, floatPrecision)

		assert.Equal("100.0000000000000", formatedValues[0])
		assert.Equal("0.2739726027397", formatedValues[1])
		assert.Equal("0.0020000000000", formatedValues[2])
		assert.Equal("20000.0000000000000", formatedValues[3])
	})
}

func TestUnit_FormatMinPrecisionToString(t *testing.T) {
	assert := assert.New(t)

	t.Run("without digits", func(t *testing.T) {
		value := float64(0)
		strFormatedValues := FormatMinPrecisionToString(value)
		assert.Equal("0", strFormatedValues)
		assert.NotContains(strFormatedValues, ".")
	})

	t.Run("with digits", func(t *testing.T) {
		t.Run("1 digit", func(t *testing.T) {
			value := float64(9.9)
			strFormatedValues := FormatMinPrecisionToString(value)
			assert.Contains(strFormatedValues, ".")

			parts := strings.Split(strFormatedValues, ".")
			assert.GreaterOrEqual(len(parts), 2) // parts: {"integer" ; "digits"

			digits := parts[1]
			assert.Equal(len(digits), 1)
		})

		t.Run("15 digits", func(t *testing.T) {
			value := float64(9.012345678901234)
			strFormatedValues := FormatMinPrecisionToString(value)
			assert.Contains(strFormatedValues, ".")

			parts := strings.Split(strFormatedValues, ".")
			assert.GreaterOrEqual(len(parts), 2) // parts: {"integer" ; "digits"

			digits := parts[1]
			assert.Equal(len(digits), 15)
		})
	})
}
