package admin

import (
	"github.com/rs/zerolog"

	"github.com/bsv-blockchain/spv-wallet/actions/v2/admin/contacts"
	"github.com/bsv-blockchain/spv-wallet/actions/v2/admin/transactions"
	"github.com/bsv-blockchain/spv-wallet/actions/v2/admin/users"
	"github.com/bsv-blockchain/spv-wallet/actions/v2/admin/webhooks"
	"github.com/bsv-blockchain/spv-wallet/engine"
)

// APIAdmin represents server with API endpoints
type APIAdmin struct {
	users.APIAdminUsers
	webhooks.APIAdminWebhooks
	transactions.APIAdminTransactions
	contacts.APIAdminContacts
}

// NewAPIAdmin creates a new APIAdmin
func NewAPIAdmin(spvWalletEngine engine.ClientInterface, logger *zerolog.Logger) APIAdmin {
	return APIAdmin{
		users.NewAPIAdminUsers(spvWalletEngine, logger),
		webhooks.NewAPIAdminWebhooks(spvWalletEngine, logger),
		transactions.NewAPIAdminTransactions(spvWalletEngine, logger),
		contacts.NewAPIAdminContacts(spvWalletEngine, logger),
	}
}
