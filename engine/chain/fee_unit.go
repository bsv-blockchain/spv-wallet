package chain

import (
	"context"

	"github.com/bsv-blockchain/spv-wallet/engine/chain/errors"
	"github.com/bsv-blockchain/spv-wallet/models/bsv"
)

// GetFeeUnit returns the current fee unit from the ARC policy.
func (s *chainService) GetFeeUnit(ctx context.Context) (*bsv.FeeUnit, error) {
	policy, err := s.GetPolicy(ctx)
	if err != nil {
		return nil, chainerrors.ErrGetFeeUnit.Wrap(err)
	}

	return &bsv.FeeUnit{
		Satoshis: policy.Content.MiningFee.Satoshis,
		Bytes:    policy.Content.MiningFee.Bytes,
	}, nil
}
