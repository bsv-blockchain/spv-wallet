package mapping

import (
	"github.com/samber/lo"

	"github.com/bsv-blockchain/spv-wallet/api"
	"github.com/bsv-blockchain/spv-wallet/engine/notifications"
)

// MapToModelsWebhooks converts a slice of ModelWebhook to ModelsWebhooks
func MapToModelsWebhooks(webhooks []notifications.ModelWebhook) api.ModelsWebhooks {
	if webhooks == nil {
		return nil
	}

	return lo.Map(webhooks, MapToModelsWebhook)
}

// MapToModelsWebhook converts a single ModelWebhook to ModelsWebhook
func MapToModelsWebhook(w notifications.ModelWebhook, _ int) api.ModelsWebhook {
	if w == nil {
		return api.ModelsWebhook{}
	}

	return api.ModelsWebhook{
		Url:    w.GetURL(),
		Banned: w.Banned(),
	}
}
