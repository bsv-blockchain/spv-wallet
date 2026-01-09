package webhooks

import (
	"context"

	"github.com/rs/zerolog"

	"github.com/bsv-blockchain/spv-wallet/engine"
	"github.com/bsv-blockchain/spv-wallet/engine/notifications"
)

type webhooksService interface {
	SubscribeWebhook(ctx context.Context, url, tokenHeader, token string) error
	UnsubscribeWebhook(ctx context.Context, url string) error
	GetWebhooks(ctx context.Context) ([]notifications.ModelWebhook, error)
}

// APIAdminWebhooks represents server with admin API endpoints
type APIAdminWebhooks struct {
	webhooks webhooksService
	logger   *zerolog.Logger
}

// NewAPIAdminWebhooks creates a new APIAdminWebhooks
func NewAPIAdminWebhooks(engine engine.ClientInterface, logger *zerolog.Logger) APIAdminWebhooks {
	log := logger.With().Str("api", "webhooks").Logger()
	return APIAdminWebhooks{
		webhooks: engine,
		logger:   &log,
	}
}
