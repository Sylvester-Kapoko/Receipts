package domain

import (
	"time"

	"github.com/shopspring/decimal"
)

type ReceiptItem struct {
	Name      string
	Qty       int
	UnitPrice decimal.Decimal
}

type Receipt struct {
	ID        string
	StoreName string
	StoreAddr string
	Items     []ReceiptItem
	TaxRate   decimal.Decimal
	Payment   struct {
		Method string
		Amount decimal.Decimal
	}
	CreatedAt time.Time
}

