package outlines

import (
	primitives "github.com/bsv-blockchain/go-sdk/primitives/ec"
	"github.com/bsv-blockchain/go-sdk/script"
	sdk "github.com/bsv-blockchain/go-sdk/transaction"
	"github.com/bsv-blockchain/go-sdk/transaction/template/p2pkh"
	pmerrors "github.com/bsv-blockchain/spv-wallet/engine/paymail/errors"
	"github.com/bsv-blockchain/spv-wallet/engine/spverrors"
	"github.com/bsv-blockchain/spv-wallet/engine/v2/keys/type42"
	"github.com/bsv-blockchain/spv-wallet/engine/v2/transaction"
	"github.com/bsv-blockchain/spv-wallet/models/bsv"
	"github.com/bsv-blockchain/spv-wallet/models/transaction/bucket"
)

func addChangeOutput(ctx *evaluationContext, outputs annotatedOutputs, change bsv.Satoshis) (annotatedOutputs, error) {
	userPubKey, err := ctx.UserPubKey()
	if err != nil {
		return nil, spverrors.Wrapf(err, "failed to get user public key")
	}

	lockingScript, customInstructions, err := lockingScriptForChangeOutput(userPubKey)
	if err != nil {
		return nil, spverrors.Wrapf(err, "failed to create locking script for change output")
	}
	changeOutput := &annotatedOutput{
		OutputAnnotation: &transaction.OutputAnnotation{
			Bucket:             bucket.BSV,
			CustomInstructions: &customInstructions,
		},
		TransactionOutput: &sdk.TransactionOutput{
			LockingScript: lockingScript,
			Satoshis:      uint64(change),
		},
	}

	return append(outputs, changeOutput), nil
}

func lockingScriptForChangeOutput(pubKey *primitives.PublicKey) (*script.Script, bsv.CustomInstructions, error) {
	dest, err := type42.NewDestinationWithRandomReference(pubKey)
	if err != nil {
		return nil, nil, pmerrors.ErrPaymentDestination.Wrap(err)
	}

	address, err := script.NewAddressFromPublicKey(dest.PubKey, true)
	if err != nil {
		return nil, nil, pmerrors.ErrPaymentDestination.Wrap(err)
	}

	lockingScript, err := p2pkh.Lock(address)
	if err != nil {
		return nil, nil, pmerrors.ErrPaymentDestination.Wrap(err)
	}

	customInstructions := bsv.CustomInstructions{
		{
			Type:        "type42",
			Instruction: dest.DerivationKey,
		},
	}

	return lockingScript, customInstructions, nil
}
