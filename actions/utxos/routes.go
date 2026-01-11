package utxos

import (
	"github.com/bsv-blockchain/spv-wallet/server/handlers"
)

// RegisterRoutes creates the specific package routes
func RegisterRoutes(handlersManager *handlers.Manager) {
	group := handlersManager.Group(handlers.GroupAPI, "/utxos")
	group.GET("", handlers.AsUser(search))
}
