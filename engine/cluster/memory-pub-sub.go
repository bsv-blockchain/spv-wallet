package cluster

import (
	"context"
	"sync"

	"github.com/rs/zerolog"
)

// MemoryPubSub struct
type MemoryPubSub struct {
	ctx       context.Context
	callbacks map[string]func(data string)
	mu        sync.RWMutex
	debug     bool
	logger    *zerolog.Logger
	prefix    string
}

// NewMemoryPubSub create a new memory pub/sub client
// this is the default (mock) implementation for the internal pub/sub communications on standalone servers
// for clusters, use another solution, like RedisPubSub
func NewMemoryPubSub(ctx context.Context) (*MemoryPubSub, error) {
	return &MemoryPubSub{
		ctx:       ctx,
		callbacks: make(map[string]func(data string)),
	}, nil
}

// Logger returns the logger to use
func (m *MemoryPubSub) Logger() *zerolog.Logger {
	return m.logger
}

// Subscribe to a channel
func (m *MemoryPubSub) Subscribe(channel Channel, callback func(data string)) (func() error, error) {
	channelName := m.prefix + string(channel)
	m.mu.Lock()
	m.callbacks[channelName] = callback
	m.mu.Unlock()

	return func() error {
		m.mu.Lock()
		delete(m.callbacks, channelName)
		m.mu.Unlock()
		return nil
	}, nil
}

// Publish to a channel
func (m *MemoryPubSub) Publish(channel Channel, data string) error {
	channelName := m.prefix + string(channel)
	m.mu.RLock()
	callback, ok := m.callbacks[channelName]
	m.mu.RUnlock()

	if ok {
		callback(data)
	}

	return nil
}
