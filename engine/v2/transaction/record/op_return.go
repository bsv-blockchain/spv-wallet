package record

import (
	"github.com/bsv-blockchain/go-sdk/script"
	trx "github.com/bsv-blockchain/go-sdk/transaction"

	"github.com/bsv-blockchain/spv-wallet/conv"
	"github.com/bsv-blockchain/spv-wallet/engine/v2/transaction"
	txerrors "github.com/bsv-blockchain/spv-wallet/engine/v2/transaction/errors"
	"github.com/bsv-blockchain/spv-wallet/engine/v2/transaction/txmodels"
	"github.com/bsv-blockchain/spv-wallet/models/bsv"
	"github.com/bsv-blockchain/spv-wallet/models/transaction/bucket"
)

func getDataFromOpReturn(lockingScript *script.Script) ([]byte, error) {
	if !lockingScript.IsData() {
		return nil, txerrors.ErrAnnotationMismatch
	}

	chunks, err := lockingScript.Chunks()
	if err != nil {
		return nil, txerrors.ErrParsingScript.Wrap(err)
	}

	// Find the OP_RETURN chunk
	for _, chunk := range chunks {
		if chunk.Op != script.OpRETURN {
			continue
		}

		// Since go-sdk v1.6.0 the OP_RETURN chunk's Data holds every byte after the OP_RETURN opcode
		// (e.g. PUSH_LENGTH + DATA), so parse it as a regular sequence of operations.
		ops, err := script.NewFromBytes(chunk.Data).ParseOps()
		if err != nil {
			return nil, txerrors.ErrParsingScript.Wrap(err)
		}

		var bytes []byte
		for _, op := range ops {
			if op.Op > script.OpPUSHDATA4 || op.Op == script.OpZERO {
				return nil, txerrors.ErrOnlyPushDataAllowed
			}
			bytes = append(bytes, op.Data...)
		}
		return bytes, nil
	}

	return nil, txerrors.ErrAnnotationMismatch
}

func processDataOutputs(tx *trx.Transaction, userID string, annotations *transaction.Annotations) ([]txmodels.NewOutput, error) {
	txID := tx.TxID().String()

	var dataOutputs []txmodels.NewOutput

	for vout, annotation := range annotations.Outputs {
		if annotation.Bucket != bucket.Data {
			continue
		}

		if len32, err := conv.IntToUint32(len(tx.Outputs)); err != nil {
			return nil, txerrors.ErrAnnotationIndexOutOfRange.Wrap(err)
		} else if vout >= len32 {
			return nil, txerrors.ErrAnnotationIndexOutOfRange
		}
		outpoint := bsv.Outpoint{TxID: txID, Vout: vout}

		lockingScript := tx.Outputs[vout].LockingScript

		data, err := getDataFromOpReturn(lockingScript)
		if err != nil {
			return nil, err
		}
		dataOutputs = append(dataOutputs, txmodels.NewOutputForData(outpoint, userID, data))
	}

	return dataOutputs, nil
}
