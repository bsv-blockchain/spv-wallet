package logging

import (
	"os"
	"sync"

	"github.com/rs/zerolog"
	"go.elastic.co/ecszerolog"
)

var (
	defaultLogger     *zerolog.Logger
	defaultLoggerOnce sync.Once
)

// GetDefaultLogger returns a singleton default logger instance.
// Uses sync.Once to ensure the logger is only created once, avoiding race conditions
// with ecszerolog.New() which modifies global zerolog state.
func GetDefaultLogger() *zerolog.Logger {
	defaultLoggerOnce.Do(func() {
		logger := ecszerolog.New(os.Stdout, ecszerolog.Level(zerolog.InfoLevel)).
			With().
			Timestamp().
			Caller().
			Str("application", "spv-wallet-default").
			Logger()
		defaultLogger = &logger
	})
	return defaultLogger
}
