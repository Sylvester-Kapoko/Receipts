package domain

import (
	"fmt"

	"github.com/shopspring/decimal"
)

// Money wraps a decimal for currency display.
type Money struct {
	Amount decimal.Decimal
}

// Display formats the amount with exactly two decimal places.
func (m Money) Display() string {
	return fmt.Sprintf("$%s", m.Amount.StringFixed(2))
}


// DisplayCurrency formats a decimal as a currency string.
func DisplayCurrency(d decimal.Decimal) string {
	return fmt.Sprintf("$%s", d.StringFixed(2))
}

