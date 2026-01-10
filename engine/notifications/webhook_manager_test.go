package notifications

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/jarcoal/httpmock"
	"github.com/stretchr/testify/assert"
)

type mockRepository struct {
	webhooks []ModelWebhook
	mu       sync.Mutex
}

func (r *mockRepository) Create(_ context.Context, url, tokenHeader, tokenValue string) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	model := newMockWebhookModel(url, tokenHeader, tokenValue)
	r.webhooks = append(r.webhooks, model)
	return nil
}

func (r *mockRepository) Save(_ context.Context, model ModelWebhook) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, w := range r.webhooks {
		if w.GetURL() == model.GetURL() {
			r.webhooks[i] = model
			return nil
		}
	}
	r.webhooks = append(r.webhooks, model)
	return nil
}

func (r *mockRepository) Delete(_ context.Context, model ModelWebhook) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	for i, w := range r.webhooks {
		if w.GetURL() == model.GetURL() {
			webhook := r.webhooks[i].(*mockModelWebhook)
			webhook.deleted = true
			r.webhooks[i] = webhook
			return nil
		}
	}
	return nil
}

func (r *mockRepository) GetAll(_ context.Context) ([]ModelWebhook, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	result := make([]ModelWebhook, len(r.webhooks))
	copy(result, r.webhooks)
	return result, nil
}

func (r *mockRepository) GetByURL(_ context.Context, url string) (ModelWebhook, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	for _, w := range r.webhooks {
		if w.GetURL() == url {
			return w, nil
		}
	}
	return nil, nil
}

func TestWebhookManager(t *testing.T) {
	t.Run("one webhook notifier previously subscribed", func(t *testing.T) {
		httpmock.Reset()
		httpmock.Activate()
		defer httpmock.Deactivate()

		client := newMockClient("http://localhost:8080")

		ctx, cancel := context.WithCancel(context.Background())

		n := NewNotifications(ctx, &nopLogger)
		repo := &mockRepository{webhooks: []ModelWebhook{newMockWebhookModel(client.url, "", "")}}

		manager := NewWebhookManager(ctx, &nopLogger, n, repo)
		time.Sleep(100 * time.Millisecond) // wait for manager to update notifiers
		defer manager.Stop()

		expected := []string{}
		for i := 0; i < 10; i++ {
			msg := fmt.Sprintf("msg-%d", i)
			n.Notify(newMockEvent(msg))
			expected = append(expected, msg)
		}

		time.Sleep(100 * time.Millisecond)
		cancel()

		client.assertEvents(t, expected)
		client.assertEventsWereSentInBatches(t, true)
	})

	t.Run("one webhook notifier - subscribe", func(t *testing.T) {
		httpmock.Reset()
		httpmock.Activate()
		defer httpmock.Deactivate()

		client := newMockClient("http://localhost:8080")

		ctx, cancel := context.WithCancel(context.Background())

		n := NewNotifications(ctx, &nopLogger)
		repo := &mockRepository{webhooks: []ModelWebhook{newMockWebhookModel(client.url, "", "")}}

		manager := NewWebhookManager(ctx, &nopLogger, n, repo)
		time.Sleep(100 * time.Millisecond)
		defer manager.Stop()

		_ = manager.Subscribe(ctx, client.url, "", "")
		time.Sleep(100 * time.Millisecond) // wait for manager to update notifiers

		expected := []string{}
		for i := 0; i < 10; i++ {
			msg := fmt.Sprintf("msg-%d", i)
			n.Notify(newMockEvent(msg))
			expected = append(expected, msg)
		}

		time.Sleep(100 * time.Millisecond)
		cancel()

		client.assertEvents(t, expected)
		client.assertEventsWereSentInBatches(t, true)
	})
}

func TestWebhookManager_StopCancelsContexts(t *testing.T) {
	httpmock.Activate()
	defer httpmock.DeactivateAndReset()

	t.Run("stop cancels all notifier contexts before stopping", func(t *testing.T) {
		client := newMockClient("http://example.com")

		ctx := context.Background()
		n := NewNotifications(ctx, &nopLogger)
		repo := &mockRepository{webhooks: []ModelWebhook{newMockWebhookModel(client.url, "", "")}}

		manager := NewWebhookManager(ctx, &nopLogger, n, repo)
		time.Sleep(100 * time.Millisecond)

		// Subscribe to create a notifier
		_ = manager.Subscribe(ctx, client.url, "", "")
		time.Sleep(100 * time.Millisecond)

		// Send some events
		for i := 0; i < 5; i++ {
			n.Notify(newMockEvent(fmt.Sprintf("msg-%d", i)))
		}

		time.Sleep(100 * time.Millisecond)

		// Stop should complete quickly (contexts are canceled before waiting)
		start := time.Now()
		manager.Stop()
		duration := time.Since(start)

		assert.Less(t, duration, 1*time.Second, "Stop should complete quickly with context cancellation")
	})
}
