// Package tests provides the base test suite for the entire package
package tests

import (
	"context"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/suite"

	"github.com/bsv-blockchain/spv-wallet/config"
	"github.com/bsv-blockchain/spv-wallet/engine"
	"github.com/bsv-blockchain/spv-wallet/engine/datastore"
	"github.com/bsv-blockchain/spv-wallet/engine/tester"
	"github.com/bsv-blockchain/spv-wallet/initializer"
	"github.com/bsv-blockchain/spv-wallet/logging"
	"github.com/bsv-blockchain/spv-wallet/server/middleware"
)

// TestSuite is for testing the entire package using real/mocked services
type TestSuite struct {
	AppConfig       *config.AppConfig      // App config
	Router          *gin.Engine            // Gin router with handlers
	Logger          zerolog.Logger         // Logger
	SpvWalletEngine engine.ClientInterface // SPV Wallet Engine
	suite.Suite                            // Extends the suite.Suite package
}

// BaseSetupSuite runs at the start of the suite
func (ts *TestSuite) BaseSetupSuite() {
	cfg := config.GetDefaultAppConfig()
	cfg.DebugProfiling = false
	cfg.Logging.Level = zerolog.LevelDebugValue
	cfg.Logging.Format = "console"
	cfg.CustomFeeUnit = &config.FeeUnitConfig{
		Satoshis: 1,
		Bytes:    1000,
	}
	cfg.Notifications.Enabled = false

	cfg.Db.Datastore.Engine = datastore.SQLite
	cfg.Db.SQLite.Shared = false
	cfg.Db.SQLite.MaxIdleConnections = 1
	cfg.Db.SQLite.MaxOpenConnections = 1
	cfg.Db.SQLite.DatabasePath = "file:spv-wallet-suite-test.db?mode=memory"

	ts.AppConfig = cfg
}

// BaseTearDownSuite runs after the suite finishes
func (ts *TestSuite) BaseTearDownSuite() {
	ts.T().Cleanup(func() {
		_ = os.Remove("datastore.db")
		_ = os.Remove("spv-wallet.db")
	})
}

// BaseSetupTest runs before each test
func (ts *TestSuite) BaseSetupTest() {
	// Load the services
	var err error
	ts.Logger = tester.Logger(ts.T())

	opts, err := initializer.ToEngineOptions(ts.AppConfig, ts.Logger)
	ts.Require().NoError(err)

	ts.SpvWalletEngine, err = engine.NewClient(context.Background(), opts...)
	ts.Require().NoError(err)

	logging.SetGinMode(gin.ReleaseMode)
	ginEngine := gin.New()
	ginEngine.Use(logging.GinMiddleware(ts.Logger), gin.Recovery())
	ginEngine.Use(middleware.AppContextMiddleware(ts.AppConfig, ts.SpvWalletEngine, ts.Logger))
	ginEngine.Use(middleware.CorsMiddleware())

	ts.Router = ginEngine
	ts.Require().NotNil(ts.Router)

	ts.Require().NoError(err)
}

// BaseTearDownTest runs after each test
func (ts *TestSuite) BaseTearDownTest() {
	if ts.SpvWalletEngine != nil {
		err := ts.SpvWalletEngine.Close(context.Background())
		ts.Require().NoError(err)
	}
}
