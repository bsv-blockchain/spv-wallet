package server

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rs/zerolog"
	"github.com/stretchr/testify/suite"

	"github.com/bsv-blockchain/spv-wallet/config"
	"github.com/bsv-blockchain/spv-wallet/models"
	"github.com/bsv-blockchain/spv-wallet/tests"
)

const (
	testXpubAuth = "xpub661MyMwAqRbcGpZVrSHU7EZ5Zwx5cNZmD5iLHPcg8MPnVcPdsApRi4Z27Mg3Zy53XYMKuJC5GiwECCFVNkhNgrBrfcA22YoJhasH7GcArNX"
)

// TestSuite is for testing the entire package using real/mocked services
type TestSuite struct {
	tests.TestSuite
}

// SetupSuite runs at the start of the suite
func (ts *TestSuite) SetupSuite() {
	ts.BaseSetupSuite()
}

// TearDownSuite runs after the suite finishes
func (ts *TestSuite) TearDownSuite() {
	ts.BaseTearDownSuite()
}

// SetupTest runs before each test
func (ts *TestSuite) SetupTest() {
	ts.BaseSetupTest()

	logger := zerolog.Nop()
	setupServerRoutes(ts.AppConfig, ts.SpvWalletEngine, ts.Router, &logger)
}

// TearDownTest runs after each test
func (ts *TestSuite) TearDownTest() {
	ts.BaseTearDownTest()
}

func TestTestSuite(t *testing.T) {
	suite.Run(t, new(TestSuite))
}

func (ts *TestSuite) TestAdminAuthentication() {
	ts.Run("no value", func() {
		w := httptest.NewRecorder()

		req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/api/"+config.APIVersion+"/admin/status", nil)
		ts.Require().NoError(err)
		ts.Require().NotNil(req)

		ts.Router.ServeHTTP(w, req)

		ts.Equal(http.StatusUnauthorized, w.Code)
	})

	ts.Run("false value", func() {
		w := httptest.NewRecorder()

		req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/api/"+config.APIVersion+"/admin/status", nil)
		ts.Require().NoError(err)
		ts.Require().NotNil(req)

		req.Header.Set(models.AuthHeader, testXpubAuth)

		ts.Router.ServeHTTP(w, req)

		ts.Equal(http.StatusUnauthorized, w.Code)
	})

	ts.Run("admin key", func() {
		w := httptest.NewRecorder()

		req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/api/"+config.APIVersion+"/admin/status", bytes.NewReader([]byte("test")))
		ts.Require().NoError(err)
		ts.Require().NotNil(req)

		req.Header.Set(models.AuthHeader, ts.AppConfig.Authentication.AdminKey)

		ts.Router.ServeHTTP(w, req)

		ts.Equal(http.StatusOK, w.Code)
	})
}

func (ts *TestSuite) TestApiAuthentication() {
	ts.Run("no value", func() {
		w := httptest.NewRecorder()

		req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/api/"+config.APIVersion+"/transactions", nil)
		ts.Require().NoError(err)
		ts.Require().NotNil(req)

		ts.Router.ServeHTTP(w, req)

		ts.Equal(http.StatusUnauthorized, w.Code)
	})

	ts.Run("false value", func() {
		w := httptest.NewRecorder()

		req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/api/"+config.APIVersion+"/transactions", nil)
		ts.Require().NoError(err)
		ts.Require().NotNil(req)

		req.Header.Set(models.AuthHeader, testXpubAuth)

		ts.Router.ServeHTTP(w, req)

		ts.Equal(http.StatusUnauthorized, w.Code)
	})

	ts.Run("valid value", func() {
		w := httptest.NewRecorder()

		xpub, err := ts.SpvWalletEngine.NewXpub(context.Background(), testXpubAuth)
		ts.Require().NoError(err)
		ts.Require().NotNil(xpub)

		req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/api/"+config.APIVersion+"/transactions", bytes.NewReader([]byte("test")))
		ts.Require().NoError(err)
		ts.Require().NotNil(req)

		req.Header.Set(models.AuthHeader, xpub.RawXpub())

		ts.Router.ServeHTTP(w, req)

		ts.Equal(http.StatusOK, w.Code)
	})
}

func (ts *TestSuite) TestBasicAuthentication() {
	ts.Run("no value", func() {
		w := httptest.NewRecorder()

		req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/", nil)
		ts.Require().NoError(err)
		ts.Require().NotNil(req)

		ts.Router.ServeHTTP(w, req)

		ts.Equal(http.StatusOK, w.Code)
	})

	ts.Run("non existing xpub", func() {
		w := httptest.NewRecorder()

		req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/", nil)
		ts.Require().NoError(err)
		ts.Require().NotNil(req)

		req.Header.Set(models.AuthHeader, testXpubAuth)

		ts.Router.ServeHTTP(w, req)

		ts.Equal(http.StatusOK, w.Code)
	})

	ts.Run("valid value", func() {
		w := httptest.NewRecorder()

		xpub, err := ts.SpvWalletEngine.NewXpub(context.Background(), testXpubAuth)
		ts.Require().NoError(err)
		ts.Require().NotNil(xpub)

		key, err := ts.SpvWalletEngine.NewAccessKey(context.Background(), xpub.RawXpub())
		ts.Require().NoError(err)
		ts.Require().NotNil(key)

		req, err := http.NewRequestWithContext(context.Background(), http.MethodGet, "/", bytes.NewReader([]byte("test")))
		ts.Require().NoError(err)
		ts.Require().NotNil(req)

		req.Header.Set(models.AuthHeader, xpub.RawXpub())

		ts.Router.ServeHTTP(w, req)

		ts.Equal(http.StatusOK, w.Code)
	})
}
