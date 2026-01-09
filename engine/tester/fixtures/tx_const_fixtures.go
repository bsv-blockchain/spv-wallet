package fixtures

import "github.com/bsv-blockchain/spv-wallet/models/bsv"

// DefaultFeeUnit is the default fee unit used in the tests.
var DefaultFeeUnit = bsv.FeeUnit{
	Satoshis: 1,
	Bytes:    1000,
}
