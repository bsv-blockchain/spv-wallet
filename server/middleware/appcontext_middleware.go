package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"

	"github.com/bsv-blockchain/spv-wallet/config"
	"github.com/bsv-blockchain/spv-wallet/engine"
	"github.com/bsv-blockchain/spv-wallet/server/reqctx"
)

// AppContextMiddleware is a middleware that sets the appConfig, engine and logger in the request context
func AppContextMiddleware(appConfig *config.AppConfig, engine engine.ClientInterface, logger zerolog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		reqctx.SetAppConfig(c, appConfig)
		reqctx.SetEngine(c, engine)
		reqctx.SetLogger(c, &logger)

		c.Next()
	}
}
