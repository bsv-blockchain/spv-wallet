package operations

import (
	"context"

	"github.com/bsv-blockchain/spv-wallet/engine/v2/operations/operationsmodels"
	"github.com/bsv-blockchain/spv-wallet/models"
	"github.com/bsv-blockchain/spv-wallet/models/filter"
)

// Repo is an interface for operations repository.
type Repo interface {
	PaginatedForUser(ctx context.Context, userID string, page filter.Page) (*models.PagedResult[operationsmodels.Operation], error)
}
