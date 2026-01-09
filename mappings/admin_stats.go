package mappings

import (
	"github.com/bsv-blockchain/spv-wallet/engine"
	"github.com/bsv-blockchain/spv-wallet/models/response"
)

// MapToAdminStatsContract will map the model from spv-wallet to the spv-wallet-models contract
func MapToAdminStatsContract(s *engine.AdminStats) *response.AdminStats {
	if s == nil {
		return nil
	}

	return &response.AdminStats{
		Balance:            s.Balance,
		Destinations:       s.Destinations,
		PaymailAddresses:   s.PaymailAddresses,
		Transactions:       s.Transactions,
		TransactionsPerDay: s.TransactionsPerDay,
		Utxos:              s.Utxos,
		UtxosPerType:       s.UtxosPerType,
		XPubs:              s.XPubs,
	}
}
