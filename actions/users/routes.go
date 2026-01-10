package users

import (
	"github.com/bsv-blockchain/spv-wallet/server/handlers"
)

// RegisterRoutes creates the specific package routes in RESTful style
func RegisterRoutes(handlersManager *handlers.Manager) {
	group := handlersManager.Group(handlers.GroupAPI, "/users/current")
	group.GET("", handlers.AsUser(get))
	group.PATCH("", handlers.AsUser(update))
}
