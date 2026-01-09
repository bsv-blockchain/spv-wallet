package chain

import (
	"context"
	"net/url"

	"github.com/bsv-blockchain/go-paymail/spv"
	sdk "github.com/bsv-blockchain/go-sdk/transaction"

	chainmodels "github.com/bsv-blockchain/spv-wallet/engine/chain/models"
	"github.com/bsv-blockchain/spv-wallet/models"
	"github.com/bsv-blockchain/spv-wallet/models/bsv"
)

// ARCService for querying ARC server.
type ARCService interface {
	QueryTransaction(ctx context.Context, txID string) (*chainmodels.TXInfo, error)
	GetFeeUnit(ctx context.Context) (*bsv.FeeUnit, error)
	Broadcast(ctx context.Context, tx *sdk.Transaction) (*chainmodels.TXInfo, error)
}

// BHSService for querying BHS server.
type BHSService interface {
	GetMerkleRoots(ctx context.Context, query url.Values) (*models.MerkleRootsBHSResponse, error)
	VerifyMerkleRoots(ctx context.Context, merkleRoots []*spv.MerkleRootConfirmationRequestItem) (bool, error)
	HealthcheckBHS(ctx context.Context) error
}

// Service related to the chain.
type Service interface {
	ARCService
	BHSService
}
