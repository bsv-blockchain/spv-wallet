package utxo

import (
	"gorm.io/gorm"

	"github.com/bsv-blockchain/spv-wallet/engine/v2/transaction/outlines"
	"github.com/bsv-blockchain/spv-wallet/engine/v2/transaction/outlines/utxo/internal/sql"
	"github.com/bsv-blockchain/spv-wallet/models/bsv"
)

// NewSelector creates a new instance of UTXOSelector.
func NewSelector(db *gorm.DB, feeUnit bsv.FeeUnit) outlines.UTXOSelector {
	if db == nil {
		panic("db is required")
	}

	if !feeUnit.IsValid() {
		panic("valid fee unit is required")
	}

	return sql.NewUTXOSelector(db, feeUnit)
}
