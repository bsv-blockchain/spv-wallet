package data

import (
	"context"

	"github.com/bsv-blockchain/spv-wallet/engine/v2/data/datamodels"
)

// Repo is the interface that wraps the basic operations with data.
type Repo interface {
	FindForUser(ctx context.Context, id string, userID string) (*datamodels.Data, error)
}
