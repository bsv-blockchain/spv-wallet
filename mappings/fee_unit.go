package mappings

import (
	"github.com/bsv-blockchain/spv-wallet/models/bsv"
	"github.com/bsv-blockchain/spv-wallet/models/response"
)

// MapToFeeUnitContract will map the fee-unit model from spv-wallet to the spv-wallet-models contract
func MapToFeeUnitContract(fu *bsv.FeeUnit) (fc *response.FeeUnit) {
	if fu == nil {
		return nil
	}

	return &response.FeeUnit{
		Satoshis: fu.Satoshis,
		Bytes:    fu.Bytes,
	}
}

// MapFeeUnitModelToEngine will map the fee-unit model from spv-wallet-models to the spv-wallet contract
func MapFeeUnitModelToEngine(fu *response.FeeUnit) (fc *bsv.FeeUnit) {
	if fu == nil {
		return nil
	}

	return &bsv.FeeUnit{
		Satoshis: fu.Satoshis,
		Bytes:    fu.Bytes,
	}
}
