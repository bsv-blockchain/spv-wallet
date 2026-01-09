package data

import (
	"github.com/rs/zerolog"

	"github.com/bsv-blockchain/spv-wallet/engine"
)

// APIData represents server with API endpoints
type APIData struct {
	engine engine.ClientInterface
	logger *zerolog.Logger
}

// NewAPIData creates a new server with API endpoints
func NewAPIData(engine engine.ClientInterface, log *zerolog.Logger) APIData {
	logger := log.With().Str("api", "data").Logger()

	return APIData{
		engine: engine,
		logger: &logger,
	}
}
