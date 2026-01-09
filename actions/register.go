package actions

import (
	accesskeys "github.com/bsv-blockchain/spv-wallet/actions/access_keys"
	"github.com/bsv-blockchain/spv-wallet/actions/admin"
	"github.com/bsv-blockchain/spv-wallet/actions/base"
	"github.com/bsv-blockchain/spv-wallet/actions/contacts"
	"github.com/bsv-blockchain/spv-wallet/actions/merkleroots"
	"github.com/bsv-blockchain/spv-wallet/actions/paymails"
	"github.com/bsv-blockchain/spv-wallet/actions/sharedconfig"
	"github.com/bsv-blockchain/spv-wallet/actions/transactions"
	"github.com/bsv-blockchain/spv-wallet/actions/users"
	"github.com/bsv-blockchain/spv-wallet/actions/utxos"
	"github.com/bsv-blockchain/spv-wallet/server/handlers"
)

// Register collects all the action's routes and registers them using the handlersManager
func Register(handlersManager *handlers.Manager) {
	admin.RegisterRoutes(handlersManager)
	base.RegisterRoutes(handlersManager)
	accesskeys.RegisterRoutes(handlersManager)
	transactions.RegisterRoutes(handlersManager)
	utxos.RegisterRoutes(handlersManager)
	users.RegisterRoutes(handlersManager)
	paymails.RegisterRoutes(handlersManager)
	sharedconfig.RegisterRoutes(handlersManager)
	merkleroots.RegisterRoutes(handlersManager)
	contacts.RegisterRoutes(handlersManager)
}
