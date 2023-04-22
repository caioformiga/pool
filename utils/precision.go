package utils

import (
	"strconv"

	"github.com/shopspring/decimal"
)

func FormatPrecisionWithTruncate(value float64, precision int32) float64 {
	d := decimal.NewFromFloat(value)
	return d.Truncate(precision).InexactFloat64()
}

func FormatPrecisionToString(value float64, precision int32) string {
	d := decimal.NewFromFloat(value)
	return d.Truncate(precision).StringFixed(precision)
}

// FortmatManyPrecisionToString returns all itens on slice values as string formatted on precision decimals
// use same order as input to retrieve values
func FormatManyPrecisionToString(values []float64, floatPrecision int) (formatedValues []string) {
	formatedValues = make([]string, 0)
	for _, v := range values {
		strV := FormatPrecisionToString(v, int32(floatPrecision))
		formatedValues = append(formatedValues, strV)
	}
	return formatedValues
}

// FormatMinPrecisionToString returns value using the minimum number of necessary digits
// util 15 digits, max representation of float64
//	1.00000 = "1"
//	1.10000 = "1.1"
func FormatMinPrecisionToString(value float64) string {
	return strconv.FormatFloat(value, 'f', -1, 64)
}
