package repository

import (
	"context"
	"iter"
	"slices"

	"github.com/samber/lo"
	"gorm.io/gorm"

	"github.com/bsv-blockchain/spv-wallet/engine/spverrors"
	"github.com/bsv-blockchain/spv-wallet/engine/v2/database"
	"github.com/bsv-blockchain/spv-wallet/engine/v2/transaction/txmodels"
	"github.com/bsv-blockchain/spv-wallet/models/bsv"
)

// Outputs is a repository for outputs.
type Outputs struct {
	db *gorm.DB
}

// NewOutputsRepo creates a new repository for outputs.
func NewOutputsRepo(db *gorm.DB) *Outputs {
	return &Outputs{db: db}
}

// FindByOutpoints returns outputs from the database based on the provided outpoints.
func (o *Outputs) FindByOutpoints(ctx context.Context, outpoints iter.Seq[bsv.Outpoint]) ([]txmodels.TrackedOutput, error) {
	outpointsClause := slices.Collect(func(yield func(sqlPair []any) bool) {
		for outpoint := range outpoints {
			yield([]any{outpoint.TxID, outpoint.Vout})
		}
	})

	var outputs []*database.TrackedOutput

	if err := o.db.WithContext(ctx).
		Model(&database.TrackedOutput{}).
		Where("(tx_id, vout) IN ?", outpointsClause).
		Find(&outputs).Error; err != nil {
		return nil, spverrors.Wrapf(err, "failed to get outputs")
	}

	return lo.Map(outputs, func(output *database.TrackedOutput, _ int) txmodels.TrackedOutput {
		return txmodels.TrackedOutput{
			TxID:       output.TxID,
			Vout:       output.Vout,
			SpendingTX: output.SpendingTX,
			UserID:     output.UserID,
			Satoshis:   output.Satoshis,
			CreatedAt:  output.CreatedAt,
			UpdatedAt:  output.UpdatedAt,
		}
	}), nil
}
