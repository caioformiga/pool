package src

import (
	"fmt"
	"log"

	"github.com/caioformiga/pool/utils"
)

type Validator struct {
	pool IPool
}

func NewValidator(p IPool) IValidator {
	return &Validator{
		pool: p,
	}
}

// ValidatePool returns true if pool.K is equal to product of base and quote amount
// it considering informed precisions as work-around to handle numeric issues related
// to float representation; for highest precision of 9 (base precision + quote precision)
// 10 == 10.000000001
func (v *Validator) ValidatePool(basePrecision, quotePrecision int32) bool {
	// define the highest precision (base + quote) to any product of base and quote
	p := basePrecision + quotePrecision

	expectedK := utils.FormatPrecisionWithTruncate(v.pool.CalcK(), p)
	match := expectedK == v.pool.GetK()

	if !match {
		strMsgExpectedK := utils.FormatPrecisionToString(expectedK, p)
		logMsg := fmt.Sprintf("K is not equal to current calculated K is '%s'", strMsgExpectedK)
		log.Default().Println(logMsg)
	}
	return match
}
