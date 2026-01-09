package mapping

import (
	"github.com/bsv-blockchain/spv-wallet/api"
	"github.com/bsv-blockchain/spv-wallet/engine/v2/transaction/txmodels"
)

// RecordedOutline maps domain RecordedOutline to api.ModelsRecordedOutline.
func RecordedOutline(r *txmodels.RecordedOutline) api.ModelsRecordedOutline {
	return api.ModelsRecordedOutline{
		TxID: r.TxID,
	}
}
