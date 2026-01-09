package utils

import (
	ec "github.com/bsv-blockchain/go-sdk/primitives/ec"
	trx "github.com/bsv-blockchain/go-sdk/transaction"
	sighash "github.com/bsv-blockchain/go-sdk/transaction/sighash"
	"github.com/bsv-blockchain/go-sdk/transaction/template/p2pkh"
	"github.com/bsv-blockchain/spv-wallet/engine/spverrors"
)

// GetUnlockingScript will generate an unlocking script
func GetUnlockingScript(tx *trx.Transaction, inputIndex uint32, privateKey *ec.PrivateKey) (*p2pkh.P2PKH, error) {
	sigHashFlags := sighash.AllForkID

	sc, err := p2pkh.Unlock(privateKey, &sigHashFlags)
	if err != nil {
		return nil, spverrors.Wrapf(err, "failed to create unlocking script")
	}

	return sc, nil
}
