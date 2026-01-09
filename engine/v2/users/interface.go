package users

import (
	"context"

	"github.com/bsv-blockchain/spv-wallet/engine/v2/users/usersmodels"
	"github.com/bsv-blockchain/spv-wallet/models/bsv"
	"github.com/bsv-blockchain/spv-wallet/models/transaction/bucket"
)

// UserRepo is an interface for users repository.
type UserRepo interface {
	Exists(ctx context.Context, userID string) (bool, error)
	GetIDByPubKey(ctx context.Context, pubKey string) (string, error)
	Get(ctx context.Context, userID string) (*usersmodels.User, error)
	Create(ctx context.Context, newUser *usersmodels.NewUser) (*usersmodels.User, error)
	GetBalance(ctx context.Context, userID string, name bucket.Name) (bsv.Satoshis, error)
	Delete(ctx context.Context, userID string) error
}
