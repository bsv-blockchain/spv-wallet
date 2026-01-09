package logging

import (
	"os"
	"sync"

	"github.com/rs/zerolog"
	"go.elastic.co/ecszerolog"
)

var (
	loggerMutex sync.Mutex
)

// GetDefaultLogger generates and returns a new default logger instance.
// Uses a mutex to protect ecszerolog.New() which has non-thread-safe global state.
func GetDefaultLogger() *zerolog.Logger {
	loggerMutex.Lock()
	defer loggerMutex.Unlock()

	logger := ecszerolog.New(os.Stdout, ecszerolog.Level(zerolog.InfoLevel)).
		With().
		Timestamp().
		Caller().
		Str("application", "spv-wallet-default").
		Logger()

	return &logger
}
