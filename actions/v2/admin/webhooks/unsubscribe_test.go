package webhooks_test

import (
	"testing"

	"github.com/bsv-blockchain/spv-wallet/actions/testabilities"
	"github.com/bsv-blockchain/spv-wallet/engine/spverrors"
	testengine "github.com/bsv-blockchain/spv-wallet/engine/testabilities"
)

func TestUnsubscribeWebhookHappyPath(t *testing.T) {
	t.Run("unsubscribe webhook", func(t *testing.T) {
		// given:
		given, then := testabilities.NewOf(testabilities.Given(t), t)
		cleanup := given.StartedSPVWalletWithConfiguration(
			testengine.WithNotificationsEnabled(),
			testengine.WithV2(),
		)
		defer cleanup()
		client := given.HttpClient().ForAdmin()

		// when:
		res, _ := client.R().
			SetBody(
				map[string]string{
					urlField:      "http://localhost:8080",
					"tokenHeader": "Auth1",
					"tokenValue":  "123",
				},
			).
			Post(webhookAPIURL)

		then.Response(res).IsOK()

		res, _ = client.R().
			Get(webhookAPIURL)

		then.Response(res).
			IsOK().
			WithJSONf(`[
                {"url": "http://localhost:8080", "banned": false}
            ]`)

		res, _ = client.
			R().
			SetBody(
				map[string]string{
					urlField: "http://localhost:8080",
				},
			).
			Delete(webhookAPIURL)

		then.Response(res).IsOK()

		res, _ = client.R().
			Get(webhookAPIURL)

		// then:
		then.Response(res).
			IsOK().
			WithJSONf(`[]`)
	})
}

func TestUnsubscribeWebhookErrorPath(t *testing.T) {
	t.Run("unsubscribe with invalid JSON returns bad request", func(t *testing.T) {
		// given:
		given, then := testabilities.NewOf(testabilities.Given(t), t)
		cleanup := given.StartedSPVWalletWithConfiguration(
			testengine.WithNotificationsEnabled(),
			testengine.WithV2(),
		)
		defer cleanup()
		client := given.HttpClient().ForAdmin()

		// when:
		res, _ := client.
			R().
			SetBody("{invalid json}").
			Delete(webhookAPIURL)

		// then:
		then.Response(res).
			HasStatus(400).
			WithJSONMatching(`{
				"code": "{{ .code }}",
				"message": "{{ .message }}"
			}`, map[string]any{
				codeField:    spverrors.ErrCannotBindRequest.Code,
				messageField: spverrors.ErrCannotBindRequest.Message,
			})
	})

	t.Run("unsubscribe with notification disabled returns error 404", func(t *testing.T) {
		// given:
		given, then := testabilities.NewOf(testabilities.Given(t), t)
		cleanup := given.StartedSPVWalletWithConfiguration(
			testengine.WithV2(),
		)
		defer cleanup()
		client := given.HttpClient().ForAdmin()

		// when:
		res, _ := client.
			R().
			SetBody(
				map[string]string{
					urlField: "http://localhost:8080",
				},
			).
			Delete(webhookAPIURL)

		// then:
		then.Response(res).
			HasStatus(404).
			WithJSONMatching(`{
				"code": "{{ .code }}",
				"message": "{{ .message }}"
			}`, map[string]any{
				codeField:    spverrors.ErrNotificationsDisabled.Code,
				messageField: spverrors.ErrNotificationsDisabled.Message,
			})
	})

	t.Run("unsubscribe non-existent webhook returns internal error", func(t *testing.T) {
		// given:
		given, then := testabilities.NewOf(testabilities.Given(t), t)
		cleanup := given.StartedSPVWalletWithConfiguration(
			testengine.WithNotificationsEnabled(),
			testengine.WithV2(),
		)
		defer cleanup()
		client := given.HttpClient().ForAdmin()

		// when:
		res, _ := client.
			R().
			SetBody(
				map[string]string{
					urlField: "http://nonexistent.com",
				},
			).
			Delete(webhookAPIURL)

		// then:
		then.Response(res).
			HasStatus(404).
			WithJSONMatching(`{
				"code": "{{ .code }}",
				"message": "{{ .message }}"
			}`, map[string]any{
				codeField:    spverrors.ErrWebhookSubscriptionNotFound.Code,
				messageField: spverrors.ErrWebhookSubscriptionNotFound.Message,
			})
	})

	t.Run("unsubscribe with missing URL returns bad request", func(t *testing.T) {
		// given:
		given, then := testabilities.NewOf(testabilities.Given(t), t)
		cleanup := given.StartedSPVWalletWithConfiguration(
			testengine.WithNotificationsEnabled(),
			testengine.WithV2(),
		)
		defer cleanup()
		client := given.HttpClient().ForAdmin()

		// when:
		res, _ := client.
			R().
			SetBody(map[string]string{}).
			Delete(webhookAPIURL)

		// then:
		then.Response(res).
			HasStatus(400).
			WithJSONMatching(`{
				"code": "{{ .code }}",
				"message": "{{ .message }}"
			}`, map[string]any{
				codeField:    spverrors.ErrWebhookUrlRequired.Code,
				messageField: spverrors.ErrWebhookUrlRequired.Message,
			})
	})

	t.Run("unsubscribe with incorrect URL returns bad request", func(t *testing.T) {
		// given:
		given, then := testabilities.NewOf(testabilities.Given(t), t)
		cleanup := given.StartedSPVWalletWithConfiguration(
			testengine.WithNotificationsEnabled(),
			testengine.WithV2(),
		)
		defer cleanup()
		client := given.HttpClient().ForAdmin()

		// when:
		res, _ := client.
			R().
			SetBody(map[string]string{
				urlField: "http://test.com/%",
			}).
			Delete(webhookAPIURL)

		// then:
		then.Response(res).
			HasStatus(400).
			WithJSONMatching(`{
				"code": "{{ .code }}",
				"message": "{{ .message }}"
			}`, map[string]any{
				codeField:    spverrors.ErrWebhookUrlInvalid.Code,
				messageField: spverrors.ErrWebhookUrlInvalid.Message,
			})
	})
}
