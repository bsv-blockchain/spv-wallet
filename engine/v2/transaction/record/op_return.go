package record

import (
	"github.com/bsv-blockchain/go-sdk/script"
	trx "github.com/bsv-blockchain/go-sdk/transaction"
	"github.com/bsv-blockchain/spv-wallet/conv"
	"github.com/bsv-blockchain/spv-wallet/engine/v2/transaction"
	"github.com/bsv-blockchain/spv-wallet/engine/v2/transaction/errors"
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

	var bytes []byte

	// Find the OP_RETURN chunk
	for i, chunk := range chunks {
		if chunk.Op == script.OpRETURN {
			// The OP_RETURN chunk.Data contains: OP_RETURN_OPCODE + PUSH_LENGTH + DATA
			// Example: 0x6a 0x0b "hello world"
			// We need to skip the first two bytes and extract the data
			if len(chunk.Data) > 2 {
				// chunk.Data[0] = 0x6a (OP_RETURN opcode)
				// chunk.Data[1] = length of data
				// chunk.Data[2...] = actual data
				pushLength := int(chunk.Data[1])
				if len(chunk.Data) >= pushLength+2 {
					bytes = chunk.Data[2 : pushLength+2]
				}
			}

			// Also check for subsequent chunks (alternative format)
			for j := i + 1; j < len(chunks); j++ {
				if chunks[j].Op > script.OpPUSHDATA4 || chunks[j].Op == script.OpZERO {
					return nil, txerrors.ErrOnlyPushDataAllowed
				}
				bytes = append(bytes, chunks[j].Data...)
			}
			return bytes, nil
		}
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
