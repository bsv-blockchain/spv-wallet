package outlines

import (
	sdk "github.com/bsv-blockchain/go-sdk/transaction"
	"github.com/bsv-blockchain/spv-wallet/engine/spverrors"
	"github.com/bsv-blockchain/spv-wallet/engine/v2/transaction"
	txerrors "github.com/bsv-blockchain/spv-wallet/engine/v2/transaction/errors"
)

// TransactionSpec represents client provided specification for a transaction outline.
type TransactionSpec struct {
	Outputs OutputsSpec
	UserID  string
	Inputs  InputsSpec
}

func (t *TransactionSpec) evaluate(ctx *evaluationContext) (*sdk.Transaction, transaction.Annotations, error) {
	outputs, err := t.Outputs.evaluate(ctx)
	if err != nil {
		return nil, transaction.Annotations{}, spverrors.Wrapf(err, "failed to evaluate outputs")
	}

	inputs, change, err := t.Inputs.evaluate(ctx, outputs)
	if err != nil {
		return nil, transaction.Annotations{}, err
	}

	if change > 0 {
		outputs, err = addChangeOutput(ctx, outputs, change)
		if err != nil {
			return nil, transaction.Annotations{}, txerrors.ErrOutlineAddChangeOutput.Wrap(err)
		}
	}

	txOuts, outputsAnnotations := outputs.splitIntoTransactionOutputsAndAnnotations()
	txIns, inputsAnnotations := inputs.splitIntoTransactionInputsAndAnnotations()

	tx := sdk.NewTransaction()
	tx.Inputs = txIns
	tx.Outputs = txOuts

	annotations := transaction.Annotations{
		Inputs:  inputsAnnotations,
		Outputs: outputsAnnotations,
	}

	return tx, annotations, nil
}
