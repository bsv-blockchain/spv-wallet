package notifications

import (
	"context"
	"sync"
	"time"

	"github.com/rs/zerolog"

	"github.com/bsv-blockchain/spv-wallet/models"
)

const lengthOfInputChannel = 100

// Notifications - service for sending events to multiple notifiers
type Notifications struct {
	inputChannel   chan *models.RawEvent
	outputChannels *sync.Map //[string, chan *Event]
	burstLogger    *zerolog.Logger
	mu             sync.RWMutex
	closed         bool
	wg             sync.WaitGroup
}

// AddNotifier - add notifier by key
func (n *Notifications) AddNotifier(key string, ch chan *models.RawEvent) {
	n.outputChannels.Store(key, ch)
}

// RemoveNotifier - remove notifier by key
func (n *Notifications) RemoveNotifier(key string) {
	n.outputChannels.Delete(key)
}

// Notify - send event to all notifiers
func (n *Notifications) Notify(event *models.RawEvent) {
	n.mu.RLock()
	defer n.mu.RUnlock()
	if n.closed {
		return
	}
	n.inputChannel <- event
}

// exchange - exchange events between input and output channels, uses fan-out pattern
func (n *Notifications) exchange(ctx context.Context) {
	defer n.wg.Done()
	for {
		select {
		case event, ok := <-n.inputChannel:
			if !ok {
				// Channel closed, exit goroutine
				return
			}
			n.outputChannels.Range(func(_, value any) bool {
				ch := value.(chan *models.RawEvent)
				n.sendEventToChannel(ch, event)
				return true
			})
		case <-ctx.Done():
			return
		}
	}
}

// sendEventToChannel - non blocking send event to channel
func (n *Notifications) sendEventToChannel(ch chan *models.RawEvent, event *models.RawEvent) {
	select {
	case ch <- event:
		// Successfully sent event
	default:
		n.burstLogger.Warn().Msg("Failed to send event to channel")
	}
}

// Close - stops the notification service and cleans up resources
func (n *Notifications) Close() error {
	n.mu.Lock()
	if n.closed {
		n.mu.Unlock()
		return nil
	}
	n.closed = true
	if n.inputChannel != nil {
		close(n.inputChannel)
	}
	n.mu.Unlock()

	// Wait for the exchange goroutine to finish with timeout
	// Channel is closed, so goroutine should exit quickly
	done := make(chan struct{})
	go func() {
		n.wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		return nil
	case <-time.After(200 * time.Millisecond):
		// Log warning but don't return error - allow cleanup to continue
		n.burstLogger.Warn().Msg("timeout waiting for notification goroutines to finish")
		return nil
	}
}

// NewNotifications - creates a new instance of Notifications
func NewNotifications(ctx context.Context, parentLogger *zerolog.Logger) *Notifications {
	burstLogger := parentLogger.With().Logger().Sample(&zerolog.BurstSampler{
		Burst:  3,
		Period: 30 * time.Second,
	})
	n := &Notifications{
		inputChannel:   make(chan *models.RawEvent, lengthOfInputChannel),
		outputChannels: new(sync.Map),
		burstLogger:    &burstLogger,
	}

	n.wg.Add(1)
	go n.exchange(ctx)

	return n
}
